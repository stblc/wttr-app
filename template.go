package main

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Weather Forecast</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Arial', sans-serif;
            background: linear-gradient(135deg, #74b9ff, #0984e3);
            min-height: 100vh;
            padding: 20px;
        }

        .container {
            max-width: 800px;
            margin: 0 auto;
            background: rgba(255, 255, 255, 0.95);
            border-radius: 20px;
            padding: 30px;
            box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
        }

        .header {
            text-align: center;
            margin-bottom: 30px;
        }

        .header h1 {
            color: #2d3436;
            font-size: 2.5rem;
            margin-bottom: 10px;
        }

        .location-input {
            margin-bottom: 20px;
            text-align: center;
        }

        .location-input input {
            padding: 12px 20px;
            font-size: 1rem;
            border: 2px solid #ddd;
            border-radius: 25px;
            width: 300px;
            margin-right: 10px;
            outline: none;
        }

        .location-input input:focus {
            border-color: #74b9ff;
        }

        .location-input button {
            padding: 12px 25px;
            font-size: 1rem;
            background: #74b9ff;
            color: white;
            border: none;
            border-radius: 25px;
            cursor: pointer;
            transition: background 0.3s;
        }

        .location-input button:hover {
            background: #0984e3;
        }

        .current-weather {
            text-align: center;
            margin-bottom: 30px;
            padding: 20px;
            background: rgba(116, 185, 255, 0.1);
            border-radius: 15px;
        }

        .weather-icon {
            width: 80px;
            height: 80px;
            margin: 0 auto 15px;
        }

        .temperature {
            font-size: 3rem;
            color: #2d3436;
            font-weight: bold;
            margin-bottom: 10px;
        }

        .description {
            font-size: 1.2rem;
            color: #636e72;
            margin-bottom: 15px;
        }

        .details {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 15px;
            margin-top: 20px;
        }

        .detail-item {
            background: white;
            padding: 15px;
            border-radius: 10px;
            text-align: center;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.08);
        }

        .detail-label {
            color: #636e72;
            font-size: 0.9rem;
            margin-bottom: 5px;
        }

        .detail-value {
            color: #2d3436;
            font-size: 1.1rem;
            font-weight: bold;
        }

        .forecast {
            margin-top: 30px;
        }

        .forecast h3 {
            color: #2d3436;
            margin-bottom: 20px;
            text-align: center;
        }

        .forecast-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
            gap: 15px;
        }

        .forecast-item {
            background: white;
            padding: 20px;
            border-radius: 10px;
            text-align: center;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.08);
        }

        .forecast-day {
            font-weight: bold;
            color: #2d3436;
            margin-bottom: 10px;
        }

        .forecast-icon {
            width: 40px;
            height: 40px;
            margin: 10px auto;
        }

        .forecast-temp {
            color: #2d3436;
            font-weight: bold;
        }

        .forecast-desc {
            font-size: 0.8rem;
            color: #636e72;
            margin-top: 5px;
        }

        .error {
            background: #ff6b6b;
            color: white;
            padding: 15px;
            border-radius: 10px;
            text-align: center;
            margin-bottom: 20px;
        }

        @media (max-width: 600px) {
            .location-input input {
                width: 200px;
                margin-bottom: 10px;
            }
            
            .temperature {
                font-size: 2rem;
            }
            
            .header h1 {
                font-size: 2rem;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🌤️ Weather Forecast</h1>
            <p>Get current weather and forecast for any location</p>
        </div>

        <form class="location-input" method="POST" action="/weather">
            <input type="text" name="location" placeholder="Enter city name (e.g., London, New York)" 
                   value="{{if not .HasData}}London{{end}}" required>
            <button type="submit">Get Weather</button>
        </form>

        {{if .Error}}
        <div class="error">{{.Error}}</div>
        {{end}}

        {{if .HasData}}
        <div class="current-weather">
            <div class="weather-icon">{{.WeatherIcon}}</div>
            <div class="temperature">{{.Temperature}}</div>
            <div class="description">{{.Description}}</div>
            <div style="font-size: 1.1rem; color: #636e72;">{{.Location}}</div>
            
            <div class="details">
                <div class="detail-item">
                    <div class="detail-label">Feels Like</div>
                    <div class="detail-value">{{.FeelsLike}}</div>
                </div>
                <div class="detail-item">
                    <div class="detail-label">Humidity</div>
                    <div class="detail-value">{{.Humidity}}</div>
                </div>
                <div class="detail-item">
                    <div class="detail-label">Wind</div>
                    <div class="detail-value">{{.Wind}}</div>
                </div>
                <div class="detail-item">
                    <div class="detail-label">Visibility</div>
                    <div class="detail-value">{{.Visibility}}</div>
                </div>
            </div>
        </div>

        <div class="forecast">
            <h3>3-Day Forecast</h3>
            <div class="forecast-grid">
                {{range .Forecast}}
                <div class="forecast-item">
                    <div class="forecast-day">{{.Day}}</div>
                    <div class="forecast-icon">{{.Icon}}</div>
                    <div class="forecast-temp">{{.Temperature}}</div>
                    <div class="forecast-desc">{{.Description}}</div>
                </div>
                {{end}}
            </div>
        </div>
        {{end}}
    </div>
</body>
</html>`
