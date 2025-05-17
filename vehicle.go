package main

type VehicleType int

const (
	CAR VehicleType = iota
	MOTORCYCLE
	TRUCK
)

type Vehicle interface{
	GetLicensePlate() string
	GetVehicleType() VehicleType
}

type BaseVehicle struct{
	licensePlate string
	vehicleType VehicleType
}

func(b *BaseVehicle)GetLicensePlate() string{
	return b.licensePlate
}

func(b *BaseVehicle)GetVehicleType() VehicleType{
	return b.vehicleType
}