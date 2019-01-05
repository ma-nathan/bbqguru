package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// KDB_BASE_URL     = "http://bbq.iot.fumanchu.com:8080"
	KDB_BASE_URL     = "http://db:8080"
	KDB_HTTP_TIMEOUT = 10
)

func time_in_milliseconds() int64 {
	return time.Now().Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func kdb_add_json_element(in, in_name string, in_val int, tag string) (out string) {

	// This lead to weird graph tails instead of the gaps I was expecting - try again with kairos/grafana

	/*
		if in_val == NOT_RECORDED {
			out = in
			return
		}
	*/

	tags := fmt.Sprintf("\"tags\": {\"name\": \"%s\"}", tag)

	out = in + fmt.Sprintf("{\"name\": \"%s\",\"type\": \"long\",\"timestamp\": \"%d\", \"value\": %d, %s},",
		in_name, time_in_milliseconds(), in_val, tags)

	return
}

func http_call_kairos(json_str string) {

	var resp *http.Response
	var req *http.Request

	url := KDB_BASE_URL + "/api/v1/datapoints"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(json_str)))

	if err != nil {

		fmt.Printf("http_call_kairos: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{} // {Timeout: KDB_HTTP_TIMEOUT}
	resp, err = client.Do(req)

	if err != nil {

		fmt.Printf("http_call_kairos: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.Status != "200 " && resp.Status != "204 No Content" {

		fmt.Printf("Status: \"%s\"\n", resp.Status)
		fmt.Println("Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Body:", string(body))
		fmt.Printf("We sent: %s\n", json_str)
	}
}

func deliver_stats_to_kairos() {

	var json_str string

	json_str = "[" +
		kdb_add_json_element(json_str, "bbq.cook_target", bbq.CookTarget.Reading, bbq.CookName) +
		kdb_add_json_element(json_str, "bbq.pit_temp", bbq.PitTemp.Reading, bbq.CookName) +
		kdb_add_json_element(json_str, "bbq.probe1.temp", bbq.Probe1Temp.Reading, bbq.Probe1Name) +
		kdb_add_json_element(json_str, "bbq.probe2.temp", bbq.Probe2Temp.Reading, bbq.Probe2Name) +
		kdb_add_json_element(json_str, "bbq.probe3.temp", bbq.Probe3Temp.Reading, bbq.Probe3Name) +
		kdb_add_json_element(json_str, "bbq.probe1.target", bbq.Probe1Target.Reading, bbq.Probe1Name) +
		kdb_add_json_element(json_str, "bbq.probe2.target", bbq.Probe2Target.Reading, bbq.Probe2Name) +
		kdb_add_json_element(json_str, "bbq.probe3.target", bbq.Probe3Target.Reading, bbq.Probe3Name) +
		kdb_add_json_element(json_str, "bbq.fan_percent", bbq.FanPercent.Reading, bbq.CookName) +
		"]"

	json_str = strings.Replace(json_str, "},]", "}]", -1) // Final fixup

	// fmt.Printf("Payload: %s\n", json_str)

	http_call_kairos(json_str)
}
