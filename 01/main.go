package main

import (
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    content, error := os.ReadFile("list.text")
    if error != nil {
        fmt.Println(error)
    }
    str := string(content)
    arrayOfString := getArray(str)
    leftArray, rightArray := splitArray(arrayOfString)

    leftArrayOfInt := toInt(leftArray)
    rightArrayOfInt := toInt(rightArray)

    leftArraySorted := sortArray(leftArrayOfInt)
    rightArraySorted := sortArray(rightArrayOfInt)

    diff := getDiff(leftArraySorted, rightArraySorted)
    diffSum := arraySum(diff)

    similarArray := countSimilar(leftArraySorted, rightArraySorted)
    similarSum := arraySum(similarArray)

    fmt.Println(diffSum)
    fmt.Println(similarSum)
}

func getArray(str string) []string {
    a := strings.Split(str, "\n")
    return a
}

func splitArray(arrayOfString []string) ([]string, []string) {
    var leftArray []string
    var rightArray []string
    for i := 0; i < len(arrayOfString); i++ {
        a := strings.Split(arrayOfString[i], "   ")
        leftArray = append(leftArray, a[0])
        rightArray = append(rightArray, a[1])
    }
    return leftArray, rightArray
}

func toInt(stringArray []string) []int {
    var b []int
    for i := 0; i < len(stringArray); i++ {
        j, _ := strconv.Atoi(stringArray[i])
        b = append(b, j)
    }
    return b
}

func sortArray(arrayOfInt []int) []int {
    sort.Slice(arrayOfInt, func(i, j int) bool {
        return arrayOfInt[i] < arrayOfInt[j]
    })
    return arrayOfInt
}

func getDiff(leftArray []int, rightArray []int) []int {
    var diffArray []int
    var diff int
    for i := 0; i < len(leftArray); i++ {
        if leftArray[i] > rightArray[i] {
            diff = leftArray[i] - rightArray[i]
        } else {
            diff = rightArray[i] - leftArray[i]
        }
        diffArray = append(diffArray, diff)
    }
    return diffArray
}

func arraySum(array []int) int {
    sum := 0
    for i := 0; i < len(array); i++ {
        sum = sum + array[i]
    }

    return sum
}

func countSimilar(leftArray []int, rightArray []int) []int {
    var similarArray []int
    for i := 0; i < len(leftArray); i++ {
        similar := 0
        for j := 0; j < len(rightArray); j++ {
            if leftArray[i] == rightArray[j] {
                similar += 1
            }
        }
        similarArray = append(similarArray, similar*leftArray[i])
    }

    return similarArray
}
