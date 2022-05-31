package main

import (
	"l2_http/control"
	"log"
	"net/http"
)

type myMux struct {
	*http.ServeMux
	strictSlash bool
}

func (m *myMux)StrictSlash(val bool){
	m.strictSlash=true
}


func main() {
	r:=myMux{http.NewServeMux(),
		false}
	r.StrictSlash(true)

	EventCreateHandler:=http.HandlerFunc(control.EventCreateHandler)
	EventUpdateHandler:=http.HandlerFunc(control.EventUpdateHandler)
	EventDeleteHandler:=http.HandlerFunc(control.EventDeleteHandler)
	EventsAllHandler:=http.HandlerFunc(control.EventsAllHandler)
	EventDayHandler:=http.HandlerFunc(control.EventsDayHandler)
	EventWeekHandler:=http.HandlerFunc(control.EventsWeekHandler)
	EventMonthHandler:=http.HandlerFunc(control.EventsMonthHandler)

	r.Handle("/events_for_day",control.MiddlewareLog(EventDayHandler))
	r.Handle("/events_for_week",control.MiddlewareLog(EventWeekHandler))
	r.Handle("/events_for_month",control.MiddlewareLog(EventMonthHandler))
	r.Handle("/create_event",control.MiddlewareJSONCheck(control.MiddlewareLog(EventCreateHandler)))
	r.Handle("/update_event",control.MiddlewareJSONCheck(control.MiddlewareLog(EventUpdateHandler)))
	r.Handle("/delete_event",control.MiddlewareLog(EventDeleteHandler))
	r.Handle("/events_all",control.MiddlewareLog(EventsAllHandler))

	log.Fatal(http.ListenAndServe("localhost:2222",r))
}
