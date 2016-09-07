package muninplugin

// Metric is a struct that holds all of the configuration data for a
// single metric tracked in a Munin plugin.
type Metric struct {
	Value interface{}

	// Field descriptions from
	// <http://guide.munin-monitoring.org/en/latest/reference/plugin.html>

	// A CDEF statement is a Reverse Polish Notation statement.
	// It can be used here to modify the value(s) before graphing.
	CDEF string `munin:".cdef"`

	// Custom specification of colour for drawing curve.
	Colour string `munin:".colour"`

	// Used by munin-limits to submit an error code indicating
	// critical state if the value fetched is above the maximum.
	Critical interface{} `munin:".critical"`

	// Determines how the data points are displayed in the graph.
	// The “LINE” takes an optional width suffix, commonly “LINE1”,
	// “LINE2”, etc…
	// The *STACK values are specific to munin and makes the first
	// a LINE, LINE[n] or AREA datasource, and the rest as STACK.
	Draw string `munin:".draw"`

	// Extended information that is included in alert messages (see
	// warning and critical) and HTML pages.
	ExtInfo string `munin:".extinfo"`

	// Determines if the data source should be visible in the
	// generated graph.
	Graph bool `munin:".graph"`

	// Explanation on the data source in this field. The Info is
	// displayed in the field description table on the detail web
	// page of the graph.
	Info string `munin:".info"`

	// Adds a horizontal line with the fieldname’s colour (HRULE) at
	// the value defined. Will not show if outside the graph’s scale.
	Line string `munin:".line"`

	// Sets a maximum value. If the fetched value is above “max”,
	// it will be discarded.
	Max interface{} `munin:".max"`

	// Sets a minimum value. If the fetched value is below “min”,
	// it will be discarded.
	Min interface{} `munin:".min"`

	// You need this for a “mirrored” graph. Values of the named
	// field will be drawn below the X-axis then (e.g. plugin if_
	// that shows traffic going in and out as mirrored graph).
	Negative string `munin:".negative"`

	// List of field declarations referencing the data sources
	// from other plugins by their virtual path.
	Stack string `munin:".stack"`

	// List of fields to summarize. If the fields are loaned from
	// other plugins they have to be referenced by their virtual path.
	Sum string `munin:".sum"`

	// Sets the RRD Data Source Type for this field. The values must
	// be written in capitals. The type used may introduce restrictions
	// for {fieldname.value}.
	Type string `munin:".type"`

	// Used by munin-limits to submit an
	// error code indicating warning state if the value fetched is
	// higher than the given number.
	Warning interface{} `munin:".warning"`
}

func newMetric() *Metric {
	return &Metric{
		Graph: true,
	}
}
