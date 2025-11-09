package main

import (
	"fmt"
	"os"

	core "github.com/HypoxiE/go-colors-picker/pkg/core"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("\033[31mError: enter the file name\033[0m")
		return
	}

	for _, name := range os.Args[1:] {
		px, err := core.GetPixels(name)
		if err != nil {
			fmt.Printf("\033[33mWarning: %v has been skipped (%v)\033[0m\n", name, err.Error())
			continue
		}
		px = core.Merge(px, 0.1, 10)
		config := core.GetConfig(px)

		core.SaveConfig(name, config)
		fmt.Printf("\033[1;34mINFO: image %v has been successfully processed\033[0m\n", name)
	}

}
