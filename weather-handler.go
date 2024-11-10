package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var weatherConfig WeatherConfigData
var cities []CityInfo

var infoLogger *log.Logger
var warnLogger *log.Logger
var errorLogger *log.Logger

func main() {
	configureLogging()
	loadConfig()

	http.HandleFunc("/weather/dayforecast", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		servResponse, err := getDayForecast()
		if err != nil {
			warnLogger.Println("Cannot retrieve forecast:", err)
			http.Error(w, "Cannot retrieve forecast", http.StatusInternalServerError)
		}

		n, err := w.Write(servResponse)
		if err != nil {
			warnLogger.Println("Cannot write forecast:", err, " Code: ", n)
			http.Error(w, "Cannot write forecast output", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/citylookup/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.TrimPrefix(r.URL.Path, "/citylookup/")
		w.Header().Set("Content-Type", "application/json")
		servResponse, err := lookupCity(&city)
		if err != nil {
			warnLogger.Println(w, "Cannot retrieve city")
			http.Error(w, "Cannot retrieve city codes", http.StatusInternalServerError)
		}
		n, err := w.Write(servResponse)
		if err != nil {
			warnLogger.Println("Cannot write forecast:", err, " Code: ", n)
			http.Error(w, "Cannot write forecast output", http.StatusInternalServerError)
		}
	})

	fmt.Println("Running, waiting for request...")

	if weatherConfig.Port != "" {
		infoLogger.Println("Weather server listening on http://localhost:" + weatherConfig.Port)
		err := http.ListenAndServe("0.0.0.0:"+weatherConfig.Port, nil)
		if err != nil {
			errorLogger.Fatal("Could not start server: ", err)
		}
	} else {
		infoLogger.Println("Weather server listening on http://localhost:8080")
		err := http.ListenAndServe("0.0.0.0:8080", nil)
		if err != nil {
			errorLogger.Fatal("Could not start server: ", err)
		}
	}
}

func loadConfig() {
	args := os.Args
	var configPath string
	if len(args) < 2 {
		configPath = "config.json"
	} else {
		configPath = args[1]
	}

	jsonFile, err := os.Open(configPath)
	if err != nil {
		errorLogger.Fatal("Config load error: ", err)
	}
	jsonByteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(jsonByteValue, &weatherConfig)
	if err != nil {
		errorLogger.Fatal("Config load error: ", err)
	}
}

func configureLogging() {
	file, _ := os.Create("weatherApp.log")
	flags := log.Ldate | log.Lshortfile
	infoLogger = log.New(file, "INFO: ", flags)
	warnLogger = log.New(file, "WARN: ", flags)
	errorLogger = log.New(file, "ERROR: ", flags)
}

func getDayForecast() ([]byte, error) {

	infoLogger.Println("Requesting 5 day forecast")
	req, err := http.NewRequest("GET", "http://dataservice.accuweather.com/forecasts/v1/daily/5day/"+weatherConfig.Areacode, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("apikey", weatherConfig.Apikey)
	query.Add("language", "en-GB")
	query.Add("details", "true")
	query.Add("metric", "true")
	req.URL.RawQuery = query.Encode()

	body, err := sendHttpRequest(req)
	if err != nil {
		return nil, err
	}

	var forecast Forecast
	err = json.Unmarshal([]byte(body), &forecast)
	if err != nil {
		return nil, err
	}

	infoLogger.Println("Responding with 5 day weather forecast from date: ", forecast.DailyForecasts[0].Date)

	var returnDayWeather ReturnDayWeatherData
	for _, day := range forecast.DailyForecasts {
		var returnDayData DayData
		returnDayData.MaximumTemp = day.Temperature.Maximum.Value
		returnDayData.MinimumTemp = day.Temperature.Minimum.Value
		returnDayData.Unit = day.Temperature.Maximum.Unit
		returnDayData.IconRef = day.Day.Icon
		returnDayData.IconUrl = fmt.Sprintf("https://developer.accuweather.com/sites/default/files/%02d-s.png", day.Day.Icon)
		returnDayData.IconPhrase = day.Day.IconPhrase
		returnDayData.RainProbability = day.Day.RainProbability
		returnDayWeather.DailyForecasts = append(returnDayWeather.DailyForecasts, returnDayData)
	}
	jsonResponse, err := json.Marshal(returnDayWeather)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil

}

func lookupCity(city *string) ([]byte, error) {
	infoLogger.Println("Requesting city code")
	req, err := http.NewRequest("GET", "http://dataservice.accuweather.com/locations/v1/cities/search", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("apikey", weatherConfig.Apikey)
	query.Add("language", "en-GB")
	query.Add("details", "false")
	query.Add("q", *city)
	req.URL.RawQuery = query.Encode()

	body, err := sendHttpRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &cities)
	if err != nil {
		return nil, err
	}

	infoLogger.Println("Responding with city lookups with seach string: ", *city)

	var citySearchResults []CitySearchResults
	for _, city := range cities {
		var returnCity CitySearchResults
		returnCity.LocationKey = city.Key
		returnCity.Type = city.Type
		returnCity.Country = city.Country.EnglishName
		returnCity.Region = city.Region.EnglishName
		citySearchResults = append(citySearchResults, returnCity)
	}

	jsonResponse, err := json.Marshal(citySearchResults)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func sendHttpRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		errorLogger.Println("Error sending server request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorLogger.Println("Error reading server response:", err)
		return nil, err
	}

	return body, nil
}
