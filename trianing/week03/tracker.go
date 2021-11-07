package week03

import (
	"log"
	"net/http"
	"time"
)

// Tracker track the event of application
type Tracker struct {}

// Event record an event to a database or stream
func (t *Tracker) Event(data string) {
	time.Sleep(time.Millisecond)
	log.Println(data)
}

// App an application
type App struct {
	track Tracker
}

func (a *App) Handle(w http.ResponseWriter, req *http.Request)  {
	w.WriteHeader(http.StatusCreated)

	go a.track.Event("create")
}

// func runApp() {
// 	var a App
// 	a.Handle()
// }
