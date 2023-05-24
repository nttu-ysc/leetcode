package main

import (
	"fmt"
	"math"
)

func discountPrices(sentence string, discount int) string {
	var ans string

	tmp := []rune{}
	for _, r := range sentence + " " {
		if r != ' ' {
			tmp = append(tmp, r)
			continue
		}

		if len(tmp) > 0 {
			if tmp[0] != '$' || len(tmp) == 1 {
				ans += string(tmp) + " "
				tmp = []rune{}
				continue
			}

			var number float64
			for i := 1; i < len(tmp); i++ {
				if tmp[i]-'0' > 9 || tmp[i]-'0' < 0 {
					ans += string(tmp) + " "
					tmp = []rune{}
					goto h
				}

				number += float64(tmp[i]-'0') * (math.Pow10(len(tmp) - i - 1))
			}

			ans += fmt.Sprintf("$%.2f ", number*(1-float64(discount)/100))
			tmp = []rune{}
		}
	h:
	}

	return ans[:len(ans)-1]
}
