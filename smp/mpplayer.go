package main

import (
	lib2 "GoStudyNote/smp/lib"
	"GoStudyNote/smp/mp"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lib *lib2.MusicManager
var id int = 1
var ctrl1, signal chan int

//对命令行的处理
func handleLibCommands(tokens []string) {

	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&lib2.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3],tokens[4],tokens[5]})

		}else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3{
			lib.Remove(3)

		}else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:",tokens[1])
	}

}

func handlePlayCommand(tokens []string) {

	if len(tokens) != 2 {
		fmt.Print("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])

	if e == nil {
		fmt.Println("the music", tokens[1], "does not exist.")
		return
	}

	mp.Play(e.Source, e.Type)

}


func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> --play the specified music
	`)

	lib = lib2.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("Enter command -> ")
		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		}else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		}else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}


