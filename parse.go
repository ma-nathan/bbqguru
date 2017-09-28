package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

const ()

func parse_and_update(payload string) {

	v := XMLResult{}

	err := xml.Unmarshal([]byte(payload), &v)
	if err != nil {
		fmt.Printf("parse_and_update: %v", err)
		return
	}

	bbq.CookTarget.Reading = 0
	bbq.PitTemp.Reading = 0
	bbq.Probe1Temp.Reading = 0
	bbq.Probe1Target.Reading = 0
	bbq.Probe2Temp.Reading = 0
	bbq.Probe2Target.Reading = 0
	bbq.Probe3Temp.Reading = 0
	bbq.Probe3Target.Reading = 0
	bbq.FanPercent.Reading = 0

	// We don't map these directly as we may have to do some massaging of the data

	bbq.CookName = v.CookName
	bbq.Probe1Name = v.Food1Name
	bbq.Probe2Name = v.Food2Name
	bbq.Probe3Name = v.Food3Name

	if len(v.FanPercent) > 0 && v.FanPercent != "OPEN" {
		bbq.FanPercent.Reading, _ = strconv.Atoi(v.FanPercent)
		bbq.FanPercent.Last = time.Now()
	}

	if len(v.PitTemp) > 0 && v.PitTemp != "OPEN" {
		bbq.PitTemp.Reading, _ = strconv.Atoi(v.PitTemp)
		bbq.PitTemp.Reading /= 10
		bbq.PitTemp.Last = time.Now()
	}

	if len(v.CookTarget) > 0 && v.CookTarget != "OPEN" {
		bbq.CookTarget.Reading, _ = strconv.Atoi(v.CookTarget)
		bbq.CookTarget.Reading /= 10
		bbq.CookTarget.Last = time.Now()
	}

	if len(v.Food1Temp) > 0 && v.Food1Temp != "OPEN" {
		bbq.Probe1Temp.Reading, _ = strconv.Atoi(v.Food1Temp)
		bbq.Probe1Temp.Reading /= 10
		bbq.Probe1Temp.Last = time.Now()
	}

	if len(v.Food1Target) > 0 && v.Food1Target != "OPEN" {
		bbq.Probe1Target.Reading, _ = strconv.Atoi(v.Food1Target)
		bbq.Probe1Target.Reading /= 10
		bbq.Probe1Target.Last = time.Now()
	}

	if len(v.Food2Temp) > 0 && v.Food2Temp != "OPEN" {
		bbq.Probe2Temp.Reading, _ = strconv.Atoi(v.Food2Temp)
		bbq.Probe2Temp.Reading /= 10
		bbq.Probe2Temp.Last = time.Now()
	}

	if len(v.Food2Target) > 0 && v.Food2Target != "OPEN" {
		bbq.Probe2Target.Reading, _ = strconv.Atoi(v.Food2Target)
		bbq.Probe2Target.Reading /= 10
		bbq.Probe2Target.Last = time.Now()
	}

	if len(v.Food3Temp) > 0 && v.Food3Temp != "OPEN" {
		bbq.Probe3Temp.Reading, _ = strconv.Atoi(v.Food3Temp)
		bbq.Probe3Temp.Reading /= 10
		bbq.Probe3Temp.Last = time.Now()
	}

	if len(v.Food3Target) > 0 && v.Food3Target != "OPEN" {
		bbq.Probe3Target.Reading, _ = strconv.Atoi(v.Food3Target)
		bbq.Probe3Target.Reading /= 10
		bbq.Probe3Target.Last = time.Now()
	}

	// fmt.Printf("Data now: %v\n", bbq)
}
