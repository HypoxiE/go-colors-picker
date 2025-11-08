package gocolorspicker

import (
	"sort"

	"github.com/lucasb-eyer/go-colorful"
)

type Color struct {
	Color   colorful.Color
	Matches int
}

type HexColor struct {
	Color   string `json:"color"`
	Matches int    `json:"match"`
}

type HyprlandConfig struct {
	ActiveBorderColor1   string `json:"active_border_color_1"`
	ActiveBorderColor2   string `json:"active_border_color_2"`
	InactiveBorderColor1 string `json:"inactive_border_color_1"`
	InactiveBorderColor2 string `json:"inactive_border_color_2"`
}

type EwwConfig struct {
	MainColor      string `json:"main_color"`
	SecondaryColor string `json:"secondary_color"`
	TextColor      string `json:"text_color"`
	IconsColor     string `json:"icons_color"`
}

type SwayncConfig struct {
	MainColor      string `json:"main_color"`
	SecondaryColor string `json:"secondary_color"`
	TextColor      string `json:"text_color"`
	IconsColor     string `json:"icons_color"`
}

type Configuration struct {
	Colors   []HexColor     `json:"all_colors"`
	Hyprland HyprlandConfig `json:"hyprland"`
	Eww      EwwConfig      `json:"eww"`
	Swaync   SwayncConfig   `json:"swaync"`
}

func SortByDistance(all_pixels map[colorful.Color]int) []colorful.Color {

	var primary_color colorful.Color
	var colors []colorful.Color
	for i, matches := range all_pixels {
		if all_pixels[primary_color] < matches {
			primary_color = i
		}
		colors = append(colors, i)
	}

	sort.Slice(colors, func(i, j int) bool {
		di := colors[i].DistanceCIEDE2000(primary_color)
		dj := colors[j].DistanceCIEDE2000(primary_color)
		return di < dj
	})

	return colors
}

func MaxColor(colors []colorful.Color, matches_map map[colorful.Color]int) colorful.Color {
	max_col := colors[0]
	for _, color := range colors[1:] {
		if matches_map[max_col] < matches_map[color] {
			max_col = color
		}
	}
	return max_col
}

func GetConfig(all_pixels map[colorful.Color]int) Configuration {
	var result Configuration

	sorted_colors := SortByDistance(all_pixels)
	for _, color := range sorted_colors {
		hex_color := HexColor{Color: color.Hex(), Matches: all_pixels[color]}
		result.Colors = append(result.Colors, hex_color)
	}

	if len(all_pixels) < 10 {
		//hyprland
		result.Hyprland.InactiveBorderColor1 = sorted_colors[0].Hex()
		result.Hyprland.InactiveBorderColor2 = sorted_colors[0].Hex()
		result.Hyprland.ActiveBorderColor1 = sorted_colors[len(sorted_colors)-1].Hex()
		result.Hyprland.ActiveBorderColor2 = sorted_colors[len(sorted_colors)-1].Hex()

		//eww
		result.Eww.MainColor = sorted_colors[0].Hex()
		result.Eww.TextColor = sorted_colors[len(sorted_colors)-1].Hex()
		result.Eww.IconsColor = result.Eww.TextColor
		result.Eww.SecondaryColor = sorted_colors[len(sorted_colors)/2].Hex()

		//swaync
		result.Swaync.MainColor = result.Eww.MainColor
		result.Swaync.TextColor = result.Eww.TextColor
		result.Swaync.IconsColor = result.Eww.IconsColor
		result.Swaync.SecondaryColor = result.Eww.SecondaryColor

		return result
	}

	//hyprland
	result.Hyprland.InactiveBorderColor1 = sorted_colors[0].Hex()
	result.Hyprland.InactiveBorderColor2 = sorted_colors[1].Hex()
	result.Hyprland.ActiveBorderColor1 = MaxColor(sorted_colors[3:7], all_pixels).Hex()
	result.Hyprland.ActiveBorderColor2 = MaxColor(sorted_colors[7:10], all_pixels).Hex()

	//eww
	result.Eww.MainColor = sorted_colors[0].Hex()
	result.Eww.TextColor = MaxColor(sorted_colors[7:10], all_pixels).Hex()
	result.Eww.IconsColor = result.Eww.TextColor
	result.Eww.SecondaryColor = MaxColor(sorted_colors[3:7], all_pixels).Hex()

	//swaync
	result.Swaync.MainColor = result.Eww.MainColor
	result.Swaync.TextColor = result.Eww.TextColor
	result.Swaync.IconsColor = result.Eww.IconsColor
	result.Swaync.SecondaryColor = result.Eww.SecondaryColor

	return result
}
