package main

import (
	"bubblecli/cli"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store := &cli.Store{}
	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)
	}
	mod := cli.NewModel(store)
	p := tea.NewProgram(mod)
	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run tui: %v", err)
	}
}
