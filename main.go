package main

// A simple example demonstrating the use of multiple text input components from the Bubbles component library. import (
import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#0C0C0C", Dark: "#F0F0F0"})
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#2F2F2F", Dark: "#DFDFDF"})
	entryStyle   = focusedStyle.Copy().BorderStyle(lipgloss.RoundedBorder()).Padding(1).Width(40)
	itemStyle    = entryStyle.Copy()
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	itemList   ItemList
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "item"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "amount"
			t.CharLimit = 64
		}

		m.inputs[i] = t
	}

	return m
}
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs)-1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	name := m.inputs[0].Value()
	amount, _ := strconv.ParseFloat(m.inputs[1].Value(), 64)

	m.itemList = Materials(name, amount)
	done := false
	for !done {
		m.itemList, done = Simplify(m.itemList)
	}

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var entryB strings.Builder

	for i := range m.inputs {
		entryB.WriteString(m.inputs[i].View())
		if i != len(m.inputs)-1 {
			entryB.WriteRune('\n')
		}
	}

	entryView := entryStyle.Render(entryB.String()) + "\n"

	keys := make([]string, 0, len(m.itemList))
	for key := range m.itemList {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var itemB strings.Builder

	for _, key := range keys {
		s := fmt.Sprintf("%s: %f", key, m.itemList[key])
		itemB.WriteString(itemStyle.Render(s))
		itemB.WriteRune('\n')
	}

	return entryView + itemB.String()
}

func main() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
