package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	pix := make([][]uint8, dx)
	
	for x := range pix {
		pix[x] = make([]uint8, dy)
		
		for y := range pix[x] {
			// pix[x][y] = uint8((x+y)/2)
			// pix[x][y] = uint8(x*y)
			pix[x][y] = uint8(x^y)
		}
	}
	
	return pix
}

func main() {
	pic.Show(Pic)
}
