package devices

import (
	"github.com/dansouza/gizmo/pkg/devices/v0001p0001v0001"
	"github.com/dansouza/gizmo/pkg/devices/v0001p0003v0001"
	"github.com/dansouza/gizmo/pkg/devices/v106Bp1570v00F1"
	"github.com/dansouza/gizmo/pkg/registry"
)

// GetDevice looks up a device by vendor/product/version and returns it.
// If not found, returns an registry.ErrDeviceNotFound
func GetDevice(vendorID, productID, version uint16) (registry.Device, error) {
	return registry.DeviceRegistry.Lookup(vendorID, productID, version)
}

// Register registers all supported devices with the registry
func Register() {
	registry.DeviceRegistry.Register(v0001p0001v0001.Device{})
	registry.DeviceRegistry.Register(v0001p0003v0001.Device{})
	registry.DeviceRegistry.Register(v106Bp1570v00F1.Device{})
}
