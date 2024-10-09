package pkg

import (
	"github.com/kouhenrui/GGCaptcha"
	"time"
)

type Captcha struct {
	captcha *GGCaptcha.GGCaptcha
	//driver inter.Driver
	//store  inter.Store
}

func NewCaptcha() *Captcha {
	driver := GGCaptcha.NewDriverString()
	localStore := GGCaptcha.NewLocalStore()
	//redisOption := GGCaptcha.RedisOptions{
	//	Host:     "192.168.245.22",
	//	Port:     "6379",
	//	Db:       4,
	//	PoolSize: 10,
	//	MaxRetry: 5,
	//}
	//redisStore := GGCaptcha.NewRediStore(redisOption)
	ggcaptcha := GGCaptcha.NewGGCaptcha(driver, localStore, time.Minute, 10*time.Minute, 50)
	return &Captcha{
		captcha: ggcaptcha,
		//driver: driver,
		//store:  redisStore,
	}
}
func (c *Captcha) GenerateCaptcha() (string, string, error) {
	//ggcaptcha := GGCaptcha.NewGGCaptcha(c.driver, c.store, time.Minute, 10*time.Minute, 50)
	return c.captcha.GenerateGGCaptcha()
}
func (c *Captcha) VerifyGGCaptcha(id, answer string, true bool) bool {
	return c.captcha.VerifyGGCaptcha(id, answer, true)
}
