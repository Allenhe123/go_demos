package testing

import (
    "fmt"
)

func TestingRange() {
    nums := []int{1,2,3,4,5,6}
    for i := range nums {   // i is index not value
        if i == 3 {
            nums[i] |= i
        }
    }
    fmt.Println(nums)

    for i, v := range nums {
        fmt.Println(i, v)
    }

    nums2 := [...]int{1,2,3,4,5,6}
    maxIdx := len(nums2) - 1
    for i, e := range nums2 {
        if i == maxIdx {
            nums2[0] += e
        } else {
            nums2[i+1] += e
        }
    }
    fmt.Println(nums2)

    nums3 := []int{1,2,3,4,5,6}
    maxIdx3 := len(nums3) - 1
    for i, e := range nums3 {
        if i == maxIdx3 {
            nums3[0] += e
        } else {
            nums3[i+1] += e
        }
    }
    fmt.Println(nums3)
}
