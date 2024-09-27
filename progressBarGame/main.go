package main

/// SPAM A AND D FOR FILLING THE BAR AND WHEN REACH 100% STOP

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type model struct {
	progressBar progress.Model
	progress    float64
	maxProgress float64
	lastKey     string
	decayAmount float64
}

func initialModel() model {
	// Set the progress bar width to 40 characters
	pb := progress.New(
		progress.WithDefaultScaledGradient(),
		progress.WithWidth(70), // Adjust this value to make the bar wider or narrower
	)
	return model{
		progressBar: pb,
		progress:    0,
		maxProgress: 1.0,   // Max is 100%
		lastKey:     "",    // No key pressed yet
		decayAmount: 0.005, // Smaller constant decay amount for smoother decay
	}
}

func (m model) Init() tea.Cmd {
	// Start the ticking for decay, more frequent (every 100 milliseconds)
	return tea.Tick(time.Millisecond*5, func(t time.Time) tea.Msg {
		return t
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// Alternating between 'a' and 'd'
		if (msg.String() == "a" && m.lastKey != "a") || (msg.String() == "d" && m.lastKey != "d") {
			if m.progress < m.maxProgress {
				m.progress += 0.1 // Increment progress by 10% per valid key press
				if m.progress > m.maxProgress {
					m.progress = m.maxProgress
				}
				m.lastKey = msg.String() // Update last key pressed
			}
		}

	case time.Time:
		// Constant decay of 0.005 every 100ms, but no decay if progress is 100%
		if m.progress < m.maxProgress && m.progress > 0 {
			m.progress -= m.decayAmount
			if m.progress < 0 {
				m.progress = 0
			}
		}
		// Schedule the next tick for decay
		return m, tea.Tick(time.Millisecond*5, func(t time.Time) tea.Msg {
			return t
		})
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf(
		"\nAlternate between 'a' and 'd' to increase progress\n%s\nProgress: %.2f%%\n\nPress 'q' to quit.",
		m.progressBar.ViewAs(m.progress),
		m.progress*100,
	)
}

func main() {
    p := tea.NewProgram(initialModel())
    p.Run()
}
