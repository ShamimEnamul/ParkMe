package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Reservation struct {
	Id            int64     `orm:"auto; pk"`
	ParkingLotId  int       `json:"parkingLotId" orm:"column(parking_lot_id)"`
	ParkingSlotId int       `json:"parkingSlotId" orm:"column(parking_slot_id)"`
	VehicleId     int       `json:"vehicleId" orm:"column(vehicle_id)"`
	InTime        time.Time `json:"-" orm:"auto_now_add; type(datetime); column(in_time)"`
	OutTime       time.Time `json:"-" orm:"column(out_time)"`
	Active        bool      `json:"-" orm:"column(active)"`
	//CreateDate    time.Time `json:"-" orm:"auto_now_add; type(datetime); column(create_date)"`
	//UpdateDate    time.Time `json:"-" orm:"column(update_date)"`
}

func (n *Reservation) TableName() string {
	return "reservations"
}

func init() {
	orm.RegisterModel(new(Reservation))
}

func NewReservation(parkingLotId, parkingSlotId, vehicleId int) *Reservation {
	return &Reservation{
		ParkingLotId:  parkingLotId,
		ParkingSlotId: parkingSlotId,
		VehicleId:     vehicleId,
		InTime:        time.Now(),
		Active:        true,
		//CreateDate:    time.Now(),
	}
}

// AddReservation insert a new Reservation into database and returns
// last inserted Id on success.
func AddReservation(m *Reservation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReservationById retrieves Reservation by Id. Returns error if
// Id doesn't exist
func GetReservationById(id int64) (v *Reservation, err error) {
	o := orm.NewOrm()
	v = &Reservation{Id: id}
	if err = o.QueryTable(new(Reservation)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReservation retrieves all Reservation matches certain condition. Returns empty list if
// no records exist
func GetAllReservation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Reservation))
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

	var l []Reservation
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

// UpdateReservation updates Reservation by Id and returns error if
// the record to be updated doesn't exist
func UpdateReservationById(m *Reservation) (err error) {
	o := orm.NewOrm()

	// Define the SQL query
	sql := `
        UPDATE reservations
        SET active = false,
            out_time = now()
        WHERE parking_lot_id = ? 
          AND parking_slot_id = ?
          AND vehicle_id = ?
          AND active = true
        ORDER BY id DESC
        LIMIT 1
    `

	// Execute the raw SQL query with parameters
	_, err = o.Raw(sql, m.ParkingLotId, m.ParkingSlotId, m.VehicleId).Exec()
	if err != nil {
		return err
	}

	return nil
}

// DeleteReservation deletes Reservation by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReservation(id int64) (err error) {
	o := orm.NewOrm()
	v := Reservation{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Reservation{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetDurationOfParking(parkingLotId int, parkingSlotId int, vehicleId int) (float32, error) {
	query := `SELECT in_time, out_time
	FROM reservations where parking_lot_id = ? and  parking_slot_id = ? and vehicle_id = ? and active = 0 order by id desc limit 1 ;`
	var in_time time.Time
	var out_time time.Time
	o := orm.NewOrm()
	err := o.Raw(query, parkingLotId, parkingSlotId, vehicleId).QueryRow(&in_time, &out_time)

	if err != nil {
		return 0, errors.New("DB query error occurred")
	}
	d := out_time.Sub(in_time).Hours()
	return float32(d), nil
}
