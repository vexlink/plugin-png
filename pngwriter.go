package main

import (
	"github.com/vexlink/vexlink"
	"image"
	"image/png"
	"os"
)

type PNGWriter struct{}

func (this *PNGWriter) GetInformation() vexlink.NodeInformation {
	return vexlink.NodeInformation{
		Options: map[string]interface{}{
			"path": "/tmp/image.png",
		},
		Inputs: []string{
			"image",
		},
		Outputs: []string{},
	}
}

func (this *PNGWriter) Run(options map[string]interface{}, input chan map[string]interface{}, output chan map[string]interface{}, stop chan struct{}) {
	path := options["path"].(string)

	for {
		// Process channels
		select {
		case <-stop: // Break when stop signal is received
			return

		case in := <-input: // Handle input
			if in["image"] != nil {
				img := in["image"].(*image.Image)
				f, _ := os.Create(path)
				png.Encode(f, *img)
				f.Close()
			}

			output <- nil
		}

	}
}
