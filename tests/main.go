package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	url := "http://localhost:8080/api/v1/register"

	type RegisterPayload struct {
		Username  string   `json:"username" binding:"required"`
		Email     string   `json:"email" binding:"required"`
		Password  string   `json:"password" binding:"required"`
		Gender    string   `json:"gender" binding:"required"`
		Label     []string `json:"label"`
		Biography string   `gorm:"default:''" json:"biography"`
	}

	var gender = []string{"male", "female"}
	var labels = []string{
		"football", "basketball", "volleyball",
		"painting", "music", "photography",
		"programming", "gaming", "robotics",
		"physics", "chemistry", "biology",
		"movies", "books", "travel",
	}

	for i := 1; i <= 10; i++ {
		// choose random 3 labels
		var userLabels []string
		for j := 0; j < 3; j++ {
			userLabels = append(userLabels, labels[(i+j)%len(labels)])
		}

		var number = strconv.Itoa(i)

		payload := RegisterPayload{
			Username:  "user" + number,
			Email:     "user" + number + "@gmail.com",
			Password:  "string",
			Gender:    gender[i%2],
			Biography: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Label:     userLabels,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Request %d failed: %v", i, err)
			continue
		}
		defer resp.Body.Close()

		log.Printf("Request %d completed with status: %s", i, resp.Status)
	}
}
