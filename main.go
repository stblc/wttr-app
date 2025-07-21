package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// WeatherData represents the structure from wttr.in API
type WeatherData struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
	NearestArea      []NearestArea      `json:"nearest_area"`
	Weather          []Weather          `json:"weather"`
}

type CurrentCondition struct {
	TempC       string `json:"temp_C"`
	TempF       string `json:"temp_F"`
	FeelsLikeC  string `json:"FeelsLikeC"`
	FeelsLikeF  string `json:"FeelsLikeF"`
	Humidity    string `json:"humidity"`
	WindspeedKmph string `json:"windspeedKmph"`
	Winddir16Point string `json:"winddir16Point"`
	Visibility  string `json:"visibility"`
	WeatherDesc []WeatherDesc `json:"weatherDesc"`
}

type WeatherDesc struct {
	Value string `json:"value"`
}

type NearestArea struct {
	AreaName []AreaName `json:"areaName"`
	Country  []Country  `json:"country"`
}

type AreaName struct {
	Value string `json:"value"`
}

type Country struct {
	Value string `json:"value"`
}

type Weather struct {
	Date     string   `json:"date"`
	MaxtempC string   `json:"maxtempC"`
	MintempC string   `json:"mintempC"`
	Hourly   []Hourly `json:"hourly"`
}

type Hourly struct {
	WeatherDesc []WeatherDesc `json:"weatherDesc"`
}

// Template data structure
type PageData struct {
	Location        string
	Temperature     string
	Description     string
	WeatherIcon     template.HTML
	FeelsLike       string
	Humidity        string
	Wind            string
	Visibility      string
	Forecast        []ForecastDay
	Error          string
	HasData        bool
}

type ForecastDay struct {
	Day         string
	Icon        template.HTML
	Temperature string
	Description string
}

// App holds the application configuration and dependencies
type App struct {
	tmpl   *template.Template
	client *http.Client
}

// NewApp creates a new application instance with proper configuration
func NewApp() (*App, error) {
	// Parse template once at startup
	tmpl, err := template.New("weather").Parse(htmlTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: APITimeout,
	}

	return &App{
		tmpl:   tmpl,
		client: client,
	}, nil
}

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	r := mux.NewRouter()
	
	// Serve static files (CSS, etc.)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(StaticPath))))
	
	// Routes
	r.HandleFunc("/", app.homeHandler).Methods("GET")
	r.HandleFunc("/weather", app.weatherHandler).Methods("POST")
	
	fmt.Printf("Server starting on %s...\n", ServerPort)
	log.Fatal(http.ListenAndServe(ServerPort, r))
}

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{HasData: false}
	app.renderTemplate(w, data)
}

func (app *App) weatherHandler(w http.ResponseWriter, r *http.Request) {
	location := strings.TrimSpace(r.FormValue("location"))
	if location == "" {
		data := PageData{Error: ErrEmptyLocation, HasData: false}
		app.renderTemplate(w, data)
		return
	}

	weatherData, err := app.fetchWeatherData(r.Context(), location)
	if err != nil {
		log.Printf("Error fetching weather data for %q: %v", location, err)
		data := PageData{Error: ErrFetchWeatherData, HasData: false}
		app.renderTemplate(w, data)
		return
	}

	data := app.processWeatherData(weatherData)
	app.renderTemplate(w, data)
}

func (app *App) fetchWeatherData(ctx context.Context, location string) (*WeatherData, error) {
	encodedLocation := url.QueryEscape(location)
	apiURL := fmt.Sprintf(WeatherAPIURL, encodedLocation)
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	resp, err := app.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return &data, nil
}

func (app *App) processWeatherData(data *WeatherData) PageData {
	if err := app.validateWeatherData(data); err != nil {
		log.Printf("Invalid weather data: %v", err)
		return PageData{Error: ErrInvalidWeatherData, HasData: false}
	}

	current := data.CurrentCondition[0]
	location := data.NearestArea[0]
	
	// Build location string
	locationName := fmt.Sprintf("%s, %s", location.AreaName[0].Value, location.Country[0].Value)
	
	// Convert temperatures
	tempC, _ := strconv.Atoi(current.TempC)
	tempF, _ := strconv.Atoi(current.TempF)
	temperature := fmt.Sprintf("%d°C (%d°F)", tempC, tempF)
	
	// Weather description
	description := ""
	if len(current.WeatherDesc) > 0 {
		description = current.WeatherDesc[0].Value
	}
	
	// Feels like temperature
	feelsLikeC := current.FeelsLikeC
	feelsLikeF := current.FeelsLikeF
	feelsLike := fmt.Sprintf("%s°C (%s°F)", feelsLikeC, feelsLikeF)
	
	// Wind information
	wind := fmt.Sprintf("%s km/h %s", current.WindspeedKmph, current.Winddir16Point)
	
	// Process forecast
	forecast := app.processForecast(data.Weather)

	return PageData{
		Location:     locationName,
		Temperature:  temperature,
		Description:  description,
		WeatherIcon:  template.HTML(getWeatherIcon(description)),
		FeelsLike:    feelsLike,
		Humidity:     current.Humidity + "%",
		Wind:         wind,
		Visibility:   current.Visibility + " km",
		Forecast:     forecast,
		HasData:      true,
	}
}

// validateWeatherData validates that the weather data has required fields
func (app *App) validateWeatherData(data *WeatherData) error {
	if len(data.CurrentCondition) == 0 {
		return fmt.Errorf("missing current condition data")
	}
	if len(data.NearestArea) == 0 {
		return fmt.Errorf("missing location data")
	}
	if len(data.Weather) == 0 {
		return fmt.Errorf("missing forecast data")
	}
	return nil
}

// processForecast processes the forecast data and returns up to MaxForecastDays
func (app *App) processForecast(weatherData []Weather) []ForecastDay {
	forecast := make([]ForecastDay, 0, MaxForecastDays)
	for i, day := range weatherData {
		if i >= MaxForecastDays {
			break
		}
		
		dayName := "Today"
		if i > 0 {
			date, err := time.Parse("2006-01-02", day.Date)
			if err == nil {
				dayName = date.Format("Mon")
			}
		}
		
		maxTemp, _ := strconv.Atoi(day.MaxtempC)
		minTemp, _ := strconv.Atoi(day.MintempC)
		temp := fmt.Sprintf("%d° / %d°", maxTemp, minTemp)
		
		// Get weather condition from hourly data (midday)
		condition := ""
		if len(day.Hourly) > 0 {
			midIndex := len(day.Hourly) / 2
			if len(day.Hourly[midIndex].WeatherDesc) > 0 {
				condition = day.Hourly[midIndex].WeatherDesc[0].Value
			}
		}
		
		forecast = append(forecast, ForecastDay{
			Day:         dayName,
			Icon:        template.HTML(getWeatherIcon(condition)),
			Temperature: temp,
			Description: condition,
		})
	}
	
	return forecast
}

func (app *App) renderTemplate(w http.ResponseWriter, data PageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	if err := app.tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, ErrTemplateExecution, http.StatusInternalServerError)
	}
}

