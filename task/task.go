package task

import (
	//"go-cms/models"
	"github.com/astaxie/beego/toolbox"
)

// 参考文档 https://beego.me/docs/module/toolbox.md#task
func StartTask() {
	// 每天0点执行
	//orderTask := toolbox.NewTask("orderTask", "0 0 17 * * *   ", models.CloseFailureOrder)
	//toolbox.AddTask("orderTask", orderTask)
	//toolbox.StartTask()

}

func StopTask() {
	toolbox.StopTask()
}
