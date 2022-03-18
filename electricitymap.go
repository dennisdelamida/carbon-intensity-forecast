package main

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"net/http"
)

type Params struct {
	Zone string

	Lon float32

	Lat float32
}

type Response struct {
	Zone string `json:"zone"`

	Forecast []Forecast `json:"forecast"`
}

type Forecast struct {
	CarbonIntensity int `json:"carbonIntensity"`

	DateTime string `json:"datetime"`
}

func GetForecastedCarbonIntensity(params Params) Response {

	fmt.Println("Calling GetForecastedCarbonIntensity API.....")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.electricitymap.org/v3/carbon-intensity/forecast", nil)

	if err != nil {

		fmt.Print(err.Error())

	}

	q := req.URL.Query()

	q.Add("zone", params.Zone)

	if params.Lon != 0 {

		q.Add("lon", fmt.Sprintf("%v", params.Lon))

	}

	if params.Lat != 0 {

		q.Add("lat", fmt.Sprintf("%v", params.Lat))

	}

	req.URL.RawQuery = q.Encode()

	req.Header.Add("auth-token", "3bhtgXSayVvgmuwEHry6zYYr")

	resp, err := client.Do(req)

	if err != nil {

		fmt.Print(err.Error())

	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		fmt.Print(err.Error())

	}

	var responseObject Response

	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Printf("API Response as struct %+v\n", responseObject)

	return responseObject

}

func main() {
	var params Params
	params.Zone = "US-CAL-CISO"
	GetForecastedCarbonIntensity(params)
	print()
}
