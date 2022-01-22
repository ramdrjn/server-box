package main

import "github.com/ramdrjn/serverbox/pkgs/state"

func main() {
	state.Initialize(true, "./internal/sample.conf")
}
