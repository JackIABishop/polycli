package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// --- types and helpers (exported) ---

// MainData holds temperature info.
type MainData struct {
	Temp float64 `json:"temp"`
}

// WeatherElement is one entry in the "weather" array.
type WeatherElement struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

// SysData holds system info, like the country code.
type SysData struct {
	Country string `json:"country"`
}

// WeatherResp is the trimmed JSON we care about.
type WeatherResp struct {
	Name    string           `json:"name"`
	Main    MainData         `json:"main"`
	Weather []WeatherElement `json:"weather"`
	Sys     SysData          `json:"sys"`
}

// Fetch contacts OpenWeatherMap and returns a nice string.
func Fetch(city string) (string, error) {
	key := os.Getenv("OPENWEATHER_KEY")
	cityEscaped := url.QueryEscape(city)
	reqURL := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s",
		cityEscaped, key,
	)

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(reqURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}

	var w WeatherResp
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return "", err
	}

	// Select weather emoji based on primary condition
	emoji := ""
	if len(w.Weather) > 0 {
		emoji = getWeatherEmoji(w.Weather[0].Main)
	}
	// Select country flag emoji
	flag := getCountryFlag(w.Sys.Country)
	return fmt.Sprintf("%s %s %s: %.1f°C", emoji, flag, w.Name, w.Main.Temp), nil
}

// getWeatherEmoji returns an emoji for a given weather condition string.
func getWeatherEmoji(condition string) string {
	switch condition {
	case "Clear":
		return "☀️"
	case "Clouds":
		return "☁️"
	case "Rain":
		return "🌧️"
	case "Drizzle":
		return "🌦️"
	case "Thunderstorm":
		return "⛈️"
	case "Snow":
		return "❄️"
	case "Mist", "Smoke", "Haze", "Dust", "Fog", "Sand", "Ash", "Squall", "Tornado":
		return "🌫️"
	default:
		return ""
	}
}

// getCountryFlag turns "US" into 🇺🇸, etc.
func getCountryFlag(cc string) string {
	flag := ""
	for _, r := range cc {
		if r >= 'A' && r <= 'Z' {
			flag += string(rune(0x1F1E6 + (r - 'A')))
		}
	}
	return flag
}
