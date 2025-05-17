package main

type ParkingSpot struct{
	vehicleType VehicleType
	spotNumber int
	parkedVehicle Vehicle
}

func NewParkingSpot(spotNumber int, vehicleType VehicleType) *ParkingSpot{
	return &ParkingSpot{
		spotNumber: spotNumber,
		vehicleType: vehicleType,
	}
}

func(ps *ParkingSpot)IsAvailable() bool{
	return ps.parkedVehicle == nil
}

func(ps *ParkingSpot)ParkVehicle(vehicle Vehicle){
	if ps.IsAvailable() && ps.vehicleType == vehicle.GetVehicleType() {
		ps.parkedVehicle = vehicle
	}
}

func(ps *ParkingSpot)UnParkVehicle(){
	ps.parkedVehicle = nil
}

func(ps *ParkingSpot)GetSpotNumber() int{
	return ps.spotNumber
}

func(ps *ParkingSpot)GetVehicleType() VehicleType{
	return ps.vehicleType
}

func(ps *ParkingSpot)GetParkedVehicle() Vehicle{
	return ps.parkedVehicle
}
