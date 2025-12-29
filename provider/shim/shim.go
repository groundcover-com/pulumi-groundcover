package shim

import (
	p "github.com/groundcover-com/terraform-provider-groundcover/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider(version string) provider.Provider {
	return p.New(version)()
}
