package muninplugin

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	c.Graph = true
	t.Log(c)
}

func TestNewPlugin(t *testing.T) {
	p := NewPlugin()
	t.Log(p)
}

func TestNewMetric(t *testing.T) {
	m := NewMetric()
	t.Log(m)
}

func TestNewMetrics(t *testing.T) {
	m := NewMetrics()
	m["test"] = NewMetric()
	m["test"].Graph = false
	t.Log(m["test"])
	m["test"].Graph = true
	t.Log(m["test"])
}

func TestPrintConfig(t *testing.T) {
	c := NewConfig()
	c.GraphTitle = "Test Title"
	c.GraphHeight = 600
	c.GraphWidth = 800
	c.GraphVLabel = "Vertical"
	c.GraphCategory = "Filesystem"
	t.Log(c.Output())
}

func TestMetricsOutput(t *testing.T) {
	m := NewMetrics()
	m["test"] = NewMetric()
	m["test"].Graph = false
	t.Log(m.Output())
}
