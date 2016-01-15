package mqtt

import (
	//"fmt"
	//MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"paybox-test/models"
	"encoding/json"
	//"os"
)


func sub(h string) (val string){

	b := models.TransJSON{
		BoxID : 1,
		CardID : 1,
		Cash : 100,
		Change : 20,
		JobID : 1,
		SiteID : 1,
		TimeStamp : "",
		Value : 80,
		VendorID : 1,
	}


	ret ,err :=  json.Marshal(b)
	val = string(ret)
	if err != nil {}
	return val
}

