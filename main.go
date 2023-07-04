package main
import (
	"strings"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type apiConfigData struct{
	OpenWeatherApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct{
	Name string `json:"Name"`
	Main struct{
		Kelvin float64 `json:"temp"`
	}`json:"main"`
}