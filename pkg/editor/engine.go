package editor

func (b *Buffer) Input(key string) {
	if key == "" {
		return
	}
	b.Keystrokes++
	b.KeyLog = append(b.KeyLog, key)
	b.StatusMsg = ""

	if b.recordingReg != 0 && !(b.Mode == ModeNormal && key == "q" && b.pendingKind == 0 && b.pendingOp == 0 && b.count == "") {
		b.macroBuf.WriteString(key + "\x00")
	}

	switch b.Mode {
	case ModeNormal:
		b.handleNormal(key)
	case ModeInsert:
		b.handleInsert(key)
	case ModeVisual, ModeVisualLine:
		b.handleVisual(key)
	case ModeCommand:
		b.handleCommandMode(key)
	}
	b.clampCursor()
}
