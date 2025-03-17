package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	speed        int
	batteryDrain int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(spd, battDrain int) Car {
	out := Car{
		speed:        spd,
		battery:      100,
		distance:     0,
		batteryDrain: battDrain,
	}
	return out
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	out := Track{
		distance: distance,
	}
	return out
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	drives := car.battery / car.batteryDrain
	driveDistance := drives * car.speed

	if driveDistance >= track.distance {
		return true
	} else {
		return false
	}
}
