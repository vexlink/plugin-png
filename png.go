package main

import "github.com/vexlink/vexlink"

var Version string

var Plugin = vexlink.Plugin{
	Name:    "png",
	Version: Version,
	Nodes: []vexlink.Node{
		&PNGWriter{},
		&PNGReader{},
	},
}
