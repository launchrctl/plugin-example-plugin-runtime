// Package launchrctlexample implements a example launchr plugin
package launchrctlexample

import (
	"context"
	_ "embed"

	"github.com/launchrctl/keyring"

	"github.com/launchrctl/launchr"
	"github.com/launchrctl/launchr/pkg/action"
)

//go:embed action.yaml
var actionYaml []byte

func init() {
	launchr.RegisterPlugin(&Plugin{})
}

// Plugin is [launchr.Plugin] providing example action.
type Plugin struct {
	k keyring.Keyring
}

// PluginInfo implements [launchr.Plugin] interface.
func (p *Plugin) PluginInfo() launchr.PluginInfo {
	return launchr.PluginInfo{
		Weight: 10,
	}
}

// OnAppInit implements [launchr.OnAppInitPlugin] interface.
func (p *Plugin) OnAppInit(app launchr.App) error {
	app.GetService(&p.k)
	return nil
}

// DiscoverActions implements [launchr.ActionDiscoveryPlugin] interface.
func (p *Plugin) DiscoverActions(_ context.Context) ([]*action.Action, error) {
	a := action.NewFromYAML("example", actionYaml)
	a.SetRuntime(action.NewFnRuntime(func(_ context.Context, a *action.Action) error {
		input := a.Input()
		username := input.Opt("username").(string)
		password := input.Opt("password").(string)
		return example(username, password, p.k)
	}))
	return []*action.Action{a}, nil
}

func example(username, password string, k keyring.Keyring) error {
	// Code to execute goes here
	println("Hello world")
}
