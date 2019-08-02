package main

import (
	"fmt"
	"os"

	"github.com/dansouza/gizmo/pkg/devices"
	"github.com/dansouza/gizmo/pkg/registry"
)

func main() {
	devices.Register()
	fmt.Printf("recognized devices:\n")
	registry.DeviceRegistry.Enumerate(func(device registry.Device) bool {
		fmt.Printf("[>] vendor=%04X\tproductID=%04X\tversion=%04X\t%s\n",
			device.GetVendorID(),
			device.GetProductID(),
			device.GetVersionID(),
			device.GetName())
		return true
	})

	doubleEndedLightSaber, err := registry.DeviceRegistry.Lookup(0x106b, 0x1570, 0xf1)
	if err != nil {
		fmt.Printf("error looking up double ended lightsaber: %s\n", err)
		os.Exit(-1)
	}

	fmt.Printf("[*] found preferred lightsaber: %s\n", doubleEndedLightSaber.GetName())
	fmt.Printf("[*] reading from device...\n")

	input, err := os.Open("/proc/version")
	if err != nil {
		fmt.Printf("error opening input: %s\n", err)
		os.Exit(-1)
	}
	wrote, err := doubleEndedLightSaber.Parse(input, os.Stdout)

	fmt.Printf("[*] device wrote %d bytes.\n", wrote)

	return
}
