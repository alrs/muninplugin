package muninplugin

type Metric struct {
	Def *MetricDefinition
	Val float32 `munin:".value"`
}

func NewMetric() *Metric {
	return &Metric{
		Def: NewMetricDefinition(),
	}
}
