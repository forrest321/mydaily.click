package handlers

import (
	"net/http"
	"github.com/forrest321/mydaily.click/waterdata"
	"github.com/Sirupsen/logrus"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {

	testData, err := waterdata.GetWaterData()

	if err != nil {
		logrus.Infoln(err)
	}

	w.Write([]byte(testData))
}
