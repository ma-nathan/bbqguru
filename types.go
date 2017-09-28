package main

import (
	"encoding/xml"
	"time"
)

type Measurement struct {
	Reading int
	Last    time.Time
}

type BbqData struct {
	CookTarget   Measurement
	CookName     string
	PitTemp      Measurement
	Probe1Temp   Measurement
	Probe1Target Measurement
	Probe1Name   string
	Probe2Temp   Measurement
	Probe2Target Measurement
	Probe2Name   string
	Probe3Temp   Measurement
	Probe3Target Measurement
	Probe3Name   string
	FanPercent   Measurement
}

type XMLResult struct {
	XMLName     xml.Name `xml:"nutcallstatus"`
	CookName    string   `xml:"COOK>COOK_NAME"`
	PitTemp     string   `xml:"COOK>COOK_TEMP"`
	CookTarget  string   `xml:"COOK>COOK_SET"`
	Food1Name   string   `xml:"FOOD1>FOOD1_NAME"`
	Food1Temp   string   `xml:"FOOD1>FOOD1_TEMP"`
	Food1Target string   `xml:"FOOD1>FOOD1_SET"`
	Food2Name   string   `xml:"FOOD2>FOOD2_NAME"`
	Food2Temp   string   `xml:"FOOD2>FOOD2_TEMP"`
	Food2Target string   `xml:"FOOD2>FOOD2_SET"`
	Food3Name   string   `xml:"FOOD3>FOOD3_NAME"`
	Food3Temp   string   `xml:"FOOD3>FOOD3_TEMP"`
	Food3Target string   `xml:"FOOD3>FOOD3_SET"`
	FanPercent  string   `xml:"OUTPUT_PERCENT"`
}
