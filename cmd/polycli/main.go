package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/JackIABishop/polycli/pkg/weather"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ no .env found:", err)
	}

	cities := []string{"London", "New York", "Tokyo", "Paris"}
	var wg sync.WaitGroup
	results := make(chan string, len(cities))

	wg.Add(len(cities))
	for _, c := range cities {
		go func(city string) {
			defer wg.Done()
			res, err := weather.Fetch(city)
			if err != nil {
				results <- fmt.Sprintf("❌ %s: %v", city, err)
			} else {
				results <- res
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}
