package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/adk-go/adk"
)

// getWeather retrieves the current weather report for a specified city.
func getWeather(city string) map[string]string {
	if strings.ToLower(city) == "new york" {
		return map[string]string{
			"status": "success",
			"report": "The weather in New York is sunny with a temperature of 25 degrees Celsius (77 degrees Fahrenheit).",
		}
	}
	return map[string]string{
		"status":        "error",
		"error_message": fmt.Sprintf("Weather information for '%s' is not available.", city),
	}
}

// getCurrentTime returns the current time in a specified city.
func getCurrentTime(city string) map[string]string {
	var loc *time.Location
	var err error

	if strings.ToLower(city) == "new york" {
		loc, err = time.LoadLocation("America/New_York")
	} else {
		return map[string]string{
			"status":        "error",
			"error_message": fmt.Sprintf("Sorry, I don't have timezone information for %s.", city),
		}
	}

	if err != nil {
		log.Printf("Error loading location for %s: %v", city, err)
		return map[string]string{
			"status":        "error",
			"error_message": fmt.Sprintf("Sorry, I could not load timezone information for %s.", city),
		}
	}

	now := time.Now().In(loc)
	report := fmt.Sprintf("The current time in %s is %s", city, now.Format("2006-01-02 15:04:05 MST"))
	return map[string]string{
		"status": "success",
		"report": report,
	}
}

func main() {
	agent := adk.NewAgent(
		adk.WithName("weather_time_agent"),
		adk.WithModel("gemini-2.0-flash"),
		adk.WithDescription("Agent to answer questions about the time and weather in a city."),
		adk.WithInstruction("You are a helpful agent who can answer user questions about the time and weather in a city."),
		adk.WithTools(getWeather, getCurrentTime),
	)

	if err := adk.Run(agent); err != nil {
		log.Fatalf("Error running agent: %v", err)
	}
}
