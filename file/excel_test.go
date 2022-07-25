package file

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"arctron.lib/timex"
	"github.com/xuri/excelize/v2"
)

func Test_readDataFormat(t *testing.T) {
	http.ListenAndServe(":8080", http.HandlerFunc(Listen))
}

func Listen(w http.ResponseWriter, req *http.Request) {
	mf, _, err := req.FormFile("file") //(r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	if err != nil {
		fmt.Println(err)
	}

	defer mf.Close()

	f, err := excelize.OpenReader(mf)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}

	ans := ""
	year := time.Now().Year()
	yprefix := strconv.Itoa(year)[:2]
	_, err = time.ParseInLocation(rows[0][0], timex.DateFormat, time.Local)
	{
		if err != nil {
			slc := strings.Split(rows[0][0], "-")
			if len(slc) == 3 {
				ans = fmt.Sprintf("%s%s-%s-%s", yprefix, slc[2], slc[0], slc[1])
			}
		}
	}

	fmt.Println(ans)
}
