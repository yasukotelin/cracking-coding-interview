package main

func main() {
}

// paddingZero はM*Nの行列について、要素が0であれば、その行と列のすべてを0にする
func paddingZero(matrix [][]int) [][]int {
	zeroPoints := searchZeroPoints(matrix)

	for _, point := range zeroPoints {

		// 左方向
		x, y := point.x, point.y
		for x >= 0 {
			matrix[y][x] = 0
			x--
		}

		// 右方向
		x, y = point.x, point.y
		for x < len(matrix[0]) {
			matrix[y][x] = 0
			x++
		}

		// 上方向
		x, y = point.x, point.y
		for y >= 0 {
			matrix[y][x] = 0
			y--
		}

		// 下方向
		x, y = point.x, point.y
		for y < len(matrix) {
			matrix[y][x] = 0
			y++
		}
	}
	return matrix
}

type point struct {
	x int
	y int
}

func searchZeroPoints(matrix [][]int) []point {
	var points []point
	for y, row := range matrix {
		for x, elm := range row {
			if elm == 0 {
				points = append(points, point{x: x, y: y})
			}
		}
	}
	return points
}
