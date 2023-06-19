package problem2739

func distanceTraveled(mainTank int, additionalTank int) int {
	var dis int
	for mainTank > 0 {
		if mainTank >= 5 {
			dis += 5 * 10
			mainTank -= 5
			if additionalTank > 0 {
				additionalTank--
				mainTank++
			}
			continue
		}
		dis += mainTank * 10
		mainTank -= mainTank
	}
	return dis
}
