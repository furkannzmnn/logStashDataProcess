package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
)

var logs map[int]string

func main() {
	divideTask()
}

func divideTask() {
	RunAndExecuteJobsMap(task)
	Router()
}

func handleLargeRequest(w http.ResponseWriter, r *http.Request) {
	var events event
	err := json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if events.DeviceName == "MACOS" {
		sendMail(events.DeviceName, w)
	}
	handleApiRequest(w, r)
}

func handleApiRequest(w http.ResponseWriter, r *http.Request) {
	var api apiUrl
	err := json.NewDecoder(r.Body).Decode(&api)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: SPLIT REQUEST TYPE AND SAVE TO DB AND RUN JOBS IN BACKGROUND
	// contains string in the url
	if strings.Contains(api.Url, "/api/v1/") {
		fmt.Println("api url initalized")
		_, err = w.Write([]byte("Success"))
		if err != nil {
			return
		}
		logs[rand.Intn(100)] = api.Url
	}
}

func sendMail(deviceName string, w http.ResponseWriter) {
	from := "ozmenf97@gmail.com"
	pass := "password"
	to := []string{
		"ozmenf97@gmail.com", // TODO: sql connection implementation
	}
	msg := "This is a test email:" + deviceName

	smtHost := "smtp.gmail.com"
	smtPort := ":587"

	auth := smtp.PlainAuth("", from, pass, smtHost)

	err := smtp.SendMail(smtHost+smtPort, auth, from, to, []byte(msg))

	if err != nil {
		fmt.Println("Error: ", err)
		_, err := w.Write([]byte("Error: " + err.Error()))
		if err != nil {
			return
		}
		return
	}
	fmt.Println("Mail sent")
}

func findByUserEmail(userId string) string {

	db, err := sql.Open("mysql", "root:root@tcp("+"localhost:3306"+")/logstash")
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	var email string

	err = db.QueryRow("SELECT * from dual", userId).Scan(&email) // todo: update query

	if err != nil {
		panic(err.Error())
	}
	return userId
}

type event struct {
	DeviceName string `json:"device_name"`
	UserID     string `json:"user_id"`
}

type apiUrl struct {
	Url        string `json:"url"`
	MethodName string `json:"method_name"`
	Date       string `json:"date"`
}
type requestType struct {
	RequestType string `json:"type"`
}

var task = func() {
	fmt.Println("task is running")
	for _, value := range logs {
		fmt.Println("executed task: " + value)
		// TODO: SEND TO ELASTİC TO API_URL INDEX
		_, err := http.PostForm("http://localhost:8080/create/index/count_api", nil)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		logs = make(map[int]string)
	}
}

func getRequestType(w http.ResponseWriter, r *http.Request) {
	var args requestType
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
