package problem1603

type ParkingSystem struct {
	BigReaming    int
	MediumReaming int
	SmallReaming  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		BigReaming:    big,
		MediumReaming: medium,
		SmallReaming:  small,
	}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if p.BigReaming > 0 {
			p.BigReaming--
			return true
		}
		return false
	case 2:
		if p.MediumReaming > 0 {
			p.MediumReaming--
			return true
		}
		return false
	case 3:
		if p.SmallReaming > 0 {
			p.SmallReaming--
			return true
		}
		return false
	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
