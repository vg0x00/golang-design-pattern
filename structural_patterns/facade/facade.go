package facade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// mask outside API to these two acitons
type CurrentWeatherGetter interface {
	GetByCountryAndCityCode(countryCode, cityCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

// API return OBJ defined by remote API
type Weather struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float32 `json:"temp"`
		Pressure  float32 `json:"pressure"`
		Humidity  float32 `json:"humidity"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		SeaLevel  float32 `json:"sea_level"`
		GrntLevel float32 `json:"grnd_level"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All float32 `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Snow struct {
		ThreeHours float32 `json:"3h"`
	} `json:"snow"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		Id      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int32   `json:"sunrise"`
		Sunset  int32   `json:"sunset"`
	} `json:"sys"`
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

type CurrentWeatherData struct {
	APIkey string
}

func (c *CurrentWeatherData) GetByCountryAndCityCode(countryCode, city string) (*Weather, error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s", city, countryCode, c.APIkey))
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (*Weather, error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, c.APIkey))
}

// return a pointer to be more efficient
func (c *CurrentWeatherData) ResponseParser(r io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(r).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (c *CurrentWeatherData) doRequest(uri string) (*Weather, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	// request method is GET, so all content locate in URL
	// request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(response.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		return nil, fmt.Errorf("response status error: %d\n error message: %s\n",
			response.StatusCode, errMsg)
	}

	weather, err := c.ResponseParser(response.Body)
	response.Body.Close()
	return weather, err

}

// use mock data test for local services
func getMockData() io.Reader {
	data := `{"coord":
{"lon":145.77,"lat":-16.92},
"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],
"base":"cmc stations",
"main":{"temp":293.25,"pressure":1019,"humidity":83,"temp_min":289.82,"temp_max":295.37},
"wind":{"speed":5.1,"deg":150},
"clouds":{"all":75},
"rain":{"3h":3},
"dt":1435658272,
"sys":{"type":1,"id":8166,"message":0.0166,"country":"AU","sunrise":1435610796,"sunset":1435650870},
"id":2172797,
"name":"Cairns",
"cod":200}
	`
	r := bytes.NewReader([]byte(data))
	return r
}
