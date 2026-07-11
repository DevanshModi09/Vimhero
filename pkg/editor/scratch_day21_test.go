package editor

import (
	"strings"
	"testing"
)

func typeStrS(s string) []string {
	out := make([]string, 0, len(s))
	for _, r := range s {
		out = append(out, string(r))
	}
	return out
}

func run(start []string, cursor Pos, keys []string) *Buffer {
	b := New(start, cursor)
	for _, k := range keys {
		b.Input(k)
	}
	return b
}

func TestScratchDay21(t *testing.T) {
	cases := []struct {
		name   string
		start  []string
		cursor Pos
		keys   []string
		want   []string
	}{
		{
			name:   "ch1",
			start:  []string{"wor1d needs a fix", "jump target line stays put", "extra line to remove completely", "footer text to duplicate", "last line needs a suffix"},
			cursor: Pos{0, 0},
			keys: append(append([]string{"l", "l", "l", "x", "i", "l", "esc", "j", "j", "d", "d", "y", "y", "p", "G", "A"}, typeStrS(" - done")...), "esc"),
			want: []string{"world needs a fix", "jump target line stays put", "footer text to duplicate", "footer text to duplicate", "last line needs a suffix - done"},
		},
	}

	for _, tc := range cases {
		b := run(tc.start, tc.cursor, tc.keys)
		got := b.Text()
		want := strings.Join(tc.want, "\n")
		if got != want {
			t.Errorf("%s: keystrokes=%d\n got: %q\nwant: %q\ncursor: %+v", tc.name, len(tc.keys), b.Lines, tc.want, b.Cursor)
		} else {
			t.Logf("%s OK, %d keystrokes", tc.name, len(tc.keys))
		}
	}
}
