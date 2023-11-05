package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func checkErr(cause string, err error) {
	if err != nil {
		log.Fatalln(cause, err)
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
	checkErr("Fail to get a", err)
	b, err = strconv.Atoi(r.FormValue("b"))
	checkErr("Fail to get b", err)
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
	hostname, err := os.Hostname()
	checkErr("Fail to get hostname", err)
	ip, err := net.ResolveIPAddr("ip", hostname)
	checkErr("Fail to get ip", err)
	return []byte(ip.String() + `:{"result": ` + strconv.Itoa(result) + `}`)
}

func main() {
	http.Handle("/plus", calcHandler('+'))
	http.Handle("/minus", calcHandler('-'))
	http.Handle("/multiply", calcHandler('*'))
	http.Handle("/divide", calcHandler('/'))
	log.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
