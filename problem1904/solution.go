package problem1904

func numberOfRounds(loginTime, logoutTime string) int {
	loginH := convStr2Int(loginTime[:2])
	loginM := convStr2Int(loginTime[3:])
	logoutH := convStr2Int(logoutTime[:2])
	logoutM := convStr2Int(logoutTime[3:])
	var round int
	if (loginH == logoutH && loginM > logoutM) || loginH > logoutH {
		logoutH += 24
	}

	// same hour
	if loginH == logoutH {
		if loginM%15 == 0 {
			return logoutM/15 - loginM/15
		}
		if logoutM-loginM <= 15 {
			return 0
		}
		return ((logoutM - logoutM%15) / 15) - ((loginM - loginM%15) / 15) - 1
	}

	var nextHour = loginH
	if loginM != 0 {
		round += (60 - loginM) / 15
		nextHour++
	}
	if logoutM != 0 {
		round += logoutM / 15
	}
	round += 4 * (logoutH - nextHour)
	return round
}

func convStr2Int(s string) int {
	b := 1
	var num int
	for i := len(s) - 1; i >= 0; i-- {
		num += int(s[i]-'0') * b
		b *= 10
	}
	return num
}
