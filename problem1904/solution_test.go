package problem1904

import "testing"

func TestNumberOfRounds(t *testing.T) {
	testCases := []struct {
		loginTime  string
		logoutTime string
		want       int
	}{
		{
			loginTime:  "20:00",
			logoutTime: "20:16",
			want:       1,
		},
		{
			loginTime:  "09:00",
			logoutTime: "09:58",
			want:       3,
		},
		{
			loginTime:  "12:08",
			logoutTime: "12:57",
			want:       2,
		},
		{
			loginTime:  "06:13",
			logoutTime: "06:19",
			want:       0,
		},
		{
			loginTime:  "00:47",
			logoutTime: "00:57",
			want:       0,
		},
		{
			loginTime:  "00:01",
			logoutTime: "00:00",
			want:       95,
		},
		{
			loginTime:  "09:31",
			logoutTime: "10:14",
			want:       1,
		},
		{
			loginTime:  "21:30",
			logoutTime: "03:00",
			want:       22,
		},
	}

	for _, tC := range testCases {
		if got := numberOfRounds(tC.loginTime, tC.logoutTime); got != tC.want {
			t.Errorf("Test failed. The input loginTime = \"%s\", logoutTime = \"%s\" is expected to %d but return %d", tC.loginTime, tC.logoutTime, tC.want, got)
		}
	}
}
