package main

import (
    "fmt"
    "math/rand"
)

type Stats struct {
    Answer int
    CorrectAnswer int
    WrongAnswer map[string]bool
}

func(stats Stats) StatsToString()string{
   return  fmt.Sprintf("\nAnswer: %d \nCorrect Answer: %d\n",
                        stats.Answer,stats.CorrectAnswer)
}

func randKey(kana map[string]string)string{
    randKey := rand.Intn(len(kana))
    i := 0
    for key := range kana{
        if (i == randKey){
            return key
        }
        i++
    }
    return "no"
}

func main(){
    var answer string
    kana := map[string]string{
        "a": "あ",
        "i": "い",
        "u": "う",
        "e": "え",
        "o": "お",
    }
    stats := Stats{
        Answer: 0,
        CorrectAnswer: 0,
        WrongAnswer: make(map[string]bool),
    }

    for{
        key := randKey(kana)
        fmt.Println("What this kana: ", kana[key], "?")
        fmt.Scan(&answer)

        if(answer == "exit"){
            fmt.Print("Wrong kana:\n")
            for key:= range stats.WrongAnswer{
                fmt.Println(key, " ", kana[key])
            }
            return
        }

        if(answer == key){
            fmt.Println("yes")
            stats.CorrectAnswer++
        } else {
            fmt.Println("no, this: ", key)
            stats.WrongAnswer[key] = true
        }

        stats.Answer++
        fmt.Println(stats.StatsToString())
    }
}   



