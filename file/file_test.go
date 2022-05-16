package file

import (
	"context"
	"fmt"
	"testing"
	"time"

	"arctron.lib/conf"
	"arctron.lib/miniox"
)

func init() {
	miniox.Register(conf.APP_MINIO_DSN.Value("admin:p1ssw0rd@10.243.11.35:9000"))
	miniox.GetInstance().MakeBuckets(context.Background(), "oss-file")
}

func Test_upload(t *testing.T) {
	r, err := miniox.GetInstance().FPutObject(context.Background(), "oss-file",
		"bim/"+time.Now().Format("2006-01-02")+"/", "/home/jw/haha.sh", miniox.PutObjectOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}
