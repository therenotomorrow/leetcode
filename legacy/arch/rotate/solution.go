package rotate

func rotate(nums []int, k int) {
	k %= len(nums) // don't reverse the whole circle

	origin := make([]int, len(nums)-k)
	for i := 0; i < len(nums)-k; i++ {
		origin[i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if i < k {
			nums[i] = nums[len(nums)-k+i]
		} else {
			nums[i] = origin[i-k]
		}
	}
}

func rotate2(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[n-1-j][i]

			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]

			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]

			matrix[j][n-1-i] = matrix[i][j]

			matrix[i][j] = tmp
		}
	}
}
