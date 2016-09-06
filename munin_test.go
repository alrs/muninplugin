package muninplugin

import (
	"testing"
)

func TestNewPlugin(t *testing.T) {
	p := NewPlugin()
	switch interface{}(p).(type) {
	case *Plugin:
		t.Logf("NewPlugin() created a *Plugin: %v\n", p)
	default:
		t.Fatalf("NewPlugin() did not create a *Plugin")
	}
}

func TestNewMetric(t *testing.T) {
	m := NewMetric()
	switch interface{}(m).(type) {
	case *Metric:
		t.Logf("NewMetric() created a *Metric: %v\n", m)
	default:
		t.Fatalf("NewMetric() did not create a *Metric")
	}
}

func TestNewMetrics(t *testing.T) {
	m := NewMetrics()
	switch interface{}(m).(type) {
	case Metrics:
		t.Logf("NewMetrics() created a Metrics: %v\n", m)
	default:
		t.Fatalf("NewMetrics() did not create a Metrics")
	}
}

func TestNewMetricDefinition(t *testing.T) {
	md := NewMetricDefinition()
	switch interface{}(md).(type) {
	case *MetricDefinition:
		t.Logf("NewMetricDefinition() created a *MetricDefinition: %v\n", md)
	default:
		t.Fatalf("NewMetricDefinition() did not create a *MetricDefinition.")
	}
}

func TestNewMetricDefinitionDefaults(t *testing.T) {
	md := NewMetricDefinition()
	if md.Graph == true {
		t.Log("Graph member correctly set to 'true' by default.")
	} else {
		t.Fatal("Graph member does not default to 'true.'")
	}
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
	t.Log(m.Config())
}

func TestNonNumberValue(t *testing.T) {
	p := NewPlugin()
	p.Metrics["test"] = NewMetric()
	p.Metrics["test"].Val = "This isn't a number"
	values := p.Metrics.Values()
	if values != "test.value U" {
		t.Fatalf("Set metric value to a non-number, output should have been the letter U\n%s\n",
			values)
	} else {
		t.Log(values)
	}
}
