package authlogin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/session"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	//"time"
)

type ActionMethodBefoerFunc func(c *beego.Controller, method string, action string) (abort bool)

var (
	globalSessions *session.Manager

	actionMethodBefoerFunc ActionMethodBefoerFunc

	cnf config.ConfigContainer
	ops options
)

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
