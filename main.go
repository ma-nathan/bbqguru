package main

import (
	"fmt"
)

func main() {

    fmt.Println("bbq-data-collector polls a BBQguru local network device.")

	var config = ReadConfig()
	influx_client := influxDBClient(config)

    go update_datastore(influx_client, config)
    watch_http_endpoint(config)
}

