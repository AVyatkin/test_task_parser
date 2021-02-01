package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
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
	q := getInputFromFile()
	e, result := parse(q)
	if e != nil {
		//todo
	}

	fmt.Printf("result: %s\n\n", result)
}

func getInputFromFile() string {
	content, err := ioutil.ReadFile("../../source/inputExample.json")
	if err != nil {
		return ""
	}

	return string(content)
}

func parse(data string) (e error, result string) {
	var theData inputData
	err := json.Unmarshal([]byte(data), &theData)
	if err != nil {
		return e, result
	}

	// Keeps different states
	itemsStates := []string{}
	// Keeps itemsByCode data
	itemsByCode := make(map[string][]address)

	// Parse addresses
	for _, v := range theData.Data {
		if v.Item == "" {
			continue
		}
		addr, e := parseAddress(v.Item)
		if e != nil {
			println("Parse address error: " + e.Error())
			continue
		}

		if _, ok := itemsByCode[addr.StateCode]; !ok {
			itemsByCode[addr.StateCode] = []address{}
			itemsStates = append(itemsStates, addr.StateCode)
		}
		itemsByCode[addr.StateCode] = append(itemsByCode[addr.StateCode], addr)
	}

	// Sort by states
	sort.Slice(itemsStates, func(i, j int) bool {
		return itemsStates[i] < itemsStates[j]
	})

	resultStrings := []string{}
	for _, code := range itemsStates {
		resultString := statesList[code] + "\n"

		items := itemsByCode[code]
		peopleList := []string{}
		for _, v := range items {
			peopleList = append(peopleList, "..... "+v.Name+" "+v.Street+" "+v.City+" "+statesList[code])
		}

		// Sort by address string
		sort.Slice(peopleList, func(i, j int) bool {
			return peopleList[i] < peopleList[j]
		})
		resultString += strings.Join(peopleList, "\n")

		resultStrings = append(resultStrings, resultString)
	}

	var resultData outputData
	resultData.ResType = theData.ReqType
	resultData.Result = "success"
	resultData.Data = strings.Join(resultStrings, "\n")

	resultBytes, _ := json.Marshal(resultData)

	return e, string(resultBytes)
}

type inputData struct {
	ReqType string     `json:"req_type"`
	Data    []dataItem `json:"data"`
}

type dataItem struct {
	Item string `json:"item"`
}

type outputData struct {
	ResType string `json:"res_type"`
	Result  string `json:"result"`
	Data    string `json:"data"`
}

type address struct {
	Name      string
	Street    string
	City      string
	StateCode string
}

func parseAddress(s string) (a address, e error) {
	parts := strings.Split(s, ",")
	if len(parts) < 3 {
		return address{}, errors.New("incorrect item format")
	}

	a.Name = strings.Trim(parts[0], " ")
	a.Street = strings.Trim(parts[1], " ")

	stateString := strings.Trim(parts[2], " ")
	stateParts := strings.Split(stateString, " ")
	if len(stateParts) < 2 {
		return address{}, errors.New("incorrect item format")
	}
	a.City = strings.Join(stateParts[0:len(stateParts)-1], " ")
	a.StateCode = stateParts[len(stateParts)-1]
	if _, ok := statesList[a.StateCode]; !ok {
		return address{}, errors.New("incorrect City code")
	}

	return a, nil
}
