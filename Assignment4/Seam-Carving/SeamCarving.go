package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"
	"image/color"
	"image/jpeg"
)

func main() {
	// Decode the JPEG data. If reading from file, create a reader with
	//
	reader, err := os.Open("testdata/1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	mColor, row, col := imageToArr(m)
	for i := 0; i < col/2; i++ {
		fmt.Println(i)
		disMap := calcDis((energyMapGen(mColor)))
		mColor = removePix(mColor, disMap)
	}

	newI := arrToImage(mColor, image.Rect(0, 0, col/2, row-1))

	f, err := os.OpenFile("testout/1.jpeg", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	jpeg.Encode(f, newI, nil)
}

func imageToArr(m image.Image) ([][]color.Color, int, int) {
	row := 0
	colR := 0
	bounds := m.Bounds()
	re := [][]color.Color{}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		row++
		re = append(re, []color.Color{})
		//fmt.Print("y:", y, "    row:", y-bounds.Min.Y, "  len:", len(re), "\n")
		col := 0
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			col++
			re[y] = append(re[y-bounds.Min.Y], m.At(x, y))
		}
		colR = col

	}
	return re, row, colR
}

func arrToImage(m [][]color.Color, bounds image.Rectangle) image.Image {

	newI := image.NewRGBA(bounds)
	for y, row := range m {
		for x, c := range row {
			newI.Set(x, y, c)
		}
	}
	return newI
}

func energyMapGen(m [][]color.Color) [][]int {
	energyMap := [][]int{}
	m = append(m, m[len(m)-1])
	if len(m) < 1 || len(m[0]) < 1 {
		return energyMap
	}

	for r := 0; r < len(m); r++ {
		energyMap = append(energyMap, make([]int, len(m[0])))
		for c := 0; c < len(m[0]); c++ {
			energyMap[r][c] = calcEnergy(m, r, c)
		}
	}
	return energyMap
}

func calcEnergy(m [][]color.Color, r int, c int) int {
	count := 0
	e := 0
	r0, g0, b0, a0 := m[r][c].RGBA()
	rn := r - 1
	cn := c
	if rn >= 0 && rn < len(m) && cn >= 0 && cn < len(m[0]) {
		r1, g1, b1, a1 := m[rn][cn].RGBA()
		e += int(listAbs([]uint32{r0 - r1, g0 - g1, b0 - b1, a0 - a1}))
		count++
	}
	rn = r + 1
	cn = c
	if rn >= 0 && rn < len(m) && cn >= 0 && cn < len(m[0]) {
		r1, g1, b1, a1 := m[rn][cn].RGBA()
		e += int(listAbs([]uint32{r0 - r1, g0 - g1, b0 - b1, a0 - a1}))
		count++
	}
	rn = r
	cn = c - 1
	if rn >= 0 && rn < len(m) && cn >= 0 && cn < len(m[0]) {
		r1, g1, b1, a1 := m[rn][cn].RGBA()
		e += int(listAbs([]uint32{r0 - r1, g0 - g1, b0 - b1, a0 - a1}))
		count++
	}
	rn = r
	cn = c + 1
	if rn >= 0 && rn < len(m) && cn >= 0 && cn < len(m[0]) {
		r1, g1, b1, a1 := m[rn][cn].RGBA()
		e += int(listAbs([]uint32{r0 - r1, g0 - g1, b0 - b1, a0 - a1}))
		count++
	}
	return e / count
}

func listAbs(l []uint32) uint32 {
	var r uint32 = 0
	for _, v := range l {
		if v < 0 {
			r = r - v
		} else {
			r = r + v
		}
	}
	return r
}

func calcDis(eMap [][]int) [][]int {
	disMap := make([][]int, len(eMap)+1)
	if len(eMap) < 1 || len(eMap[0]) < 1 {
		return disMap
	}
	for i, _ := range disMap {
		disMap[i] = make([]int, len(eMap[0])+2)
		disMap[i][0] = math.MaxInt
		disMap[i][len(eMap[0])+2-1] = math.MaxInt
	}

	for r := 1; r <= len(eMap); r++ {
		for c := 1; c <= len(eMap[0]); c++ {
			d := disMap[r-1][c-1]
			if d > disMap[r-1][c] {
				d = disMap[r-1][c]
			}
			if d > disMap[r-1][c+1] {
				d = disMap[r-1][c+1]
			}
			//fmt.Println(len(eMap[0]))
			disMap[r][c] = d + eMap[r-1][c-1]
		}
	}
	return disMap
}

func removePix(m [][]color.Color, disMap [][]int) [][]color.Color {
	c := 0
	for i, v := range disMap[len(disMap)-1] {
		if v < disMap[len(disMap)-1][c] {
			c = i
		}
	}
	c--

	mN := make([][]color.Color, len(m))
	for i := 0; i < len(mN); i++ {
		mN[i] = make([]color.Color, len(m[0]))
		copy(mN[i], m[i])
	}

	for cn := c; cn < len(m[0])-1; cn++ {
		mN[len(m)-1][cn] = m[len(m)-1][cn+1]
	}

	for r := len(m) - 2; r >= 0; r-- {
		d1 := disMap[r+1][c-1+1]
		d2 := disMap[r+1][c+1]
		d3 := disMap[r+1][c+1+1]
		cn := 0
		switch {
		case d1 < d2 && d1 < d3:
			cn = c - 1
		case d2 < d1 && d2 < d3:
			cn = c
		default:
			cn = c + 1
		}
		c = cn
		//fmt.Println("Cut in col:", cn)
		for ; cn < len(m[0])-1; cn++ {
			mN[r][cn] = m[r][cn+1]
		}
	}

	return mN
}
