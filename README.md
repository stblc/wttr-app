# Weather Forecast App - Go Version

A simple weather forecast web application built with Go that displays current weather conditions and a 3-day forecast for any location.

## Features

- 🌤️ Current weather conditions (temperature, description, feels like, humidity, wind, visibility)
- 📅 3-day weather forecast
- 🎨 Clean, responsive web interface
- ⚡ Server-side rendering with Go templates
- 🔌 No client-side JavaScript required

## Technology Stack

- **Backend**: Go 1.21 with Gorilla Mux router
- **Frontend**: HTML templates with embedded CSS
- **API**: wttr.in weather service
- **Icons**: Inline SVG weather icons

## Project Structure

```
.
├── main.go          # Main HTTP server and request handlers
├── icons.go         # Weather icon SVG definitions and mapping logic
├── go.mod           # Go module dependencies
├── go.sum           # Go dependency checksums (auto-generated)
└── README.md        # This documentation
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Internet connection (for weather API calls)

### Installation & Running

1. Clone or navigate to the project directory
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run .
   ```
4. Open your browser and visit: `http://localhost:8080`

### Usage

1. Enter a city name (e.g., "London", "New York", "Tokyo")
2. Click "Get Weather" or press Enter
3. View current weather conditions and 3-day forecast

## API Endpoints

- `GET /` - Home page with weather form
- `POST /weather` - Submit location and get weather data

## Key Changes from JavaScript Version

### What was simplified:

✅ **Eliminated all client-side JavaScript** - No more complex DOM manipulation or AJAX calls
✅ **Server-side data processing** - Weather API calls now handled by Go backend  
✅ **HTML form submission** - Simple POST request instead of JavaScript event handlers
✅ **Template-based rendering** - Server generates complete HTML with data
✅ **Embedded SVG icons** - Icons served directly from Go instead of separate JS file

### Architecture improvements:

- **Single-page application** → **Traditional web app** with form submissions
- **Client-side API calls** → **Server-side API integration**
- **3 separate files** (HTML/CSS/JS) → **2 Go files** with embedded templates
- **Runtime DOM updates** → **Server-side template rendering**

## Error Handling

The application handles various error scenarios:
- Invalid location names
- API service unavailable
- Network connectivity issues
- Malformed API responses

## Performance & Reliability

- Server-side caching opportunities (can be added)
- Reduced client-side complexity
- Better SEO compatibility
- Works without JavaScript enabled
- Lower bandwidth usage (no separate JS/CSS files)

## Development

To modify the application:

- **UI/Styling**: Edit the HTML template in `main.go` (renderTemplate function)
- **Weather Logic**: Modify handlers in `main.go`
- **Icons**: Update SVG definitions in `icons.go`
- **Dependencies**: Modify `go.mod`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Copyright Notice

```
MIT License

Copyright (c) 2025 [Your Name]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
