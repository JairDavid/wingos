package option

import (
	"fmt"
	"log"
)
import "golang.org/x/sys/windows/registry"

func GetOsSystem() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	eid, _, err := k.GetStringValue("EditionID")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(eid, " : ", k)
}
