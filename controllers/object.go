package controllers

import (
	"ParkMe/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

//
//
//
//====================
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//package controllers
//
//import (
//"ParkMe/dtos"
//"ParkMe/models"
//"ParkMe/services"
//"ParkMe/utils"
//_ "ParkMe/utils"
//"ParkMe/utils/handler"
//"encoding/json"
//"fmt"
//beego "github.com/beego/beego/v2/server/web"
//"log"
//"strconv"
//)
//
//
//// URLMapping ...
//func (c *ParkingLotController) URLMapping() {
//	c.Mapping("Post", c.Post)
//	c.Mapping("GetOne", c.GetOne)
//	c.Mapping("GetAll", c.GetAll)
//	c.Mapping("Put", c.Put)
//	c.Mapping("Delete", c.Delete)
//}
//
//// Post ...
//// @Title Create
//// @Description create ParkingLot
//// @Param	body		body 	models.ParkingLot	true		"body for ParkingLot content"
//// @Success 201 {object} models.ParkingLot
//// @Failure 403 body is empty
//// @router / [post]
//func (c *ParkingLotController) Post() {
//
//}
//
//// GetOne ...
//// @Title GetOne
//// @Description get ParkingLot by id
//// @Param	id		path 	string	true		"The key for staticblock"
//// @Success 200 {object} models.ParkingLot
//// @Failure 403 :id is empty
//// @router /:id [get]
//func (c *ParkingLotController) GetOne() {
//	defer handler.PanicRecovery(c.Controller)
//	parkingLotId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
//	val, err := models.GetParkingLotById(int64(parkingLotId))
//
//	if err != nil {
//		c.Data["json"] = utils.GetCode("failed")
//		_ = c.ServeJSON()
//		return
//	}
//	c.Data["json"] = val
//	c.ServeJSON()
//}
//
//// GetOne ...
//// @Title GetDetails
//// @Description get ParkingLot by id
//// @Param	id		path 	string	true		"The key for staticblock"
//// @Success 200 {object} models.ParkingLot
//// @Failure 403 :id is empty
//// @router /details/:id [get]
//func (c *ParkingLotController) GetDetails() {
//	defer handler.PanicRecovery(c.Controller)
//	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
//	startTime := c.Ctx.Input.Query("startTime")
//	endTime := c.Ctx.Input.Query("endTime")
//	val, err := models.GetParkingLotDetailsByRange(int64(id), startTime, endTime)
//	if err != nil {
//		c.Data["json"] = utils.GetCode("failed")
//		c.ServeJSON()
//		return
//	}
//	c.Data["json"] = val
//	c.ServeJSON()
//
//}
//
//// GetAll ...
//// @Title GetAll
//// @Description get ParkingLot
//// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
//// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
//// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
//// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
//// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
//// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
//// @Success 200 {object} models.ParkingLot
//// @Failure 403
//// @router / [get]
//func (c *ParkingLotController) GetAll() {
//	defer handler.PanicRecovery(c.Controller)
//	userId, _ := strconv.Atoi(c.Ctx.Input.Query("user_id"))
//	fmt.Println(userId)
//	val, err := models.GetAllParkingLotByUserId(int64(userId))
//	if err != nil {
//		c.Data["json"] = utils.GetCode("failed")
//		c.ServeJSON()
//		return
//	}
//	c.Data["json"] = val
//	c.ServeJSON()
//}
//
//// Put ...
//// @Title Put
//// @Description update the ParkingLot
//// @Param	id		path 	string	true		"The id you want to update"
//// @Param	body		body 	models.ParkingLot	true		"body for ParkingLot content"
//// @Success 200 {object} models.ParkingLot
//// @Failure 403 :id is not int
//// @router /:id [put]
//func (c *ParkingLotController) Put() {
//	defer handler.PanicRecovery(c.Controller)
//	parkingLotId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
//	var payload dtos.UpdateParkingLot
//	json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
//	err := handler.BodyParamCheck(&payload)
//
//	if err != nil {
//		fmt.Println(err)
//		c.Data["json"] = utils.GetCode("validation_error")
//		c.ServeJSON()
//		return
//	}
//
//	err = services.UpdateParkingLot(parkingLotId, payload)
//	if err != nil {
//		log.Println("Error occurred while updating data into creating Parking lot.")
//		c.Data["json"] = utils.GetCode("failed")
//		c.ServeJSON()
//		return
//	}
//
//	c.Data["json"] = utils.GetCode("success")
//	c.ServeJSON()
//}
//
//// Delete ...
//// @Title Delete
//// @Description delete the ParkingLot
//// @Param	id		path 	string	true		"The id you want to delete"
//// @Success 200 {string} delete success!
//// @Failure 403 id is empty
//// @router /:id [delete]
//func (c *ParkingLotController) Delete() {
//
//}
