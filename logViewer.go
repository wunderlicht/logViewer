package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/rivo/tview"
)

const (
	colorStd   = "[white:black]"
	colorError = "[red]"
	colorFatal  = "[yellow:red]"
)

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetBorder(true).SetTitle("Stdin")

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			switch {
			case strings.Contains(line, "ERROR"):
				line = colorError + line + colorStd
			case strings.Contains(line, "FATAL"):
				line = colorFatal + line + colorStd
			}
			textView.Write([]byte(line + "\n"))
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
