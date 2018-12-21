package config

import "flag"

// FIXME bind to config
func init() {
	flag.String("master", "localhost", "master to connect to")
	flag.Parse()
}
