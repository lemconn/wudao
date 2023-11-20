package main

import "github.com/alecthomas/kingpin/v2"

// program version
var version = "0.0.0"

func main() {
	app := kingpin.New("Wudao", "Wudao api gateway controller")

	app.Version(version)
}
