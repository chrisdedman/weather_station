package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	pressure   interface{}
	humidity   interface{}
	city       interface{}
	visibility interface{}
	weather    interface{}
	wind       interface{}

	feels_like         float64
	sunrise            float64
	sunset             float64
	currentTemperature float64
}

// Retrieve and return the data from the interface
func retreiveData(data map[string]interface{}) Data {
	mainData, ok := data["main"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unable to access 'main' data")
		os.Exit(1)
	}
	windData, ok := data["wind"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unable to access 'wind' data")
		os.Exit(1)
	}
	weatherData, ok := data["weather"].([]interface{})[0].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unable to access 'weather' data")
		os.Exit(1)
	}
	sysData, ok := data["sys"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unable to access 'weather' data")
		os.Exit(1)
	}

	weatherInfo := Data{
		city:               data["name"],
		pressure:           mainData["pressure"],
		humidity:           mainData["humidity"],
		visibility:         data["visibility"],
		weather:            weatherData["description"],
		wind:               windData["speed"],
		currentTemperature: mainData["temp"].(float64),
		feels_like:         mainData["feels_like"].(float64),
		sunrise:            sysData["sunrise"].(float64),
		sunset:             sysData["sunset"].(float64),
	}
	return weatherInfo
}

// Print out the fetched data to the console
func printData(data Data) {
	fmt.Println()
	fmt.Println("\tWeather Report for", data.city)
	fmt.Println("----------------------------------------------------")
	fmt.Printf(" Sky Description         | %s\n", data.weather)
	fmt.Printf(" Current Temperature     | %.0f°F\n", convertToFahrenheit(data.currentTemperature))
	fmt.Printf(" Feels Like              | %.0f°F\n", convertToFahrenheit(data.feels_like))
	fmt.Println("----------------------------------------------------")
	fmt.Printf(" Visibility              | %v m\n", data.visibility)
	fmt.Printf(" Pressure                | %v hPa\n", data.pressure)
	fmt.Printf(" Humidity                | %v%%\n", data.humidity)
	fmt.Printf(" Wind Speed              | %.1f m/s\n", data.wind)
	fmt.Println("----------------------------------------------------")
	fmt.Printf(" Sunrise Time            | %s\n", formatTime(data.sunrise))
	fmt.Printf(" Sunset Time             | %s\n", formatTime(data.sunset))
	fmt.Println("----------------------------------------------------")
	fmt.Println()
}

// Convert the temperature from Kelvin to fahrenheit
func convertToFahrenheit(kelvin float64) float64 {
	fahrenheit := kelvin - 273.15
	return fahrenheit*9/5 + 32
}

// Format Unix timestamp to time in 24-hour format with timezone
func formatTime(timestamp float64) string {
	return time.Unix(int64(timestamp), 0).Format("15:04 MST")
}

// Error handler
func errorHandler(err error) {
	log.SetPrefix("Error: ")
	log.SetFlags(0)
	if err != nil {
		log.Fatal(err)
	}
}

// Fetch the weather data from the API Station,
// return a mapped interface with the data
func getAPIData(url string) map[string]interface{} {
	api, err := http.Get(url)
	errorHandler(err)

	defer api.Body.Close()
	body, err := ioutil.ReadAll(api.Body)
	errorHandler(err)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	errorHandler(err)

	return data
}

func main() {
	api_key := // ------- add your key API here ------- //

	fmt.Printf("Enter a location: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		city_name := scanner.Text()
		complete_url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%v&q=%v", api_key, city_name)
		data := getAPIData(complete_url)
		weatherInfo := retreiveData(data)

		printData(weatherInfo)
	} else {
		fmt.Println("Error scanning input:", scanner.Err())
	}
}
