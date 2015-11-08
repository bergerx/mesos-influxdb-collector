package collector

import (
	"fmt"
	"github.com/kpacha/mesos-influxdb-collector/config"
)

func ExampleCollectorFromConfig() {
	txtConfig := `Master {
		host = "localhost"
		port = 5050
		leader = true
	}
	Master {
		host = "localhost"
		port = 15050
	}
	Slave {
		host = "localhost"
		port = 5051
	}
	Slave {
		host = "localhost"
		port = 5052
	}`
	cp := config.ConfigParser{}
	c, _ := cp.ParseConfig(txtConfig)
	collector := NewCollectorFromConfig(c)
	fmt.Println(collector)

	// Output:
	// {[{http://localhost:5050/metrics/snapshot {localhost true}} {http://localhost:15050/metrics/snapshot {localhost false}} {http://localhost:5051/metrics/snapshot {localhost}} {http://localhost:5052/metrics/snapshot {localhost}}]}
}