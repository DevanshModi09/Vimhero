package ui

import "testing"

func TestCurriculumSolvable(t *testing.T) {
	cases := []struct {
		day, challenge int
		keys           []string
	}{
		{1, 0, []string{"j", "j", "j", "l", "l", "l", "l"}},
		{1, 1, []string{"j", "j", "j", "j", "j", "j", "l", "l", "l", "l", "l", "l", "l"}},
		{1, 2, []string{"j", "j", "j", "j", "j", "j", "j", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l"}},
		{1, 3, []string{"k", "k", "k", "k", "k", "h", "h", "h", "h", "h"}},
		{1, 4, []string{"k", "k", "k", "k", "k", "k", "k", "k", "k", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l", "l"}},

		{2, 0, []string{"w", "w", "w", "w", "w", "w", "w", "w"}},
		{2, 1, []string{"b", "b", "b", "b", "b", "b"}},
		{2, 2, []string{"e", "e", "e", "e", "e", "e", "e", "e", "e"}},
		{2, 3, []string{"w", "w", "w", "w", "w", "w", "w", "w"}},
		{2, 4, []string{"b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}},

		{3, 0, append([]string{"i"}, append(typeStr("welcome to vim insert mode"), "esc")...)},
		{3, 1, append([]string{"A"}, append(typeStr(" world of vim editing"), "esc")...)},
		{3, 2, append([]string{"o"}, append(typeStr("second line goes right here"), "esc")...)},
		{3, 3, append([]string{"O"}, append(typeStr("first line"), "esc")...)},
		{3, 4, concatKeys(
			append([]string{"O"}, append(typeStr("appetizer"), "esc")...),
			[]string{"j"},
			append([]string{"A"}, append(typeStr(" dish"), "esc")...),
			append([]string{"o"}, append(typeStr("dessert"), "esc")...),
		)},

		{4, 0, []string{"g", "g"}},
		{4, 1, []string{"G"}},
		{4, 2, []string{"x", "j", "x", "j", "x"}},
		{4, 3, []string{"7", "g", "g"}},
		{4, 4, []string{"k", "k", "x", "j", "j", "j", "x", "G", "x"}},

		{5, 0, []string{"j", "d", "d", "j", "j", "d", "d"}},
		{5, 1, []string{"k", "y", "y", "p", "p"}},
		{5, 2, []string{"y", "y", "P", "P"}},
		{5, 3, []string{"d", "d", "G", "p"}},
		{5, 4, []string{"j", "d", "d", "p", "G", "d", "d", "k", "p"}},

		{6, 0, []string{"d", "w", "d", "w"}},
		{6, 1, concatKeys(
			append([]string{"c", "w"}, append(typeStr("quick"), "esc")...),
			[]string{"w"},
			append([]string{"c", "w"}, append(typeStr("brown"), "esc")...),
		)},
		{6, 2, concatKeys(
			append([]string{"c", "i", "w"}, append(typeStr("word"), "esc")...),
			[]string{"w", "w"},
			append([]string{"c", "i", "w"}, append(typeStr("time"), "esc")...),
		)},
		{6, 3, []string{"d", "i", "w"}},
		{6, 4, concatKeys(
			[]string{"d", "w"},
			[]string{"j", "w"},
			append([]string{"c", "w"}, append(typeStr("quick"), "esc")...),
			[]string{"j", "b"},
			append([]string{"c", "i", "w"}, append(typeStr("word"), "esc")...),
		)},

		{7, 0, []string{"d", "a", "w", "w", "d", "a", "w"}},
		{7, 1, []string{"D", "j", "D"}},
		{7, 2, concatKeys(
			append([]string{"C"}, append(typeStr("you"), "esc")...),
			[]string{"j", "h", "h", "h", "h"},
			append([]string{"C"}, append(typeStr("know you"), "esc")...),
		)},
		{7, 3, []string{"d", "a", "w", "w", "w", "D"}},
		{7, 4, []string{"d", "a", "w", "j", "b", "D"}},

		{8, 0, []string{"c", "w", "q", "u", "i", "c", "k", "esc", "j", "w", "c", "w", "o", "v", "e", "r", "esc"}},
		{8, 1, []string{"d", "d", "g", "g", "P"}},
		{8, 2, []string{"A", " ", "w", "o", "r", "l", "d", "esc", "G", "o", "S", "e", "e", " ", "y", "o", "u", " ", "s", "o", "o", "n", "esc"}},
		{8, 3, []string{"c", "w", "q", "u", "i", "c", "k", "esc", "w", "w", "w", "w", "c", "w", "o", "v", "e", "r", "esc"}},
		{8, 4, concatKeys(
			[]string{"d", "d", "p"},
			[]string{"j"},
			append([]string{"c", "w"}, append(typeStr("quick"), "esc")...),
			[]string{"j", "w"},
			[]string{"d", "a", "w"},
			[]string{"j", "w", "w"},
			[]string{"D"},
		)},
	}

	m := NewModel()
	for _, tc := range cases {
		dayIdx := tc.day - 1
		day := m.days[dayIdx]
		mm := m.enterChallenge(dayIdx, tc.challenge)
		for _, k := range tc.keys {
			mm.buf.Input(k)
		}
		if !mm.checkWin() {
			t.Errorf("day %d challenge %d (%s) not solved by scripted keys %v\n  got lines: %q\n  got cursor: %+v\n  want: kind=%v target=%v goal=%+v",
				tc.day, tc.challenge, day.Challenges[tc.challenge].Title, tc.keys,
				mm.buf.Lines, mm.buf.Cursor,
				mm.challenge.Kind, mm.challenge.Target, mm.challenge.GoalPos)
			continue
		}
		want := day.Challenges[tc.challenge].Par
		if len(tc.keys) != want {
			t.Errorf("day %d challenge %d (%s): Par=%d but scripted minimal solve took %d keystrokes",
				tc.day, tc.challenge, day.Challenges[tc.challenge].Title, want, len(tc.keys))
		}
	}
}

func concatKeys(groups ...[]string) []string {
	out := []string{}
	for _, g := range groups {
		out = append(out, g...)
	}
	return out
}

func typeStr(s string) []string {
	out := make([]string, 0, len(s))
	for _, r := range s {
		out = append(out, string(r))
	}
	return out
}
