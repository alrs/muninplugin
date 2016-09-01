package muninplugin

type Metric struct {
	Def *MetricDefinition
	Val float32 `munin:".value"`
}

// NewMetric returns a pointer to a Metric
func NewMetric() *Metric {
	return &Metric{
		Def: NewMetricDefinition(),
	}
}
