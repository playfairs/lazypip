package ui

import (
	"strings"

	"github.com/jroimartin/gocui"
)

func setStatus(g *gocui.Gui, msg string) {
	if v, err := g.View("status"); err == nil {
		v.Clear()
		v.Write([]byte(msg))
	}
}

func enterCommandLineMode(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	boxWidth := maxX * 2 / 3
	if boxWidth < 30 {
		boxWidth = maxX - 2
	}
	startX := (maxX - boxWidth) / 2
	endX := startX + boxWidth
	startY := maxY - 5
	endY := maxY - 2

	if v, err := g.SetView("cmdline", startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Command"
		v.Editable = true
		v.Clear()
		g.SetCurrentView("cmdline")
	}
	if err := g.SetKeybinding("cmdline", gocui.KeyEnter, gocui.ModNone, executeCommand); err != nil {
		return err
	}
	if err := g.SetKeybinding("cmdline", gocui.KeyEsc, gocui.ModNone, cancelCommandMode); err != nil {
		return err
	}
	return nil
}

func cancelCommandMode(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView("cmdline")
	g.SetCurrentView("main")
	return nil
}

func executeCommand(g *gocui.Gui, v *gocui.View) error {
	cmd := strings.TrimSpace(v.Buffer())
	g.DeleteView("cmdline")
	g.SetCurrentView("main")

	if cmd == "q" || cmd == ":q" {
		g.Update(func(g *gocui.Gui) error {
			return gocui.ErrQuit
		})
		return gocui.ErrQuit
	}

	setStatus(g, "Unknown command: "+cmd)
	return nil
}
