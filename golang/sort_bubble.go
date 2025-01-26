package main

import "fmt"

func sort_bubble(a []int) {
    l := len(a)
    
    for i := 0; i < l-1; i++ {
        for j := 0; j < l-i-1; j++ {
            if a[j] > a[j+1] {
                temp := a[j]
                a[j] = a[j+1]
                a[j+1] = temp
            }
        }
    }
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    fmt.Println("Before Sorting:", a)
    sort_bubble(a)
    fmt.Println("After Sorting:", a)
}
