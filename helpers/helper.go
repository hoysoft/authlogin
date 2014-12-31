package helpers

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"io"
	"github.com/hoysoft/authlogin/models"
	"github.com/astaxie/beego/orm"
)

type  Select struct {
	Id         string
	IsSelected bool
	Value      string
}

func GetAuthModes() []*Select {
	authModes := []*Select{}
	ldaps := models.GetAllLdapConnector_sm()
	if ldaps != nil {
		for _, ldap := range *ldaps {
				authModes = append(authModes, &Select{strconv.Itoa(ldap.Id), Ops.AuthMode == ldap.Id, ldap.Name})
		}
	}
	return authModes
}

//空表时增加默认管理帐号
func AddUserDefaultData(m *models.User) {
	ldapcnn := AddLdapConnectorDefaultData()
	m.Password =   Sha1(m.Password)
	m.Ldap = ldapcnn

	_, er := models.AddUser(m)
	if er != nil {
		fmt.Println("add user error:%s", er)
	}

}

//空表时增加默认LdapConnector(本地连接)
func AddLdapConnectorDefaultData() *models.LdapConnector {
	localuser_name:=  Cnf.DefaultString("ldap::local","Local Users")
	u := models.LdapConnector{Name: localuser_name}
	o := orm.NewOrm()
	err := o.Read(&u, "Name")
	if err != nil {
		_, er := models.AddLdapConnector(&u)
		if er != nil {
			fmt.Println("add LdapConnector error:%s", er)
			return nil
		}  
	}
	return &u
}

//对字符串进行SHA1哈希
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}