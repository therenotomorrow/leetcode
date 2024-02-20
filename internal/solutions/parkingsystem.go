package solutions

type ParkingSystem struct {
	cars [3]int
}

func ParkingSystemConstructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{cars: [3]int{big, medium, small}}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	carType--
	ps.cars[carType]--

	return ps.cars[carType] >= 0
}
