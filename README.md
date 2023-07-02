## Go Weather Station

This Go project allows you to retrieve weather data from the OpenWeatherMap API based on a user-provided city name. It fetches various weather information such as current temperature, pressure, humidity, visibility, wind speed, sunrise time, and sunset time. The program demonstrates how to make HTTP requests, handle JSON data, and work with time in a Go application.

## Features

- Retrieve weather data for a specific city by making requests to the OpenWeatherMap API.
- Display the fetched weather information in a formatted and readable manner.
- Supports retrieval of current temperature, pressure, humidity, visibility, wind speed, sunrise time, and sunset time.
- Demonstrates how to handle HTTP requests, JSON parsing, and time manipulation in Go.

## Installation

1. Clone the repository:

```shell
git clone https://github.com/your-username/weather-data-retrieval.git
```

2. Navigate to the project directory:
```shell
cd weather-data-retrieval
```
3. Build the Go program:
```shell
go build
```
4. Run the program:
```shell
./weather-data-retrieval
```

## Usage

1. Upon running the program, you will be prompted to enter a location (city name).
2. Enter the desired city name and press Enter.
3. The program will make an API request to retrieve the weather data for the specified city.
4. The fetched weather information, including current temperature, pressure, humidity, visibility, wind speed, sunrise time, and sunset time, will be displayed on the console.

### Station Report Example
```shell
Enter a location: quebec

        Weather Report for Québec
----------------------------------------------------
 Sky Description         | mist
 Current Temperature     | 78°F
 Feels Like              | 79°F
----------------------------------------------------
 Visibility              | 10000 m
 Pressure                | 1010 hPa
 Humidity                | 75%
 Wind Speed              | 2.1 m/s
----------------------------------------------------
 Sunrise Time            | 04:54 EDT
 Sunset Time             | 20:42 EDT
----------------------------------------------------
```
## Configuration

Before running the program, make sure you have an API key from OpenWeatherMap. Follow these steps to set up your API key:

1. Visit the OpenWeatherMap website: [openweathermap](https://openweathermap.org/).
2. Sign up for a free account and obtain your API key.
3. Open the main() function in the main.go file.
4. Locate the api_key variable declaration and assign your API key to it:
```go
api_key := "YOUR_API_KEY"
```
5. Save the change

## Dependencies

This project relies on the following Go packages:

* net/http - For making HTTP requests.
* encoding/json - For handling JSON data.
* fmt - For formatted output.
* os - For accessing command-line arguments and standard input/output.
* time - For working with time and date.
* Make sure you have Go installed on your system along with the necessary packages before running the program.

## Acknowledgments

* The project uses the OpenWeatherMap API to retrieve weather data. Visit their website for more information: [openweathermap.org](https://openweathermap.org/).
* This project was developed as a learning exercise for working with APIs, JSON parsing, and time manipulation in Go.
* Feel free to customize this README description according to your project's specific details and requirements.
