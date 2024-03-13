package services

import (
	"ParkMe/dtos"
	"ParkMe/models"
	"errors"
)

func CreateParkingSlot(request dtos.CreateParkingSlotRequest) (id int64, err error) {
	// check the parking lots capacity overflowed or not
	if getCapacityOfAParkingLots(int64(request.ParkingLotId)) <= getTheNumberOfSlots(int64(request.ParkingLotId)) {
		return 0, errors.New("capacity overflowed")
	}

	parkingSlot := models.NewParkingSlot(request.ParkingLotId)
	id, err = models.AddParkingSlot(parkingSlot)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateStatus(parkingSlotId int, request dtos.UpdateParkingSlotStatus) error {
	// need to check if the slot is occupied or not
	parkingSlot := models.NewParkingSlot(request.ParkingLotId)
	parkingSlot.Id = int64(parkingSlotId)
	parkingSlot.IsMaintenance = request.IsMaintenance
	parkingSlot.IsOccupied = request.IsOccupied

	return models.UpdateParkingSlotById(parkingSlot)
}

func getTheNumberOfSlots(parkingLotId int64) int {
	count, _ := models.GetCountOfAllParkingSlotsByParkingId(parkingLotId)
	return count
}

func getAvailableParkingSlotByParkingLotId(parkingLotId int64) int {
	parkingSlotId, _ := models.GetAvailAbleParkingSlotByParkingLotId(int(parkingLotId))

	return parkingSlotId
}
