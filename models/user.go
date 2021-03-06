package models

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
	//"strconv"
	"strings"
	"time"
)

// 用户表
type User struct {
	Id            int            `orm:"auto;PK"`                                     // 用户ID
	Email         string         `orm:"size(32);unique" valid:"Email; MaxSize(100)"` // 用户Email
	LastName      string         `orm:"size(32)"`                                    //姓氏
	FirstName     string         `orm:"size(32)"`                                    //名字
	Password      string         `orm:"size(32)`                                     // 密码；加密形式
	Account       string         `orm:"size(32)`                                     //用户登录名
	MobileNumber  string         `orm:"size(12);null" valid:"Mobile"`                //手机号
	Remark        string         `orm:"size(100);null"`                              //备注
	Question      string         `orm:"size(100);null"`                              //提示问题，找回密码用
	Answer        string         `orm:"size(100);null"`                              //信息答案，找回密码用
	Status        int            //状态 0-未激活 1-在线 2-禁言
	IsOnline      bool           //是否在线
	Ldap          *LdapConnector `orm:"rel(one)"`                                  //LDAP连接id
	Createdtime   time.Time      `orm:"auto_now_add;type(datetime)"`               // 用户创建时间
	Updatedtime   time.Time      `orm:"auto_now;type(datetime)"`                   // 用户最后修改时间
	Lastlogintime time.Time      `orm:"column(lastlogintime);type(datetime);null"` // 用户最后登录时间
}

func init() {
	orm.RegisterModel(new(User))
}



func GetUserCount() (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("user").Count()
	fmt.Println("count--", err)
	return
}

//增加用户
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//修改用户
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//删除用户
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err := o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllUse_sm() *[]User {
	var l []User
	o := orm.NewOrm()
	o.QueryTable(new(User)).All(&l)
	return &l
}

// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//记录数
	count, _ = qs.Count()
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
					return nil, count, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, count, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, count, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, count, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
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
		return ml, count, nil
	}
	return nil, count, err
}
