package graphics

import (
	"github.com/gonutz/d3d9"
	"log"
)

type DX9Graphics struct {
	D3D9 *d3d9.Direct3D
}

func (dx9 *DX9Graphics) Create() (err error) {
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
	
	return err
}

func (dx9 *DX9Graphics) Release() {
	defer dx9.D3D9.Release()
}
