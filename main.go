package main

import (
	"github.com/JairDavid/wingo/config/keys"
	"github.com/JairDavid/wingo/pkg/option"
)

func main() {

	keys.Load()
	option.GetOsSystem()
}
