package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/sys/windows/registry"
)

//OEMINFO Struct to store oem information
type OEMINFO struct {
	Logo         string
	Manufacturer string
	Model        string
	Hours        string
	Phone        string
	URL          string
}

func main() {
	regkey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\OEMInformation`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Println("reg")
		log.Fatal(err)
	}
	defer regkey.Close()

	logo, _, err := regkey.GetStringValue("Logo")
	manufacturer, _, err := regkey.GetStringValue("Manufacturer")
	model, _, err := regkey.GetStringValue("Model")
	hours, _, err := regkey.GetStringValue("SupportHours")
	phone, _, err := regkey.GetStringValue("SupportPhone")
	url, _, err := regkey.GetStringValue("SupportURL")

	OI := OEMINFO{logo, manufacturer, model, hours, phone, url}

	jsOutput, err := json.Marshal(OI)

	err = ioutil.WriteFile("out.json", jsOutput, 0644)
	fmt.Println(string(jsOutput))

}
