package secrets

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	// Setup
	NetAppBaseURL = "netapp_base_url"
	NetAppKey     = "netapp_key"

	// Deploy
	InstanceUsername = "instance_username"

	PemFilePath = "pem_file_path"
)

func GetValueFromSecrets(key string) string {
	secretsPath := os.Getenv("SECRETS_PATH")
	secrets, _ := ioutil.ReadFile(secretsPath)
	data := make(map[string]string)
	json.Unmarshal(secrets, &data)
	return data[key]
}
