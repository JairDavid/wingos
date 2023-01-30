package keys

import (
	"encoding/json"
	"os"
)

func Load() (map[string]string, error) {
	keysJson, _ := os.ReadFile("./config/keys/keys.json")
	keys := make(map[string]string)

	err := json.Unmarshal(keysJson, &keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
