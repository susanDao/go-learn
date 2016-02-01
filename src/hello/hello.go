package main

import (
    "fmt"
    "more"
    "unicode/utf8"
)


const(
    x = iota
    y = iota
)

const c_i int = 2
func p(arr []int) (x int) {
    fmt.Println(arr)
    x=0
    return
}

func main() {
    var arr [5]int 
    arr = [5]int{1,2,3,4,5}
    for i:=0; i<len(arr); i++ {
        fmt.Println(arr[i])
    }

    var pt *int
    pt = &arr[0]
    fmt.Println(*pt)

    var sl []int = arr[0:1]
    fmt.Println(sl)

    ap := new([5]int)
    fmt.Println(*ap)
    fmt.Println(ap[1])

    mk := make([]int,0)
    mk = append(mk, 100)
    fmt.Println(mk)

    for _,item := range arr {
        item *= 2
    }
    fmt.Println(arr)

    sa := []int{1,2}
    sa = append(sa, 3)
    fmt.Println(sa)

    s := "hello"
    fmt.Println(s[0])
    c := []byte(s)
    c[0] = 'H'
    fmt.Println(c[0])
    p(sa)
    more.Test()
    more.Test1()

    s_r := "你好"
    fmt.Println(len(s_r))
    fmt.Println(s_r[0])

    k_r := []byte("墨迹阿婆") 
    slen := 0
    sizet := 0
    for i:=0;i<len(k_r);i=i+sizet {
        _, sizet = utf8.DecodeLastRune(k_r)
        slen++ 
    }
    fmt.Println(slen)
}
