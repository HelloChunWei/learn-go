package martix

func reverseSingleRow(list []int) {
	i, j := 0, len(list)-1
	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
}

// 逆時鐘旋轉
func reverseMartix(matrix [][]int) [][]int {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		reverseSingleRow(row)
	}

	return matrix
}

// 順時針旋轉

func reverseMartix2(matrix [][]int) [][]int {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		reverseSingleRow(row)
	}
	return matrix
}
