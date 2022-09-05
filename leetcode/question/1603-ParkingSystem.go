package question

//1603. 设计停车系统

//PakingSystem 停车系统
type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

//Constructor1 结构体
func Constructor1(big int, medium int, small int) ParkingSystem {
	ps := new(ParkingSystem)
	ps.Big = big
	ps.Medium = medium
	ps.Small = small
	return *ps
}

//AddCar 停车
func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if ps.Big > 0 {
			ps.Big--
			return true
		}
		return false
	case 2:
		if ps.Medium > 0 {
			ps.Medium--
			return true
		}
		return false
	case 3:
		if ps.Small > 0 {
			ps.Small--
			return true
		}
		return false
	default:
		return false
	}
}
