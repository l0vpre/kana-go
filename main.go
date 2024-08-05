package main

import "fmt"

func main(){
    kana := map[string]string{
        "a": "あ",
        "i": "い",
        "u": "う",
        "e": "え",
        "o": "お",
    }

    for key, value := range kana{
        fmt.Println(key, ": ", value)
    }
}
