package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Board struct {
}

func InitialBoard() *Board {
	return &Board{}
}

func (m *Board) Init() tea.Cmd {
	return nil
}

func (m *Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Board) View() string {
	return ""
}
