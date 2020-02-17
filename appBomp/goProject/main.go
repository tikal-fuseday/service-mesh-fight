package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

const PRODUCT_PAGE_URL = "http://www.google.com"

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

/**
expected response:
{"bombId": 123 }
*/
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

	query := r.URL.Query()
	location := query.Get("url")

	//timeInSeconds = timeInSeconds + 1
	//concurrentThreads = concurrentThreads + 1

	go startSendingRequests(timeInSeconds, concurrentThreads, location)

	bombId := 3

	w.Write([]byte(fmt.Sprintf(`{"bombId": %d }`, bombId)))
}

/**
expected response:
	{status: 'running' | 'done', completed: Number(between 0 to 1), grafanaUrl: string}
*/
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

	percent := float32(countOK) / 1000.0

	status := "running"
	completedPercent := percent
	grafanaUrl := "http://www.tikalk.com"

	responseBody := fmt.Sprintf(`{"status": "%s", "completed": %f, "grafanaUrl": "%s"}`, status, completedPercent, grafanaUrl)

	w.Write([]byte(responseBody))
}

func startSendingRequests(timeInSeconds int, concurrentThreads int, urlEncoded string) {
	println("Starting ", timeInSeconds, " seconds of requests on ", concurrentThreads, " concurrent threads ...")

	if (concurrentThreads > 200) || (concurrentThreads < 1) {
		println("change concurrentThreads from ", concurrentThreads, " to 1")
		concurrentThreads = 1
	}

	responses := make(chan job, 1000)
	go aggregateStatus(responses)
	for i := 0; i < concurrentThreads; i++ {
		go sendManyRequests(i, timeInSeconds, urlEncoded, responses)
	}

	println("completed!")
}

var countOK int64
var countErr int64
var percent float32

func aggregateStatus(statusCodes chan job) {
	for j := range statusCodes {
		percent = j.percent
		if j.statusCode == 200 {
			countOK = countOK + 1
		} else {
			countErr++
		}
	}
}

type Pair struct {
	a, b interface{}
}
type job struct {
	statusCode int
	percent    float32
}

func sendManyRequests(threadNumber int, timeInSeconds int, urlEncoded string, responses chan job) {
	startTime := time.Now().Unix()
	fmt.Printf("currentTime = %v\n", startTime)

	runUntil := startTime + int64(timeInSeconds)
	fmt.Printf("runUntil = %v\n", runUntil)
	for {
		statusCode := work(threadNumber, urlEncoded)
		current := 100 * (time.Now().Unix() - startTime) / int64(timeInSeconds)
		responses <- job{statusCode, float32(current)}
		//responses<-Pair{statusCode, current}
		if time.Now().Unix() > runUntil {
			break
		}
	}
}

func work(threadNumber int, urlEncoded string) int {
	println("[thread ", threadNumber, "] send ", urlEncoded, "  ")
	resp, err := http.Get(urlEncoded)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	return resp.StatusCode
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

	api.HandleFunc("/bomb/{timeInSeconds}/{concurrentThreads}", startSending1).Methods(http.MethodPost) // + ?url=http://www.google.com
	api.HandleFunc("/bomb/{bombId}/status", findStatus).Methods(http.MethodGet)                         // ?timeInSeconds=3&concurrentThreads=4

	log.Fatal(http.ListenAndServe(":30001", r))
}
