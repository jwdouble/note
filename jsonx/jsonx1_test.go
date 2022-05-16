package jsonx

import (
	"fmt"
	"testing"

	"arctron.lib/codec/json"
)

type EnergyLongDayTrendData struct {
	Electricity         float64 `json:"electricity"`
	ElectricityFourteen []struct {
		TheTableWithReading float64 `json:"theTableWithReading"`
		RecordTime          string  `json:"recordTime"`
	} `json:"electricityFourteen"`
	Water         float64 `json:"water"`
	WaterFourteen []struct {
		TheTableWithReading float64 `json:"theTableWithReading"`
		RecordTime          string  `json:"recordTime"`
	} `json:"waterFourteen"`
}

func Test_unmarshalSA(t *testing.T){
	resp:=&EnergyLongDayTrendData{}

	err := json.Unmarshal([]byte(strArray),&resp)
	if err!=nil{
		fmt.Println(err)
	}


	fmt.Println(resp)
}
