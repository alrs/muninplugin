package muninplugin

import (
	"fmt"
	"reflect"
	"strings"
)

// Plugin is a struct that contains the global configuration for a Munin
// plugin, as well as a Metrics collection that contains all of the points
// of data tracked by the the plugin.
type Plugin struct {
	Metrics Metrics
	// Field descriptions from:
	// <http://guide.munin-monitoring.org/en/latest/reference/plugin.html>

	// Whether to draw the graph.
	Graph bool `munin:"graph"`

	// Arguments for the rrd grapher.
	// This is used to control how the generated graph looks,
	// and how values are interpreted or presented.
	GraphArgs interface{} `munin:"graph_args"`

	// Category used to sort the graph on the generated index web page.
	GraphCategory string `munin:"graph_category"`

	// The height of the graph.
	// Note that this is only the graph’s height and not the height of
	// the whole PNG image.
	GraphHeight interface{} `munin:"graph_height"`

	// Provides general information on what the graph shows.
	GraphInfo string `munin:"graph_info"`

	// Ensures that the listed fields are displayed in specified order.
	// Any additional fields are added in the order of appearance after
	// fields appearing on this list.
	graphOrder []string

	// Controls the time unit munin (actually rrd) uses to calculate
	// the average rates of change.
	GraphPeriod interface{} `munin:"graph_period"`

	// Controls the format munin (actually rrd) uses to display data
	// source values in the graph legend.
	GraphPrintf string `munin:"graph_printf"`

	// Per default the unit written on the graph will be scaled.
	// So instead of 1000 you will see 1k or 1M for 1000000.
	// You may disable autoscale by setting this to ‘no’.
	GraphScale interface{} `munin:"graph_scale"`

	// Sets the title of the graph.
	GraphTitle string `munin:"graph_title"`

	// If set, summarizes all the data sources’ values and reports the
	// results in an extra row in the legend beneath the graph.
	// The value you set here is used as label for that line.
	GraphTotal string `munin:"graph_total"`

	// Label for the vertical axis of the graph. Don’t forget to also
	// mention the unit.
	GraphVLabel string `munin:"graph_vlabel"`

	// The width of the graph. Note that this is only the graph’s
	// width and not the width of the whole PNG image.
	GraphWidth interface{} `munin:"graph_width"`

	// Override the host name for which the plugin is run.
	HostName string `munin:"host_name"`

	// Herewith the plugin tells that it delivers a hierarchy of
	// graphs. The attribute will show up multiple times in the config
	// section, once for each graph that it contains. It announces the
	// name of the graph for which the further configuration
	// attributes then follow.
	MultiGraph string `munin:"multi_graph"`

	// Decides whether munin-update should fetch data for the graph.
	Update bool `munin:"update"`

	// Sets the update_rate used by the Munin master when it creates
	// the RRD file. The update rate is the interval at which the RRD
	// file expects to have data.
	UpdateRate interface{} `munin:"update_rate"`
}

// NewPlugin instantiates a new Plugin struct, and sets some options
// to reasonable default values.
func NewPlugin() *Plugin {
	p := &Plugin{
		Metrics: newMetrics(),
	}
	p.Graph = true
	p.GraphScale = false
	p.Update = true
	return p
}

//func (p *Plugin) buildGraphOrderSlice() {
//	p.graphOrder = []string{}
//	for k, _ := range p.Metrics {
//		p.graphOrder = append(p.graphOrder, k)
//	}
//}

// MakeMetric adds a new metric to the plugin and maintains the
// order in which metrics were added.
func (p *Plugin) MakeMetric(m string) {
	p.graphOrder = append(p.graphOrder, m)
	p.Metrics[m] = newMetric()
}

// Config returns global configuration options for the plugin
// that are collected by the Munin server on its first run.
func (p *Plugin) Config() string {
	var result []string

	// Populate graphOrder slice by listing keys of Metrics in the
	// order they were added. Add the formatted string to the result
	// slice.
	if len(p.Metrics) > 0 {
		result = append(result,
			fmt.Sprintf("graph_order %s\n", strings.Join(p.graphOrder, " ")))
	}

	// Iterate through every member of the struct. Use the "munin" tag
	// to determine the field name expected by Munin. Use reflection to
	// add the type-dependent formatted line to the result slice.
	val := reflect.ValueOf(*p)
	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		kind := value.Kind()
		fieldType := val.Type().Field(i)
		tags := fieldType.Tag
		muninTag := tags.Get("munin")
		switch kind {
		case reflect.Interface:
			switch value.Interface().(type) {
			case int, int8, int32, int64:
				result = append(result,
					fmt.Sprintf("%s %d\n", muninTag, value.Interface()))
			case float32, float64:
				result = append(result,
					fmt.Sprintf("%s %.2f\n", muninTag, value.Interface()))
			}
		case reflect.String:
			if value.String() != "" {
				result = append(result,
					fmt.Sprintf("%s %s\n", muninTag, value.String()))
			}
		case reflect.Int:
			result = append(result,
				fmt.Sprintf("%s %d\n", muninTag, value.Int()))
		case reflect.Bool:
			result = append(result, fmt.Sprintf("%s %s\n", muninTag, toYN(value.Bool())))
		}
	}

	result = append(result, p.Metrics.Config())

	return strings.Join(result, "")
}
