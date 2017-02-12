package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shaunthium/honeydew/deploy"
	"github.com/shaunthium/honeydew/setup"
	"github.com/shaunthium/honeydew/undeploy"
)

func getValueFromData(data map[string]string, key string) string {
	if v, ok := data[key]; !ok {
		return ""
	} else {
		return v
	}
}

func Setup(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	data := make(map[string]string)
	if err := json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if err := setup.Setup(getValueFromData(data, "volume_name")); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		panic(err)
	}
}

func Deploy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	data := make(map[string]string)
	if err := json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	volumeMountHostname := getValueFromData(data, "volume_mount_hostname")
	volumeName := getValueFromData(data, "volume_name")
	targetDirectory := getValueFromData(data, "target_directory")
	imageTagName := getValueFromData(data, "image_tag_name")
	instanceHostname := getValueFromData(data, "instance_hostname")

	if err := deploy.Deploy(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname, true); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		panic(err)
	}
}

func Undeploy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	data := make(map[string]string)
	if err := json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	targetDirectory := getValueFromData(data, "target_directory")
	instanceHostname := getValueFromData(data, "instance_hostname")

	if err := undeploy.Undeploy(targetDirectory, instanceHostname); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		panic(err)
	}
}
