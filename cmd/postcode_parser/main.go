package main

import (
    "io"
    "io/ioutil"
    "log"
    "net/http"
)

var statesList = map[string]string{
    "AZ": "Arizona",
    "CA": "California",
    "ID": "Idaho",
    "IN": "Indiana",
    "MA": "Massachusetts",
    "OK": "Oklahoma",
    "PA": "Pennsylvania",
    "VA": "Virginia",
}

func main() {
    println("Start ParseServer ...")
    http.HandleFunc("/", ParseServer)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func ParseServer(writer http.ResponseWriter, request *http.Request) {
    if request.Method != "POST" {
        io.WriteString(writer, "{expected POST}")
        return
    }

    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        io.WriteString(writer, "{some error}")
        return
    }

    err, output := parse(string(body))
    if err != nil {
        io.WriteString(writer, "{some error}")
        return
    }

    io.WriteString(writer, output)
}
