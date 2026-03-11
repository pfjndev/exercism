package jedlik

import "fmt"

/* type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
} */

// Drive updates the car's battery and distance.
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance returns the distance as a string.
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns the battery level as a string.
func (car Car) DisplayBattery() string {
	// %% is the correct way to print a literal %
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if the car can cover the track distance.
func (car Car) CanFinish(trackDistance int) bool {
	numDrives := car.battery / car.batteryDrain
	totalPossibleDistance := numDrives * car.speed
	return totalPossibleDistance >= trackDistance
}