package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//structure Event used for deserialization of json string into struct we can use
type Event struct {
	ResizeFrom		Dimension
	ResizeTo			Dimension
	WebsiteUrl 		string
	Pasted				bool
	Time 					int
	SessionId 		string
	FormId 				string
	EventType 		string
}
//server struct that will be HTTP request handler
type server struct{}

// map for storing session events from client
var clientSessions map[string]*Data

//main function for creating a server and assigning request handlers
func main() {
	s := &server{}
	http.Handle("/api", s)
	clientSessions = make(map[string]*Data)
	http.HandleFunc("/api/events", eventApi)
	
	//helper endpoint for testing
	http.HandleFunc("/api/sessions/{id}", sessionApi)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

/*
	method on server struct that for serving http requests
	it needs to implement ServerHTTP function in order to 
	be considered a Handler interface
	*/
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		
		w.Write([]byte(`{"message": "Test project /api endoints to get more info"}`))
		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(`{"message": "Use POST request on /api/event API endpoint to send Event data \n Use GET request on /api/session/{session-id} to get full session stored by now"`)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(`{"message": "not found"}`)
		}
		
}

/*
  ---------API ENDPOINTS-----------------
*/
/*
responds to [POST] requests on /event api
accepts as params event object as json
*/
func eventApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var e Event
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&e)
		
		//checking for errors
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(`{"message": "Error parsing JSON"`)
		}
		if e.SessionId == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid session id")
		}

		s := clientSessions[e.SessionId]
		if s == nil {
			s = NewData(e)
		}

		err = s.updateSession(e)
		if err != nil {
			
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} 
		clientSessions[e.SessionId] = s
		if s.isCompleted() {
			fmt.Println("Form submitted, struct completed")
			//probably at this point we should clear the session from the map
			//in full app I presume we would save this data or send it somewhere and then clear it from local storage
			// clientSessions[s.SessionId] = nil
		}
		s.printDataStruct()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(`{"message": "not supported"}`)
	}
}

// helper function I used to check struct and api from postman
func sessionApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		sid := strings.TrimPrefix(r.URL.Path, "/api/")
    if sid == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(`{"message": "Provide proper session id"}`)
		}
		s := clientSessions[sid]
		if s == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(`{"message": "Session with that id not found"}`)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*s)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(`{"message": "not supported"}`)		
	}
}