package main

import (
	"fmt"
	"github.com/alrs/muninplugin"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const loadFile = "/proc/loadavg"

func load() (l float32, err error) {
	loadSlice, err := ioutil.ReadFile(loadFile)
	if err != nil {
		return
	}
	loadString := string(loadSlice)
	values := strings.Split(loadString, " ")
	floatLoad, err := strconv.ParseFloat(values[1], 32)
	return float32(floatLoad), err
}

func main() {
	p := muninplugin.NewPlugin()
	p.GraphTitle = "Example Golang Load"

	p.Metrics["load"] = muninplugin.NewMetric()
	systemLoad, err := load()
	if err != nil {
		log.Fatal(err)
	}

	p.Metrics["load"].Val = systemLoad
	p.Metrics["load"].Def.Critical = 25.0
	p.Metrics["load"].Def.Min = 0.0
	p.Metrics["load"].Def.Max = 500.00

	if len(os.Args) > 1 && os.Args[1] == "config" {
		fmt.Println(p.ConfigOutput())
		os.Exit(0)
	}

	fmt.Println(p.Metrics.Values())
}
