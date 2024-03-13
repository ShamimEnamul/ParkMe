package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ParkingSlot struct {
	Id            int64      `orm:"auto"`
	ParkingLotID  int        `json:"parkingLotID" orm:"column(parking_lot_id)"`
	IsOccupied    bool       `json:"isOccupied" orm:"column(is_occupied);default(false)"`
	IsMaintenance bool       `json:"IsMaintenance" orm:"column(is_maintenance);default(false)"`
	CreateDate    time.Time  `json:"-" orm:"auto_now_add; type(datetime); column(create_date)"`
	UpdateDate    *time.Time `json:"-" orm:"column(update_date)"`
}

func (n *ParkingSlot) TableName() string {
	return "parking_slots"
}
func init() {
	orm.RegisterModel(new(ParkingSlot))
}

func NewParkingSlot(parkingLotID int) *ParkingSlot {
	return &ParkingSlot{
		ParkingLotID:  parkingLotID,
		IsOccupied:    false,
		IsMaintenance: false,
		CreateDate:    time.Now(),
	}
}

// AddParkingSlot insert a new ParkingSlot into database and returns
// last inserted Id on success.
func AddParkingSlot(m *ParkingSlot) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)

	if err != nil {
		fmt.Println(err)
		return
	}
	return id, nil
}

// GetParkingSlotById retrieves ParkingSlot by Id. Returns error if
// Id doesn't exist
func GetParkingSlotById(id int64) (v *ParkingSlot, err error) {
	o := orm.NewOrm()
	v = &ParkingSlot{Id: id}
	if err = o.QueryTable(new(ParkingSlot)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllParkingSlot retrieves all ParkingSlot matches certain condition. Returns empty list if
// no records exist
func GetAllParkingSlot(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ParkingSlot))
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

	var l []ParkingSlot
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

// UpdateParkingSlot updates ParkingSlot by Id and returns error if
// the record to be updated doesn't exist
func UpdateParkingSlotById(m *ParkingSlot) (err error) {
	o := orm.NewOrm()
	v := ParkingSlot{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteParkingSlot deletes ParkingSlot by Id and returns error if
// the record to be deleted doesn't exist
func DeleteParkingSlot(id int64) (err error) {
	o := orm.NewOrm()
	v := ParkingSlot{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ParkingSlot{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetCountOfAllParkingSlotsByParkingId(id int64) (int, error) {
	query := "select count(id) from parking_slots where parking_lot_id = ?;"

	var count int
	o := orm.NewOrm()
	err := o.Raw(query, id).QueryRow(&count)

	if err != nil {
		return 0, errors.New("DB query error occurred")
	}

	return count, nil
}

func GetAvailAbleParkingSlotByParkingLotId(parkingLotId int) (int, error) {
	query := "select id as slot_no from parking_slots where parking_lot_id = ? and is_occupied = 0 and is_maintenance = 0 order by id asc limit 1;"

	var parkingSlotId int
	o := orm.NewOrm()
	err := o.Raw(query, parkingLotId).QueryRow(&parkingSlotId)

	if err != nil {
		return 0, errors.New("DB query error occurred")
	}

	return parkingSlotId, nil
}
