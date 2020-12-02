package main

import (
	"bufio"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"log"
	"os"
)

func main() {
	a:=app.New()
	w := a.NewWindow("LinkCreator")


	linkIn := widget.NewEntry()
	nameIn := widget.NewEntry()

	linkText := widget.NewLabel("Enter Link: ")
	nameText := widget.NewLabel("Enter Link Name: ")

	btn := widget.NewButton("Create", func() {
		write(linkIn.Text,nameIn.Text)
		linkIn.SetText("")
		nameIn.SetText("")
	})
	hor1 := widget.NewHBox(linkText,linkIn)
	hor2 := widget.NewHBox(nameText,nameIn)

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),hor2,hor1,layout.NewSpacer(),btn)
	w.SetContent(container)

	w.Resize(fyne.NewSize(500,150))
	w.ShowAndRun()

}


func write(linkIn string, nameIn string){
	filename := nameIn + ".bat"
	file, err := os.Create(filename)

	content := "start Microsoft-Edge:\""+linkIn + "\""
	writer := bufio.NewWriter(file)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	_ = writer.Flush()




}
