/*
    For building into c-shared lib
    > go build -o mylib.so -buildmode=c-shared mylib.go
*/

package main

/*
typedef struct { long long x; long long y; } Coords;
*/
import "C"
import "fmt"

//export PrintGoStr
func PrintGoStr(s string) {
    fmt.Println(s)
}

//export PrintInt
func PrintInt(s int) {
    fmt.Println(s)
}

//export PrintCStr
func PrintCStr(s *C.char) {
    fmt.Println(C.GoString(s))
}

//export PrintStruct
func PrintStruct(c C.Coords) {
    fmt.Println(c)
}

//export PrintStructMeth
func (c C.Coords) PrintStructMeth() {
    fmt.Println(c)
}

func main() {}

// v-- Not used by C yet --v

//export PrintMap
func PrintMap(m map[int]int) {
    fmt.Println(m)
}
