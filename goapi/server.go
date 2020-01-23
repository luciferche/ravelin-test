package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"errors"
	"io/ioutil"
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

var clientSessions map[string]*Data


// updates session and returns flag if there was an error with params
func (d *Data) updateSession(ev Event) error {
	
	switch ev.EventType {
	case "copyAndPaste":
		d.CopyAndPaste[ev.FormId] = true
	case "screenResize":
		d.ResizeFrom = ev.ResizeFrom
		d.ResizeTo = ev.ResizeTo
	case "timeTaken":
		d.FormCompletionTime = ev.Time
		//this means session finished? = submitted
	default :
		return errors.New("Event type not supported " + ev.EventType)
	}
	return nil
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	clientSessions = make(map[string]*Data)
	router.HandleFunc("/api", homeHello)
	router.HandleFunc("/api/event", createEvent).Methods("POST")
	router.HandleFunc("/api/session/{id}", getOneEvent).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}


/*
  ---------API ENDPOINTS-----------------
	**/
//add postEvent
//	homeHello function responds to the root / request, prints some help
func homeHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use POST request on /api/event API endpoint to send Event data \n Use GET request on /api/session/{session-id} to get full session stored by now")
}


/*
responds to [POST] requests on /event api
accepts as params event object as json
*/
func createEvent (w http.ResponseWriter, r *http.Request) {
	var e Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter proper parameters for event")
	}
	json.Unmarshal(reqBody, &e)

	if e.SessionId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid session id")
		return
	}
	s := clientSessions[e.SessionId]
	if s != nil {
		fmt.Println("Found session already present")
	} else {
		s = NewData(e)
	}
	err = s.updateSession(e)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	clientSessions[e.SessionId] = s
	if s.FormCompletionTime > 0 {
		fmt.Println("Form submitted, struct completed")
	}
	s.printDataStruct()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("")
	// json.NewEncoder(w).Encode(e)
}

// helper function I used to check struct and api from postman
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	sessionId := mux.Vars(r)["id"]

	session := clientSessions[sessionId]
	json.NewEncoder(w).Encode(*session)
}