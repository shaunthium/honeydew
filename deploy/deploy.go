package deploy

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shaunthium/honeydew/secrets"
)

const (
	netappBaseURL = "netapp_base_url"
	netappKey     = "netapp_key"

	// Job types
	netappJobTypeClone = "clone"
)

func getUrl(jobType string) string {
	url := secrets.GetValueFromSecrets(netappBaseURL)
	authKey := secrets.GetValueFromSecrets(netappKey)
	return url + authKey + "/jobs/" + jobType
}

func Deploy() error {
	fmt.Println("in deploy")
	data := []byte(`
	  {"volume_clone_name": "hi_delin_and_jeremy"}
	`)
	req, err := http.NewRequest("POST", getUrl(netappJobTypeClone), bytes.NewBuffer(data))
	if err != nil {
		panic(err)
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
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
