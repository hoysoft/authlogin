package authlogin

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type options struct {
	authMode    string //认证方式，空-本地用户，其他-ldap方式
	nameFromart int    //姓名显示格式
}

type AdminController struct {
	beego.Controller
}

func init() {
	//增加路由
	beego.Router("/admin", &AdminController{})
}

type selects struct {
	Id         string
	IsSelected bool
	Value      string
}

func (this *AdminController) Get() {
	action := this.Ctx.Input.Param(":action")
	this.Data["Ops"] = ops
	authModes := []*selects{
		&selects{"0", ops.authMode == "", "本地用户"},
	}
	ldaps := GetAllLdapConnector_sm()
	if ldaps != nil {
		for _, ldap := range *ldaps {
			authModes = append(authModes, &selects{ldap.Name, ops.authMode == ldap.Name, ldap.Name})
		}
	}
	this.Data["authModes"] = authModes
	nameFromarts := []*selects{
		&selects{"0", ops.nameFromart == 0, "姓+名"},
		&selects{"1", ops.nameFromart == 1, "名+姓"},
		&selects{"2", ops.nameFromart == 2, "名"},
		&selects{"3", ops.nameFromart == 3, "姓"},
	}
	this.Data["nameFromarts"] = nameFromarts
	this.Data["Title"] = cnf.String("options::title")
	this.TplNames = "authlogin/options.html"
	if runActionMethodBefoer(&this.Controller, "Get", action) {
		return
	}

	return
}

func (this *AdminController) Post() {
	this.Ctx.Request.ParseForm()
	ops.authMode = this.GetString("authMode")

	ops.nameFromart, _ = this.GetInt("nameFromart")

	ops.Write()

	this.Get()
}

func (this *options) Read() {
	fmt.Printf("bbbbbbbbbbbbbbbbbbbbbbbbbb")
	ops.authMode = cnf.DefaultString("options::authMode", "")
	ops.nameFromart = cnf.DefaultInt("options::nameFromart", 0)
	fmt.Println(ops.authMode)
}

func (this *options) Write() {
	cnf.Set("options::authMode", ops.authMode)
	cnf.Set("options::nameFromart", strconv.Itoa(ops.nameFromart))
	cnf.SaveConfigFile("conf/authlogin.conf")
}
