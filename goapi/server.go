package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int // Seconds
}

type event struct {
	ResizeFrom		Dimension `json:"resizeFrom"`
	ResizeTo			Dimension `json:"resizeTo"`
	WebsiteUrl 		string 		`json:"siteUrl"`
	Pasted				bool 			`json:"pasted"`
	Time 					int 			`json:"time"`
	SessionId 		string 		`json:"sessionId"`
	FormId 				string 		`json:"formId"`
  EventType 		string 		`json:"eventType"`
}
type Dimension struct {
	Width  string
	Height string
}
var clientSessions map[string]*Data

func homeHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use POST request on /event API endpoint to send event data")
}
// updates session and returns flag if there was an error with params
func updateSession(session *Data, ev event) bool {
	
	switch ev.EventType {
	case "copyAndPaste":
		session.CopyAndPaste[ev.FormId] = true
		// session.WebsiteUrl = ev.WebsiteUrl
	case "screenResize":
		session.ResizeFrom = ev.ResizeFrom
		session.ResizeTo = ev.ResizeTo
		// session.WebsiteUrl = ev.WebsiteUrl
	case "timeTaken":
		session.FormCompletionTime = ev.Time
	default :
		return true;
	}
	return false;
}

//constructor for mapping event from params to session data
func newData(ev event) *Data {
	d := new(Data)
	d.ResizeTo = ev.ResizeTo
	d.ResizeFrom = ev.ResizeFrom
	d.WebsiteUrl = ev.WebsiteUrl 	//change to hash
	d.CopyAndPaste = make(map[string]bool)
	if ev.Pasted {
		d.CopyAndPaste[ev.FormId] = true
	}
	d.FormCompletionTime = ev.Time
	d.SessionId = ev.SessionId
	return d
}

/*
	responds to [POST] requests on /event api
	accepts as params event object as json
	*/
func createEvent (w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter proper parameters for event")
	}
	json.Unmarshal(reqBody, &newEvent)
	fmt.Fprintf(w, "Evo ga novi " + newEvent.SessionId)
	session := clientSessions[newEvent.SessionId]
	if session.SessionId != "" {
		err := updateSession(session, newEvent)
		if err {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Wrong params")
			return
		}
	} else {
		session = newData(newEvent)
	}
	clientSessions[newEvent.SessionId] = session;
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("")
	// json.NewEncoder(w).Encode(newEvent)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", homeHello)
	router.HandleFunc("/api/event", createEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}