package main

import "github.com/rivo/tview"

func DrawBox() {
    box := tview.NewBox().
        SetBorder(true).
        SetTitle("Box Demo")
    if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
        panic(err)
    }
}

func DrawList() {
    app := tview.NewApplication()
    list := tview.NewList().
        AddItem("List item 1", "Some explanatory text", 'a', nil).
        AddItem("List item 2", "Some explanatory text", 'b', nil).
        AddItem("List item 3", "Some explanatory text", 'c', nil).
        AddItem("List item 4", "Some explanatory text", 'd', nil).
        AddItem("Quit", "Press to exit", 'q', func() {
            app.Stop()
        })
    if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
        panic(err)
    }
}
