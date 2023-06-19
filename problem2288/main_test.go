package main

import "testing"

func TestDiscountPrices(t *testing.T) {
	tests := []struct {
		sentence string
		discount int
		expected string
	}{
		{
			sentence: "there are $1 $2 and 5$ candies in the shop",
			discount: 50,
			expected: "there are $0.50 $1.00 and 5$ candies in the shop",
		},
		{
			sentence: "1 2 $3 4 $5 $6 7 8$ $9 $10$",
			discount: 100,
			expected: "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$",
		},
		{
			sentence: "$76111 ab $6 $",
			discount: 48,
			expected: "$39577.72 ab $3.12 $",
		},
	}

	for _, test := range tests {
		if res := discountPrices(test.sentence, test.discount); res != test.expected {
			t.Errorf("Test failed. The input sentence \"%s\", discount %d \nis expected to \"%s\",\n but response \"%s\"", test.sentence, test.discount, test.expected, res)
		}
	}

}
