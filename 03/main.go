package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    content, err := os.ReadFile("list.text")
    if err != nil {
        fmt.Println(err)
    }
    str := string(content)
    arrayOfString := getArray(str)

    result := getResult(arrayOfString)

    fmt.Println("result: ", result)
}

func getArray(str string) []string {
    a := strings.Split(str, "\n")
    return a
}

func getResult(array []string) int {
    var result int
    for i := 0; i < len(array); i++ {
        matches := filterByRegex(array[i])
        for j := 0; j < len(matches); j++ {
            result += countRegexWithRules(matches[j])
        }
    }
    return result
}

func filterByRegex(str string) []string {
    r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
    return r.FindAllString(str, -1)
}

func countRegexWithRules(str string) int {
    numbersInStr := getNumberInStr(str)
    numberInStringSplitted := getNumberInStringSplitted(numbersInStr[0])
    product := numberInStringSplitted[0] * numberInStringSplitted[1]
    return product
}

func getNumberInStr(str string) []string {
    r, _ := regexp.Compile(`[0-9]{1,3},[0-9]{1,3}`)
    return r.FindAllString(str, -1)
}

func getNumberInStringSplitted(str string) []int {
    stringArray := strings.Split(str, ",")
    var b []int
    for i := 0; i < len(stringArray); i++ {
        j, _ := strconv.Atoi(stringArray[i])
        b = append(b, j)
    }
    return b
}
