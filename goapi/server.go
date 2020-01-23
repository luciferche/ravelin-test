package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
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

type Event struct {
	ResizeFrom		Dimension
	ResizeTo			Dimension
	WebsiteUrl 		string 		`json:"siteUrl"`
	Pasted				bool
	Time 					int
	SessionId 		string
	FormId 				string
	EventType 		string
}
type Dimension struct {
	Width  string
	Height string
}
var clientSessions map[string]*Data

func homeHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use POST request on /api/event API endpoint to send Event data \n Use GET request on /api/session/{session-id} to get full session stored by now")
}

// updates session and returns flag if there was an error with params
func updateSession(session *Data, ev Event) bool {
	
	switch ev.EventType {
	case "copyAndPaste":
		session.CopyAndPaste[ev.FormId] = true
	case "screenResize":
		session.ResizeFrom = ev.ResizeFrom
		session.ResizeTo = ev.ResizeTo
	case "timeTaken":
		session.FormCompletionTime = ev.Time
		//this means session finished? = submitted
	default :
		fmt.Printf("Event type is %v\n", ev.EventType)
		return true;
	}
	return false;
}

//constructor for mapping event from params to session data
func newData(ev Event) *Data {
	d := new(Data)
	d.ResizeTo = ev.ResizeTo
	d.ResizeFrom = ev.ResizeFrom
	d.WebsiteUrl = mainHash(ev.WebsiteUrl) 	//change to hash
	fmt.Println("HASHED URL : %v", d.WebsiteUrl)
	
	d.CopyAndPaste = make(map[string]bool,3)
	if ev.Pasted {
		d.CopyAndPaste[ev.FormId] = true
	}
	d.FormCompletionTime = ev.Time
	d.SessionId = ev.SessionId
	return d
}

func printDataStruct(data *Data) {
	v := reflect.ValueOf(*data)
	typeOfS := v.Type()

	for i := 0; i< v.NumField(); i++ {
			fmt.Printf("\t%s\t\t: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
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
	session := clientSessions[e.SessionId]
	if session != nil {
		fmt.Println("Found session already present")
	} else {
		session = newData(e)
	}
	hasErrors := updateSession(session, e)
	if hasErrors {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong params")
		return
	}
	clientSessions[e.SessionId] = session
	if session.FormCompletionTime > 0 {
		fmt.Println("Form submitted, struct completed")
	}
	printDataStruct(session)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("")
	// json.NewEncoder(w).Encode(e)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	clientSessions = make(map[string]*Data)
	router.HandleFunc("/api", homeHello)
	router.HandleFunc("/api/event", createEvent).Methods("POST")
	router.HandleFunc("/api/session/{id}", getOneEvent).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	sessionId := mux.Vars(r)["id"]

	session := clientSessions[sessionId]
	json.NewEncoder(w).Encode(*session)
}