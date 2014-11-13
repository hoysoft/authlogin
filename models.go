package authlogin

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	//	"time"
	"html/template"
	//"log"
	//"strconv"
	"bytes"
)

type MethodFunc func(*beego.Controller, string, string, bool)

var (
	globalSessions *session.Manager

	loginHtmlString  string
	mLayout, LoginT  *template.Template
	actionMethodFunc MethodFunc
)

type UserController struct {
	beego.Controller
}

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "booksmanage", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()

	//var err error

	mLayout, _ = template.New("layout").Parse(`<!DOCTYPE html>
<html>
<head><title>{{.Status}} Not Authorized</title></head>
<body>
 {{.LayoutContent}}
</body>
</html>`)

	LoginT, _ = template.New("LayoutContent").Parse(`
<div class="login">
<h1>{{.Status}} Not Authorized - {{.Message}}</h1>
<form action="{{.Url}}" method="post">
<label for="username">username: </label>
<input type="text" name="username" id="username" value="" placeholder="(Enter your username)"><br>
<label for="password">Password: </label>
<input type="password" name="password" id="password" value="" placeholder="(Enter your password)"><br>
<input type='checkbox' name='save_login' value='1' checked='checked'/> 记住我的登录信息<br>
<br>
<a href="/user/reset-pwd">忘记登录密码？</a> 
<input type='button' value='注册'/> 
<input type="submit" value="登录">

</form>
</div>
`)

	//增加路由
	beego.Router("/user", &UserController{})
	beego.Router("/user/:action", &UserController{})

}

func ActionMethod(methodFunc MethodFunc) {
	actionMethodFunc = methodFunc
}

// 检测用户是否登录
func LoginUser(this *beego.Controller, autoRedirect bool) *User {
	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userid := sess.Get("Uid")
	if userid != nil {
		user, e := GetUserById(userid.(int))
		if e == nil {
			return user //用户已经登录，返回用户信息
		}
	}

	if this.Ctx.Request.RequestURI != "/user/login" && autoRedirect { // 用户未登录，自动跳转
		this.Ctx.Redirect(302, "/user/login") // 跳转到用户登录页面
	}
	return nil
}

func (this *UserController) Get() {
	//this.Layout = tpLayout                    // 后台管理模板布局文件
	action := this.Ctx.Input.Param(":action") // 用户的添加、修改或删除

	switch action {
	case "": //用户列表
		users := GetAllUse_sm()
		this.Data["Users"] = users

	case "login": // 用户登录
		data := make(map[string]interface{})
		data["Status"] = "YYYY"
		data["Url"] = "/user/login"

		buff := bytes.NewBufferString("")
		LoginT.Execute(buff, data)

		if actionMethodFunc != nil {
			defaultContent := true
			actionMethodFunc(&this.Controller, "Get", "login", defaultContent)
			if !defaultContent {
				return
			}

			this.HtmlContent = buff.String()
			return
		}

		mLayout.Execute(this.Ctx.ResponseWriter, data)

		//this.Ctx.WriteString(loginT)

		//case "logout": // 用户退出
		//	this.Layout = "layout_one.tpl"     // 用户登录模板布局文件
		//	this.DelSession("account")         // 删除session中的用户登录信息
		//	this.TplNames = "admin/logout.tpl" // 页面模板文件
	case "add": //注册用户
	//this.Ctx.WriteString(loginT)
	case "reset-pwd": //密码复位
	case "logout": // 用户退出
		sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		defer sess.SessionRelease(this.Ctx.ResponseWriter)
		sess.Delete("Uid")
		this.Ctx.Redirect(302, "/")
	}

}

