package devices

import (
	"github.com/dansouza/gizmo/pkg/devices/v0001p0001v0001"
	"github.com/dansouza/gizmo/pkg/devices/v0001p0003v0001"
	"github.com/dansouza/gizmo/pkg/devices/v106Bp1570v00F1"
	"github.com/dansouza/gizmo/pkg/registry"
)

// Register registers all supported devices with the registry
func Register() {
	registry.DeviceRegistry.Register(v0001p0001v0001.Device)
	registry.DeviceRegistry.Register(v0001p0003v0001.Device)
	registry.DeviceRegistry.Register(v106Bp1570v00F1.Device)
}
