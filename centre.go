package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t.AppendRows([]table.Row{
		{1, "Jon", "Snow", 30},
		{300, "Tyrion", "Lannister", 5000},
	})
	t.Render()
}
