package main

import (
    "fmt"
    "os"
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
    arrayOfIntSplitted := splitArray(arrayOfString)

    safeNum := getSafe(arrayOfIntSplitted)

    fmt.Println("safeNum: ", safeNum)
}

func getArray(str string) []string {
    a := strings.Split(str, "\n")
    return a
}

func splitArray(arrayOfString []string) [][]int {
    var reports [][]int
    for i := 0; i < len(arrayOfString); i++ {
        a := strings.Split(arrayOfString[i], " ")
        reports = append(reports, toInt(a))
    }
    return reports
}

func toInt(stringArray []string) []int {
    var b []int
    for i := 0; i < len(stringArray); i++ {
        j, _ := strconv.Atoi(stringArray[i])
        b = append(b, j)
    }
    return b
}

func getSafe(array [][]int) int {
    safe := 0
    for i := 0; i < len(array); i++ {
        isSafe, indexes := isEachSafe(array[i])
        indexes = removeDuplicates(indexes)
        if !isSafe {
            for j := 0; j < len(array[i]); j++ {
                var isBonusSafe bool
                arrayBonus := removeIndex(array[i], j)
                isBonusSafe, _ = isEachSafe(arrayBonus)
                if isBonusSafe {
                    isSafe = true
                    break
                }
                fmt.Println(array[i], arrayBonus, isBonusSafe, indexes)
            }
        } else {
            isSafe = true
        }
        if isSafe {
            safe += 1
        }
    }
    return safe
}

func removeIndex(array []int, index int) []int {
    newArray := make([]int, len(array))
    copy(newArray, array)

    return append(newArray[:index], newArray[index+1:]...)
}

func removeDuplicates(input []int) []int {
    seen := make(map[int]bool)
    result := []int{}

    for _, value := range input {
        if !seen[value] {
            seen[value] = true
            result = append(result, value)
        }
    }

    return result
}

func isEachSafe(array []int) (bool, []int) {
    increaseOrDecrease, index := isIncreasingOrDecreasing(array)
    boolean := true
    var indexes []int
    if !increaseOrDecrease {
        boolean = false
        indexes = append(indexes, index...)
    }
    for i := 0; i < len(array)-1; i++ {
        if isEqual(array[i], array[i+1]) {
            /*fmt.Println(array, "isEqual", array[i], array[i+1])*/
            boolean = false
            indexes = append(indexes, i+1)
            indexes = append(indexes, i)
        }
        if !isTolerence(array[i], array[i+1]) {
            /*fmt.Println(array, "isNotTolerence", array[i], array[i+1])*/
            boolean = false
            indexes = append(indexes, i+1)
            indexes = append(indexes, i)
        }
    }
    return boolean, indexes
}

func isTolerence(a int, b int) bool {
    maxDiff := 3
    minDif := 1

    return (a-b <= maxDiff && a-b >= minDif) || (b-a <= maxDiff && b-a >= minDif)
}

func isEqual(a int, b int) bool {
    return a == b
}

func isIncreasingOrDecreasing(array []int) (bool, []int) {
    increasing := 0
    decreasing := 0
    boolean := true
    var indexes []int
    for i := 0; i < len(array)-1; i++ {
        if array[i] < array[i+1] && decreasing == 0 {
            increasing += 1
        } else if array[i] > array[i+1] && increasing == 0 {
            decreasing += 1
        } else {
            /*fmt.Println(array, "isNotIncreasingOrDecreasing", array[i], array[i+1], increasing, decreasing)*/
            boolean = false
            indexes = append(indexes, i+1)
            indexes = append(indexes, i)
        }
    }
    return boolean, indexes
}
