package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorAccent  = lipgloss.Color("99")
	colorGood    = lipgloss.Color("42")
	colorWarn    = lipgloss.Color("214")
	colorMuted   = lipgloss.Color("240")
	colorLocked  = lipgloss.Color("237")
	colorCursor  = lipgloss.Color("235")
	colorGoal    = lipgloss.Color("214")
	colorCommand = lipgloss.Color("39")
	colorVisual  = lipgloss.Color("135")
	colorStar    = lipgloss.Color("220")
	colorTip     = lipgloss.Color("45")
	colorStreak  = lipgloss.Color("208")

	titleStyle             lipgloss.Style
	subtitleStyle          lipgloss.Style
	dayLineStyle           lipgloss.Style
	dayLineSelected        lipgloss.Style
	markerSelectedStyle    lipgloss.Style
	dayLineCleared         lipgloss.Style
	lockedStyle            lipgloss.Style
	challengeTitleStyle    lipgloss.Style
	challengeSelectedStyle lipgloss.Style
	starStyle              lipgloss.Style
	dimStyle               lipgloss.Style
	goodStyle              lipgloss.Style
	boxStyle               lipgloss.Style
	winBoxStyle            lipgloss.Style
	helpStyle              lipgloss.Style
	streakStyle            lipgloss.Style
	modeNormal             lipgloss.Style
	modeInsert             lipgloss.Style
	modeVisual             lipgloss.Style
	modeCmd                lipgloss.Style
	cursorCellStyle        lipgloss.Style
	goalCellStyle          lipgloss.Style
	tipStyle               lipgloss.Style
	newKeyStyle            lipgloss.Style
)

func init() {
	buildStyles()
}

func buildStyles() {
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(colorAccent).
		Padding(0, 2)

	subtitleStyle = lipgloss.NewStyle().Foreground(colorMuted).Italic(true)

	dayLineStyle = lipgloss.NewStyle().PaddingLeft(2)
	dayLineSelected = lipgloss.NewStyle().Bold(true).Foreground(colorAccent)
	markerSelectedStyle = lipgloss.NewStyle().Bold(true).Foreground(colorAccent)
	dayLineCleared = lipgloss.NewStyle().Bold(true)
	lockedStyle = lipgloss.NewStyle().Foreground(colorLocked)

	challengeTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	challengeSelectedStyle = lipgloss.NewStyle().Bold(true).Foreground(colorAccent)

	starStyle = lipgloss.NewStyle().Foreground(colorStar).Bold(true)
	dimStyle = lipgloss.NewStyle().Foreground(colorMuted)
	goodStyle = lipgloss.NewStyle().Foreground(colorGood).Bold(true)
	boxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(colorAccent)
	winBoxStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(0, 2).BorderForeground(colorGood)
	helpStyle = lipgloss.NewStyle().Foreground(colorMuted)
	streakStyle = lipgloss.NewStyle().Bold(true).Foreground(colorStreak)

	modeNormal = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255")).Background(lipgloss.Color("62")).Padding(0, 1)
	modeInsert = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("0")).Background(colorGood).Padding(0, 1)
	modeVisual = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255")).Background(colorVisual).Padding(0, 1)
	modeCmd = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255")).Background(colorCommand).Padding(0, 1)

	cursorCellStyle = lipgloss.NewStyle().Reverse(true)
	goalCellStyle = lipgloss.NewStyle().Foreground(colorGoal).Bold(true)

	tipStyle = lipgloss.NewStyle().Foreground(colorTip)
	newKeyStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("0")).Background(colorAccent).Padding(0, 1)
}
