package conf

import (
	"fmt"
	"testing"

	"arctron.lib/conf"
)

func Test_conf(t *testing.T) {
	dsnDefault := "host=10.243.11.35 user=postgres password=p1ssw0rd dbname=arc-util port=5432 sslmode=disable TimeZone=Asia/Shanghai statement_cache_mode=describe"
	fmt.Println(conf.APP_PG_DSN.Value(dsnDefault))
}
