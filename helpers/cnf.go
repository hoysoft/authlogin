package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/middleware"
)

type options struct {
	AuthMode    int //认证方式
	NameFromart int //姓名显示格式
}

var (
	Cnf  config.ConfigContainer
	Ops  options
	I18N *middleware.Translation
)

func init() {
	//读取authlogin配置
	Cnf, _ = config.NewConfig("ini", "conf/authlogin.conf")
	Ops = options{}
	Ops.Read()

	I18N = middleware.NewLocale("conf/i18n_authlogin.conf", beego.AppConfig.String("language"))
	//I18N.Translate("username", "vn")
}

func (this *options) Read() {
	Ops.AuthMode = Cnf.DefaultInt("options::authmode", 1)
	Ops.NameFromart = Cnf.DefaultInt("options::namefromart", 0)

}

func (this *options) Write() {
	Cnf.Set("options::authmode", strconv.Itoa(Ops.AuthMode))
	Cnf.Set("options::namefromart", strconv.Itoa(Ops.NameFromart))
	Cnf.SaveConfigFile("conf/authlogin.conf")
}
