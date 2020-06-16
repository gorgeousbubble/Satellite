// +build windows

package graphics

import (
	"github.com/gonutz/d3d9"
	"log"
)

type DX9Graphics struct {
	D3D9   *d3d9.Direct3D
	Device *d3d9.Device
	Handle d3d9.HWND
	Width  uint32
	Height uint32
}

// DX9Graphics Create function
// this function is mainly used to Create an instance for DX9 Graphics
// when you use DX9, this function should be called before other operations
// first you should create a DX9Graphics instance, then call 'Create' function
// return err indicate the success or failure function execute
func (dx9 *DX9Graphics) Create() (err error) {
	var vp uint32
	// create d3d9 instance
	dx9.D3D9, err = d3d9.Create(d3d9.SDK_VERSION)
	if err != nil {
		log.Println("Error create d3d9 instance:", err)
		return err
	}
	// get d3d9 hardware caps
	caps, err := dx9.D3D9.GetDeviceCaps(d3d9.ADAPTER_DEFAULT, d3d9.DEVTYPE_HAL)
	if err != nil {
		log.Println("Error get device caps:", err)
		return err
	}
	// check d3d9 hardware calc capacity
	if caps.DevCaps&d3d9.DEVCAPS_HWTRANSFORMANDLIGHT != uint32(0) {
		vp = d3d9.CREATE_HARDWARE_VERTEXPROCESSING
	} else {
		vp = d3d9.CREATE_SOFTWARE_VERTEXPROCESSING
	}
	// create d3d9 device
	dx9.Device, _, err = dx9.D3D9.CreateDevice(d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		dx9.Handle,
		vp,
		d3d9.PRESENT_PARAMETERS{
			BackBufferWidth:        dx9.Width,
			BackBufferHeight:       dx9.Height,
			BackBufferFormat:       d3d9.FMT_A8R8G8B8,
			BackBufferCount:        1,
			MultiSampleType:        d3d9.MULTISAMPLE_NONE,
			MultiSampleQuality:     0,
			SwapEffect:             d3d9.SWAPEFFECT_DISCARD,
			HDeviceWindow:          dx9.Handle,
			Windowed:               1,
			EnableAutoDepthStencil: 1,
			AutoDepthStencilFormat: d3d9.FMT_D24S8,
			Flags:                  0,
			PresentationInterval:   d3d9.PRESENT_INTERVAL_DEFAULT,
		})
	if err != nil {
		log.Println("Error create device:", err)
		return err
	}
	return err
}

// DX9Graphics Release function
// this function is mainly used to Release an instance for DX9 Graphics
// when you use DX9, this function should be called as last function to release all resources
// first you should release all DX9Graphics resources, then call 'Release' function
// return err indicate the success or failure function execute
func (dx9 *DX9Graphics) Release() {
	defer dx9.D3D9.Release()
}
