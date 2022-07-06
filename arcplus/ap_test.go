package arcplus

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"arctron.lib/conf"
)

var (
	cf               = conf.GetEnv("app.msg.sm.url", "http://10.243.11.35:8888/arcplus/sms/sendByTemplate")
	cacheDuration, _ = time.ParseDuration(conf.GetEnv("app.msg.cache.duration", "30s"))
	restyDebug, _    = strconv.ParseBool(conf.GetEnv("app.resty.debug"))
)

func Test_env(t *testing.T) {
	_ = cacheDuration
	_ = restyDebug
	r := cf
	fmt.Println(r)
}
