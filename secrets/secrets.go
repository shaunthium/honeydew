package secrets

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// Assume that secrets are placed in the same "secrets" package
	secretsPath := filepath.Join(cwd, "secrets", "secrets.json")
	secrets, _ := ioutil.ReadFile(secretsPath)
	data := make(map[string]string)
	json.Unmarshal(secrets, &data)
	return data[key]
}
