package services

import (
	"ParkMe/dtos"
	"ParkMe/models"
	"fmt"
)

func CreateParkingLot(request dtos.CreateParkingLotRequest) (id int64, err error) {
	parkingLot := models.NewParkingLot(request.Name, request.Capacity, request.CreatedBy)
	id, err = models.AddParkingLot(parkingLot)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateParkingLot(parkingLotId int, request dtos.UpdateParkingLot) error {
	// need to check if the slot is occupied or not
	parkingLot := models.ParkingLot{
		Id: int64(parkingLotId),
	}
	if request.Active {
		parkingLot.Active = request.Active
	}

	if request.CreatedBy > 0 {
		parkingLot.CreatedBy = request.CreatedBy
	}

	if request.Capacity > 0 {
		parkingLot.Capacity = request.Capacity
	}
	if request.Name != "" {
		parkingLot.Name = request.Name
	}
	return models.UpdateParkingLotById(&parkingLot)
}

func getCapacityOfAParkingLots(id int64) int {
	capacity, _ := models.GetParkingLotCapacityById(id)
	fmt.Println("c:", capacity)
	return capacity
}
