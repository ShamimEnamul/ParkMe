package dtos

// ParkVehicleRequest represents the request body for parking a vehicle

type ParkVehicleRequest struct {
	ParkingLotId int `valid:"Required;" json:"parkingLotId"`
	//ParkingSlotId int `valid:"Required;" json:"parkingSlotId"`
	VehicleId int `valid:"Required;" json:"vehicleId"`
}

type UnParkVehicleRequest struct {
	ParkingLotId  int `valid:"Required;" json:"parkingLotId"`
	ParkingSlotId int `valid:"Required;" json:"parkingSlotId"`
}

type RegisterVehicleRequest struct {
	UserId         int    `json:"userId"`
	RegistrationNo string `json:"registrationNo"`
}
