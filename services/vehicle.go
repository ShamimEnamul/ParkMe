package services

import (
	"ParkMe/dtos"
	"ParkMe/models"
	"errors"
	"fmt"
	"time"
)

func Park(request dtos.ParkVehicleRequest) (id int64, err error) {
	// check parking slot available or not
	parkingSlotId := getAvailableParkingSlotByParkingLotId(int64(request.ParkingLotId))
	if parkingSlotId < 1 {
		return 0, errors.New("no available slot for this parking lot, try another")
	}

	// Need to check the vehicle is active with another slot or not, For now I am skipping this portion as I have a concern

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
func UnPark(vehicleId int, request dtos.UnParkVehicleRequest) (float64, error) {
	// update slot table with is_occupied true
	err := UpdateStatus(request.ParkingSlotId, dtos.UpdateParkingSlotStatus{
		ParkingLotId:  request.ParkingLotId,
		IsMaintenance: false,
		IsOccupied:    false,
	})
	if err != nil {
		return 0, err
	}

	//update reservation table and calculate
	err = UpdateReservation(&models.Reservation{
		ParkingSlotId: request.ParkingSlotId,
		ParkingLotId:  request.ParkingLotId,
		VehicleId:     vehicleId,
		OutTime:       time.Now(),
		Active:        false,
	})

	if err != nil {
		return 0, err
	}

	// calculate the price
	timeDuration, _ := models.GetDurationOfParking(request.ParkingLotId, request.ParkingSlotId, vehicleId)
	fmt.Println("d:", timeDuration)
	calculatedPrice := CalculatePayment(float64(timeDuration))

	return calculatedPrice, nil
}

func RegisterVehicle(request *dtos.RegisterVehicleRequest) (id int64, err error) {
	reg := models.NewRegister(request.UserId, request.RegistrationNo)
	id, err = models.AddVehicle(reg)

	if err != nil {
		return 0, err
	}

	return id, nil
}

//if err := json.Unmarshal(p.Ctx.Input.RequestBody, &request); err != nil {
//p.Data["json"] = ErrorResponse{Error: "Invalid request body"}
//p.ServeJSON()
//return
//}
