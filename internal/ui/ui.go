package ui

import (
	"github.com/cajun-code/go-resume/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

type ResumeApp struct {
	Resume *models.Resume
	File   string
}

func NewResumeApp() *ResumeApp {
	return &ResumeApp{}
}

func (r ResumeApp) Init() tea.Cmd {

	return nil
}

func (r ResumeApp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return r, nil
}

func (r ResumeApp) View() string {

	return ""
}
