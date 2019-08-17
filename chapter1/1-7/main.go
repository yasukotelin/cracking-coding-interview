package main

func main() {
}

func createMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		layer := make([]int, n)
		for j := 0; j < n; j++ {
			layer[j] = j + (i * 10)
		}
		matrix[i] = layer
	}
	return matrix
}

// rotate90 はN*Nの行列に描かれた、1つのピクセルが4バイト四方の画像を90度回転させる
func rotate90(matrix [][]int) {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return
	}

	n := len(matrix)

	// layer -> 行列の1行ごと
	for layer := 0; layer < n/2; layer++ {

		// layer内の最初
		first := layer
		// layer内の末尾
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]

			// 左下 → 左上
			matrix[first][i] = matrix[last-offset][first]
			// 右下 → 左下
			matrix[last-offset][first] = matrix[last][last-offset]
			// 右上 →右下
			matrix[last][last-offset] = matrix[i][last]
			// 左上 → 右上
			matrix[i][last] = top
		}
	}
}
