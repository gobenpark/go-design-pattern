package pipeline

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type WeatherData struct {
	ID       string
	Country  CountryData   `json:"parent"`
	City     string        `json:"title"`
	Forecast []*WeatherDay `json:"consolidated_weather"`
	Error    error
}

type WeatherDay struct {
	Type    string  `json:"weather_state_name"`
	Date    string  `json:"applicable_date"`
	MinTemp float64 `json:"min_temp"`
	MaxTemp float64 `json:"max_temp"`
}

type CountryData struct {
	Name string `json:"title"`
}

const BASE_URL = "https://www.metaweather.com/api/location/%s/"

func getWeatherData(ids ...string) <-chan WeatherData {
	var wg sync.WaitGroup

	out := make(chan WeatherData)
	wg.Add(len(ids))

	for _, id := range ids {
		go func(cityId string) {
			var data WeatherData
			data.ID = cityId

			url := fmt.Sprintf(BASE_URL, cityId)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				data.Error = errors.New("Error making request")
				out <- data
				wg.Done()
				return
			}

			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				data.Error = errors.New("Error calling client")
				out <- data
				wg.Done()
				return
			}

			defer resp.Body.Close()

			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				data.Error = errors.New("error decoding message")
				out <- data
				wg.Done()
				return
			}
			out <- data
			wg.Done()
		}(id)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}

func convertWeatherData(weatherData <-chan WeatherData) <-chan WeatherData {
	out := make(chan WeatherData)

	go func() {
		for data := range weatherData {
			for _, day := range data.Forecast {
				day.MaxTemp = conversionCtoF(day.MaxTemp)
				day.MinTemp = conversionCtoF(day.MinTemp)
			}
			out <- data
		}

		fmt.Println("out close")
		close(out)
	}()

	fmt.Println("out !!!!")
	return out
}

func conversionCtoF(temp float64) float64 {
	return temp*1.8 + 32
}
