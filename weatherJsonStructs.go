package main

type Forecast struct {
	DailyForecasts []DayForecast `json:"DailyForecasts"`
}

type DayForecast struct {
	Date        string      `json:"Date"`
	Temperature Temperature `json:"Temperature"`
	Day         Day         `json:"Day"`
}

type Temperature struct {
	Minimum Minimum `json:"Minimum"`
	Maximum Maximum `json:"Maximum"`
}

type Minimum struct {
	Value float32 `json:"value"`
	Unit  string  `json:"unit"`
}

type Maximum struct {
	Value float32 `json:"value"`
	Unit  string  `json:"unit"`
}

type Day struct {
	Icon            int    `json:"Icon"`
	IconPhrase      string `json:"IconPhrase"`
	RainProbability int    `json:"RainProbability"`
}

type CityInfo struct {
	Key     string  `json:"Key"`
	Type    string  `json:"Type"`
	Region  Region  `json:"Region"`
	Country Country `json:"Country"`
}

type Region struct {
	EnglishName string `json:"EnglishName"`
}

type Country struct {
	EnglishName string `json:"EnglishName"`
}

// Return structs for weather data
type DayData struct {
	MaximumTemp     float32 `json:"MaximumTemp"`
	MinimumTemp     float32 `json:"MinimumTemp"`
	Unit            string  `json:"Unit"`
	IconRef         int     `json:"IconRef"`
	IconUrl         string  `json:"IconUrl"`
	IconPhrase      string  `json:"IconPhrase"`
	RainProbability int     `json:"RainProbability"`
}

type ReturnDayWeatherData struct {
	DailyForecasts []DayData `json:"DailyForecasts"`
}

// Structs for config data
type WeatherConfigData struct {
	Port     string `json:"port"`
	Apikey   string `json:"apikey"`
	Areacode string `json:"areacode"`
}

// Return struct for City search results
type CitySearchResults struct {
	LocationKey string `json:"LocationKey"`
	Type        string `json:"Type"`
	Country     string `json:"Country"`
	Region      string `json:"Region"`
}
