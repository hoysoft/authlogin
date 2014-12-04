package authlogin

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/session"
	"github.com/hoysoft/authlogin/models"
	//"time"
)

type LoginUser struct {
	*models.User
	IsAdmin     bool
	DisplayName string //显示名称
}

type ActionMethodBefoerFunc func(c *beego.Controller, method string, action string) (abort bool)

var (
	globalSessions *session.Manager

	actionMethodBefoerFunc ActionMethodBefoerFunc

	cnf config.ConfigContainer
	ops options
)

type BaseController struct {
	beego.Controller
	user *LoginUser
}

func (this *BaseController) Prepare() {
	beego.ReadFromRequest(&this.Controller)
	this.user = CheckLogin(&this.Controller)
	if this.user == nil {
		if this.Ctx.Request.RequestURI != "/user/login" {
			this.Redirect("/user/login", 302)
		}
	}
	this.Data["User"] = &this.user
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
}

func ActionMethodBefoer(methodFunc ActionMethodBefoerFunc) {
	actionMethodBefoerFunc = methodFunc
}

func CheckLogin(this *beego.Controller) *LoginUser {
	LUser := LoginUser{}

	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userid := sess.Get("Uid")
	if userid != nil {
		var e error
		LUser.User, e = models.GetUserById(userid.(int))
		if e == nil {
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
			return &LUser //用户已经登录，返回用户信息
		}
	}

	//if this.Ctx.Request.RequestURI != "/user/login" { // 用户未登录，自动跳转
	//	this.Ctx.Redirect(302, "/user/login") // 跳转到用户登录页面
	//}
	return nil
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
