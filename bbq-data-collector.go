package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/influxdata/influxdb1-client/v2"
)

const (
	ASSUME_GONE    = -1 * time.Minute
	ENDPOINT_PAUSE = time.Second * 25
	HTTP_TIMEOUT   = time.Second * 8
	DATA_UPDATE    = time.Second * 30
	NOT_RECORDED   = 0
	NO_DATA        = "NO_DATA"
	// NOT_RECORDED     = -1
)

var bbq BbqData

func get_bbqguru_payload(config Config) (payload string, err error) {

	var resp *http.Response
	var req *http.Request
	var http_err error
	var data []byte
	var url = "http://" + config.BBQHost + "/all.xml"

	client := &http.Client{Timeout: HTTP_TIMEOUT}
	req, http_err = http.NewRequest("GET", url, nil)

	if http_err != nil {
		return "", http_err
	}

	req.SetBasicAuth(config.BBQUser, config.BBQPassword)
	resp, http_err = client.Do(req)

	if http_err != nil {
		return NO_DATA, http_err
	}

	defer resp.Body.Close()
	data, http_err = ioutil.ReadAll(resp.Body)

	payload = string(data)
	return
}

func watch_http_endpoint(config Config) {

	// Treat the BBQguru unit like a serial endpoint over which we have no control on
	// the sending side.  Keep polling to see what it currently has to say and update
	// our tracking to match.  Availability will come and go.

	for {

		payload, err := get_bbqguru_payload(config)

		if err != nil {

			fmt.Printf("Error fetching data from HTTP endpoint: %v\n", err)

		} else {

			// Send it over to get parsed

			parse_and_update(payload)
		}

		time.Sleep(ENDPOINT_PAUSE)
	}
}

func update_datastore(c client.Client, config Config) {

	// Every DATA_UPDATE interval:
	// 1. Check if our data is stale, zero it out if so
	// 2. Write what we have to datastore

	for {

		time.Sleep(DATA_UPDATE) // don't deliver first thing before we have data

		if bbq.CookTarget.Last.Before(time.Now().Add(ASSUME_GONE)) {
			bbq.CookTarget.Reading = NOT_RECORDED
		}

		if bbq.PitTemp.Last.Before(time.Now().Add(ASSUME_GONE)) {
			bbq.PitTemp.Reading = NOT_RECORDED
		}

		if bbq.Probe1Temp.Last.Before(time.Now().Add(ASSUME_GONE)) {
			bbq.Probe1Temp.Reading = NOT_RECORDED
		}

		if bbq.Probe2Temp.Last.Before(time.Now().Add(ASSUME_GONE)) {
			bbq.Probe2Temp.Reading = NOT_RECORDED
		}

		if bbq.Probe3Temp.Last.Before(time.Now().Add(ASSUME_GONE)) {
			bbq.Probe3Temp.Reading = NOT_RECORDED
		}

		if bbq.FanPercent.Last.Before(time.Now().Add(ASSUME_GONE)) {
			bbq.FanPercent.Reading = NOT_RECORDED
		}

		/*
			fmt.Printf("AirTempF: %d\n", pool.AirTempF.Reading)
			fmt.Printf("PoolTempF: %d\n", pool.PoolTempF.Reading)
			fmt.Printf("FilterSpeedRPM: %d\n", pool.FilterSpeedRPM.Reading)
			fmt.Printf("SaltPPM: %d\n", pool.SaltPPM.Reading)
			fmt.Printf("ChlorinatorPct: %d\n", pool.ChlorinatorPct.Reading)
			fmt.Printf("FilterOn: %d\n", pool.FilterOn.Reading)
			fmt.Printf("CleanerOn: %d\n", pool.CleanerOn.Reading)
			fmt.Printf("LightOn: %d\n", pool.LightOn.Reading)
		*/

		//deliver_stats_to_kairos()

		// Don't fill up the DB with null datapoints

		if	bbq.PitTemp.Reading == NOT_RECORDED &&
			bbq.Probe1Temp.Reading == NOT_RECORDED && 
			bbq.Probe2Temp.Reading == NOT_RECORDED &&
			bbq.Probe3Temp.Reading == NOT_RECORDED &&
			bbq.FanPercent.Reading == NOT_RECORDED {

			fmt.Println("Skip inserting all-null data row into DB.")

		} else {
			deliver_stats_to_influxdb(c, config)
		}
	}
}

