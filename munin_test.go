package muninplugin

import (
	"testing"
)

func TestNewPlugin(t *testing.T) {
	p := NewPlugin()
	p.Graph = true
	t.Log(p)
}

func TestNewMetric(t *testing.T) {
	m := NewMetric()
	t.Log(m)
}

func TestNewMetrics(t *testing.T) {
	m := NewMetrics()
	m["test"] = NewMetric()
	m["test"].Def.Graph = false
	t.Log(m["test"])
	m["test"].Def.Graph = true
	t.Log(m["test"])
}

func TestNewMetricDefinition(t *testing.T) {
	m := NewMetricDefinition()
	t.Log(m)
}

func TestPrintConfig(t *testing.T) {
	p := NewPlugin()
	p.GraphTitle = "Test Title"
	p.GraphHeight = 600
	p.GraphWidth = 800
	p.GraphVLabel = "Vertical"
	p.GraphCategory = "Filesystem"
	t.Log(p.Config())
}

func TestMetricsOutput(t *testing.T) {
	m := NewMetrics()
	m["test"] = NewMetric()
	m["test"].Def.Graph = false
	t.Log(m.Config())
}
