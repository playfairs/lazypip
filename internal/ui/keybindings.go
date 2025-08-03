package ui

import (
	"github.com/jroimartin/gocui"
)

func installPackage(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func installRequirements(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func upgradePackage(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func upgradeAllPackages(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func showPackageInfo(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func RegisterKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'i', gocui.ModNone, installPackage); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlI, gocui.ModNone, installRequirements); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'u', gocui.ModNone, upgradePackage); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlU, gocui.ModNone, upgradeAllPackages); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 's', gocui.ModNone, showPackageInfo); err != nil {
		return err
	}
	if err := g.SetKeybinding("", ':', gocui.ModNone, enterCommandLineMode); err != nil {
		return err
	}
	return nil
}