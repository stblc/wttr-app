package main

import "strings"

// Weather icon SVGs (converted from JavaScript)
var weatherIcons = map[string]string{
	"clear": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<circle cx="50" cy="50" r="20" fill="#ffeaa7" stroke="#fdcb6e" stroke-width="2"/>
		<g stroke="#fdcb6e" stroke-width="3" stroke-linecap="round">
			<line x1="50" y1="10" x2="50" y2="20"/>
			<line x1="50" y1="80" x2="50" y2="90"/>
			<line x1="10" y1="50" x2="20" y2="50"/>
			<line x1="80" y1="50" x2="90" y2="50"/>
			<line x1="25.86" y1="25.86" x2="32.93" y2="32.93"/>
			<line x1="67.07" y1="67.07" x2="74.14" y2="74.14"/>
			<line x1="74.14" y1="25.86" x2="67.07" y2="32.93"/>
			<line x1="32.93" y1="67.07" x2="25.86" y2="74.14"/>
		</g>
	</svg>`,

	"partlyCloudy": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<circle cx="35" cy="35" r="12" fill="#ffeaa7" stroke="#fdcb6e" stroke-width="1.5"/>
		<g stroke="#fdcb6e" stroke-width="2" stroke-linecap="round">
			<line x1="35" y1="15" x2="35" y2="20"/>
			<line x1="15" y1="35" x2="20" y2="35"/>
			<line x1="23.64" y1="23.64" x2="27.07" y2="27.07"/>
			<line x1="42.93" y1="42.93" x2="46.36" y2="46.36"/>
		</g>
		<path d="M25 55 C15 55, 15 70, 25 70 L65 70 C75 70, 75 55, 65 55 C65 45, 45 45, 45 55 Z" 
			  fill="#b2bec3" stroke="#636e72" stroke-width="1.5"/>
	</svg>`,

	"cloudy": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<path d="M20 60 C10 60, 10 75, 20 75 L70 75 C80 75, 80 60, 70 60 C70 50, 50 50, 50 60 Z" 
			  fill="#b2bec3" stroke="#636e72" stroke-width="2"/>
		<path d="M15 45 C5 45, 5 60, 15 60 L60 60 C70 60, 70 45, 60 45 C60 35, 40 35, 40 45 Z" 
			  fill="#ddd" stroke="#636e72" stroke-width="1.5"/>
	</svg>`,

	"overcast": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<path d="M25 65 C15 65, 15 80, 25 80 L75 80 C85 80, 85 65, 75 65 C75 55, 55 55, 55 65 Z" 
			  fill="#95a5a6" stroke="#636e72" stroke-width="2"/>
		<path d="M20 50 C10 50, 10 65, 20 65 L65 65 C75 65, 75 50, 65 50 C65 40, 45 40, 45 50 Z" 
			  fill="#b2bec3" stroke="#636e72" stroke-width="1.5"/>
		<path d="M15 35 C5 35, 5 50, 15 50 L60 50 C70 50, 70 35, 60 35 C60 25, 40 25, 40 35 Z" 
			  fill="#ddd" stroke="#636e72" stroke-width="1"/>
	</svg>`,

	"rain": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<path d="M20 45 C10 45, 10 60, 20 60 L70 60 C80 60, 80 45, 70 45 C70 35, 50 35, 50 45 Z" 
			  fill="#95a5a6" stroke="#636e72" stroke-width="2"/>
		<g stroke="#74b9ff" stroke-width="2.5" stroke-linecap="round">
			<line x1="25" y1="65" x2="20" y2="80"/>
			<line x1="35" y1="70" x2="30" y2="85"/>
			<line x1="45" y1="65" x2="40" y2="80"/>
			<line x1="55" y1="70" x2="50" y2="85"/>
			<line x1="65" y1="65" x2="60" y2="80"/>
		</g>
	</svg>`,

	"heavyRain": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<path d="M20 45 C10 45, 10 60, 20 60 L70 60 C80 60, 80 45, 70 45 C70 35, 50 35, 50 45 Z" 
			  fill="#636e72" stroke="#2d3436" stroke-width="2"/>
		<g stroke="#0984e3" stroke-width="3" stroke-linecap="round">
			<line x1="20" y1="65" x2="15" y2="80"/>
			<line x1="25" y1="70" x2="20" y2="85"/>
			<line x1="30" y1="65" x2="25" y2="80"/>
			<line x1="35" y1="70" x2="30" y2="85"/>
			<line x1="40" y1="65" x2="35" y2="80"/>
			<line x1="45" y1="70" x2="40" y2="85"/>
			<line x1="50" y1="65" x2="45" y2="80"/>
			<line x1="55" y1="70" x2="50" y2="85"/>
			<line x1="60" y1="65" x2="55" y2="80"/>
			<line x1="65" y1="70" x2="60" y2="85"/>
			<line x1="70" y1="65" x2="65" y2="80"/>
		</g>
	</svg>`,

	"thunderstorm": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<path d="M20 45 C10 45, 10 60, 20 60 L70 60 C80 60, 80 45, 70 45 C70 35, 50 35, 50 45 Z" 
			  fill="#2d3436" stroke="#000" stroke-width="2"/>
		<g stroke="#0984e3" stroke-width="3" stroke-linecap="round">
			<line x1="25" y1="65" x2="20" y2="80"/>
			<line x1="35" y1="70" x2="30" y2="85"/>
			<line x1="45" y1="65" x2="40" y2="80"/>
			<line x1="65" y1="65" x2="60" y2="80"/>
		</g>
		<path d="M50 62 L45 72 L48 72 L44 82 L52 70 L49 70 L50 62 Z" 
			  fill="#f1c40f" stroke="#e17055" stroke-width="1"/>
	</svg>`,

	"snow": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<path d="M20 45 C10 45, 10 60, 20 60 L70 60 C80 60, 80 45, 70 45 C70 35, 50 35, 50 45 Z" 
			  fill="#b2bec3" stroke="#636e72" stroke-width="2"/>
		<g fill="white" stroke="#ddd" stroke-width="1">
			<circle cx="25" cy="70" r="3"/>
			<circle cx="35" cy="75" r="2.5"/>
			<circle cx="45" cy="70" r="3"/>
			<circle cx="55" cy="75" r="2.5"/>
			<circle cx="65" cy="70" r="3"/>
			<circle cx="30" cy="82" r="2"/>
			<circle cx="50" cy="83" r="2"/>
			<circle cx="60" cy="82" r="2"/>
		</g>
	</svg>`,

	"fog": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<g stroke="#b2bec3" stroke-width="3" stroke-linecap="round" opacity="0.7">
			<line x1="15" y1="45" x2="85" y2="45"/>
			<line x1="20" y1="55" x2="80" y2="55"/>
			<line x1="25" y1="65" x2="75" y2="65"/>
			<line x1="30" y1="75" x2="70" y2="75"/>
		</g>
		<g stroke="#ddd" stroke-width="2" stroke-linecap="round" opacity="0.5">
			<line x1="18" y1="40" x2="82" y2="40"/>
			<line x1="22" y1="50" x2="78" y2="50"/>
			<line x1="28" y1="60" x2="72" y2="60"/>
			<line x1="32" y1="70" x2="68" y2="70"/>
		</g>
	</svg>`,

	"wind": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<g stroke="#74b9ff" stroke-width="3" stroke-linecap="round" fill="none">
			<path d="M15 35 L65 35 C75 35, 75 45, 65 45 C55 45, 55 35, 65 35"/>
			<path d="M15 50 L55 50 C65 50, 65 60, 55 60 C45 60, 45 50, 55 50"/>
			<path d="M15 65 L70 65 C80 65, 80 75, 70 75 C60 75, 60 65, 70 65"/>
		</g>
		<g stroke="#0984e3" stroke-width="2" stroke-linecap="round" fill="none">
			<path d="M20 42 L60 42"/>
			<path d="M20 57 L50 57"/>
			<path d="M20 72 L65 72"/>
		</g>
	</svg>`,

	"default": `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
		<circle cx="50" cy="50" r="30" fill="#ddd" stroke="#636e72" stroke-width="2"/>
		<text x="50" y="55" text-anchor="middle" font-family="Arial" font-size="20" fill="#636e72">?</text>
	</svg>`,
}

