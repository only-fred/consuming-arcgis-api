package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Features []Features `json:"features"`
}

type Features struct {
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	OBJECTID      int    `json:"OBJECTID"`
	ProvinceState string `json:"Province_State"`
	CountryRegion string `json:"Country_Region"`
	Confirmed     int    `json:"Confirmed"`
	Recovered     int    `json:"Recovered"`
	Deaths        int    `json:"Deaths"`
	Active        int    `json:"Active"`
}

func main() {
	url := "https://services1.arcgis.com/0MSEUqKaxRlEPj5g/ArcGIS/rest/services/ncov_cases/FeatureServer/1/query?where=1%3D1&outFields=*&f=pjson"
	result, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	resultData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resultObj Response
	json.Unmarshal(resultData, &resultObj)

	for i := 0; i < len(resultObj.Features); i++ {
		fmt.Printf("\n> Country: %s\n", resultObj.Features[i].Attributes.CountryRegion)
		fmt.Printf("> State: %s\n", resultObj.Features[i].Attributes.ProvinceState)
		fmt.Printf("> Confirmed: %d\n", resultObj.Features[i].Attributes.Confirmed)
		fmt.Printf("> Recovered: %d\n", resultObj.Features[i].Attributes.Recovered)
		fmt.Printf("> Deaths: %d\n", resultObj.Features[i].Attributes.Deaths)
		fmt.Printf("> Active: %d\n", resultObj.Features[i].Attributes.Active)
	}
}
