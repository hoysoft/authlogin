package authlogin

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/coscms/xweb/validation"
	"github.com/hoysoft/authlogin/models"
	//"github.com/coscms/forms"

	"strings"
	"time"
	//"log"
	"strconv"
	//"bytes"
	//	"io"
	"io/ioutil"
	"net/url"
	//	"mime"
	"os"
	"path"
	"runtime"
)

var (
//globalSessions *session.Manager

//loginHtmlString        string
//mLayout, LoginT        *template.Template
//actionMethodBefoerFunc ActionMethodBefoerFunc
//AuthViewPath           string //
//cnf                    config.ConfigContainer
)

type UserController struct {
	AdminController
}

func init() {

	//增加路由
	beego.Router("/user", &UserController{})
	beego.Router("/user/:action", &UserController{})
	beego.Router("/user/:action/:id", &UserController{})

	//Before_auth(beego.UrlFor("UserController.Get", ":action", "login")).UnLogin()

	Before_auth("/user/login").UnLogin()
}

func readFile(pathfilename string) string {
	//获取源文件代码路径 (ref)：
	_, filename, _, _ := runtime.Caller(1)
	fi, err := os.Open(path.Join(path.Dir(filename), pathfilename))

	//	fi, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

// 检测用户是否登录
//func LoginUser(this *beego.Controller, autoRedirect bool) *models.User {
//	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
//	defer sess.SessionRelease(this.Ctx.ResponseWriter)
//	userid := sess.Get("Uid")
//	if userid != nil {
//		user, e := models.GetUserById(userid.(int))
//		if e == nil {
//			return user //用户已经登录，返回用户信息
//		}
//	}

//	if this.Ctx.Request.RequestURI != "/user/login" && autoRedirect { // 用户未登录，自动跳转
//		this.Ctx.Redirect(302, "/user/login") // 跳转到用户登录页面
//	}
//	return nil
//}

func (this *UserController) SetPaginator(per int, nums int64) *Paginator {
	p := NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

type category struct {
	Id         string
	IsSelected bool
	Value      string
}

type Form_User_login struct {
	Username  string `form_label:"用户:" form_value:"test" form_fieldset:"fff" form_id:"us"`
	Password1 string `form_widget:"password" form_label:"密码:" form_fieldset:"fff" form_id:"pa"`
	SkipThis  int    `form_options:"-"`
}

func (this *UserController) Get() {
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)

	action := this.Ctx.Input.Param(":action") // 用户的添加、修改或删除
	this.Data["cnf"] = cnf
	switch action {
	case "": //用户列表
		p := this.Ctx.Request.URL.Query().Get("p")
		pageNo, _ := strconv.Atoi(p)
		if pageNo == 0 {
			pageNo = 1
		}
		var fields []string
		var sortby []string
		var order []string
		//	fields := []string{"id", "email", "username", "nickname", "status", "Createdtime", "Lastlogintime"}
		var limit int64 = 6 //每页10行显示
		var offset int64 = (int64(pageNo) - 1) * limit
		//var query map[string]string = map[string]string{"source": "0"}
		var query map[string]string = map[string]string{}
		users, count, _ := models.GetAllUser(query, fields, sortby, order, offset, limit)
		//fmt.Println("user:", users)
		this.Data["Users"] = &users
		//count, _ := GetUserCount()
		_ = this.SetPaginator(int(limit), count)

		usersources := []*category{
			&category{"0", true, "本地用户"},
			&category{"1", false, "LDAP用户"},
		}
		this.Data["usersources"] = usersources

		userstates := []*category{
			&category{"-1", true, "全部"},
			&category{"0", true, "注册"},
			&category{"1", false, "激活"},
			&category{"2", false, "锁定"},
		}
		this.Data["userstates"] = userstates
		this.Data["Title"] = cnf.String("user_all::title")
		this.TplNames = "authlogin/user_all.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
		//s := readFile("test.txt")
		//this.Ctx.WriteString(s)
		return
	case "export": //导出用户列表
		users := models.GetAllUse_sm()
		lang, err := json.Marshal(users)
		if err == nil {
			userAgent := strings.ToLower(this.Ctx.Request.UserAgent())
			newName := time.Now().Format("2006-01-02_15:04:05") + ".users"
			filename := ""

			switch {
			case strings.Index(userAgent, "msie") != -1: // IE浏览器，只能采用URLEncoder编码
				filename = "=" + url.QueryEscape(newName)
				break
			case strings.Index(userAgent, "firefox") != -1: // FireFox浏览器，可以使用MimeUtility或filename*或ISO编码的中文输出
				filename = "*=UTF-8''" + url.QueryEscape(newName)
				break
			case strings.Index(userAgent, "applewebkit") != -1: // Chrome浏览器，只能采用MimeUtility编码或ISO编码的中文输出
				//  new_filename = MimeUtility.encodeText(filename, "UTF8", "B");
				//  rtn = "filename=\"" + new_filename + "\"";
				//  rtn = "filename=\"" + new String(filename.getBytes("UTF-8"),"ISO8859-1") + "\"";
				filename = `="` + url.QueryEscape(newName) + `"`
				break
			case strings.Index(userAgent, "safari") != -1: // Safari浏览器，只能采用ISO编码的中文输出
				filename = `="` + newName + `"`
				break
			case strings.Index(userAgent, "opera") != -1: // Opera浏览器只能采用filename*
				filename = `*="UTF-8''` + url.QueryEscape(newName) + `"`
				break
			default:
				filename = "=" + url.QueryEscape(newName)
			}

			this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/octet-stream")
			this.Ctx.ResponseWriter.Header().Set("Content-Disposition:", "attachment;filename"+filename)
			this.Ctx.WriteString(string(lang))
		}
	case "edit": //用户编辑
		s := this.Ctx.Input.Param(":id")
		uid, err := strconv.Atoi(s)
		if err != nil {

		}

		eUser, err := models.GetUserById(uid)
		if err != nil {

		}
		this.Data["eUser"] = &eUser
		this.Data["Title"] = cnf.String("user_edit::title")
		this.TplNames = "authlogin/user_edit.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
	case "import": //导入用户列表
		this.Data["Title"] = cnf.String("user_import::title")
		this.TplNames = "authlogin/user_import.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
	case "login": // 用户登录
		if this.LoginUser != nil {
			flash := beego.NewFlash()
			flash.Notice("您已经登录，不需要再次登录")
			flash.Store(&this.Controller)
			this.Redirect("/", 302)
		}
		//_, err := cnf.GetSection("login")

		Http_referer := this.Ctx.Request.Header.Get("HTTP_REFERER")

		if Http_referer != "" {
			setSessions(&this.Controller, "http_referer", Http_referer)
		}
		this.Data["Title"] = cnf.String("login::title")

		this.TplNames = "authlogin/login.html"
		//this.RenderHtml("authlogin/login.html")
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
		//rs, _ := this.RenderString()
		//mContent,_:=template.ParseFiles("authlogin/login.html")
		//this.AdminController.mTemplate.

		//mt, _ := this.AdminController.mTemplate.ParseFiles("authlogin/login.html")
		//rs, _ := this.RenderString()
		//var buf bytes.Buffer
		//mt.Parse(rs).ExecuteTemplate(&buf, "LayoutContent", this.Data)
		//this.RenderBytes(&buf)
		break
	case "add": //注册用户

		this.Data["Title"] = cnf.String("user_add::title")
		this.TplNames = "authlogin/user_add.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
	case "delete": // 删除用户
		s := this.Ctx.Input.Param(":id")
		uid, err := strconv.Atoi(s)
		if err != nil {

		}

		models.DeleteUser(uid)

		this.Ctx.WriteString("<script> history.back(1); </script>")
	case "reset-pwd": //密码复位
	case "logout": // 用户退出
		sess.Delete("Uid")
		this.TplNames = "authlogin/logout.html"
		if runActionMethodBefoer(&this.AdminBaseController, "Get", action) {
			return
		}
	}

}

func (this *UserController) Post() {
	valid := validation.Validation{}
	//	this.Layout = "layout_admin.tpl"           // 模板布局文件
	action := this.Ctx.Input.Param(":action") // 用户的添加或修改
	switch action {
	case "add": // 添加用户
		user := models.User{}
		user.Email = this.Input().Get("email")        // 用户E-mail
		user.Account = this.Input().Get("name")       // 用户名
		password := this.Input().Get("password")      // 密码
		rePassword := this.Input().Get("re-password") // 重复输入的密码

		this.Data["eUser"] = &user
		this.Data["Title"] = cnf.String("user_add::title")
		this.TplNames = "authlogin/user_add.html"
		// 检测E-mail或密码是否为空
		if user.Email == "" || user.Account == "" {
			this.Data["Message"] = "E-mail或用户名为空"
			return
		}

		// 如果两次输入的密码不一致，需重新填写
		if password != rePassword {
			this.Data["Message"] = "两次输入的密码不一致"
			return
		}

		b, err := valid.Valid(&user)
		if err != nil {
			this.Data["Message"] = "数据验证失败"
			return
		}
		if !b {
			// validation does not pass
			// blabla...
			this.Data["Message"] = "数据验证未通过"
			for _, err := range valid.Errors {
				fmt.Println(err.Key, err.Message)
			}
			return
		}
		// 检查E-mail或用户名是否已存在
		//orm = InitDb()
		//user := User{}
		//err = orm.Where("email=? or name=?", email, name).Find(&user)
		//if err == nil {
		//	this.Data["Message"] = "E-mail或用户名已存在"
		//	this.Data["Email"] = email
		//	this.Data["Name"] = name
		//	this.Data["Password"] = password
		//	this.TplNames = "admin/add_user.tpl"
		//	return
		//}

		//保存用户
		models.AddUser(&user)

		this.Ctx.Redirect(302, "/user/") // 返回用户列表页面
	case "edit": // 修改用户
		s := this.Ctx.Input.Param(":id") // 用户ID
		uid, err := strconv.Atoi(s)
		if err != nil {

		}
		user := models.User{}
		user.Id = uid
		user.Email = this.Input().Get("email")        // 用户E-mail
		user.Account = this.Input().Get("name")       // 用户名
		password := this.Input().Get("password")      // 密码
		rePassword := this.Input().Get("re-password") // 重复输入的密码

		this.TplNames = "authlogin/user_edit.html"
		this.Data["eUser"] = &user
		this.Data["Title"] = cnf.String("user_edit::title")
		if runActionMethodBefoer(&this.AdminBaseController, "Post", action) {
			return
		}
		// 检测E-mail或密码是否为空
		if user.Email == "" || user.Account == "" {
			this.Data["Message"] = "E-mail或用户名为空"
			return
		}

		// 如果两次输入的密码不一致，需重新填写
		if password != rePassword {
			this.Data["Message"] = "两次输入的密码不一致"
			return
		}

		// 获得当前用户

		// 更新用户信息
		if password != "" {
			user.Password = models.Sha1(password)
		}
		user.Updatedtime = time.Now()

		// 保存用户信息
		models.UpdateUserById(&user)

		this.Ctx.Redirect(302, "/user/") // 返回用户列表页面
	case "import": //导入用户列表
		f, _, err := this.GetFile("user_file")
		if err != nil {
			fmt.Println(err)
		}
		fd, err := ioutil.ReadAll(f)
		//fmt.Println(fd)
		u := []models.User{}
		//var u []User
		//	json.Unmarshal(fd, &u)
		if err := json.Unmarshal(fd, &u); err == nil {
			//处理导入用户列表
			addCount := 0
			for _, r := range u {
				_, err = models.AddUser(&r)
				if err != nil {
					fmt.Println(err)
				} else {
					addCount++
				}
			}
			fmt.Println("addcount:", addCount)
		}

	//	fmt.Println(handler.Filename)
	//fmt.Println(u)
	//	this.SaveToFile("user_file", "./static/files/"+"uploaded_file.txt")

	case "login": // 用户登录

		account := this.Ctx.Request.FormValue("account")   // 用户名
		password := this.Ctx.Request.FormValue("password") // 用户密码
		this.Data["Title"] = cnf.String("login::title")
		// 检测用户名或密码是否为空
		if account == "" || password == "" {
			this.Data["Message"] = "用户名或密码为空"
			this.Ctx.Redirect(302, "/user/login")
			return
		}
		user := AuthLogin(&this.AdminBaseController, account, password)
		if user == nil {
			this.Data["Message"] = "用户名或密码错误！"
			this.Ctx.Redirect(302, "/user/login")
			return
		} else {
			this.Data["Message"] = "登录成功！"
			// 保存用户登录信息
			setSessions(&this.Controller, "Uid", user.Id)

			flash := beego.NewFlash()
			flash.Success("成功登录到系统")
			flash.Store(&this.Controller)
			//跳转到前面操作页面
			Redirect_HttpReferer(&this.Controller)

		}

	}

}
