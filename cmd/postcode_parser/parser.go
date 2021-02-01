package main

import (
    "encoding/json"
    "errors"
    "sort"
    "strings"
)

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
    Parts     []string
    StateCode string
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
            peopleList = append(peopleList, "..... " + strings.Join(v.Parts, " ") + " " + statesList[code])
        }

        // Sort by address string
        sort.Slice(peopleList, func(i, j int) bool {
            return peopleList[i] < peopleList[j]
        })
        resultString += strings.Join(peopleList, "\n")

        resultStrings = append(resultStrings, resultString)
    }

    resultData := outputData{
        theData.ReqType,
        "success",
        strings.Join(resultStrings, "\n "),
    }
    resultBytes, _ := json.Marshal(resultData)

    return e, string(resultBytes)
}

func parseAddress(s string) (a address, e error) {
    a.Parts = strings.Split(s, ",")
    if len(a.Parts) < 3 {
        return address{}, errors.New("incorrect item format")
    }
    for i, s := range a.Parts {
        a.Parts[i] = strings.Trim(s, " ")
    }

    stateParts := strings.Split(a.Parts[len(a.Parts) - 1], " ")
    if len(stateParts) < 2 {
        return address{}, errors.New("incorrect item format")
    }
    for i, s := range stateParts {
        stateParts[i] = strings.Trim(s, " ")
    }

    a.StateCode = stateParts[len(stateParts)-1]
    a.Parts[len(a.Parts) - 1] = strings.Join(stateParts[0:len(stateParts)-1], " ")
    if _, ok := statesList[a.StateCode]; !ok {
        return address{}, errors.New("incorrect City code")
    }

    return a, nil
}
