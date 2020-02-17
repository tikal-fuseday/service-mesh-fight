package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := -1
	var err error

	pathParams := mux.Vars(r)
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number for userID"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number for commentID"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func startSending1(w http.ResponseWriter, r *http.Request) {
	println("============= startSending1 ==========================")
	w.Header().Set("Content-Type", "application/json")
	timeInSeconds := -1
	var err error

	pathParams := mux.Vars(r)
	if val, ok := pathParams["timeInSeconds"]; ok {
		timeInSeconds, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "timeInSeconds must be a number"}`))
			return
		}
	}

	concurrentThreads := -1
	if val, ok := pathParams["concurrentThreads"]; ok {
		concurrentThreads, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number for commentID"}`))
			return
		}
	}

	timeInSeconds = timeInSeconds + 1
	concurrentThreads = concurrentThreads + 1

	go startSendingRequests(timeInSeconds, concurrentThreads)

	w.Write([]byte(fmt.Sprintf(`{"timeInSeconds": %d, "concurrentThreads": %d, " OK" }`, timeInSeconds, concurrentThreads)))
}

func findStatus(w http.ResponseWriter, r *http.Request) {
	println("============= findStatus ==========================")

	w.Header().Set("Content-Type", "application/json")
	bombId := -1
	var err error

	pathParams := mux.Vars(r)
	if val, ok := pathParams["bombId"]; ok {
		println(val)
		bombId, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "bombId must be a number"}`))
			return
		}
	}

	bombId = bombId + 1

	// TODO here: retrieve data of this bombId

	// {status: 'running' | 'done', completed: Number(between 0 to 1), grafanaUrl: string}
	status := "running"
	completedPercent := 0.5
	grafanaUrl := "http://www.tikalk.com"

	//responseBody := fmt.Sprintf(`{status: "running", completed: 0.3, grafanaUrl: "http://www.tikalk.com"}: %d, "concurrentThreads": %d, " OK" }`, timeInSeconds, concurrentThreads)
	responseBody := fmt.Sprintf(`{"status": "%s", "completed": %f, "grafanaUrl": "%s"}`, status, completedPercent, grafanaUrl)

	w.Write([]byte(responseBody))
}

func startSendingRequests(timeInSeconds int, concurrentThreads int) {
	println("Starting ...")
	busyWait()
	println("completed!")
}

func busyWait() {
	for a := 0; a <= 10000-1; a++ {
		for c := 0; c <= 1000000-1; c++ {
			b := 0
			b = b + 1
		}
	}
}

/*******************************
Clear is Better than Clever
Rob Pike
******************************** */
func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	//api.HandleFunc("", get).Methods(http.MethodGet)
	//api.HandleFunc("", post).Methods(http.MethodPost)
	//api.HandleFunc("", put).Methods(http.MethodPut)
	//api.HandleFunc("", delete).Methods(http.MethodDelete)

	//api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
	//api.HandleFunc("/bomb/{timeInSeconds}/{concurrentThreads}", startSending1).Methods(http.MethodGet)
	api.HandleFunc("/bomb/{timeInSeconds}/{concurrentThreads}", startSending1).Methods(http.MethodPost)
	api.HandleFunc("/bomb/{bombId}/status", findStatus).Methods(http.MethodGet) // ?timeInSeconds=3&concurrentThreads=4

	log.Fatal(http.ListenAndServe(":8080", r))
}
