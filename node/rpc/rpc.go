package rpc

import (
	"go.uber.org/fx"

	"github.com/celestiaorg/celestia-node/rpc"
)

// Config represents the configuration for the RPC client
// used for retrieving information from Celestia Core nodes // TODO @renaynay: update this post-devnet
type Config struct {
	Protocol   string
	RemoteAddr string
}

// Components collects all the components and services related to the RPC client.
func Components(cfg *Config) fx.Option {
	return fx.Provide(func() (*rpc.Client, error) {
		return rpc.NewClient(cfg.Protocol, cfg.RemoteAddr)
	})
}