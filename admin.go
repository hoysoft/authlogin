package authlogin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	//"time"
)

var (
	//globalSessions *session.Manager

	//loginHtmlString        string
	//mLayout, LoginT        *template.Template
	//actionMethodBefoerFunc ActionMethodBefoerFunc
	//AuthViewPath           string //
	cnf config.ConfigContainer
)

func init() {
	//复制文件到项目
	cpFile("views", path.Join(beego.ViewsPath, "authlogin"))
	cpFile("conf", "conf")

	//读取authlogin配置
	cnf, _ = config.NewConfig("ini", "conf/authlogin.conf")
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
