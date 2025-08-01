package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := range pic {
		row := make([]uint8, dx)
		for x := range row {
			row[x] = uint8((x ^ x + y ^ x) * 10) // puedes probar con x*y, (x+y)/2, etc.
		}
		pic[y] = row
	}
	// fmt.Println(pic)
	return pic
}

func main() {
	pic.Show(Pic)
}
