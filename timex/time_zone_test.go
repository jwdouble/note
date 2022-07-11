package timex

import (
	"arctron.lib/sqlkit"
	"fmt"
	"testing"
	"time"
)

func Test_timeZone(t *testing.T) {
	c := time.Now()
	tz := sqlkit.TimestampNewLocal(c)

	fmt.Println(c, tz.AsTime())
	fmt.Println(c.Hour(), tz.AsTime().Hour())
}
