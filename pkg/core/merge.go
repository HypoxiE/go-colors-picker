package gocolorspicker

import (
	"slices"

	"github.com/lucasb-eyer/go-colorful"
)

type Cluster struct {
	Color   colorful.Color
	Matches int
}

func (cluster *Cluster) ColorMerge(color colorful.Color, color_matches int) {
	cluster.Color.R = (cluster.Color.R*float64(cluster.Matches) + color.R*float64(color_matches)) / (float64(cluster.Matches) + float64(color_matches))
	cluster.Color.G = (cluster.Color.G*float64(cluster.Matches) + color.G*float64(color_matches)) / (float64(cluster.Matches) + float64(color_matches))
	cluster.Color.B = (cluster.Color.B*float64(cluster.Matches) + color.B*float64(color_matches)) / (float64(cluster.Matches) + float64(color_matches))

	cluster.Matches += color_matches
}

func Merge(all_pixels map[colorful.Color]int, luft float64, need_colors int) map[colorful.Color]int {

	var count_pixels int
	for _, matches := range all_pixels {
		count_pixels += matches
	}

	var clusters []Cluster

	for color, match := range all_pixels {
		if len(clusters) == 0 {
			clusters = append(clusters, Cluster{
				Color:   color,
				Matches: match,
			})
			continue
		}

		is_merged := false
		for index, cluster := range clusters {
			if cluster.Color.DistanceCIEDE2000(color) <= luft {
				clusters[index].ColorMerge(color, match)
				is_merged = true
				break
			}
		}
		if !is_merged {
			clusters = append(clusters, Cluster{
				Color:   color,
				Matches: match,
			})
		}
	}

	slices.SortFunc(clusters, func(a, b Cluster) int {
		return b.Matches - a.Matches
	})

	result := make(map[colorful.Color]int)
	for _, cluster := range clusters {
		if len(result) >= need_colors {
			break
		}
		result[cluster.Color] = cluster.Matches
	}
	return result
}
