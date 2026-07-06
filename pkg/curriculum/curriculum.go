package curriculum

type Kind int

const (
	KindGoal Kind = iota
	KindEdit
)

type Pos struct {
	Row, Col int
}

type Challenge struct {
	Title        string
	Instructions string
	Tip          string
	NewKeys      []string
	Start        []string
	CursorStart  Pos
	Kind         Kind
	GoalPos      Pos
	Target       []string
	Par          int
}

type Day struct {
	Number     int
	Week       string
	Title      string
	Summary    string
	Challenges []Challenge
}

func Stars(keystrokes, par int) int {
	if par <= 0 {
		return 3
	}
	switch {
	case keystrokes <= par:
		return 3
	case keystrokes <= par*2:
		return 2
	default:
		return 1
	}
}

func All() []Day {
	return days
}

func ByNumber(n int) (Day, bool) {
	for _, d := range days {
		if d.Number == n {
			return d, true
		}
	}
	return Day{}, false
}

func Total() int {
	return len(days)
}
