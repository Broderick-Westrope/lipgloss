package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/list"
)

func main() {
	t := list.New(
		"Documents",
		"Downloads",
		"Unfinished Projects",
	).Title("~").Enumerator(list.Tree)

	fmt.Println("A classic tree:\n" + t.String() + "\n")

	tr := list.New(
		"Documents",
		"Downloads",
		"Unfinished Projects",
	).Title("~").Enumerator(list.TreeRounded)

	fmt.Println("A cool, rounded tree:\n" + tr.String() + "\n")

	ti := list.New(
		list.New(
			"Important Documents",
			"Junk Drawer",
			"Books",
		).Title("Documents").Enumerator(list.Tree),
		"Downloads",
		"Unfinished Projects",
	).Title("~").Enumerator(list.Tree).Indenter(list.TreeIndenter)

	fmt.Println("A fancy, nested tree:\n" + ti.String() + "\n")
}
