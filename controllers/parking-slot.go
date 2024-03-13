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

// ParkingSlotController operations for ParkingSlot
type ParkingSlotController struct {
	beego.Controller
}

// URLMapping ...
func (c *ParkingSlotController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create ParkingSlot
// @Param	body		body 	models.ParkingSlot	true		"body for ParkingSlot content"
// @Success 201 {object} models.ParkingSlot
// @Failure 403 body is empty
// @router / [post]
func (c *ParkingSlotController) Post() {
	defer handler.PanicRecovery(c.Controller)
	var payload dtos.CreateParkingSlotRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
	if err != nil {
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}
	err = handler.BodyParamCheck(&payload)
	fmt.Println(payload)

	if err != nil {
		fmt.Println(err)
		c.Data["json"] = utils.GetCode("validation_error")
		c.ServeJSON()
		return
	}

	_, err = services.CreateParkingSlot(payload)
	if err != nil {
		log.Println("Error occurred while inserting data into creating Parking lot.")
		c.Data["json"] = utils.GetCode("instance_creation_error")
		c.ServeJSON()
		return
	}

	c.Data["json"] = utils.GetCode("success")
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get ParkingSlot by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ParkingSlot
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ParkingSlotController) GetOne() {
	defer handler.PanicRecovery(c.Controller)
	parkingSlotId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	val, err := models.GetParkingSlotById(int64(parkingSlotId))

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
// @Description get ParkingSlot
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ParkingSlot
// @Failure 403
// @router / [get]
func (c *ParkingSlotController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the ParkingSlot
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ParkingSlot	true		"body for ParkingSlot content"
// @Success 200 {object} models.ParkingSlot
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ParkingSlotController) Put() {
	defer handler.PanicRecovery(c.Controller)
	parkingSlotId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var payload dtos.UpdateParkingSlotStatus
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

	err = services.UpdateStatus(parkingSlotId, payload)
	if err != nil {
		log.Println("Error occurred while updating data into creating Parking lot.")
		c.Data["json"] = utils.GetCode("Failed during update the object")
		c.ServeJSON()
		return
	}

	c.Data["json"] = utils.GetCode("success")
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the ParkingSlot
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ParkingSlotController) Delete() {

}
