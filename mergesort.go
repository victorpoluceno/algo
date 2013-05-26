package main

import "fmt"

func merge(left []int, right []int) []int {
    var result []int
    var i, j = 0, 0
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            result = append(result, left[i])
            i += 1
        } else {
            result = append(result, right[j])
            j += 1
        }
    }
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    return result
}

func mergesort(list []int) []int {
    if len(list) <= 1 {
        return list
    }
    var middle = len(list)/2
    var left = mergesort(list[:middle])
    var right = mergesort(list[middle:])
    return merge(left, right)
}

func main() {
    var l = []int{2, 3, 15, 7, 11, 13}
    fmt.Println(mergesort(l))
    fmt.Println(mergesort([]int{}))
}