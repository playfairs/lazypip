package ui

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jroimartin/gocui"
)

var (
	overlayMutex     sync.Mutex
	overlayVisible   bool
	overlayTimer     *time.Timer
	lastOverlaySizeX int
	lastOverlaySizeY int
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	const minX, minY = 40, 10

	if maxX < minX || maxY < minY {
		if v, err := g.SetView("quarantine", 0, 0, maxX-1, maxY-1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Title = "Terminal too small!"
			v.Wrap = true
			v.Clear()
			v.Write([]byte("Please resize your terminal window to at least 40c x 10r to use lazypip."))
		}
		g.DeleteView("main")
		g.DeleteView("status")
		g.DeleteView("overlay")
		return nil
	}
	g.DeleteView("quarantine")

	overlayMutex.Lock()
	if lastOverlaySizeX != maxX || lastOverlaySizeY != maxY {
		overlayVisible = true
		lastOverlaySizeX = maxX
		lastOverlaySizeY = maxY
		if overlayTimer != nil {
			overlayTimer.Stop()
		}
		overlayTimer = time.AfterFunc(3*time.Second, func() {
			overlayMutex.Lock()
			overlayVisible = false
			overlayMutex.Unlock()
			g.Update(func(*gocui.Gui) error { return nil })
		})
	}
	visible := overlayVisible
	overlayMutex.Unlock()

	if visible {
		overlayMsg := fmt.Sprintf("%dc x %dr", maxX, maxY)
		ox := maxX/2 - len(overlayMsg)/2 - 2
		oy := maxY/2 - 1
		if ox < 0 {
			ox = 0
		}
		if oy < 0 {
			oy = 0
		}
		if v, err := g.SetView("overlay", ox, oy, ox+len(overlayMsg)+3, oy+2); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Frame = false
			v.BgColor = gocui.ColorBlack
			v.FgColor = gocui.ColorWhite
			v.Clear()
			v.Write([]byte(overlayMsg))
		}
	} else {
		g.DeleteView("overlay")
	}

	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "lazypip"
		v.Wrap = true
	}
	if v, err := g.SetView("status", 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Status"
		v.Wrap = true
		v.BgColor = gocui.ColorBlue
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Start() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	defer func() {
		fmt.Print("\x1b[2J\x1b[H\x1b[3J")
		g.Close()
	}()

	g.SetManagerFunc(layout)

	if err := RegisterKeybindings(g); err != nil {
		log.Panicln(err)
	}

	g.Mouse = false
	g.InputEsc = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
