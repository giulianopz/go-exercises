package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	outer := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		inner := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			inner[x] = uint8(x * x * y * y << 2)
		}
		outer[y] = inner
	}
	return outer
}

func main() {
	pic.Show(Pic)
}
