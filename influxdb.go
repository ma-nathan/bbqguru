package main

import (
	//	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)


// CREATE USER admin WITH PASSWORD '$the_usual' WITH ALL PRIVILEGES
// create database BLAH

func influxDBClient(config Config) client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.DatabaseURL,
		Username: config.DatabaseUser,
		Password: config.DatabasePassword,
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return c
}

func influx_push_metrics(c client.Client, config Config) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.DatabaseDatabase,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	eventTime := time.Now()

	/*
		Using "Line Protocol", eg: cpu,host=server02,region=uswest value=3 1434055562000010000
		http://goinbigdata.com/working-with-influxdb-in-go/

		key: bbq
		tags: none
		fields: cook_target=blah, etc.
		timestamp in seconds
	*/

	key := "bbq"
	tags := map[string]string{
		"cook_name":    bbq.CookName,
		"probe_1_name": bbq.Probe1Name,
		"probe_2_name": bbq.Probe2Name,
		"probe_3_name": bbq.Probe3Name,
	}
	fields := map[string]interface{}{
		"cook_target": bbq.CookTarget.Reading,
	}

	point, err := client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"pit_temp": bbq.PitTemp.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"probe_1_temp": bbq.Probe1Temp.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"probe_2_temp": bbq.Probe2Temp.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"probe_3_temp": bbq.Probe3Temp.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"probe_1_target": bbq.Probe1Target.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"probe_2_target": bbq.Probe2Target.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"probe_3_target": bbq.Probe3Target.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	fields = map[string]interface{}{
		"fan_percent": bbq.FanPercent.Reading,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	bp.AddPoint(point)

	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
}

func deliver_stats_to_influxdb(c client.Client, config Config) {

	influx_push_metrics( c, config )
}
