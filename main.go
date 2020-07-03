package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
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

	getResultObj(resultObj)

}

func getResultObj(resultObj Response) {
	answer := "YES"
	for answer == "YES" {
		var searchResponse int

		fmt.Println("What do you want to search?")
		fmt.Printf("[1]ID [2]Country [3]View All\n->")
		fmt.Scanf("%d", &searchResponse)

		switch searchResponse {
		case 1:
			var searchResponseID int

			fmt.Print("What [ID] you are looking for?\n->")
			fmt.Scanf("%d", &searchResponseID)

			for i := 0; i < len(resultObj.Features); i++ {
				if searchResponseID == resultObj.Features[i].Attributes.OBJECTID {
					getResult(resultObj, i)
				}
			}

		case 2:
			var searchResponseCountry string
			var searchResponseYesNo string
			var searchResponseState string

			fmt.Print("What [Country] you are looking for? (ex.: Brazil)\n->")
			fmt.Scanf("%s", &searchResponseCountry)
			searchResponseCountry = strings.Title(searchResponseCountry)

			fmt.Print("You are looking for a specific state? (YES/NO)\n->")
			fmt.Scanf("%s", &searchResponseYesNo)
			searchResponseYesNo = strings.ToUpper(searchResponseYesNo)

			if searchResponseYesNo == "YES" {
				fmt.Print("What [State] you are looking for? (ex.: Ceara)\n->")
				fmt.Scanf("%s", &searchResponseState)
				searchResponseState = strings.Title(searchResponseState)

				for i := 0; i < len(resultObj.Features); i++ {
					if searchResponseState == resultObj.Features[i].Attributes.ProvinceState {
						getResult(resultObj, i)
					}
				}
			} else {
				for i := 0; i < len(resultObj.Features); i++ {
					if searchResponseCountry == resultObj.Features[i].Attributes.CountryRegion {
						getResult(resultObj, i)
					}
				}
			}

		case 3:
			var sumConfirmed int
			var sumRecovered int
			var sumDeaths int
			var sumActive int

			for i := 0; i < len(resultObj.Features); i++ {
				sumConfirmed = sumConfirmed + resultObj.Features[i].Attributes.Confirmed
				sumRecovered = sumRecovered + resultObj.Features[i].Attributes.Recovered
				sumDeaths = sumDeaths + resultObj.Features[i].Attributes.Deaths
				sumActive = sumActive + resultObj.Features[i].Attributes.Active
			}

			fmt.Printf("\n> World <\n\n> Confirmed: %d <\n", sumConfirmed)
			fmt.Printf("> Recovered: %d <\n", sumRecovered)
			fmt.Printf("> Deaths: %d <\n", sumDeaths)
			fmt.Printf("> Active: %d <\n\n", sumActive)
		}
		timeNow()
		fmt.Print("Do you want do search again? (YES/NO)\n->")
		fmt.Scanf("%s", &answer)
		fmt.Print("\n")
		answer = strings.ToUpper(answer)
	}
}

func getResult(resultObj Response, position int) {
	fmt.Printf("\n> ID: %d <\n", resultObj.Features[position].Attributes.OBJECTID)
	fmt.Printf("\n> Country: %s <\n", resultObj.Features[position].Attributes.CountryRegion)
	fmt.Printf("> State: %s <\n", resultObj.Features[position].Attributes.ProvinceState)
	fmt.Printf("> Confirmed: %d <\n", resultObj.Features[position].Attributes.Confirmed)
	fmt.Printf("> Recovered: %d <\n", resultObj.Features[position].Attributes.Recovered)
	fmt.Printf("> Deaths: %d <\n", resultObj.Features[position].Attributes.Deaths)
	fmt.Printf("> Active: %d <\n\n", resultObj.Features[position].Attributes.Active)
}

func timeNow() {
	const layout = "02/01/2006"
	timeNow := time.Now()

	fmt.Print(timeNow.Format(layout), "\n\n")
}
