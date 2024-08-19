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
    wrongKanas []Kana
}

func (stats *Stats) AddCorrectAnswer() {
    stats.correctAnswers++
    stats.answers++
}

func (stats *Stats) AddWrongAnswer (kana Kana) {
    stats.wrongKanas = append(stats.wrongKanas, kana)
    stats.answers ++
}

func (stats *Stats) GetWrongAnswerCount() int {
    return stats.answers - stats.correctAnswers
}

func (stats *Stats) PrintStats() {
        fmt.Printf("\nCorrect: %d/%d\n",
        stats.correctAnswers, stats.answers)
}

func (stats *Stats) PrintWrongKanas() {
    for _, kana := range stats.wrongKanas{
        fmt.Printf(" %s (%s)\n", kana.Letter, kana.Transcription)
    }
}

func NewStats() Stats {
    return Stats{
        answers: 0,
        correctAnswers: 0,
        wrongKanas: make([]Kana, 0),
    }
}

func (stats *Stats) Reset() {
    stats.answers = 0
    stats.correctAnswers = 0
    stats.wrongKanas = make([]Kana, 0)
}

func (stats *Stats) GetWrongAnswer() []Kana {
    return stats.wrongKanas
}

type Kana struct {
    Letter string
    Transcription string
}


type Row struct {
    Name string
    Kanas []Kana
}

type GameDictionary struct {
    kanas []Kana
}


func (dict *GameDictionary) PopRandomKana() Kana {
    kana := dict.kanas[len(dict.kanas)-1]
    dict.kanas = dict.kanas[0:len(dict.kanas)-1]
    return kana
}

func (gameDict *GameDictionary) Len() int {
    return len(gameDict.kanas)
}

type Dictionary struct {
    Rows []Row
}

func (dict *Dictionary) GetRow(key string) (Row, bool) {
    for _ , element := range dict.Rows {
        if(element.Name == key) {
            return element, true
        }
    }
    return Row{}, false
}

func (dict *Dictionary) GetSelected(names []string) ([]Row, []string) {
    rows := make([]Row, len(names))
    notFoundNames := make([]string,0)

    for _, element := range names {
        r, b := dict.GetRow(element)
        if(b) {
            rows = append(rows, r)
        } else {
            notFoundNames = append(notFoundNames, element)
        }
    }
    return rows, notFoundNames
}
func NewGameDictionary(kanas []Kana) GameDictionary {

    rand.Shuffle(len(kanas), func(i, j int) { kanas[i], kanas[j] = kanas[j], kanas[i] })

    return  GameDictionary{
        kanas: kanas,
    }
}

func NewGameDictionaryFromRows(rows []Row) GameDictionary {
    kanas := make([]Kana, 0)

    for _, row := range rows{
        kanas = append(kanas, row.Kanas...)
    }

    return NewGameDictionary(kanas)
}

const(
    Menu = iota
    Game
    Statistics
)

func main(){
    var input string
    var gameDict GameDictionary
    reader := bufio.NewReader(os.Stdin)

    dictionary := Dictionary{
        Rows: []Row{
        { Name: "a", Kanas: []Kana{
             {"あ", "a",},
             {"い", "i",},
             {"う", "u",},
             {"え", "e",},
             {"お", "o",},
                },
            },
         { Name: "ka", Kanas: []Kana{
             {"か", "ka",},
             {"き", "ki",},
             {"く", "ku",},
             {"け", "ke",},
             {"こ", "ko",},
                },
            },
         { Name: "sa", Kanas: []Kana{
             {"さ", "sa",},
             {"し", "shi",},
             {"す", "su",},
             {"せ", "se",},
             {"そ", "so",},
                },
            },
        },
    }

    stats := NewStats()
    currentState := Menu

    fmt.Printf("KANA-GO\n")
    fmt.Printf("Press s for start!\n")
    input, _ = reader.ReadString('\n')
    input = strings.TrimSpace(input)

    for{
        switch currentState{
            case Menu:
                fmt.Printf("What do you want play?\nAll Kana: a\nSelect rows: s\n")

                input, _ = reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if(input == "a"){
                    gameDict = NewGameDictionaryFromRows(dictionary.Rows)
                } else if(input == "s") {
                    fmt.Printf("Available rows: %s", dictionary.Rows[0].Name)
                for i := 1 ; i < len(dictionary.Rows) ;i++{
                        fmt.Printf(", %s", dictionary.Rows[i].Name)
                    }
                    fmt.Printf(".\n")

                    input, _ = reader.ReadString('\n')
                    input = strings.TrimSpace(input)
                    args := strings.Split(input, " ")
                    rows, notFound := dictionary.GetSelected(args)

                    if(len(notFound)>0) {
                        fmt.Printf("Not found: ")
                        fmt.Print(strings.Join(notFound, ", "))
                        fmt.Printf(".\n")
                    }

                    gameDict = NewGameDictionaryFromRows(rows)
                    fmt.Printf("\n")
                    }

                currentState = Game

            case Game:
                if (gameDict.Len() == 0) {
                    currentState = Statistics
                    continue
                }

                kana := gameDict.PopRandomKana()
                fmt.Printf("What is this kana: %s? (%d left)\n",kana.Letter, gameDict.Len() + 1)

                input, _ = reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if(input == "exit"){
                    currentState = Statistics
                    continue
                }

               if(input == kana.Transcription){
                    stats.AddCorrectAnswer()
                } else {
                fmt.Println("no, it's a: ", kana.Transcription)
                    stats.AddWrongAnswer(kana)
                }

                stats.PrintStats()

            case Statistics:
                fmt.Printf("\nThe game is over. Your stats: ")
                stats.PrintStats()

                if( stats.GetWrongAnswerCount() > 0) {
                    fmt.Printf("Wrong kana:\n")
                    stats.PrintWrongKanas()
                }
                

                fmt.Printf("\nNew Game: n\nExit: exit\n")

                input, _ = reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if(input == "n"){
                    if(stats.GetWrongAnswerCount()> 0) {
                        fmt.Printf("First, correct the errors\n")
                        gameDict = NewGameDictionary(stats.GetWrongAnswer())
                        stats.Reset()
                        currentState = Game
                        continue
                    }
                    stats.Reset()
                    currentState = Menu
                } else if (input == "exit"){
                    return
                }
        }
    }

}


