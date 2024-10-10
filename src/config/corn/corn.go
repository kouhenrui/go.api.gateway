package corn

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

// CronTesk /**
var CronTesk *cron.Cron

func init() {
	//CronTesk = cron.New()

	//CronTesk.AddFunc("0 * * * * *", addCron1)
	////CronTesk.Start()
	//log.Println("定时任务初始化成功")
}
func addCron1() {
	//util.DtoToStruct(reqDto.RuleList{}, pojo.Rule{})
	fmt.Println("Task executed at", time.Now())

}
