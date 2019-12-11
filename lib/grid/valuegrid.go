package grid

type ValueGrid map[Point]interface{}

func (g ValueGrid) maxX() int {
	i := 0
	maxX := 0
	for p := range g {
		if maxX < p.X || i == 0 {
			maxX = p.X
		}
		i++
	}
	return maxX
}

func (g ValueGrid) maxY() int {
	i := 0
	maxY := 0
	for p := range g {
		if maxY < p.Y || i == 0 {
			maxY = p.Y
		}
		i++
	}
	return maxY
}

func (g ValueGrid) minX() int {
	i := 0
	minX := 0
	for p := range g {
		if minX > p.X || i == 0 {
			minX = p.X
		}
		i++
	}
	return minX
}

func (g ValueGrid) minY() int {
	i := 0
	minY := 0
	for p := range g {
		if minY > p.Y || i == 0 {
			minY = p.Y
		}
		i++
	}
	return minY
}

func (g ValueGrid) row(y, minX, maxX int) []interface{} {
	result := make([]interface{}, 0, 1+maxX-minX)
	for x := minX; x <= maxX; x++ {
		result = append(result, g[Point{X: x, Y: y}])
	}
	return result
}

func (g ValueGrid) Rows() [][]interface{} {
	var result [][]interface{}
	maxY := g.maxY()
	minY := g.minY()
	maxX := g.maxX()
	minX := g.minX()
	for y := maxY; y >= minY; y-- {
		result = append(result, g.row(y, minX, maxX))
	}
	return result
}

func (g ValueGrid) Height() int {
	rows := g.Rows()
	return len(rows)
}

func (g ValueGrid) Width() int {
	rows := g.Rows()
	if len(rows) == 0 {
		return 0
	}
	return len(rows[0])
}
