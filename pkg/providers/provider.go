package providers

import (
	"errors"

	"github.com/Jason-ZW/autok3s/pkg/providers/alibaba"
	"github.com/Jason-ZW/autok3s/pkg/types"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Provider interface {
	GetProviderName() string
	// Create command flags.
	GetCreateFlags(cmd *cobra.Command) *pflag.FlagSet
	// Join command flags.
	GetJoinFlags(cmd *cobra.Command) *pflag.FlagSet
	// Credential flags.
	GetCredentialFlags(cmd *cobra.Command) *pflag.FlagSet
	// Use this method to bind Viper, although it is somewhat repetitive.
	BindCredentialFlags() *pflag.FlagSet
	// K3s create cluster interface.
	CreateK3sCluster(ssh *types.SSH) error
	// K3s join node interface.
	JoinK3sNode(ssh *types.SSH) error
	// Rollback when error occurs.
	Rollback() error
}

func Register(provider string) (Provider, error) {
	var p Provider

	switch provider {
	case "alibaba":
		p = alibaba.NewProvider()
	default:
		return p, errors.New("not a valid provider, please run `autok3s get provider` display valid providers\n")
	}

	return p, nil
}

func SupportedProviders(provider string) [][]string {
	providers := [][]string{
		{"alibaba", "yes"},
	}
	if provider == "" {
		return providers
	}
	for _, ss := range providers {
		if ss[0] == provider {
			return [][]string{ss}
		}
	}

	return [][]string{}
}