package util

import (
	"context"
	"fmt"
	"testing"

	"arctron.lib/ctxkit"
	"arctron.lib/token"
)

func Test_TokenMaker(t *testing.T) {
	//ctx := ctxkit.SYSTEMWithProject("1")
	fmt.Println("--> ", token.SYSTEM().GrantToTenant("jwtest")) // 只能本地测试用， 线上需要去网页里拿token
	//fmt.Println(token.Sign("31e4e0a89e8d4ef38a0eacbf99c55dac", "*", "SYSTEM", "", 0, map[string]interface{}{"privilege": true}))
}

func Test_ctxMaker(t *testing.T) {
	tid := "31e4e0a89e8d4ef38a0eacbf99c55dac"
	uid := "220394560432374782"
	tokenStr := token.Sign(tid, "*", uid, "", 0, map[string]interface{}{"privilege": true})

	c, _ := token.Parse(tokenStr)

	ctx := ctxkit.FromContext(c.WithContext(context.Background()), "token", c.Raw())
	token.MustFromContext(ctx)
	fmt.Println(c.GetTenantId(), c.GetUserId())
}
