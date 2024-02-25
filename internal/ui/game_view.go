package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type GameView struct {
}

func InitialGameView() *GameView {
	return &GameView{}
}

func (m *GameView) Init() tea.Cmd {
	return nil
}

func (m *GameView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *GameView) View() string {
	return ""
}
