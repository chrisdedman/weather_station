package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
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

//Retrieve and return the data from the interface
func retreiveData(data map[string]interface{}) Data {
	mainData, ok := data["main"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unable to access 'main' data")
		os.Exit(1)
	}

	pressure := mainData["pressure"]
	humidity := mainData["humidity"]
	feels_like := mainData["feels_like"].(float64)
	currentTemperature := mainData["temp"].(float64)
	city := data["name"]
	visibility := data["visibility"]

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
		city:               city,
		pressure:           pressure,
		humidity:           humidity,
		visibility:         visibility,
		currentTemperature: currentTemperature,
		weather:            weatherData["description"],
		wind:               windData["speed"],
		feels_like:         feels_like,
		sunrise:            sysData["sunrise"].(float64),
		sunset:             sysData["sunset"].(float64),
	}
	return weatherInfo
}

// Print out to the console the data fetched
func printData(data Data) {
	fmt.Println()
	fmt.Printf("City                -> %v\n", data.city)
	fmt.Printf("Current Temperature -> %vF\n", math.Ceil((data.currentTemperature-273.15)*9/5+32))
	fmt.Printf("Feels Like          -> %vF\n", math.Ceil((data.feels_like-273.15)*9/5+32))
	fmt.Printf("Visibility          -> %vm\n", data.visibility)
	fmt.Printf("Pressure            -> %vhPa\n", data.pressure)
	fmt.Printf("Humidity            -> %v%%\n", data.humidity)
	fmt.Printf("Wind Speed          -> %vm/sec\n", data.wind)
	fmt.Printf("Sky Description     -> %v\n", data.weather)
	fmt.Printf("Sunrise             -> %v\n", time.Unix(int64(data.sunrise), 0).Format("15:04 MST"))
	fmt.Printf("Sunset              -> %v\n", time.Unix(int64(data.sunset), 0).Format("15:04 MST"))
	fmt.Println()
}

// Error handler
func errorHandler(err error) {
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
	log.SetPrefix("Error: ")
	log.SetFlags(0)
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
