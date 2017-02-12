package secrets

import (
	"encoding/json"
	"io/ioutil"
)

const (
	// Hardcode keypath for now
	keyPath = "/Users/Shaun/go/src/github.com/shaunthium/honeydew/secrets/secrets.json"
)

func GetValueFromSecrets(key string) string {
	secrets, _ := ioutil.ReadFile(keyPath)
	data := make(map[string]string)
	json.Unmarshal(secrets, &data)
	return data[key]
}
