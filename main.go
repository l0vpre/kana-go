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

func(stats Stats) StatsToString() string {
   return  fmt.Sprintf("\nAnswer: %d \nCorrect Answer: %d\n",
                        stats.Answer,stats.CorrectAnswer)
}

const(
    Main = iota
    Game
    Statistics
)

func randKey(kana map[string]string) string {
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

    currentState := Main

    fmt.Println("KANA-OO\n Press s for start!")
    for{
        switch currentState{
            case Main:
                fmt.Scan(&answer)
                if( answer == "s"){
                    currentState = Game
                }
            case Game:
                key := randKey(kana)
                fmt.Println("What this kana: ", kana[key], "?")
                fmt.Scan(&answer)
                if(answer == "exit"){
                    currentState = Statistics
                    continue
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
            case Statistics:
                fmt.Print("Wrong kana:\n")
                for key:= range stats.WrongAnswer{
                    fmt.Println(key, " ", kana[key])
                }
                fmt.Println("New Game: n\nExit: exit")
                fmt.Scan(&answer)
                if(answer == "n"){
                    stats.Answer = 0
                    stats.CorrectAnswer = 0
                    stats.WrongAnswer = make(map[string]bool)
                    currentState = Game
                } else if (answer == "exit"){
                    return
                }
        }
    }
}



