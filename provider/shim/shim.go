package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	p "github.com/thulasirajkomminar/terraform-provider-influxdb/v2/internal/provider"
)

func NewProvider(version string) provider.Provider {
	return p.New(version)()
}
