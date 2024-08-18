package main

import (
    "fmt" 
    "math/rand"
    "bufio"
    "os"
    "strings"
)

type Stats struct {
    answers int
    correctAnswers int
    wrongKanas map[string]string
}

func (s *Stats) InceaseCorrectAnswer() {
    s.correctAnswers ++
}

func (s *Stats) InceaseAnswer () {
    s.answers ++
}

func (s *Stats) AddWrongKana(key string, value string) {
    s.wrongKanas[key] = value
}

func(s *Stats) DeleteWrongKana(key string) {
    delete(s.wrongKanas, key)
}

func(s Stats) GetLenghthWrongKanas() int{
    return len(s.wrongKanas)
}

func (s Stats) GetWrongKanas() map[string]string {
    return s.wrongKanas
}

func (s Stats) MakeGameDictFromWrongKanas()  GameDict {
    gameDict := GameDict{
        kanas: make(map[string]string),
    }

    for key, value := range s.wrongKanas{
        gameDict.kanas[key] = value
    }

    return gameDict
}

func (stats Stats) PrintStats() {
        fmt.Printf("\nCorrect: %d/%d\n",
        stats.correctAnswers, stats.answers)
}

func NewStats(answers int, correctAnswers int, wrongKanas map[string]string) Stats {
    return Stats{
        answers: answers,
        correctAnswers: correctAnswers,
        wrongKanas: wrongKanas,
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

func (row Row) GetKana(key string) string {
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
        if (i == randKey) {
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

func (dict Dictionary) MakeGameDictWithSelected(packs []string) GameDict {
    for i := range packs{
        if(packs[i] == "all"){
            return dict.MakeGameDict()
        }
    }

    gameDict := GameDict{
        kanas: make(map[string]string),
    }

    for i := range packs{
        for key, value := range dict.row{
            if(key == packs[i]){
                value.AppendKanasTo(gameDict.kanas)
            }
        }
    }

    return gameDict
}

func (dict Dictionary) GetRows () string {
    var bilder strings.Builder

    for key := range dict.row {
        bilder.WriteString(key)
        bilder.WriteString(", ")
    }
    result := bilder.String()
    return result
}

const(
    Menu = iota
    Game
    Statistics
)

func main(){
    var input string
    var gameDict GameDict
    reader := bufio.NewReader(os.Stdin)

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

    fmt.Printf("KANA-GO\n")
    fmt.Printf("Press s for start!\n")
    input, _ = reader.ReadString('\n')
    input = strings.TrimSpace(input)
    
    for{
        switch currentState{
            case Menu:
                fmt.Printf("Select packs: all, %s\n", dictionary.GetRows())
                input, _ = reader.ReadString('\n')
                input = strings.TrimSpace(input)

                args := strings.Split(input, " ")
                gameDict = dictionary.MakeGameDictWithSelected(args)

                currentState = Game

            case Game:
                if (gameDict.Len() == 0) {
                    currentState = Statistics
                    continue
                }

                key, value := gameDict.PopRandomKana()
                fmt.Printf("What is this kana: %s? (%d left)\n", value, gameDict.Len() + 1)

                input, _ = reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if(input == "exit"){
                    currentState = Statistics
                    continue
                }

               if(input == key){
                    stats.InceaseCorrectAnswer()
                    stats.DeleteWrongKana(key)
                } else {
                fmt.Println("no, it's a: ", key)
                    stats.AddWrongKana(key, value)
                }

                stats.InceaseAnswer()
                stats.PrintStats()

            case Statistics:
                fmt.Printf("\nThe game is over. Your stats: ")
                stats.PrintStats()
                if( stats.GetLenghthWrongKanas() > 0) {
                    fmt.Printf("Wrong kana:\n")
                    for key, value := range stats.GetWrongKanas(){
                        fmt.Printf("  %s (%s)\n", value, key)
                    }
                }

                fmt.Printf("\nNew Game: n\nExit: exit\n")

                input, _ = reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if(input == "n"){
                    if(stats.GetLenghthWrongKanas()> 0) {
                        fmt.Printf("First, correct the errors\n")
                        stats = NewStats(0,0,stats.GetWrongKanas())
                        gameDict = stats.MakeGameDictFromWrongKanas()
                        currentState = Game
                        continue
                    }
                    stats = NewStats(0,0,make(map[string]string))
                    currentState = Menu
                } else if (input == "exit"){
                    return
                }
        }
    }
}



