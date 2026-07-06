package main

import (
	"fmt"

	"vimhero/internal/ui"
)

func main() {
	m := ui.NewModel()
	fmt.Println(m.View())
}
