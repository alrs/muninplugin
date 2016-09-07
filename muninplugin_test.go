package muninplugin

import (
	"strings"
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
	m := newMetrics()
	switch interface{}(m).(type) {
	case Metrics:
		t.Logf("NewMetrics() created a Metrics: %v\n", m)
	default:
		t.Fatalf("NewMetrics() did not create a Metrics")
	}
}

func TestNewPluginDefaults(t *testing.T) {
	p := NewPlugin()
	if p.Graph == true {
		t.Log("Plugin Graph member correctly set to true by default.")
	} else {
		t.Fatalf("Plugin Graph member not set true by default.")
	}
	if p.GraphScale == false {
		t.Log("Plugin GraphScale correctly set to false by default.")
	} else {
		t.Fatalf("Plugin GraphScale not set false by default.")
	}
	if p.Update == true {
		t.Log("Plugin Update correctly set to true by default.")
	} else {
		t.Fatalf("Plugin Update not set true by default.")
	}

	expectedGraphWidth := 400
	expectedGraphHeight := 180
	if p.GraphWidth == expectedGraphWidth {
		t.Logf("Plugin GraphWidth correctly set to %d\n", p.GraphWidth)
	} else {
		t.Fatalf("Plugin GraphWidth should be %d, found %d\n",
			expectedGraphWidth, p.GraphWidth)
	}
	if p.GraphHeight == expectedGraphHeight {
		t.Logf("Plugin GraphHeight correctly set to %d\n", p.GraphHeight)
	} else {
		t.Fatalf("Plugin GraphHeight should be %d, found %d\n",
			expectedGraphHeight, p.GraphHeight)
	}
}

func TestPluginConfig(t *testing.T) {
	expectedDirectives := map[string]bool{
		"graph yes":                 false,
		"graph_category Filesystem": false,
		"graph_height 600":          false,
		"graph_period 0":            false,
		"graph_scale no":            false,
		"graph_title Test Title":    false,
		"graph_width 800":           false,
		"update yes":                false,
		"graph_vlabel Vertical":     false,
	}

	p := NewPlugin()
	p.GraphTitle = "Test Title"
	p.GraphHeight = 600
	p.GraphWidth = 800
	p.GraphVLabel = "Vertical"
	p.GraphCategory = "Filesystem"

	foundDirectives := strings.Split(p.Config(), "\n")
	t.Log(p.Config())
	for _, d := range foundDirectives {
		if len(d) > 0 {
			if _, ok := expectedDirectives[d]; ok {
				expectedDirectives[d] = true
			} else {
				t.Logf("Could not find: %s in Plugin Config output\n", d)
			}
		}
	}
	for dir, fnd := range expectedDirectives {
		if fnd == true {
			t.Logf("Found expected: %s\n", dir)
		} else {
			t.Fatalf("Found unexpected: %s\n", dir)
		}
	}
}

func TestMetricsConfig(t *testing.T) {
	expectedDirectives := map[string]bool{
		"test.label test":      false,
		"test.critical 190.00": false,
		"test.graph yes":       false,
		"test.max 200":         false,
		"test.min 0":           false,
		"test.warning 120":     false,
		"test.info extra":      false,
	}
	m := newMetrics()
	m["test"] = NewMetric()
	m["test"].Value = 100
	m["test"].Min = 0
	m["test"].Max = 200
	m["test"].Critical = 190.00
	m["test"].Warning = 120
	m["test"].Info = "extra"
	t.Log(m.Config())
	foundDirectives := strings.Split(m.Config(), "\n")

	for _, d := range foundDirectives {
		if len(d) > 0 {
			if _, ok := expectedDirectives[d]; ok {
				expectedDirectives[d] = true
			} else {
				t.Fatalf("Could not find: %s in Metrics output\n", d)
			}
		}
	}
	for dir, fnd := range expectedDirectives {
		if fnd == true {
			t.Logf("Found expected: %s\n", dir)
		} else {
			t.Fatalf("Not Found: %s\n", dir)
		}
	}

}

func TestMetricsValues(t *testing.T) {
	expectedDirectives := map[string]bool{
		"float.value 3.14": true,
		"int.value 3":      true,
		"nonumber.value U": true,
	}
	m := newMetrics()
	m["float"] = NewMetric()
	m["float"].Value = 3.14
	m["int"] = NewMetric()
	m["int"].Value = 3
	m["nonumber"] = NewMetric()
	m["nonumber"].Value = "this is not a number"

	foundDirectives := strings.Split(m.Values(), "\n")

	for _, d := range foundDirectives {
		if len(d) > 0 {
			if _, ok := expectedDirectives[d]; ok {
				expectedDirectives[d] = true
			} else {
				t.Fatalf("Could not find: %s in Metrics output\n", d)
			}
		}
	}
	for dir, fnd := range expectedDirectives {
		if fnd == true {
			t.Logf("Found expected: %s\n", dir)
		} else {
			t.Fatalf("Found unexpected: %s\n", dir)
		}
	}

}
