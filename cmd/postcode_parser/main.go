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
        writer.WriteHeader(http.StatusMethodNotAllowed)
        io.WriteString(writer, "{\"result\":\"fail\",\"data\":\"POST request expected\"}")
        return
    }

    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        io.WriteString(writer, "{\"result\":\"fail\",\"data\":\""+err.Error()+"\"}")
        return
    }

    err, output := parse(string(body))
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        io.WriteString(writer, "{\"result\":\"fail\",\"data\":\""+err.Error()+"\"}")
        return
    }

    io.WriteString(writer, output)
}
