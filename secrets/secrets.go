package secrets

import (
	"encoding/json"
	"io/ioutil"
)

const (
	// Setup
	NetAppBaseURL = "netapp_base_url"
	NetAppKey     = "netapp_key"

	// Deploy
	Instance1Hostname = "instance_1_hostname"
	InstanceUsername  = "instance_username"

	PemFilePath = "pem_file_path"

	// Hardcode path to secrets for now
	secretsPath = "/Users/Shaun/go/src/github.com/shaunthium/honeydew/secrets/secrets.json"
)

func GetValueFromSecrets(key string) string {
	secrets, _ := ioutil.ReadFile(secretsPath)
	data := make(map[string]string)
	json.Unmarshal(secrets, &data)
	return data[key]
}
