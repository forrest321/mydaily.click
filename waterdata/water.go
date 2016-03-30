package waterdata

import (
	"github.com/gorilla/schema"
	"net/http"
	//"encoding/json"
	"io/ioutil"
)

func GetWaterData() (*WaterServiceResponse, error) {
	url := WebServiceTemplate

	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)



	riverData := new(WaterServiceResponse)
	decoder := schema.NewDecoder()

	err = decoder.Decode(riverData, body)
	if err != nil {
		panic(err)
	}
	return riverData, err
}
