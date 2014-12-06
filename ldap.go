package authlogin

import (
	"fmt"
	//	"github.com/mavricknz/ldap"
	"github.com/hoysoft/authlogin/models"
	//"reflect"
	//	"encoding/json"
	"errors"
	"strconv"

	"github.com/astaxie/beego"
)

//type UserAuth struct {
//	Login    bool
//	HasThumb bool
//	Account  string
//	Mail     string
//	Thumb    string
//	Name     string
//	Err      error
//}

type LdapController struct {
	AdminBaseController
}

func init() {
	//增加路由
	beego.Router("/ldap", &LdapController{})
	beego.Router("/ldap/:action", &LdapController{})
	beego.Router("/ldap/:action/:id", &LdapController{})
}

func (this *LdapController) SetPaginator(per int, nums int64) *Paginator {
	p := NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *LdapController) Get() {
	action := this.Ctx.Input.Param(":action")
	switch action {
	case "": //LDAP列表
		p := this.Ctx.Request.URL.Query().Get("p")
		pageNo, _ := strconv.Atoi(p)
		if pageNo == 0 {
			pageNo = 1
		}
		var fields []string
		var sortby []string
		var order []string
		var limit int64 = 6 //每页6行显示
		var offset int64 = (int64(pageNo) - 1) * limit
		//var query map[string]string = map[string]string{"source": "0"}
		var query map[string]string = map[string]string{}
		ldaps, count, _ := models.GetAllLdapConnector(query, fields, sortby, order, offset, limit)
		//ldaps, count, _ := Table_GetAll(&LDAPConnector{}, query, fields, sortby, order, offset, limit)
		this.Data["Ldaps"] = &ldaps
		//count, _ := GetUserCount()
		_ = this.SetPaginator(int(limit), count)

		this.Data["Title"] = cnf.String("ldap_all::title")
		this.TplNames = "authlogin/ldap_all.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
		//s := readFile("test.txt")
		//this.Ctx.WriteString(s)
		return
	case "add": //LDAP新增
		ld := models.LdapConnector{}
		this.Data["ldap"] = ld
		this.Data["Title"] = cnf.String("ldap_add::title")
		this.TplNames = "authlogin/ldap_edit.html"

		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}

		return
	case "edit": //LDAP修改
		s := this.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(s)
		if err != nil {
			return
		}

		ld, err := models.GetLdapConnectorById(id)
		if err != nil {
			return
		}
		fmt.Println("ldd:", ld)
		this.Data["ldap"] = ld
		this.Data["Title"] = cnf.String("ldap_edit::title")
		this.TplNames = "authlogin/ldap_edit.html"

		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}

		return
	}
}

func post_LdapConnector(this *LdapController, ld *models.LdapConnector) (err error) {
	if err = this.ParseForm(ld); err != nil {
		this.Data["Message"] = err
		return
	}
	if ld.Name == "" {
		err = errors.New("标识不能为空")
		this.Data["Message"] = err

		return
	}
	return
}

func (this *LdapController) Post() {
	action := this.Ctx.Input.Param(":action")
	switch action {
	case "add": //LDAP新增
		ld := models.LdapConnector{}
		err := post_LdapConnector(this, &ld)
		if err != nil {
			this.Render()
			return
		}
		models.AddLdapConnector(&ld)

		this.Redirect("/ldap", 302)

		return
	case "edit": //LDAP修改

		s := this.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(s)
		if err != nil {
			this.Render()
			return
		}

		ld, err := models.GetLdapConnectorById(id)
		if err != nil {
			return
		}
		err = post_LdapConnector(this, ld)
		if err != nil {
			this.Get()
			return
		}
		err = models.UpdateLdapConnectorById(ld)
		if err != nil {
			fmt.Println("err:", err)
		}
		this.Redirect("/ldap", 302)

		return
	}
}

//login check, input account, domain, passwd, server, port and base_dn for search
//func saAuthCheck(
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

//func main() {
//	SaAuthCheck("cn=hqj,dc=hz,dc=com","","","192.168.8.1",389,"dc=hz,dc=com")

//}
