package weather

import "testing"

func TestGetWeatherEmoji(t *testing.T) {
	cases := []struct {
		cond, want string
	}{
		{"Clear", "☀️"},
		{"Clouds", "☁️"},
		{"Rain", "🌧️"},
		{"Drizzle", "🌦️"},
		{"Thunderstorm", "⛈️"},
		{"Snow", "❄️"},
		{"Fog", "🌫️"},
		{"Unknown", ""},
	}
	for _, c := range cases {
		got := getWeatherEmoji(c.cond)
		if got != c.want {
			t.Errorf("getWeatherEmoji(%q) = %q; want %q", c.cond, got, c.want)
		}
	}
}

func TestGetCountryFlag(t *testing.T) {
	cases := []struct {
		code, want string
	}{
		{"US", "🇺🇸"},
		{"GB", "🇬🇧"},
		{"JP", "🇯🇵"},
		{"", ""},
		{"ZZ", "🇿🇿"},
	}
	for _, c := range cases {
		if got := getCountryFlag(c.code); got != c.want {
			t.Errorf("getCountryFlag(%q) = %q; want %q", c.code, got, c.want)
		}
	}
}
