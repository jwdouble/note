package arcplus

import (
	"arctron.lib/conf"
	"arctron.lib/miniox"
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"
)

func Test_minioDownload(t *testing.T) {
	miniox.Register(conf.APP_MINIO_DSN.Value("admin:p1ssw0rd@10.243.11.35:9000"))

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", `attachment; filename="基站导入清册.xlsx"`)
	obj, err := miniox.GetInstance().PresignedGetObject(context.Background(), "oss-file", "arc-project/place-cards/基站导入清册.xlsx", time.Minute, reqParams)
	if err != nil {
		fmt.Println(err)
	}
	miniox.ReplaceHost(obj)
	fmt.Println(obj.String())
}
