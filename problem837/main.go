package problem837

import (
	"fmt"
	"strconv"
)

// Alice plays the following game, loosely based on the card game "21".

// Alice starts with 0 points and draws numbers while she has less than k points. During each draw, she gains an integer number of points randomly from the range [1, maxPts], where maxPts is an integer. Each draw is independent and the outcomes have equal probabilities.

// Alice stops drawing numbers when she gets k or more points.

// Return the probability that Alice has n or fewer points.

// Answers within 10-5 of the actual answer are considered accepted.

func new21Game(n int, k int, maxPts int) float64 {
	var dp = make([]float64, n+1)
	dp[0] = 1

	var s float64
	if k > 0 {
		s = 1
	}

	for i := 1; i <= n; i++ {
		dp[i] = s / float64(maxPts)
		if i < k {
			s += dp[i]
		}
		if i-maxPts >= 0 && i-maxPts < k {
			s -= dp[i-maxPts]

		}
	}

	var ans float64
	for i := k; i <= n; i++ {
		ans += dp[i]
	}

	fAns, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", ans), 64)

	return fAns
}
