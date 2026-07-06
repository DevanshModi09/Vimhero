package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Name    string
	Accent  string
	Good    string
	Warn    string
	Muted   string
	Locked  string
	Cursor  string
	Goal    string
	Command string
	Visual  string
	Star    string
	Tip     string
	Streak  string
}

var themes = []Theme{
	{Name: "Violet", Accent: "99", Good: "42", Warn: "214", Muted: "248", Locked: "244", Cursor: "235", Goal: "214", Command: "39", Visual: "135", Star: "220", Tip: "45", Streak: "208"},
	{Name: "Ocean", Accent: "39", Good: "42", Warn: "214", Muted: "248", Locked: "244", Cursor: "235", Goal: "214", Command: "75", Visual: "111", Star: "220", Tip: "51", Streak: "208"},
	{Name: "Forest", Accent: "28", Good: "82", Warn: "214", Muted: "248", Locked: "244", Cursor: "235", Goal: "214", Command: "37", Visual: "108", Star: "220", Tip: "115", Streak: "166"},
	{Name: "Sunset", Accent: "208", Good: "42", Warn: "220", Muted: "248", Locked: "244", Cursor: "235", Goal: "220", Command: "209", Visual: "203", Star: "228", Tip: "216", Streak: "202"},
	{Name: "Rose", Accent: "205", Good: "42", Warn: "214", Muted: "248", Locked: "244", Cursor: "235", Goal: "214", Command: "212", Visual: "168", Star: "220", Tip: "218", Streak: "204"},
	{Name: "Monochrome", Accent: "252", Good: "255", Warn: "247", Muted: "248", Locked: "244", Cursor: "235", Goal: "255", Command: "250", Visual: "245", Star: "255", Tip: "248", Streak: "252"},
	{Name: "Dracula", Accent: "141", Good: "84", Warn: "215", Muted: "248", Locked: "244", Cursor: "235", Goal: "215", Command: "117", Visual: "212", Star: "228", Tip: "117", Streak: "203"},
	{Name: "Nord", Accent: "67", Good: "108", Warn: "179", Muted: "249", Locked: "244", Cursor: "235", Goal: "179", Command: "110", Visual: "139", Star: "222", Tip: "152", Streak: "173"},
	{Name: "Solarized", Accent: "37", Good: "64", Warn: "136", Muted: "248", Locked: "244", Cursor: "235", Goal: "136", Command: "33", Visual: "61", Star: "136", Tip: "37", Streak: "160"},
	{Name: "Cyberpunk", Accent: "201", Good: "46", Warn: "226", Muted: "248", Locked: "244", Cursor: "235", Goal: "226", Command: "51", Visual: "165", Star: "226", Tip: "51", Streak: "202"},
}

var currentThemeIdx = 0

func applyTheme(i int) {
	if i < 0 || i >= len(themes) {
		i = 0
	}
	currentThemeIdx = i
	t := themes[i]
	colorAccent = lipgloss.Color(t.Accent)
	colorGood = lipgloss.Color(t.Good)
	colorWarn = lipgloss.Color(t.Warn)
	colorMuted = lipgloss.Color(t.Muted)
	colorLocked = lipgloss.Color(t.Locked)
	colorCursor = lipgloss.Color(t.Cursor)
	colorGoal = lipgloss.Color(t.Goal)
	colorCommand = lipgloss.Color(t.Command)
	colorVisual = lipgloss.Color(t.Visual)
	colorStar = lipgloss.Color(t.Star)
	colorTip = lipgloss.Color(t.Tip)
	colorStreak = lipgloss.Color(t.Streak)
	buildStyles()
}

func themeSwatch(hex string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(hex)).Bold(true).Render("●")
}
