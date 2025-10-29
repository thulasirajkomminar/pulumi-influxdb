package main

import (
	influxdb "github.com/thulasirajkomminar/pulumi-influxdb/provider"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"
)

func main() {
	tfgen.Main("influxdb", influxdb.Provider())
}
