package main

import (
	"fmt"
	"os"

	core "github.com/HypoxiE/go-colors-picker/pkg/core"
	"github.com/lucasb-eyer/go-colorful"
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
		var config core.Configuration
		if len(px) == 1 {
			merged := core.Merge(px[0], 0.1, 10)
			config = core.GetConfig(merged)
		} else {
			all_merged := map[colorful.Color]int{}
			ch := make(chan map[colorful.Color]int)
			done := make(chan struct{})

			for _, frame := range px {
				go func(frame map[colorful.Color]int) {
					ch <- core.Merge(frame, 0.1, 100)
				}(frame)
			}

			go func() {
				for range len(px) {
					merged := <-ch
					for color, matches := range merged {
						all_merged[color] += matches
					}
				}
				close(done)
			}()
			<-done

			all_merged = core.Merge(all_merged, 0.1, 10)
			config = core.GetConfig(all_merged)
		}

		core.SaveConfig(name, config)
		fmt.Printf("\033[1;34mINFO: image %v has been successfully processed\033[0m\n", name)
	}

}
