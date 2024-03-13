package services

import (
	"ParkMe/dtos"
	"ParkMe/models"
	"errors"
	"fmt"
)

func CreateReservationForAVehicle(request dtos.ParkVehicleRequest) (id int64, err error) {
	// check parking slot available or not
	parkingSlotId := getAvailableParkingSlotByParkingLotId(int64(request.ParkingLotId))
	if parkingSlotId < 1 {
		return 0, errors.New("no available slot for this parking lot, try another")
	}

	reservation := models.NewReservation(request.ParkingLotId, parkingSlotId, request.VehicleId)

	id, err = models.AddReservation(reservation)

	if err != nil {
		return 0, err
	}

	// update slot table with is_occupied true
	err = UpdateStatus(parkingSlotId, dtos.UpdateParkingSlotStatus{
		ParkingLotId:  request.ParkingLotId,
		IsMaintenance: false,
		IsOccupied:    true,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateReservation(request *models.Reservation) error {
	fmt.Println(request)
	err := models.UpdateReservationById(request)
	if err != nil {
		return err
	}
	return nil
}
