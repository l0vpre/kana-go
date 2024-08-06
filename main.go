package main

import (
    "fmt"
    "math/rand"
)

func randKey(kana map[string]string)string{
    randKey := rand.Intn(len(kana))
    i := 0
    for key:= range kana{

        if (i == randKey){
            return key
        }
        i++
    }
    return "no"
}
func main(){
    kana := map[string]string{
        "a": "あ",
        "i": "い",
        "u": "う",
        "e": "え",
        "o": "お",
    }
    
    key := randKey(kana)
    fmt.Println(kana[key])
    var answer string
    fmt.Scan(&answer)
    if(answer == key){
        fmt.Println("yes")
    } else {
        fmt.Println("no, this: ", key)
    }

    fmt.Print(answer)

}
