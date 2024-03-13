package controllers

import (
	"ParkMe/dtos"
	"ParkMe/models"
	"ParkMe/services"
	"ParkMe/utils"
	"ParkMe/utils/handler"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"strconv"
)

// VehicleController operations for Vehicle
type VehicleController struct {
	beego.Controller
}

// URLMapping ...
func (c *VehicleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description park Vehicle
// @Param	body		body 	models.Vehicle	true		"body for Vehicle content"
// @Success 201 {object} models.Vehicle
// @Failure 403 body is empty
// @router / [post]
func (c *VehicleController) Post() {
	defer handler.PanicRecovery(c.Controller)
	var payload dtos.RegisterVehicleRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
	if err != nil {
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	err = handler.BodyParamCheck(&payload)

	if err != nil {
		fmt.Println(err)
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	_, err = services.RegisterVehicle(&payload)
	if err != nil {
		log.Println("Error occurred while inserting data into reservation table.")
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()

	}

	c.Data["json"] = utils.GetCode("success")
	c.ServeJSON()
}

// Post ...
// @Title Park
// @Description park Vehicle
// @Param	body		body 	models.Vehicle	true		"body for Vehicle content"
// @Success 201 {object} models.Vehicle
// @Failure 403 body is empty
// @router /park [post]
func (c *VehicleController) Park() {
	defer handler.PanicRecovery(c.Controller)
	var payload dtos.ParkVehicleRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
	if err != nil {
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	err = handler.BodyParamCheck(&payload)

	if err != nil {
		fmt.Println(err)
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	_, err = services.Park(payload)
	if err != nil {
		log.Println("Error occurred while inserting data into reservation table.")
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()

	}

	c.Data["json"] = utils.GetCode("success")
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Vehicle by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Vehicle
// @Failure 403 :id is empty
// @router /:id [get]
func (c *VehicleController) GetOne() {
	defer handler.PanicRecovery(c.Controller)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	startTime := c.Ctx.Input.Query("startTime")
	endTime := c.Ctx.Input.Query("endTime")
	fmt.Println("af", startTime, endTime)
	val, err := models.GetParkingLotDetailsByRange(int64(id), startTime, endTime)

	if err != nil {
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()
		return
	}
	c.Data["json"] = val
	c.ServeJSON()
}

// GetOne ...
// @Title Get Vehicle Details by parkingLot and Slot
// @Description get Vehicle by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Vehicle
// @Failure 403 :id is empty
// @router /details [get]
func (c *VehicleController) GetVehicleDetails() {
	defer handler.PanicRecovery(c.Controller)
	parkingLotId, _ := strconv.Atoi(c.Ctx.Input.Query("parkingLotId"))
	parkingSlotId, _ := strconv.Atoi(c.Ctx.Input.Query("parkingSlotId"))
	fmt.Println("..", parkingLotId, parkingSlotId)
	val, err := models.GetVehicleByParkingSlotAndParkindLotId(int64(parkingLotId), int64(parkingSlotId))

	if err != nil {
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()
		return
	}
	c.Data["json"] = val
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Vehicle
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Vehicle
// @Failure 403
// @router / [get]
func (c *VehicleController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Vehicle
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Vehicle	true		"body for Vehicle content"
// @Success 200 {object} models.Vehicle
// @Failure 403 :id is not int
// @router /:id [put]
func (c *VehicleController) Put() {
	defer handler.PanicRecovery(c.Controller)
	vehicleId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var payload dtos.UpdateParkingSlotStatus
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
	if err != nil {
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	err = handler.BodyParamCheck(&payload)
	fmt.Println(vehicleId)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}

	c.Data["json"] = utils.GetCode("success")
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Vehicle
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *VehicleController) Delete() {

}

// Put ...
// @Title UnPark
// @Description unpark the Vehicle
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Vehicle	true		"body for Vehicle content"
// @Success 200 {object} models.Vehicle
// @Failure 403 :id is not int
// @router /unpark/:id [put]
func (c *VehicleController) UnPark() {
	defer handler.PanicRecovery(c.Controller)

	vehicleId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var payload dtos.UnParkVehicleRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload)

	if err != nil {
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	err = handler.BodyParamCheck(&payload)

	if err != nil {
		fmt.Println(err)
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	calculatedPrice, err := services.UnPark(vehicleId, payload)

	if err != nil {
		fmt.Println(err)
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1000, "message": "success", "amount": calculatedPrice}
	c.ServeJSON()
}
