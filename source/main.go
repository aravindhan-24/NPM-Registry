package main

import (
	"npm-registry/cmd"
	_ "npm-registry/server"
)

func main() {
	cmd.Execute()
}
