package main

import (
	"flag"
	"fmt"
)

func main() {
	config := parseCliArgs()
	fmt.Println(config)
}

//Config is the configuration of amq
type Config struct {
	Port             *int
	ProfilingEnabled *bool
	PromMetrics      *bool
}

func parseCliArgs() Config {
	conf := Config{}
	conf.Port = flag.Int("port", 3737, "Port number for amq to run")
	conf.ProfilingEnabled = flag.Bool("profile", false, "Flag to enable pprof profiling")
	conf.PromMetrics = flag.Bool("collectmetics", false, "Flag to serve metrics endpoint for prometheus")
	flag.Parse()
	return conf
}
