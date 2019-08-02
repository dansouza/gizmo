package device

// GenericDevice encapsulates the basic information that every device shares
type GenericDevice struct {
	VendorID  uint16
	ProductID uint16
	VersionID uint16
	Name      string
}

// GetVendorID returns the device's Vendor ID
func (d GenericDevice) GetVendorID() uint16 {
	return d.VendorID
}

// GetProductID returns the device's Product ID
func (d GenericDevice) GetProductID() uint16 {
	return d.ProductID
}

// GetVersionID returns the device's version number
func (d GenericDevice) GetVersionID() uint16 {
	return d.VersionID
}

// GetName returns the consumer-friendly name of the device
func (d GenericDevice) GetName() string {
	return d.Name
}
