package main
import (
	"reflect"
	"fmt"
)

type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int // Seconds
}

type Dimension struct {
	Width  string
	Height string
}

//constructor for mapping event from params to session data
func NewData(ev Event) *Data {
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

func (d *Data) printDataStruct() {
	v := reflect.ValueOf(*d)
	typeOfS := v.Type()

	for i := 0; i< v.NumField(); i++ {
			fmt.Printf("\t%s\t\t: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}