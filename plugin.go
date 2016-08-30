package muninplugin

import "strings"

type Plugin struct {
	Config  Config
	Metrics Metrics
}

func NewPlugin() *Plugin {
	p := &Plugin{
		Config:  *NewConfig(),
		Metrics: make(Metrics),
	}

	return p
}

func (p *Plugin) Output() string {
	var results []string
	results = append(results, p.Config.Output())
	results = append(results, p.Metrics.Output())
	return strings.Join(results, "")
}
