package jsonx

import (
	"fmt"
	"testing"

	"arctron.lib/codec/json"
)

func Test_complex1(t *testing.T) {
	s := `{
    "code": 200,
    "msg": "ok",
    "data": {
        "faultPercent": "0.00",
        "fault": 0,
        "offlinePercent": "39.29",
        "offline": 2349,
        "device": {
            "": [
                {
                    "pid": "220607091309sgy9",
                    "type": "门禁管理系统-未分类",
                    "total": 358,
                    "fault": 0,
                    "offline": 180
                },
                {
                    "pid": "2206130914442bko",
                    "type": "公共照明-未分类",
                    "total": 147,
                    "fault": 0,
                    "offline": 55
                },
                {
                    "pid": "220610090949z06q",
                    "type": "三相四线制电力电表",
                    "total": 113,
                    "fault": 0,
                    "offline": 3
                },
                {
                    "pid": "220617105042oTf3",
                    "type": "数字式电表-未分类",
                    "total": 1659,
                    "fault": 0,
                    "offline": 0
                },
                {
                    "pid": "220606153732d3xV",
                    "type": "智慧教室报警机器",
                    "total": 98,
                    "fault": 0,
                    "offline": 0
                },
                {
                    "pid": "2206130915002rTm",
                    "type": "生活水泵控制柜",
                    "total": 10,
                    "fault": 0,
                    "offline": 6
                },
                {
                    "pid": "220705105123hk7Q",
                    "type": "空调系统-未分类",
                    "total": 688,
                    "fault": 0,
                    "offline": 688
                },
                {
                    "pid": "220616152719X0ED",
                    "type": "水表",
                    "total": 19,
                    "fault": 0,
                    "offline": 12
                },
                {
                    "pid": "220527154219LxpU",
                    "type": "枪式摄像头",
                    "total": 500,
                    "fault": 0,
                    "offline": 1
                },
                {
                    "pid": "220614090659Nw4O",
                    "type": "电梯系统-未分类",
                    "total": 11,
                    "fault": 0,
                    "offline": 7
                },
                {
                    "pid": "220609135644slqT",
                    "type": "快球摄像头",
                    "total": 2,
                    "fault": 0,
                    "offline": 0
                },
                {
                    "pid": "220607093404I0Jo",
                    "type": "半球摄像头",
                    "total": 935,
                    "fault": 0,
                    "offline": 1
                },
                {
                    "pid": "220613091428LPwS",
                    "type": "消防系统-未分类",
                    "total": 1413,
                    "fault": 0,
                    "offline": 1380
                },
                {
                    "pid": "220610091542rwJZ",
                    "type": "报警主机",
                    "total": 16,
                    "fault": 0,
                    "offline": 16
                },
                {
                    "pid": "220707101205XLBE",
                    "type": "环境监测设备",
                    "total": 9,
                    "fault": 0,
                    "offline": 0
                }
            ]
        }
    }
}`
	type MapDeviceData struct {
		Type  string `json:"type"`
		Total int    `json:"total"`
	}

	type DeviceData struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			MalfunctioningDevice string                     `json:"malfunctioning_device"`
			Malfunctioning       int                        `json:"malfunctioning"`
			OfflineDevice        string                     `json:"offline_device"`
			Offline              int                        `json:"offline"`
			DeviceMap            map[string][]MapDeviceData `json:"device"`
		} `json:"data"`
	}

	resp := DeviceData{}

	json.MustUnmarshal([]byte(s), &resp)

	fmt.Println(resp)
	var a, b int32
	for _, v := range resp.Data.DeviceMap[""] {
		switch v.Type {
		case "快球摄像头", "枪式摄像头", "半球摄像头":
			a += int32(v.Total)
		}
		b += int32(v.Total)
	}
	fmt.Println(a, b)
}
