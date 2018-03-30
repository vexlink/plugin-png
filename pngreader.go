package main

import (
	"github.com/vexlink/vexlink"
	"image/png"
	"os"
)

type PNGReader struct{}

func (this *PNGReader) GetInformation() vexlink.NodeInformation {
	return vexlink.NodeInformation{
		Options: map[string]interface{}{
			"path": "/tmp/image.png",
		},
		Inputs: []string{},
		Outputs: []string{
			"image",
		},
	}
}

func (this *PNGReader) Run(options map[string]interface{}, input chan map[string]interface{}, output chan map[string]interface{}, stop chan struct{}) {
	path, _ := options["path"].(string)

	for {
		// Process channels
		select {
		case <-stop: // Break when stop signal is received
			return

		case <-input: // Handle input
			f, err := os.Open(path)
			defer f.Close()

			if err == nil {
				img, _ := png.Decode(f)
				output <- map[string]interface{}{"image": img}
			} else {
				panic(err)
			}
		}

	}
}
