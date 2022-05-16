package rpc

import (
	"context"
	"fmt"
	"testing"

	"arctron.lib/conf"
	"arctron.lib/ctxkit"
	"arctron.lib/logx"
	"arctron.lib/rpcx"
	"arctron.os.api/spatial"
	"arctron.os.api/svc"
)

// fyyy  10.252.31.76:5882
// hjjt  10.10.95.115:7200   1038951302332417

func Test_grpc(t *testing.T) {
	cc, err := rpcx.Dial(conf.APP_API_ADDR_INTERNAL.Value("10.10.95.115:7200"), rpcx.WithInsecure())
	if err != nil {
		logx.Fatalf("dial client err: %v", err)
	}

	spatialCli := svc.NewArcSpatialServiceClient(cc)
	//ctxkit.SYSTEMWithProject()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIxMDM4OTUxMzAyODg5NDczIiwicHJqIjoiMTAzODk1MTMwMjMzMjQxNyIsInVpZCI6IjEwNTk3MDI3NDUzMDEwMDkiLCJqdGkiOiJUQ1MzWFdTOEpFQyIsImlhdCI6MTY0NDQ4Njg0NywidmVyIjoxfQ.7DMHo5XTMhEEmJoq84uNwhDq0yogvaUCjVh8mbEEMF4"
	ctx := context.Background()
	c := ctxkit.FromContext(ctx, "token", token)

	_, err = spatialCli.BatchDeletePoint(c, &spatial.BatchDeletePointReq{
		PointCodes: []string{
			"15-12.30.05.20-002.04F.005.101",
			"15-12.30.05.20-002.04F.005.102",
			"15-12.30.05.20-002.04F.005.103",
			"15-12.30.05.20-002.04F.005.104",
			"15-12.30.05.20-002.04F.005.105",
			"15-12.30.05.20-002.04F.005.106",
			"15-12.30.05.20-002.04F.005.107",
			"15-12.30.05.20-002.04F.005.108",
			"15-12.30.05.20-002.04F.005.109",
			"15-12.30.05.20-002.04F.005.110",
			"15-12.30.05.20-002.04F.005.111",
			"15-12.30.05.20-002.04F.005.112",
			"15-12.30.05.20-002.04F.005.113",
			"15-12.30.05.20-002.04F.005.114",
			"15-12.30.05.20-002.04F.005.115",
			"15-12.30.05.20-002.04F.005.116",
			"15-12.30.05.20-002.04F.005.117",
			"15-12.30.05.20-002.04F.005.118",
			"15-12.30.05.20-002.04F.005.119",
			"15-12.30.05.20-002.04F.005.120",
			"15-12.30.05.20-002.04F.005.121",
			"15-12.30.05.20-002.04F.005.122",
			"15-12.30.05.20-002.04F.005.123",
			"15-12.30.05.20-002.04F.005.124",
			"15-12.30.05.20-002.04F.005.125",
			"15-12.30.05.20-002.04F.005.126",
			"15-12.30.05.20-002.04F.005.127",
			"15-12.30.05.20-002.04F.005.128",
			"15-12.30.05.20-002.04F.005.129",
			"15-12.30.05.20-002.04F.005.130",
		},
		WithoutDevice: false,
	})

	fmt.Println(err)
}
