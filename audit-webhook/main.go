package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Event struct {
	Verb      string     `json:"verb"`
	Stage     string     `json:"stage"`
	ObjectRef *ObjectRef `json:"objectRef"`
}

type ObjectRef struct {
	Resource string `json:"resource"`
}

type EventList struct {
	Items []Event `json:"items"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read body: %v\n", err)
			http.Error(w, "Failed to read body", http.StatusBadRequest)
			return
		}

		var events EventList
		err = json.Unmarshal(body, &events)
		if err != nil {
			log.Printf("Failed to unmarshal audit events: %v\n", err)
			http.Error(w, "Failed to unmarshal audit events", http.StatusBadRequest)
			return
		}

		// Iterate and filter audit events
		for _, event := range events.Items {
			if isPodCreation(event) {
				log.Printf("Pod creation event detected: %+v\n", event)
			}
		}
	})

	log.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// isPodCreation returns true if the given event is of a pod creation
func isPodCreation(event Event) bool {
	return event.Verb == "create" &&
		event.Stage == "ResponseComplete" &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "pods"
}
