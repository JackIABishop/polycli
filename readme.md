# Polycli 🌦️☁️

A tiny concurrent CLI that fetches weather (and more!) in parallel. Written in Go with goroutines, channels, and a clean code structure.

---

## Features

- Fetch current temperature for multiple cities concurrently  
- 🏷️ Country‐flag and weather‐condition emojis  
- ⚙️ Configurable via `.env` (OpenWeatherMap API key)  
- ✔️ Testable helpers: emoji selectors & HTTP logic  

---

## Prerequisites

- [Go 1.20+](https://golang.org/dl/)  
- OpenWeatherMap API key (free at https://openweathermap.org/)  
- `git`, to clone the repo  

---

## Installation

```bash
git clone git@github.com:JackIABishop/polycli.git
cd polycli
go mod download
```

---
## Example API Response

```json
{
  "name": "New York",
  "main": { "temp": 28.8 },
  "weather": [
    { "main": "Clouds", "description": "scattered clouds" }
  ],
  "sys": { "country": "US" }
}
```