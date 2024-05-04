package main

import (
	"decorator/pkg"
	"fmt"
)

var (
	base = &pkg.BasePC{}
	home = &pkg.HomePC{
		CPU:          4,
		GraphicsCard: 2,
		Wrapper:      base,
	}
	server = &pkg.ServerPC{
		CPU:     8,
		Memory:  64,
		Wrapper: base,
	}
)

func main() {
	fmt.Println("Base PC price: " + fmt.Sprintf("%f", base.GetPrice()))
	fmt.Println("HomePC price: " + fmt.Sprintf("%f", home.GetPrice()))
	fmt.Println("ServerPC price: " + fmt.Sprintf("%f", server.GetPrice()))
}
