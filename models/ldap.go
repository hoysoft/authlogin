package models

import (
	//"crypto/sha1"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
	//"io"
	"reflect"
	//"strconv"
	"strings"
	"time"
)

//LDAP认证模式
// LDAP连接管理表
type LdapConnector struct {
	Id                 int       `orm:"auto;PK"`     // 连接ID
	Name               string    `orm:"size(32)"`    // 名称
	Host               string    `orm:"size(32)"`    // 主机
	Port               int       `orm:"size(12);389` // 端口
	Ldaps              bool      //ldaps 协议
	MangerAccount      string    `orm:"size(32)"`       //管理账号
	Password           string    `orm:"size(200);null"` //密码；加密形式
	BaseDN             string    `orm:"size(52)"`
	Filter             string    `orm:"size(80);null"`               //LDAP过滤器
	TimeOut            int       `orm:"size(12);0"`                  //超时（秒）
	PropUser_Account   string    `orm:"size(32)"`                    //用户登录名属性
	PropUser_FirstName string    `orm:"size(32)"`                    //名字属性
	PropUser_LastName  string    `orm:"size(32)"`                    //姓氏属性
	PropUser_Email     string    `orm:"size(32)"`                    //Email属性
	Createdtime        time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	Updatedtime        time.Time `orm:"auto_now;type(datetime)"`     // 最后修改时间
}

func init() {
	orm.RegisterModel(new(LdapConnector))
}

//空表时增加默认LdapConnector(本地连接)
func AddLdapConnectorDefaultData() *LdapConnector {
	u := LdapConnector{Name: ""}
	o := orm.NewOrm()
	err := o.Read(&u, "Name")
	if err != nil {
		l := LdapConnector{Name: ""}
		_, er := AddLdapConnector(&l)
		if er != nil {
			fmt.Println("add LdapConnector error:%s", er)
			return nil
		} else {
			return &l
		}
	}
	return &u
}

func GetLdapConnectorCount() (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("ldapConnector").Count()
	fmt.Println("count--", err)
	return
}

//增加LdapConnector
func AddLdapConnector(m *LdapConnector) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return

}

//修改LdapConnector
func UpdateLdapConnectorById(m *LdapConnector) (err error) {
	o := orm.NewOrm()
	v := LdapConnector{Id: m.Id}
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}

	}
	return
}

func GetLdapConnectorById(id int) (v *LdapConnector, err error) {
	o := orm.NewOrm()
	v = &LdapConnector{Id: id}
	if err := o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllLdapConnector_sm() *[]LdapConnector {
	var l []LdapConnector
	o := orm.NewOrm()
	o.QueryTable(new(LdapConnector)).All(&l)
	return &l
}

// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
func GetAllLdapConnector(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LdapConnector))
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

	var l []LdapConnector
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

////用户验证
//func LDAPUserAuthCheck(
//	account,
//	domain,
//	passwd,
//	ldap_server string,
//	ldap_port uint16,
//	base_dn string) UserAuth {

//	user := UserAuth{Login: false, HasThumb: false}
//	var err error

//	//connect
//	fmt.Println("Connecting")
//	l := ldap.NewLDAPConnection(ldap_server, ldap_port)
//	err = l.Connect()
//	if err != nil {
//		fmt.Printf("LDAP connectiong error: %v", err)
//		user.Err = err
//		return user
//	}
//	defer l.Close()

//	//authentification (Bind)
//	loginname := account + "@" + domain
//	err = l.Bind(loginname, passwd)
//	if err != nil {
//		fmt.Printf("error: %v", err)
//		user.Err = err
//		return user
//	}
//	user.Login = true
//	fmt.Print("Authenticated")

//	//Search, Get entries and Save entry
//	attributes := []string{}
//	filter := fmt.Sprintf(
//		"(&(objectclass=user)(samaccountname=%s))",
//		account,
//	)
//	search_request := ldap.NewSimpleSearchRequest(
//		base_dn,
//		2, //ScopeWholeSubtree 2, ScopeSingleLevel 1, ScopeBaseObject 0 ??
//		filter,
//		attributes,
//	)
//	sr, _ := l.Search(search_request)
//	user.Account = account
//	user.Name = sr.Entries[0].GetAttributeValue("name")
//	user.Mail = sr.Entries[0].GetAttributeValue("mail")
//	user.Thumb = sr.Entries[0].GetAttributeValue("thumbnailphoto")
//	if user.Thumb != "" {
//		user.HasThumb = true
//	}

//	return user
//}
