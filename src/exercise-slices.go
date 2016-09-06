package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var ret [][]uint8
	for i:=0; i<dy; i++ {
		ret = append(ret, make([]uint8, dx))
	}
	for i:=0; i<dy; i++ {
		for j:=0; j<dx; j++ {
			ret[i][j] = uint8(i^j)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
