package jsonx

import (
	"fmt"
	"strings"
	"testing"

	"arctron.lib/codec/json"
)

func Test_marshal(t *testing.T) {

}

type Info struct {
	RQ string `json:"RQ,omitempty"`
	SL int    `json:"SL,omitempty"`
}

func Test_marshal1(t *testing.T) {
	resp := make([]*Info, 2)
	for n := range resp {
		resp[n] = &Info{}
	}

	marshal1(resp)

	fmt.Println(json.MustMarshalString(resp))
	fmt.Println(resp, resp[0].RQ)

}

func marshal1(data interface{}) {
	s := `[
   "{\"RQ\":\"2022-06-03\",\"SL\":33}",
    "{\"RQ\":\"2022-05-02\",\"SL\":39}"
	]`

	var list []string
	json.Unmarshal([]byte(s), &list)

	for n := range list {
		json.Unmarshal([]byte(list[n]), data.([]*Info)[n])
	}

}

type EnergyBuildingRankData struct {
	Data []struct {
		Name   string `json:"name"`
		Energy int    `json:"energy"`
	}
}

func Test_marshal2(t *testing.T) {
	resp := &EnergyBuildingRankData{}
	marshal2(&resp.Data)
	fmt.Println(resp)
}

func marshal2(data interface{}) {
	s := `[
    {
        "name": "学生食堂1",
        "energy": 13
    },
	{
        "name": "学生食堂2",
        "energy": 12
    }
]
`
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_marshalArray(t *testing.T) {
	type T1 struct {
		Name   string  `json:"name,omitempty"`
		Energy float64 `json:"energy,omitempty"`
	}

	type T2 struct {
		Data []struct {
			Name   string  `json:"name,omitempty"`
			Energy float64 `json:"energy,omitempty"`
		}
	}

	arr := make([]*T1, 2)
	arr[0] = &T1{Name: "1", Energy: 11.0}
	arr[1] = &T1{Name: "2", Energy: 22.0}

	buf := json.MustMarshalString(arr)
	fmt.Println(string(buf))
}

func Test_marshalString(t *testing.T) {
	s := "ni shi shei? wo shi  ni ba ba"

	b := []byte(s)

	fmt.Println(json.MustMarshalString(b)) // x
	fmt.Println(json.MustMarshalString(s)) // y
	fmt.Println(string(b))                 // y

}

type Template struct {
	KK string `json:"kk,omitempty"`
	JJ string `json:"jj,omitempty"`
}

func Test_marshalToString(t *testing.T) {
	s := "ni shi shei? <kk> ni <jj>"

	tt := Template{
		KK: "wo shi",
		JJ: "ba ba",
	}

	var param []string
	var p, q int
	for n, v := range s {
		if v == '<' {
			p = n
		}
		if v == '>' {
			q = n
			param = append(param, s[p+1:q])
		}
	}

	fmt.Println(param)

	data := json.MustMarshalString(tt)
	fmt.Println(data)
	ref := map[string]string{}
	for _, v := range param {
		i := strings.Index(data, v)
		temp := data[i:]
		j := strings.Index(temp, ",")
		if j == -1 {
			j = len(temp) - 1
		}
		ref[v] = temp[5 : j-1]
	}

	fmt.Println(ref)

	var list []string
	for k, v := range ref {
		list = append(list, "<"+k+">")
		list = append(list, v)
	}

	fmt.Println(list)
	r := strings.NewReplacer(list...)
	s1 := r.Replace(s)

	fmt.Println("----> " + s1)
}

func Test_marshMap(t *testing.T) {
	m := map[string]string{
		"": "hahaha",
	}

	fmt.Println(json.MustMarshalString(m))

	delete(m, "")

	m["haha"] = "666"
	fmt.Println(json.MustMarshalString(m))
}
