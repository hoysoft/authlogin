package authlogin

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/hoysoft/authlogin/models"
)

type options struct {
	authMode    string //认证方式，空-本地用户，其他-ldap方式
	nameFromart int    //姓名显示格式
}

type AdminController struct {
	AdminBaseController
}

func (this *AdminController) Prepare() {
	this.AdminBaseController.Prepare()

	//this.LayoutSections = make(map[string]string)

	//this.LayoutSections["LayoutContent"] = "authlogin/layout.tpl"
	//this.LayoutSections["HtmlHead"] = "blogs/html_head.tpl"
}

func init() {
	//增加路由
	beego.Router("/admin", &AdminController{})
	beego.Router("/admin/:action", &AdminController{})

	Before_auth(beego.UrlFor("AdminController.Get", ":action", "new")).UnLogin()
}

type selects struct {
	Id         string
	IsSelected bool
	Value      string
}

//POST   /uri     创建
//DELETE /uri/xxx 删除
//PUT    /uri/xxx 更新或创建
//GET    /uri/xxx 查看

func (this *AdminController) Get() {
	action := this.Ctx.Input.Param(":action")
	switch action {
	case "": // 全局配置
		this.Data["Ops"] = ops
		authModes := []*selects{}
		ldaps := models.GetAllLdapConnector_sm()
		if ldaps != nil {
			for _, ldap := range *ldaps {
				if ldap.Name == "" {
					authModes = append(authModes, &selects{ldap.Name, ops.authMode == ldap.Name, "本地用户"})
				} else {
					authModes = append(authModes, &selects{ldap.Name, ops.authMode == ldap.Name, ldap.Name})
				}
			}
		}
		this.Data["authmodes"] = authModes
		nameFromarts := []*selects{
			&selects{"0", ops.nameFromart == 0, "姓+名"},
			&selects{"1", ops.nameFromart == 1, "名+姓"},
			&selects{"2", ops.nameFromart == 2, "名"},
			&selects{"3", ops.nameFromart == 3, "姓"},
		}
		this.Data["namefromarts"] = nameFromarts
		this.Data["Title"] = cnf.String("options::title")
		this.TplNames = "authlogin/options.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
		break
	case "new": // 空用户表时增加管理员帐户
		//检查是否空表
		if count, _ := models.GetUserCount(); count != 0 {
			this.Abort("404")
			return
		}
		eUser := models.User{}

		this.Data["eUser"] = &eUser
		this.Data["Title"] = cnf.String("options::newadmin")
		this.TplNames = "authlogin/user_edit.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
		break
	}

}

//POST   /uri     创建
//DELETE /uri/xxx 删除
//PUT    /uri/xxx 更新或创建
//GET    /uri/xxx 查看
func (this *AdminController) Post() {
	//	valid := validation.Validation{}
	this.Ctx.Request.ParseForm()
	action := this.Ctx.Input.Param(":action")
	switch action {
	case "": // 全局配置
		ops.authMode = this.GetString("authmode")

		ops.nameFromart, _ = this.GetInt("namefromart")

		ops.Write()

		this.Get()
		break
	case "new": // 空用户表时增加管理员帐户
		//检查是否空表
		if count, _ := models.GetUserCount(); count != 0 {
			this.Abort("404")
			return
		}
		//u * models.User
		u := models.User{}
		u.Account = this.GetString("Account")
		u.Email = this.GetString("Email")
		u.Password = this.GetString("Password")
		//if err := this.ParseForm(u); err != nil {
		//	this.Data["Message"] = err
		//	fmt.Println(err)
		//	return
		//}
		fmt.Println("ppppppppppppppppp")
		rePassword := this.Input().Get("re-password") // 重复输入的密码
		// 如果两次输入的密码不一致，需重新填写
		if u.Password != rePassword {
			this.Data["Message"] = "两次输入的密码不一致"
			return
		}
		//b, err1 := valid.Valid(&u)
		//if err1 != nil {
		//	this.Data["Message"] = "数据验证失败"
		//	return
		//}
		//if !b {
		//	this.Data["Message"] = "数据验证未通过"
		//	for _, err1 := range valid.Errors {
		//		fmt.Println(err1.Key, err1.Message)
		//	}
		//	return
		//}

		models.AddUserDefaultData(&u)
		//跳转到前面操作页面
		Redirect_HttpReferer(&this.Controller)
		break
	}
}

//POST   /uri     创建
//DELETE /uri/xxx 删除
//PUT    /uri/xxx 更新或创建
//GET    /uri/xxx 查看
func (this *AdminController) Delete() {

}

//POST   /uri     创建
//DELETE /uri/xxx 删除
//PUT    /uri/xxx 更新或创建
//GET    /uri/xxx 查看
//PATCH 局部修改
func (this *AdminController) Put() {

}

func (this *AdminController) Patch() {

}

func (this *options) Read() {
	ops.authMode = cnf.DefaultString("options::authmode", "")
	ops.nameFromart = cnf.DefaultInt("options::namefromart", 0)
	fmt.Println(ops.authMode)
}

func (this *options) Write() {
	cnf.Set("options::authmode", ops.authMode)
	cnf.Set("options::namefromart", strconv.Itoa(ops.nameFromart))
	cnf.SaveConfigFile("conf/authlogin.conf")
}

//自定义Admin渲染
func (this *AdminController) RenderHtml(tplfile string) {
	this.TplNames = "authlogin/layout.tpl"
	//f, err := ioutil.ReadFile(path.Join(beego.ViewsPath, tplfile))
	//if err != nil {
	//	fmt.Printf("%s\n", err)
	//	panic(err)
	//}

	//mContent, err := template.ParseFiles(path.Join(beego.ViewsPath, tplfile))
	//fmt.Println(mContent.Templates())
	//this.Data["AdminContent"] = template.HTML(string(f))

	//_, err := template.New("authlogin/paginator.tpl").ParseFiles(path.Join(beego.ViewsPath, "authlogin/paginator.tpl"))
	//mContent, err := template.ParseFiles(path.Join(beego.ViewsPath, tplfile))
	//if err != nil {
	//	fmt.Printf("%s\n", err)
	//	panic(err)
	//}

	var buf bytes.Buffer

	err := beego.BeeTemplates[tplfile].ExecuteTemplate(&buf, tplfile, &this.Data)
	//err = mContent.ExecuteTemplate(&buf, tplfile, &this.Data)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	this.Data["AdminContent"] = template.HTML(buf.String())
}

func (this *AdminController) Render() error {
	this.RenderHtml(this.TplNames)
	err := this.Controller.Render()
	return err
}
