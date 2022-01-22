package main

import "github.com/ramdrjn/serverbox/pkgs/statistics"

func main() {
	statistics.Initialize(true, "./internal/sample.conf")
}
