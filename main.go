package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(strings.Repeat("=", 80)) // ASCII line
		fmt.Println("NEW WEBHOOK RECEIVED")
		fmt.Println(strings.Repeat("=", 80)) // ASCII line

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		var payload map[string]interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			return
		}

		if eventResource, ok := payload["event_resource"].(string); ok {
			var parsedEventResource interface{}
			if err := json.Unmarshal([]byte(eventResource), &parsedEventResource); err != nil {
				log.Printf("Error parsing nested JSON: %v", err)
				return
			}
			payload["event_resource"] = parsedEventResource
		}

		prettyJSON, err := json.MarshalIndent(payload, "", "    ")
		if err != nil {
			log.Printf("Error pretty printing JSON: %v", err)
			return
		}
		log.Println(string(prettyJSON))
		fmt.Println(strings.Repeat("-", 80))
		fmt.Println()
	})
	fmt.Println("Listening for webhooks on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
