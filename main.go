package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func calcHandler(op byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, b := getOperands(r)
		resp := calculate(a, b, op)
		w.Header().Set("Content-type", "application/json")
		w.Write(resp)
	})
}

func getOperands(r *http.Request) (a, b int) {
	a, err := strconv.Atoi(r.FormValue("a"))
	checkErr(err)
	b, err = strconv.Atoi(r.FormValue("b"))
	checkErr(err)
	return a, b
}

func calculate(a, b int, op byte) []byte {
	var result int
	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	}
	return []byte(`{"result": ` + strconv.Itoa(result) + `}`)
}

func main() {
	http.Handle("/plus", calcHandler('+'))
	http.Handle("/minus", calcHandler('-'))
	http.Handle("/multiply", calcHandler('*'))
	http.Handle("/divide", calcHandler('/'))
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
