package authlogin

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	"github.com/hoysoft/authlogin/models"
)

type AuthFilter struct {
	login bool
	admin bool
}

type LoginUser struct {
	*models.User
	IsAdmin     bool
	DisplayName string //显示名称
}

type AdminBaseController struct {
	beego.Controller
	LoginUser *LoginUser
}

type ActionMethodBefoerFunc func(c *AdminBaseController, method string, action string) (abort bool)

var (
	globalSessions         *session.Manager
	user_count             int
	actionMethodBefoerFunc ActionMethodBefoerFunc

	cnf config.ConfigContainer
	ops options

	BeforeAuthFilter map[string]*AuthFilter
)

func Before_auth(url string) *AuthFilter {
	a := AuthFilter{login: true, admin: false}
	BeforeAuthFilter[url] = &a
	return &a
}

//不登录验证
func (a *AuthFilter) UnLogin() *AuthFilter {
	a.login = false
	return a
}

//必须管理员
func (a *AuthFilter) MustAdmin() *AuthFilter {
	a.admin = true
	a.login = true
	return a
}

func (this *AdminBaseController) Prepare() {

	var IsContinue bool
	this.LoginUser, IsContinue = CheckLogin(&this.Controller)
	fmt.Println("LoginUser:", this.LoginUser)
	if this.LoginUser != nil {
		this.Data["LoginUser"] = &this.LoginUser
	}
	if !IsContinue {
		return
	}

	aFilter, ok := BeforeAuthFilter[this.Ctx.Request.RequestURI]
	if !ok {
		aFilter = &AuthFilter{login: true, admin: false}
	}

	if !aFilter.login {
		//不需要登录
		return
	}

	flash := beego.NewFlash()
	if this.LoginUser == nil {
		//必须是登录用户
		flash.Warning("请登录后进行操作")
		flash.Store(&this.Controller)
		this.SetSession("lastAdminPage", string(this.Ctx.Request.RequestURI))
		this.Redirect("/user/login", 302)
		return
	}

	if aFilter.admin && !this.LoginUser.IsAdmin {
		//必须登录为管理员
		flash.Warning("必须登录为管理员后进行操作")
		flash.Store(&this.Controller)
		this.SetSession("lastAdminPage", string(this.Ctx.Request.RequestURI))
		this.Redirect("/user/login", 302)
		return
	}
	beego.ReadFromRequest(&this.Controller)

}

func init() {
	//复制文件到项目
	cpFile("views", path.Join(beego.ViewsPath, "authlogin"))
	cpFile("conf", "conf")

	//读取authlogin配置
	cnf, _ = config.NewConfig("ini", "conf/authlogin.conf")
	ops = options{}
	ops.Read()

	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "booksmanage", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()

	BeforeAuthFilter = make(map[string]*AuthFilter)
}

func setSessions(this *beego.Controller, key interface{}, val interface{}) {
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	sess.Set(key, val)
}

func getSessions(this *beego.Controller, key interface{}) interface{} {
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	return sess.Get(key)
}

func ActionMethodBefoer(methodFunc ActionMethodBefoerFunc) {
	actionMethodBefoerFunc = methodFunc
}

func CheckLogin(this *beego.Controller) (l *LoginUser, IsContinue bool) {

	if user_count == 0 {
		user_count, _ := models.GetUserCount()
		//如果用户表为空，增加管理员用户界面
		if user_count == 0 {
			if this.Ctx.Request.RequestURI != "/admin/new" {
				this.SetSession("lastAdminPage", string(this.Ctx.Request.RequestURI))
				this.Redirect("/admin/new", 302)
				return nil, false
			} else {
				return nil, true
			}
		}
	}

	userid := getSessions(this, "Uid")

	if userid != nil {
		var e error
		LUser := LoginUser{}
		LUser.User, e = models.GetUserById(userid.(int))
		fmt.Println("UUU:", LUser.User)
		if LUser.User != nil && e == nil {
			switch cnf.DefaultInt("options:nameFromart", 0) {
			case 0: //"姓+名"
				LUser.DisplayName = LUser.User.LastName + LUser.User.FirstName
			case 1: // "名+姓"
				LUser.DisplayName = LUser.User.FirstName + LUser.User.LastName
			case 2: //"名"
				LUser.DisplayName = LUser.User.FirstName
			case 3: //"姓"
				LUser.DisplayName = LUser.User.LastName
			}
			if LUser.DisplayName == "" {
				LUser.DisplayName = LUser.User.Account
			}
			LUser.IsAdmin = true
			return &LUser, true //用户已经登录，返回用户信息
		}
	}

	return nil, true
}

//验证用户
func AuthLogin(this *AdminBaseController, account string, password string) *models.User {
	var user models.User
	var ldap models.LdapConnector
	var err error
	o := orm.NewOrm()

	if err = o.QueryTable("ldap_connector").Filter("Name", ops.authMode).One(&ldap); err == nil {
		err = o.QueryTable("user").Filter("Ldap", &ldap).Filter("Account", account).Filter("password", models.Sha1(password)).One(&user)
		if err != nil && &user != nil {
			// 更新用户最后登录时间信息.
			user.Lastlogintime = time.Now()
			models.UpdateUserById(&user)

			return &user
		}
	}

	return &user
}

func Redirect_HttpReferer(this *beego.Controller) {
	var Http_referer string
	h := this.GetSession("lastAdminPage")
	if h != nil {
		Http_referer = h.(string)
	}
	if Http_referer == "" {
		Http_referer = "/"
	}
	this.Redirect(Http_referer, 302)
}

func cpFile(spathName, dpath string) {
	_, filename, _, _ := runtime.Caller(1)
	filepath.Walk(path.Join(path.Dir(filename), spathName), func(spath string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		copyFile(dpath, spath, info.Name())
		return nil
	})
}

func copyFile(dstPath, srcName, filename string) (written int64, err error) {
	dstName := path.Join(dstPath, filename)
	//	srcName := path.Join(srcPath, filename)
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()
	//文件存在不覆盖
	if _, err = os.Stat(dstName); err == nil {
		return
	}

	e := os.MkdirAll(path.Dir(dstName), os.ModePerm)
	if e != nil {
		fmt.Println(e)
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}

	defer dst.Close()
	fmt.Println("Copy File:", src.Name(), "\n=>", dst.Name())
	return io.Copy(dst, src)
}

func runActionMethodBefoer(c *AdminBaseController, method string, action string) bool {
	if actionMethodBefoerFunc != nil {
		return actionMethodBefoerFunc(c, method, action)
	}
	return false
}
