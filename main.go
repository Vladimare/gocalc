package main

import (
    "net/http"
    "fmt"
    "log"
    "strconv"
)

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func plusHandler(w http.ResponseWriter, r *http.Request) {
    a, err := strconv.Atoi(r.FormValue("a"))
    checkErr(err)
    b, err := strconv.Atoi(r.FormValue("b"))
    checkErr(err)
    w.Header().Set("Content-type", "application/json")
    resp := `{"result": ` + strconv.Itoa(a + b) + `}`
    w.Write([]byte(resp))
}

func minusHandler(w http.ResponseWriter, r *http.Request) {
    a, err := strconv.Atoi(r.FormValue("a"))
    checkErr(err)
    b, err := strconv.Atoi(r.FormValue("b"))
    checkErr(err)
    w.Header().Set("Content-type", "application/json")
    resp := `{"result": ` + strconv.Itoa(a - b) + `}`
    w.Write([]byte(resp))
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
    a, err := strconv.Atoi(r.FormValue("a"))
    checkErr(err)
    b, err := strconv.Atoi(r.FormValue("b"))
    checkErr(err)
    w.Header().Set("Content-type", "application/json")
    resp := `{"result": ` + strconv.Itoa(a * b) + `}`
    w.Write([]byte(resp))
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
    a, err := strconv.Atoi(r.FormValue("a"))
    checkErr(err)
    b, err := strconv.Atoi(r.FormValue("b"))
    checkErr(err)
    w.Header().Set("Content-type", "application/json")
    resp := `{"result": ` + strconv.Itoa(a / b) + `}`
    w.Write([]byte(resp))
}

func main() {
    http.HandleFunc("/plus", plusHandler)
    http.HandleFunc("/minus", minusHandler)
    http.HandleFunc("/multiply", multiplyHandler)
    http.HandleFunc("/divide", divideHandler)
    fmt.Println("Starting server at :8080")
    http.ListenAndServe(":8080", nil)
}
