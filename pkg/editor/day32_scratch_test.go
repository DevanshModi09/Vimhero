package editor

import "testing"

func TestDay32C1(t *testing.T) {
	b := New([]string{"TODO handle parsing", "TODO handle rendering"}, Pos{0, 0})
	feed(b, "m", "a")
	feed(b, "j")
	feed(b, "c", "i", "w")
	feed(b, "D", "O", "N", "E")
	feed(b, "esc")
	feed(b, "'", "a")
	feed(b, "c", "i", "w")
	feed(b, "D", "O", "N", "E")
	feed(b, "esc")
	want := "DONE handle parsing\nDONE handle rendering"
	if b.Text() != want {
		t.Fatalf("c1: got %q want %q", b.Text(), want)
	}
	t.Logf("c1 keystrokes=%d", b.Keystrokes)
}

func TestDay32C2(t *testing.T) {
	b := New([]string{"keep1 remove this whole middle part keep2"}, Pos{0, 6})
	feed(b, "m", "a")
	feed(b, "f", "k")
	feed(b, "d", "'", "a")
	want := "keep1 keep2"
	if b.Text() != want {
		t.Fatalf("c2: got %q want %q", b.Text(), want)
	}
	t.Logf("c2 keystrokes=%d cursor=%v", b.Keystrokes, b.Cursor)
}

func TestDay32C3(t *testing.T) {
	b := New([]string{"template line", "copy: "}, Pos{0, 0})
	feed(b, "w")
	feed(b, "m", "a")
	feed(b, "0")
	feed(b, "y", "'", "a")
	feed(b, "j")
	feed(b, "$")
	feed(b, "p")
	t.Logf("c3 text=%q keystrokes=%d cursor=%v", b.Text(), b.Keystrokes, b.Cursor)
}

func TestDay32C4(t *testing.T) {
	b := New([]string{"TODO alpha section", "middle stuff here", "TODO beta section"}, Pos{0, 0})
	feed(b, "m", "a")
	feed(b, "j", "j")
	feed(b, "m", "b")
	feed(b, "'", "a")
	feed(b, "c", "i", "w")
	feed(b, "D", "O", "N", "E")
	feed(b, "esc")
	feed(b, "'", "b")
	feed(b, "c", "i", "w")
	feed(b, "D", "O", "N", "E")
	feed(b, "esc")
	want := "DONE alpha section\nmiddle stuff here\nDONE beta section"
	if b.Text() != want {
		t.Fatalf("c4: got %q want %q", b.Text(), want)
	}
	t.Logf("c4 keystrokes=%d", b.Keystrokes)
}

func TestDay32C5(t *testing.T) {
	b := New([]string{
		"TODO alpha section",
		"keep1 remove all this middle junk keep2",
		"TODO beta section",
	}, Pos{0, 0})
	feed(b, "m", "a")
	feed(b, "j", "w")
	feed(b, "m", "b")
	feed(b, "2", "f", "k")
	feed(b, "d", "'", "b")
	feed(b, "0", "j")
	feed(b, "c", "i", "w")
	feed(b, "D", "O", "N", "E")
	feed(b, "esc")
	feed(b, "'", "a")
	feed(b, "c", "i", "w")
	feed(b, "D", "O", "N", "E")
	feed(b, "esc")
	t.Logf("c5 text=%q keystrokes=%d", b.Text(), b.Keystrokes)
}
