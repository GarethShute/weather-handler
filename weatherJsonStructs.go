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
