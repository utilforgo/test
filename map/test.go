package main

import(
    "fmt"
    )
func main(){
    fmt.Println("this is a test for map")
    mapTest()
}

// 初始化
// 赋值
// 遍历
// 判断元素是否存在
// 删除元素

func mapTest(){
    var mapA map[string]string
    //mapA["1"] = "2" // panic: assignment to entry in nil map
     mapA = make(map[string]string)
     mapA["2"] = "3"

    var mapB = make(map[string]string)
     mapB["a"] = "b"

    fmt.Println(mapA)
    fmt.Println(mapB)

    for key := range mapA {
        v := mapA[key]
        fmt.Println(key, v)
    }

    iterm, ok := mapA["2"]
    if (ok) {
        fmt.Println("iterm exist ", iterm)
    }else{
        fmt.Println("iterm not exist ", iterm)
    }

    iterm2, ok := mapA["1"]
    if (ok) {
        fmt.Println("iterm exist ", iterm2)
    }else{
        fmt.Println("iterm not exist ", iterm2)
    }

    delete (mapA, "2")
    fmt.Println(mapA)
}



