package v0001p0001v0001

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/dansouza/gizmo/pkg/device"
)

// DeviceHandler is a wrapper around device.Device to allow adding a custom
// Parse() method
type DeviceHandler struct {
	device.GenericDevice
}

// Device encapsulates the device information
var Device = DeviceHandler{
	device.GenericDevice{
		VendorID:  0x1,
		ProductID: 0x1,
		VersionID: 0x1,
		Name:      "HITACHI Blue Lightsaber",
	},
}

// Parse parses data according to the device
func (d DeviceHandler) Parse(r io.Reader, w io.Writer) (int, error) {
	var input []byte
	var err error

	// reads from r, writes to w
	if input, err = ioutil.ReadAll(r); err != nil {
		return 0, err
	}

	return fmt.Fprintf(w, "returned blue lightsaber stuff from device vendor=%X productID=%X version=%X from input: %s\n", d.GetVendorID(), d.GetProductID(), d.GetVersionID(), string(input))
}
