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

func(stats Stats) PrintStats() string {
   return  fmt.Sprintf("\nAnswer: %d \nCorrect Answer: %d\n",
                        stats.Answer,stats.CorrectAnswer)
}

const(
    Main = iota
    Game
    Statistics
)

func GetRandomKey(kana map[string]string) string {
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
    kana := map[string]map[string]string{
        "a": {
            "a": "あ",   
            "i": "い",
            "u": "う",
            "e": "え",
            "o": "お",
        },
        "ka": {
            "ka": "か",
            "ki": "き",
            "ku": "く",
            "ke": "け",
            "ko": "こ",
        },
        "sa": {
            "sa": "さ",
            "shi": "し",
            "su": "す",
            "se": "せ",
            "so": "そ",
        },
    }

    stats := Stats{
        Answer: 0,
        CorrectAnswer: 0,
        WrongAnswer: make(map[string]bool),
    }

    currentState := Main    
    gameKana := make(map[string]string)

    fmt.Println("KANA-OO")
    for{
        switch currentState{
            case Main:
//              fmt.Println("Select packs\nAvailable: a, ka, sa")
                for _, value:= range kana{
                        for k, v := range value{
                            gameKana[k] = v
                        }
                }
                currentState = Game
                continue

            case Game:
                key := GetRandomKey(gameKana)
                fmt.Println("What this kana: ", gameKana[key], "?")
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
                fmt.Println(stats.PrintStats())

            case Statistics:
                fmt.Print("Wrong kana:\n")
                for key:= range stats.WrongAnswer{
                    fmt.Println(key, " ", gameKana[key])
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



