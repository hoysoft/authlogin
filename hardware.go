//***************Hardware Monitor

//gopsutil 是go语言实现系统线控的开源项目，可以方便的监控服务器的cpu ，内存，硬盘等信息。json格式输出，用法简单非常实用。
//简单用法如下
//import (
//     "fmt"

//     "github.com/shirou/gopsutil"
//)

//func main() {
//     v, _ := gopsutil.VirtualMemory()

//     // almost every return value is struct
//     fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

//     // convert to JSON. String() is also implemented
//     fmt.Println(v)
//}
//*********************************************
package authlogin

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/shirou/gopsutil"
)

type HwController struct {
	AdminController
}

func init() {
	//增加路由
	beego.Router("/admin/hw", &HwController{})
	beego.Router("/admin/hw/:action", &HwController{})
	beego.Router("/admin/hw/:action/:id", &HwController{})

}

func (this *HwController) Get() {
	v, _ := gopsutil.VirtualMemory()
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Println(v)
}
