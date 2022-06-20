package main

import(
    "fmt"
    "strings"
    "strconv"
)

func main(){

    input := "1 9 3 4 -5"
    var max, min int32
    var result string

    arr := strings.Split(input, " ")

    for i, el := range arr{
        intEl, _ := strconv.Atoi(el)

        if i == 0 {
            max = int32(intEl)
            min = int32(intEl)
        }

        if int32(intEl)>max {
            max = int32(intEl)
        }

        if int32(intEl)<min {
            min = int32(intEl)
        }
    }

    if max == min {
        result = fmt.Sprintf("%d", max)
    } else {
        result = fmt.Sprintf("%d %d", max, min)
    }

    fmt.Println(result)
}