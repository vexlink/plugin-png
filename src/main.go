package main

import "gopkg.in/vexlink/vexlink.v0"

var Version string

var Plugin = vexlink.Plugin{
	Name:    "png",
	Version: Version,
	Nodes: map[string]vexlink.Node{
		"pngwriter": &PNGWriter{},
		"pngreader": &PNGReader{},
	},
}
