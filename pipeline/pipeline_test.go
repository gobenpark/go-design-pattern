package pipeline

import (
	"fmt"
	"testing"
)

func TestPipeLine(t *testing.T) {
	cities := []string{"44418", "2358820", "2471217", "2459115", "4118", "2372071", "615702", "968019", "727232", "650272"}

	fmt.Printf("Gathering weather information for %d cities\n", len(cities))
	fmt.Printf("\n#####################\n\n")
	data := convertWeatherData(getWeatherData(cities...))
	for data := range data {
		if data.Error != nil {
			fmt.Printf("Error fetching weather data for city id: %s\n", data.ID)
			continue
		}
		fmt.Printf("Weather Forcast for %s, %s\n", data.City, data.Country.Name)

		for _, day := range data.Forecast {
			fmt.Printf("\tDate: %s\n", day.Date)
			fmt.Printf("\t\t%s, High of %.2f℉, Low of %.2f℉\n\n", day.Type, day.MaxTemp, day.MinTemp)
		}
		fmt.Printf("\n#######################\n\n")
	}

	fmt.Println("Data fetch complete!")

}
