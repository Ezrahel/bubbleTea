package main

import (
	"bubblecli/cli"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// var mod cli.Model
func main() {
	store := &cli.Store{}
	if err := store.Init(); err != err {
		log.Fatalf("unable to init store: %v", err)
	}
	mod := cli.NewModel(store)
	p := tea.NewProgram(mod)
	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run tui: %v", err)
	}
}
