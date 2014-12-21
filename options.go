package authlogin

import (
	"bytes"
	"fmt"
	"html/template"

	"strings"

	"github.com/astaxie/beego"
	"github.com/hoysoft/authlogin/helpers"
	"github.com/hoysoft/authlogin/models"
)

type AdminController struct {
	AdminBaseController
}

func (this *AdminController) Prepare() {
	this.AdminBaseController.Prepare()
	this.Data["cnf"] = helpers.Cnf

	tags := strings.Split(this.Ctx.Request.URL.Path, "/")
	if len(tags) > 1 {
		this.Data["__Tag"] = tags[1]
	} else {
		this.Data["__Tag"] = ""
	}

	//this.LayoutSections = make(map[string]string)

	//this.LayoutSections["LayoutContent"] = "authlogin/layout.tpl"
	//this.LayoutSections["HtmlHead"] = "blogs/html_head.tpl"
}

func init() {
	//增加路由
	beego.Router("/admin", &AdminController{})
	beego.Router("/admin/:action", &AdminController{})

	Before_auth(beego.UrlFor("AdminController.Get", ":action", "new")).NonLogin()
}

//POST   /uri     创建
//DELETE /uri/xxx 删除
//PUT    /uri/xxx 更新或创建
//GET    /uri/xxx 查看

func (this *AdminController) Get() {
	action := this.Ctx.Input.Param(":action")
	switch action {
	case "": // 全局配置
		this.Data["Ops"] = helpers.Ops
		this.Data["authmodes"] = helpers.GetAuthModes()
		nameFromarts := []*helpers.Select{
			&helpers.Select{"0", helpers.Ops.NameFromart == 0, "姓+名"},
			&helpers.Select{"1", helpers.Ops.NameFromart == 1, "名+姓"},
			&helpers.Select{"2", helpers.Ops.NameFromart == 2, "名"},
			&helpers.Select{"3", helpers.Ops.NameFromart == 3, "姓"},
		}
		this.Data["namefromarts"] = nameFromarts
		this.Data["Title"] = helpers.Cnf.String("options::title")
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
		this.Data["Title"] = helpers.Cnf.String("options::newadmin")
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

		helpers.Ops.AuthMode, _ = this.GetInt("authmode")

		helpers.Ops.NameFromart, _ = this.GetInt("namefromart")

		helpers.Ops.Write()

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

		helpers.AddUserDefaultData(&u)
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

	if beego.BeeTemplates[tplfile] != nil {
		var buf bytes.Buffer

		err := beego.BeeTemplates[tplfile].ExecuteTemplate(&buf, tplfile, &this.Data)
		//err = mContent.ExecuteTemplate(&buf, tplfile, &this.Data)
		if err != nil {
			fmt.Printf("%s\n", err)
			panic(err)
		}
		this.Data["AdminContent"] = template.HTML(buf.String())
	}
}

func (this *AdminController) Render() error {
	this.RenderHtml(this.TplNames)
	err := this.Controller.Render()
	return err
}
