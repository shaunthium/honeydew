package setup

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/shaunthium/honeydew/secrets"
)

const (
	// Job types
	NetappJobTypeClone = "clone"
)

func getUrl(jobType string) string {
	url := secrets.GetValueFromSecrets(secrets.NetappBaseURL)
	authKey := secrets.GetValueFromSecrets(secrets.NetappKey)
	return url + authKey + "/jobs/" + jobType
}

func Setup(newVolumeName string) error {
	fmt.Println("Cloning volume with new name: " + newVolumeName)

	data := []byte(`
    {"volume_clone_name": "` + newVolumeName + `"}
	`)
	req, err := http.NewRequest("POST", getUrl(NetappJobTypeClone), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Add("authorization", "Basic YWRtaW46UGFzc3dvcmRAMTIz")
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("NetApp Response Status:", resp.Status)
	fmt.Println("NetApp Response Headers:")
	for k, v := range resp.Header {
		fmt.Printf(k + ": ")
		fmt.Printf("%v\n", v)
	}
	return nil
}
