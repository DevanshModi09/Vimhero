package editor

import "testing"

func TestDbgC2(t *testing.T) {
	b := New([]string{"keep1 remove this whole middle part keep2"}, Pos{0, 6})
	feed(b, "m", "a")
	t.Logf("after ma cursor=%v marks=%v", b.Cursor, b.marks)
	feed(b, "f", "k")
	t.Logf("after fk cursor=%v", b.Cursor)
	feed(b, "d", "'", "a")
	t.Logf("after d'a cursor=%v text=%q", b.Cursor, b.Text())
}
