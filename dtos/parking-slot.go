package dtos

type CreateParkingSlotRequest struct {
	ParkingLotId int `valid:"Required; " json:"parkingLotId"`
}

type UpdateParkingSlotStatus struct {
	ParkingLotId  int  `valid:"Required;" json:"parkingLotId"`
	IsMaintenance bool `json:"isMaintenance"`
	IsOccupied    bool `json:"isOccupied"`
}
