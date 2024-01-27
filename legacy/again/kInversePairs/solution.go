package kInversePairs

const MOD = 1000000007

func kInversePairs(n int, k int) int {
	dp := make([]int, k+1)

	for i := 1; i <= n; i++ {
		temp := make([]int, k+1)
		temp[0] = 1

		for j := 1; j <= k; j++ {
			val := dp[j] + MOD

			if (j - i) >= 0 {
				val -= dp[j-i]
			}

			val %= MOD

			temp[j] = (temp[j-1] + val) % MOD
		}
		dp = temp
	}

	res := dp[k] + MOD
	if k > 0 {
		res -= dp[k-1]
	}

	return res % MOD
}
