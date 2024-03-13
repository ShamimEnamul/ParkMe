package controllers

import (
	"ParkMe/dtos"
	"ParkMe/models"
	"ParkMe/services"
	"ParkMe/utils"
	_ "ParkMe/utils"
	"ParkMe/utils/handler"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"strconv"
)

// ParkingLotController operations for ParkingLot
type ParkingLotController struct {
	beego.Controller
}

// URLMapping ...
func (c *ParkingLotController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ParkingLot
// @Param	body		body 	models.ParkingLot	true		"body for ParkingLot content"
// @Success 201 {int} models.ParkingLot
// @Failure 403 body is empty
// @router / [post]
func (c *ParkingLotController) Post() {
	defer handler.PanicRecovery(c.Controller)
	var payload dtos.CreateParkingLotRequest
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

	id, err := services.CreateParkingLot(payload)
	if err != nil {
		log.Println("Error occurred while inserting data into creating Parking lot.")
		c.Data["json"] = utils.GetCode("instance_creation_error")
		c.ServeJSON()
	}
	fmt.Println(id)
	fmt.Println("payload")
	c.Data["json"] = utils.GetCode("success")
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ParkingLot by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ParkingLot
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ParkingLotController) GetOne() {
	defer handler.PanicRecovery(c.Controller)
	parkingLotId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	val, err := models.GetParkingLotById(int64(parkingLotId))

	if err != nil {
		c.Data["json"] = utils.GetCode("failed")
		_ = c.ServeJSON()
		return
	}
	c.Data["json"] = val
	c.ServeJSON()
}

// GetOne ...
// @Title GetDetails
// @Description get ParkingLot by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ParkingLot
// @Failure 403 :id is empty
// @router /details/:id [get]
func (c *ParkingLotController) GetDetails() {
	defer handler.PanicRecovery(c.Controller)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	startTime := c.Ctx.Input.Query("startTime")
	endTime := c.Ctx.Input.Query("endTime")
	val, err := models.GetParkingLotDetailsByRange(int64(id), startTime, endTime)
	if err != nil {
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()
		return
	}
	c.Data["json"] = val
	c.ServeJSON()

}

// GetAll ...
// @Title Get All
// @Description get ParkingLot
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ParkingLot
// @Failure 403
// @router / [get]
func (c *ParkingLotController) GetAll() {

	defer handler.PanicRecovery(c.Controller)
	userId, _ := strconv.Atoi(c.Ctx.Input.Query("user_id"))
	fmt.Println(userId)
	val, err := models.GetAllParkingLotByUserId(int64(userId))
	if err != nil {
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()
		return
	}
	c.Data["json"] = val
	c.ServeJSON()

	//var fields []string
	//var sortby []string
	//var order []string
	//var query = make(map[string]string)
	//var limit int64 = 10
	//var offset int64
	//
	//// fields: col1,col2,entity.col3
	//if v := c.GetString("fields"); v != "" {
	//	fields = strings.Split(v, ",")
	//}
	//// limit: 10 (default is 10)
	//if v, err := c.GetInt64("limit"); err == nil {
	//	limit = v
	//}
	//// offset: 0 (default is 0)
	//if v, err := c.GetInt64("offset"); err == nil {
	//	offset = v
	//}
	//// sortby: col1,col2
	//if v := c.GetString("sortby"); v != "" {
	//	sortby = strings.Split(v, ",")
	//}
	//// order: desc,asc
	//if v := c.GetString("order"); v != "" {
	//	order = strings.Split(v, ",")
	//}
	//// query: k:v,k:v
	//if v := c.GetString("query"); v != "" {
	//	for _, cond := range strings.Split(v, ",") {
	//		kv := strings.SplitN(cond, ":", 2)
	//		if len(kv) != 2 {
	//			c.Data["json"] = errors.New("Error: invalid query key/value pair")
	//			c.ServeJSON()
	//			return
	//		}
	//		k, v := kv[0], kv[1]
	//		query[k] = v
	//	}
	//}
	//
	//l, err := models.GetAllParkingLot(query, fields, sortby, order, offset, limit)
	//if err != nil {
	//	c.Data["json"] = err.Error()
	//} else {
	//	c.Data["json"] = l
	//}
	//c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ParkingLot
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ParkingLot	true		"body for ParkingLot content"
// @Success 200 {object} models.ParkingLot
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ParkingLotController) Put() {
	defer handler.PanicRecovery(c.Controller)
	parkingLotId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var payload dtos.UpdateParkingLot
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

	err = services.UpdateParkingLot(parkingLotId, payload)
	if err != nil {
		log.Println("Error occurred while updating data into creating Parking lot.")
		c.Data["json"] = utils.GetCode("failed")
		c.ServeJSON()
		return
	}

	c.Data["json"] = utils.GetCode("success")
	c.ServeJSON()
	//
	//idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	//v := models.ParkingLot{Id: id}
	//json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	//if err := models.UpdateParkingLotById(&v); err == nil {
	//	c.Data["json"] = "OK"
	//} else {
	//	c.Data["json"] = err.Error()
	//}
	//c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the ParkingLot
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ParkingLotController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteParkingLot(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
