package edition

import (
	"golang.org/x/sys/windows/registry"
)

func GetOsSystem() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	eid, _, err := k.GetStringValue("EditionID")
	if err != nil {
		return "", err
	}

	return eid, nil
}
