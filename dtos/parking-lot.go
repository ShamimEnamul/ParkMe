package dtos

import "ParkMe/models"

//type LoginWithPasswordDto struct {
//	Phone       string `valid:"Required;MinSize(11);MaxSize(11)" json:"phone"`
//	Password    string `valid:"Required;MinSize(16);MaxSize(128)" json:"password"`
//	SocialId    string `valid:"MaxSize(128)" json:"social_id"`
//	SocialName  string `valid:"MaxSize(128)" json:"social_name"`
//	SocialEmail string `valid:"MaxSize(128)" json:"social_email"`
//	LoginType   string `valid:"Required;MinSize(2);MaxSize(10)" json:"login_type"`
//}
//

type CreateParkingLotRequest struct {
	Name      string `valid:"Required; MaxSize(255)" json:"name"`
	Capacity  int    `valid:"Required; " json:"capacity"`
	CreatedBy int    `valid:"Required;" json:"createdBy"`
}

type CreateParkingLotResponse struct {
	Name     string `valid:"Required; MaxSize(255)" json:"name"`
	Capacity int    `valid:"Required;" json:"capacity"`
}

type GetAllParkingLotResponse struct {
	ParkingLots []models.ParkingLot
}

type UpdateParkingLot struct {
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	CreatedBy int    `json:"createdBy"`
	Active    bool   `json:"active"`
}
