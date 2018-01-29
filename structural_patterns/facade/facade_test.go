package facade

import (
	"fmt"
	"testing"
)

const key = "63523b10e333ea8dbf6de13a73e767d2"

var weatherMap = &CurrentWeatherData{APIkey: key}

func TestMockDataParser(t *testing.T) {
	r := getMockData()

	weatherMap := CurrentWeatherData{} // ommit APIKey
	weather, err := weatherMap.ResponseParser(r)
	if err != nil {
		t.Errorf("weather map response parser error: %s\n", err.Error())
	} else if weather.Id != 2172797 {
		t.Fatalf("weather map response parser should return id: %d but not\ngot: %d\n",
			2172797, weather.Id)
	}
}

func TestCurrentDataWithCountryAndCityCode(t *testing.T) {
	weather, err := weatherMap.GetByCountryAndCityCode("London", "uk")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("temperature in London is %.2f\n", weather.Main.Temp-273.15)
}

func TestCurrentDataWithGeoCoordiante(t *testing.T) {
	weather, err := weatherMap.GetByGeoCoordinates(35.0, 139.0)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("temperature lat:35, lon: 139 : %.2f\n", weather.Main.Temp-273.15)
}
