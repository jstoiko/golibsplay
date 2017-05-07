/*
    For building into JS with GopherJS
    > gopherjs build myjslib.go
*/

package main

import "github.com/gopherjs/gopherjs/js"

type Vertex struct {
    A, B int
}

func ReturnGoStr(s string) string {
    return s
}

func ReturnInt(s int) int {
    return s
}

func ReturnMap(m map[int]int) map[int]int {
    return m
}

func GetLen(s string) int {
    return len(s)
}

func main() {
    js.Module.Set("exports", map[string]interface{}{ // node
    // js.Global.Set("mylib", map[string]interface{}{ // browser
        "ReturnGoStr": ReturnGoStr,
        "ReturnInt": ReturnInt,
        "ReturnMap": ReturnMap,
        "GetLen": GetLen,
    })
}
