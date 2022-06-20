package main

import "fmt"

func main(){
    arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
    var result []int

    for i:=0; i<len(arr); i++ {

        isElemDuplicated := false

        for j:=0; j<len(result); j++{

            if arr[i] == result[j] {
                isElemDuplicated = true
                break
            }

        }

        if !isElemDuplicated {
            result = append(result, arr[i])
        }

    }

    fmt.Print(result)

}