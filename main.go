package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Location struct {
	Ip           string  `json:"ip"`
	Country_code string  `json:"country_code"`
	Country_name string  `json:"country_name"`
	Region_code  string  `json:"region_code"`
	Region_name  string  `json:"region_name"`
	City         string  `json:"city"`
	Zip_code     string  `json:"zip_code"`
	Time_zone    string  `json:"time_zone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Metro_code   float64 `json:"metro_code"`
}

type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func getApiKey(id string) string {
	var val string
	if apikey, exists := os.LookupEnv(id); exists {
		val = apikey
	} else {
		val = ""
		log.Fatal("api key not found")
	}

	return val
}

func getLocation() Location {
	var usrLocationInfo Location
	apiKey := getApiKey("FREEGEOIP_API_KEY")
	url := fmt.Sprintf("https://api.freegeoip.app/json?apikey=%s", apiKey)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &usrLocationInfo); err != nil { // Parse []byte to the go struct pointer
		log.Fatal("unable to parse JSON")
	}

	fmt.Printf("usr location info: %v", usrLocationInfo)

	return usrLocationInfo

}

func getWeather() {
	apikey := getApiKey("OPEN_WEATHER_API_KEY")
	location := getLocation()
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", location.Latitude, location.Longitude, apikey)
	var weatherData Weather

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &weatherData); err != nil { // Parse []byte to the go struct pointer
		log.Fatal("unable to parse JSON")
	}

	fmt.Printf("weather data: %v", weatherData)
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	getWeather()
}
