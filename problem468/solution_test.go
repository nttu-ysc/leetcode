package problem468

import "testing"

func TestValidIPAddress(t *testing.T) {
	const (
		IPv4    = "IPv4"
		IPv6    = "IPv6"
		NEITHER = "Neither"
	)
	testCases := []struct {
		queryIP string
		want    string
	}{
		{
			queryIP: "172.16.254.1",
			want:    IPv4,
		},
		{
			queryIP: "2001:0db8:85a3:0:0:8A2E:0370:7334",
			want:    IPv6,
		},
		{
			queryIP: "256.256.256.256",
			want:    NEITHER,
		},
		{
			queryIP: "12..33.4",
			want:    NEITHER,
		},
		{
			queryIP: "12.12.12",
			want:    NEITHER,
		},
		{
			queryIP: "2001:0db8:85a3:00000:0:8A2E:0370:7334",
			want:    NEITHER,
		},
		{
			queryIP: "192.0.0.1",
			want:    IPv4,
		},
	}

	for _, tC := range testCases {
		if got := validIPAddress(tC.queryIP); got != tC.want {
			t.Errorf("Test failed. The input queryIP = \"%s\" is expected to %s but return %s", tC.queryIP, tC.want, got)
		}
	}
}
