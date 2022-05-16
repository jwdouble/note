package sqlx

import (
	//"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
	//"github.com/jackc/pgx/v4"
)

func Test_pgInit(t *testing.T) {
	//c,err:=pgx.Connect(context.Background(),"postgres://postgres:p1ssw0rd@150.158.7.96:25432/logs")
	db, err := sql.Open("postgres", "host=150.158.7.96 user=postgres password=p1ssw0rd "+
		"dbname=logs port=25432 sslmode=disable TimeZone=Asia/Shanghai statement_cache_mode=describe")
	if err != nil {
		fmt.Println(err)
		return
	}

	//_,err = c.Exec(context.Background(),"insert into service (create_time, level, position, content) values($1, $2, $3, $4);",timex.Now(), 1, "11", "111")
	//if err!=nil{
	//	fmt.Println("-->",err)
	//}

	stmt, err := db.Prepare("insert into service (create_time, level, position, content) values($1, $2, $3, $4)")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = stmt.Exec(time.Now, 1, "11", "111")
	if err != nil {
		fmt.Println(err)
	}
}