func (this *UserController) Post() {
	//	this.Layout = "layout_admin.tpl"           // 模板布局文件
	action := this.Ctx.Input.Param(":action") // 用户的添加或修改
	switch action {
	//case "add": // 添加用户
	//	email := this.Input().Get("email")            // 用户E-mail
	//	name := this.Input().Get("name")              // 用户名
	//	password := this.Input().Get("password")      // 密码
	//	rePassword := this.Input().Get("re-password") // 重复输入的密码

	//	// 检测E-mail或密码是否为空
	//	if email == "" || name == "" {
	//		this.Data["Message"] = "E-mail或用户名为空"
	//		this.Data["Email"] = email
	//		this.Data["Name"] = name
	//		this.Data["Password"] = password
	//		this.TplNames = "admin/add_user.tpl"
	//		return
	//	}

	//	// 如果两次输入的密码不一致，需重新填写
	//	if password != rePassword {
	//		this.Data["Message"] = "两次输入的密码不一致"
	//		this.Data["Email"] = email
	//		this.Data["Name"] = name
	//		this.Data["Password"] = password
	//		this.TplNames = "admin/add_user.tpl"
	//		return
	//	}

	//	// 检查E-mail或用户名是否已存在
	//	orm = InitDb()
	//	user := User{}
	//	err = orm.Where("email=? or name=?", email, name).Find(&user)
	//	if err == nil {
	//		this.Data["Message"] = "E-mail或用户名已存在"
	//		this.Data["Email"] = email
	//		this.Data["Name"] = name
	//		this.Data["Password"] = password
	//		this.TplNames = "admin/add_user.tpl"
	//		return
	//	}

	//	// 保存用户
	//	orm = InitDb()
	//	user = User{}
	//	user.Email = email
	//	user.Name = name
	//	user.Password = Sha1(password)
	//	user.Created = time.Now()
	//	err = orm.Save(&user)
	//	Check(err)
	//	Debug("User `%s` added.", user)

	//	this.Ctx.Redirect(302, "/user/") // 返回用户列表页面
	//case "edit": // 修改用户
	//	id := this.Ctx.Input.Params(":id") // 用户ID

	//	email := this.Input().Get("email")            // 用户E-mail
	//	name := this.Input().Get("name")              // 用户名
	//	password := this.Input().Get("password")      // 密码
	//	rePassword := this.Input().Get("re-password") // 重复输入的密码

	//	// 检测E-mail或密码是否为空
	//	if email == "" || name == "" {
	//		this.Data["Message"] = "E-mail或用户名为空"
	//		this.Data["Email"] = email
	//		this.Data["Name"] = name
	//		this.Data["Password"] = password
	//		this.TplNames = "admin/add_user.tpl"
	//		return
	//	}

	//	// 如果两次输入的密码不一致，需重新填写
	//	if password != rePassword {
	//		this.Data["Message"] = "两次输入的密码不一致"
	//		this.Data["Email"] = email
	//		this.Data["Name"] = name
	//		this.Data["Password"] = password
	//		this.TplNames = "admin/add_user.tpl"
	//		return
	//	}

	//	// 获得当前用户
	//	orm = InitDb()
	//	user := User{}
	//	err = orm.Where("id=?", id).Find(&user)
	//	Check(err)

	//	// 更新用户信息
	//	user.Email = email
	//	user.Name = name
	//	if password != "" {
	//		user.Password = Sha1(password)
	//	}
	//	user.Updated = time.Now()

	//	// 保存用户信息
	//	err = orm.Save(&user)
	//	Check(err)

	//	this.Ctx.Redirect(302, "/user/") // 返回用户列表页面
	case "login": // 用户登录
		name := this.Ctx.Request.FormValue("username")     // 用户名
		password := this.Ctx.Request.FormValue("password") // 用户密码
		// 检测用户名或密码是否为空
		if name == "" || password == "" {
			this.Data["Message"] = "用户名或密码为空"
			this.Ctx.Redirect(302, "/user/login")
			return
		}
		user, b := AuthUser(name, password)
		if !b || user == nil {
			this.Data["Message"] = "用户名或密码错误！"
			this.Ctx.Redirect(302, "/user/login")
			return
		} else {
			this.Data["Message"] = "登录成功！"
			// 保存用户登录信息
			sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
			defer sess.SessionRelease(this.Ctx.ResponseWriter)
			sess.Set("Uid", user.Id)
			this.Ctx.Redirect(302, "/v1/site")
		}
		// 如果不存在任何用户，那么直接以admin身份登录
		//orm = InitDb()
		//users := []User{}
		//err = orm.FindAll(&users)
		//if err == nil && len(users) == 0 {
		//	this.SetSession("Uid", "0")
		//	this.Ctx.Redirect(302, "/admin/")
		//}

		//	errFlag := false // 判断登录是否出错

		//	orm = InitDb()
		//	user := User{}
		//	err = orm.Where("name=? and password=?", name, Sha1(password)).Find(&user)
		//	if err != nil {
		//		this.Data["Message"] = "用户名或密码错误"
		//		errFlag = true
		//	} else { // 用户名、密码验证成功
		//		// 保存用户登录信息
		//		sess.Set("account", name)

		//		this.Ctx.Redirect(302, "/admin/") // 跳转到管理后台首页
		//	}

		//	// 显示用户登录页面，再次登录
		//	if errFlag {
		//		this.Data["Name"] = name
		//		this.Data["Password"] = password
		//		this.Layout = "layout_one.tpl"
		//		SiteName := beego.AppConfig.String("appname") // 网站名称
		//		this.Data["SiteName"] = SiteName
		//		this.Data["Categories"] = GetCategories() // 分类列表，用于导航栏
		//		this.TplNames = "admin/login.tpl"
		//		return
		//	}

	}
}
