package main

import "fmt"

func quicksort(list []int) []int {
    var lesser, greater []int
    var pivot = 0

    if len(list) <= 1 {
        return list
    } else {
        pivot = list[0]
        list = list[1:]
        for _, x := range list {
            if x <= pivot {
                lesser = append(lesser, x)
            } else {
                greater = append(greater, x)
            }
        }
    }
    var result = append(quicksort(lesser), pivot)
    result = append(result, quicksort(greater)...)
    return result
}

func main() {
    var l = []int{2, 3, 15, 7, 11, 13}
    fmt.Println(quicksort(l))
    fmt.Println(quicksort([]int{}))
}