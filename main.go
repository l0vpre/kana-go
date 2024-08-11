package main

import (
	"fmt"
	"math/rand"
)

type Stats struct {
    Answers int
    CorrectAnswers int
    WrongKanas map[string]string
}

func (stats Stats) PrintStats() {
        fmt.Printf("\nCorrect: %d/%d\n",
        stats.CorrectAnswers, stats.Answers)
}

func NewStats(answers int, correctAnswers int, wrongKanas map[string]string) Stats {
    return Stats{
        Answers: answers,
        CorrectAnswers: correctAnswers,
        WrongKanas: wrongKanas,
    }
}

type Row struct {
    kanas map[string]string
}

func NewRow(kanas map[string]string) Row {
    row := Row{
        kanas: make(map[string]string),
    }

    for key, value := range kanas{
        row.kanas[key] = value
    }

    return row
}

func (row Row) GetKana(key string) string{
    return row.kanas[key]
}

func (row Row) AppendKanasTo(kanas map[string]string) {
    for key, value := range row.kanas {
        kanas[key] = value
    }
}

type GameDict struct {
    kanas map[string]string
}

func (gameDict GameDict) PopRandomKana() (string, string) {
    randKey := rand.Intn(len(gameDict.kanas))
    i := 0

    for key, value := range gameDict.kanas{
        if (i == randKey){
            defer delete(gameDict.kanas, key)
            return key, value
        }
        i++
    }

    return "",""
}

func (gameDict GameDict) Len() int {
    return len(gameDict.kanas)
}

type Dictionary struct {
    row map[string]Row
}

func NewDictionary(dictionary map[string]map[string]string) Dictionary {
    dict := Dictionary{
        row: make(map[string]Row),
    }

    for key, value := range dictionary{
        dict.row[key] = NewRow(value)
    }

    return dict
}

func (dict Dictionary) MakeGameDict() GameDict {
    gameDict := GameDict{
        kanas: make(map[string]string),
    }
    for _, value := range dict.row{
        value.AppendKanasTo(gameDict.kanas)
    }

    return gameDict
}

const(
    Menu = iota
    Game
    Statistics
)

func main(){
    var answer string
    var gameDict GameDict

    dictionary := NewDictionary(map[string]map[string]string{
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
    })

    stats := NewStats(0,0, make(map[string]string))
    currentState := Menu

    fmt.Println("KANA-GO")
    for{
        switch currentState{
            case Menu:
                fmt.Scan(&answer)
                if(answer == "s"){
                    gameDict = dictionary.MakeGameDict()
                    currentState = Game
                    continue
                }

            case Game:
                if (gameDict.Len() == 0) {
                    currentState = Statistics
                    continue
                }
                key, value := gameDict.PopRandomKana()
            fmt.Printf("What is this kana: %s? (%d left)\n", value, gameDict.Len() + 1)
                fmt.Scan(&answer)

                if(answer == "exit"){
                    currentState = Statistics
                    continue
                }

               if(answer == key){
                    stats.CorrectAnswers++
                } else {
                fmt.Println("no, it's a: ", key)
                    stats.WrongKanas[key] = value
                }

                stats.Answers++
                stats.PrintStats()

            case Statistics:
                fmt.Printf("\nThe game is over. Your stats: ")
                stats.PrintStats()
                if( len(stats.WrongKanas) > 0) {
                    fmt.Printf("Wrong kana:\n")
                    for key, value := range stats.WrongKanas{
                        fmt.Printf("  %s (%s)\n", value, key)
                    }
                }

                fmt.Printf("\nNew Game: n\nExit: exit")
                fmt.Scan(&answer)

                if(answer == "n"){
                    stats = NewStats(0,0,make(map[string]string))
                    currentState = Game
                } else if (answer == "exit"){
                    return
                }
        }
    }
}



