package registry

import (
	"fmt"
	"io"
	"sync"
)

var (
	// ErrAlreadyRegistered is returned when the same device is registered twice
	ErrAlreadyRegistered = fmt.Errorf("already registered")
	// ErrDeviceNotFound is returned when a device is not found in the registry
	ErrDeviceNotFound = fmt.Errorf("device not found")
)

// Device encapsulates a hardware device
type Device interface {
	GetVendorID() uint16
	GetProductID() uint16
	GetVersionID() uint16
	GetName() string
	Parse(io.Reader, io.Writer) (int, error)
}

type registry struct {
	index map[string]Device
	sync.RWMutex
}

// DeviceRegistry is the public interface for the device registry
var DeviceRegistry = newRegistry()

func newRegistry() *registry {
	r := registry{}
	r.index = make(map[string]Device)
	return &r
}

func hash(vendorID, productID, version uint16) string {
	return fmt.Sprintf("v%X_p%X_v%X", vendorID, productID, version)
}

func hashDevice(device Device) string {
	return hash(device.GetVendorID(), device.GetProductID(), device.GetVersionID())
}

func (r *registry) Register(device Device) error {
	r.Lock()
	defer r.Unlock()

	h := hashDevice(device)
	if _, ok := r.index[h]; ok {
		return ErrAlreadyRegistered
	}
	r.index[h] = device
	return nil
}

func (r *registry) Lookup(vendorID, productID, version uint16) (Device, error) {
	r.RLock()
	defer r.RUnlock()

	h := hash(vendorID, productID, version)
	if device, ok := r.index[h]; ok {
		return device, nil
	}
	return nil, ErrDeviceNotFound
}

func (r *registry) Enumerate(callback func(Device) bool) {
	r.RLock()
	defer r.RUnlock()

	for _, device := range r.index {
		if !callback(device) {
			break
		}
	}
}
