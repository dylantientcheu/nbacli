package ui

import (
	"github.com/charmbracelet/bubbles/key"
)

// game help keymap
type GameKM struct {
	Down     key.Binding
	Up       key.Binding
	Previous key.Binding
}

func (k GameKM) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func (k GameKM) ShortHelp() []key.Binding {
	return []key.Binding{k.Down, k.Up, k.Previous}
}

// standing help keymap
type StandingKM struct {
	Down     key.Binding
	Up       key.Binding
	Previous key.Binding
	NextTab  key.Binding
}

func (k StandingKM) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func (k StandingKM) ShortHelp() []key.Binding {
	return []key.Binding{k.NextTab, k.Down, k.Up, k.Previous}
}
