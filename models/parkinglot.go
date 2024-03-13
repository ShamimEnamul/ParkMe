package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ParkingLot struct {
	Id         int64      `orm:"auto; pk"`
	Name       string     `valid:"Required" json:"name" orm:"size(128); column(name)"`
	Capacity   int        `valid:"Required" json:"capacity" orm:"column(capacity)"`
	Active     bool       `valid:"Required" json:"active" orm:"column(active); default(true)"`
	CreatedBy  int        `valid:"Required" json:"createdBy" orm:"column(created_by)"`
	CreateDate time.Time  `json:"-" orm:"auto_now_add; type(datetime); column(create_date)"`
	UpdateDate *time.Time `json:"-" orm:"column(update_date)"`
}

func (n *ParkingLot) TableName() string {
	return "parking_lots"
}

func init() {
	orm.RegisterModel(new(ParkingLot))
}

func NewParkingLot(name string, capacity int, createdBy int) *ParkingLot {
	return &ParkingLot{
		Name:       name,
		Capacity:   capacity,
		Active:     true,
		CreatedBy:  createdBy,
		CreateDate: time.Now(),
	}
}

// AddParkingLot insert a new ParkingLot into database and returns
// last inserted Id on success.
func AddParkingLot(m *ParkingLot) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)

	if err != nil {
		return
	}
	return id, nil
}

// GetParkingLotById retrieves ParkingLot by Id. Returns error if
// Id doesn't exist
func GetParkingLotById(id int64) (v *ParkingLot, err error) {
	o := orm.NewOrm()
	v = &ParkingLot{Id: id}
	if err = o.QueryTable(new(ParkingLot)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetParkingLotCapacityById(id int64) (int, error) {
	query := "select `capacity` from parking_lots where id = ? and active = 1;"

	var parkingLot ParkingLot
	o := orm.NewOrm()
	err := o.Raw(query, id).QueryRow(&parkingLot)

	if err != nil {
		return 0, errors.New("DB query error occured")
	}

	return parkingLot.Capacity, err
}

// GetAllParkingLot retrieves all ParkingLot matches certain condition. Returns empty list if
// no records exist
func GetAllParkingLotByUserId(id int64) (v *[]ParkingLot, err error) {
	query := `select * from parking_lots where created_by = ? and active = 1;`

	var parkingLotList []ParkingLot

	o := orm.NewOrm()

	_, err = o.Raw(query, id).QueryRows(&parkingLotList)

	if err != nil {
		return nil, errors.New("DB query error occured")
	}

	return &parkingLotList, err
}

// GetAllParkingLot retrieves all ParkingLot matches certain condition. Returns empty list if
// no records exist
func GetAllParkingLot(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ParkingLot))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ParkingLot
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateParkingLot updates ParkingLot by Id and returns error if
// the record to be updated doesn't exist
func UpdateParkingLotById(m *ParkingLot) (err error) {
	o := orm.NewOrm()
	v := ParkingLot{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteParkingLot deletes ParkingLot by Id and returns error if
// the record to be deleted doesn't exist
func DeleteParkingLot(id int64) (err error) {
	o := orm.NewOrm()
	v := ParkingLot{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ParkingLot{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type ParkingLotDetails struct {
	TotalFee     int64   `json:"totalFee" orm:"column(total_fee)"`
	TotalHour    float32 `json:"totalHour" orm:"column(total_hour)"`
	TotalVehicle int64   `json:"totalVehicle" orm:"column(total_vehicle)"`
}

func GetParkingLotDetailsByRange(parkingLotId int64, startTime, endTime string) (*ParkingLotDetails, error) {
	//
	query := `SELECT SUM(TIMESTAMPDIFF(HOUR, in_time, out_time) * 10) AS total_fee,  SUM(TIMESTAMPDIFF(HOUR, in_time, out_time)) as total_hour, count(distinct(vehicle_id)) total_vehicle
FROM reservations
WHERE parking_lot_id = ? 
and in_time between ? and ?;`
	fmt.Println(query)
	var parkingLotDetails *ParkingLotDetails

	o := orm.NewOrm()

	err := o.Raw(query, parkingLotId, startTime, endTime).QueryRow(&parkingLotDetails)

	if err != nil {
		return nil, errors.New("DB query error occured")
	}

	return parkingLotDetails, err
}
