package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/JairDavid/wingo/config/keys"
	"github.com/JairDavid/wingo/pkg/edition"
)

func main() {

	fmt.Println("Detecting windows version...")

	keys, err := keys.Load()
	if err != nil {
		panic("we could not load the keys")
	}

	edition, err := edition.GetOsSystem()
	if err != nil {
		panic(err)
	}

	if _, ok := keys[edition]; !ok {
		panic("we could not activate your version")
	}

	fmt.Println("Windows version detected...")
	serial, _ := keys[edition]

	exec.Command("slmgr /ipk", serial)
	exec.Command("slmgr /skms", "kms.msguides.com")
	exec.Command("slmgr /ato")

	fmt.Println("¡Windows activated! :D")
	time.Sleep(time.Second * 1)
}
