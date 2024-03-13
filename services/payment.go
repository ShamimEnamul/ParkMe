package services

import (
	"ParkMe/utils"
	"fmt"
)

func CalculatePayment(duration float64) float64 {
	fmt.Println("fa ", duration)
	return utils.PRICE_PER_UNIT * duration
}