// getWeatherIcon returns the appropriate SVG icon based on weather condition
func getWeatherIcon(condition string) string {
	lowerCondition := strings.ToLower(condition)

	if strings.Contains(lowerCondition, "clear") || strings.Contains(lowerCondition, "sunny") {
		return weatherIcons["clear"]
	} else if strings.Contains(lowerCondition, "partly") || strings.Contains(lowerCondition, "partial") {
		return weatherIcons["partlyCloudy"]
	} else if strings.Contains(lowerCondition, "overcast") {
		return weatherIcons["overcast"]
	} else if strings.Contains(lowerCondition, "cloudy") || strings.Contains(lowerCondition, "cloud") {
		return weatherIcons["cloudy"]
	} else if strings.Contains(lowerCondition, "thunder") || strings.Contains(lowerCondition, "storm") {
		return weatherIcons["thunderstorm"]
	} else if strings.Contains(lowerCondition, "heavy rain") || strings.Contains(lowerCondition, "downpour") {
		return weatherIcons["heavyRain"]
	} else if strings.Contains(lowerCondition, "rain") || strings.Contains(lowerCondition, "shower") || strings.Contains(lowerCondition, "drizzle") {
		return weatherIcons["rain"]
	} else if strings.Contains(lowerCondition, "snow") || strings.Contains(lowerCondition, "blizzard") {
		return weatherIcons["snow"]
	} else if strings.Contains(lowerCondition, "fog") || strings.Contains(lowerCondition, "mist") || strings.Contains(lowerCondition, "haze") {
		return weatherIcons["fog"]
	} else if strings.Contains(lowerCondition, "wind") {
		return weatherIcons["wind"]
	}

	return weatherIcons["default"]
}
