package main
import "fmt"

type ReadWrite interface {
    Cout(s string)
}

type Person struct {
    name string
}

func (_ Person) Cout (s string) {
    fmt.Println("xxxxx")
}

func main(){
    p := new(Person)
    p.Cout("xxxx")

    rw := [1]ReadWrite{p}
    rw[0].Cout("xx")

    var fd ReadWrite
    fd = p
    fd.Cout("111")

    _, ok := fd.(*Person)
    fmt.Println(ok)
    if ok {
        panic("test")
    }


}
