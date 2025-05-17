package main

func main(){
	parkingLot := NewParkingLotInstance()
	parkingLot.AddLevel(NewLevel(1, 10))
	parkingLot.AddLevel(NewLevel(2,5))

	car1 := NewCar("PLK133")
	truck1 := NewTruck("KLK133")
	bike1 := NewMotorCycle("KKLK133")

	parkingLot.DisplayAvailability()

	println("=================================")

	parkingLot.ParkVehicle(car1)
	parkingLot.ParkVehicle(truck1)
	parkingLot.ParkVehicle(bike1)

	parkingLot.DisplayAvailability()

	println("=================================")

	parkingLot.UnParkVehicle(car1)

	parkingLot.DisplayAvailability()

	println("=================================")
}