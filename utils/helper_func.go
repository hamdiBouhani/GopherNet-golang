package utils

import (
	"math"

	"github.com/hamdiBouhani/GopherNet-golang/storage/model"
)

func CalculateVolume(burrow *model.Burrow) float64 {
	// Volume of a cylinder: V = π * r^2 * h (where r is radius, h is height)
	radius := burrow.Wide / 2
	volume := math.Pi * math.Pow(radius, 2) * burrow.Depth
	return volume
}
