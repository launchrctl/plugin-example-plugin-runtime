// Package example implements an example launchr plugin
package example

import (
	"context"
	_ "embed"

	"github.com/launchrctl/launchr"
	"github.com/launchrctl/launchr/pkg/action"
)

//go:embed action.yaml
var actionYaml []byte

func init() {
	launchr.RegisterPlugin(&Plugin{})
}

// Plugin is [launchr.Plugin] providing example action.
type Plugin struct{}

// PluginInfo implements [launchr.Plugin] interface.
func (p *Plugin) PluginInfo() launchr.PluginInfo {
	return launchr.PluginInfo{}
}

// DiscoverActions implements [launchr.ActionDiscoveryPlugin] interface.
func (p *Plugin) DiscoverActions(_ context.Context) ([]*action.Action, error) {
	a := action.NewFromYAML("example", actionYaml)
	a.SetRuntime(action.NewFnRuntime(func(_ context.Context, _ *action.Action) error {
		return example()
	}))
	return []*action.Action{a}, nil
}

func example() error {
	// Code to execute goes here
	launchr.Term().Println("Hello world")

	return nil
}
