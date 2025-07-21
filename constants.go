package main

import "time"

const (
	// Server configuration
	ServerPort    = ":8080"
	StaticPath    = "./static/"
	
	// API configuration
	WeatherAPIURL = "https://wttr.in/%s?format=j1"
	APITimeout    = 10 * time.Second
	
	// Application constants
	MaxForecastDays = 3
	DefaultLocation = "London"
	
	// Error messages
	ErrEmptyLocation     = "Please enter a location"
	ErrFetchWeatherData  = "Unable to fetch weather data. Please check the location name and try again."
	ErrInvalidWeatherData = "Invalid weather data received"
	ErrTemplateExecution = "Error rendering template"
)
