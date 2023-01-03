package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
)

var state *State

func main() {
	var err error

	state, err = NewStateFromFile("./todo.json")
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range state.GetListKeys() {
		log.Println(v)
	}

	go func() {
		w := app.NewWindow(
			app.Title("todo"),
			app.Size(unit.Dp(800), unit.Dp(600)),
		)

		todo := NewTodoApp(w)

		if err := todo.Run(); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}
