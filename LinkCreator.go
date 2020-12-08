/*
 * LinkCreator.go
 * creates batch files that open links with MS Edge
 * Author: Lukas Becker
 * Last Change: 02.12.20, 21:38
 */

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
	a := app.New()                  //new app
	w := a.NewWindow("LinkCreator") //new Window

	linkIn := widget.NewEntry() //text-box for the link
	nameIn := widget.NewEntry() //text-box for the name of the link

	linkText := widget.NewLabel("Enter Link: ")      //description Label
	nameText := widget.NewLabel("Enter Link Name: ") //description Label
	choice := 0
	radio := widget.NewRadio([]string{"Microsoft Edge", "Mozilla Firefox"}, func(s string) {
		if s == "Microsoft Edge" {
			choice = 0
		} else {
			choice = 1
		}
	})

	btn := widget.NewButton("Create", func() { //Button press function
		write(linkIn.Text, nameIn.Text, choice) //call this function, pass the values of the text-boxes
		//reset text-boxes
		linkIn.SetText("")
		nameIn.SetText("")
	})

	//Layout Design
	hor1 := widget.NewHBox(linkText, linkIn) //Line for link
	hor2 := widget.NewHBox(nameText, nameIn) //Line for link name

	//This container spaces the boxes and the button
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), hor2, hor1, layout.NewSpacer(), radio, layout.NewSpacer(), btn)
	w.SetContent(container)

	w.Resize(fyne.NewSize(500, 150)) //fixed size
	w.ShowAndRun()                   //actually show the window

}

func write(linkIn string, nameIn string, browser int) {
	//This function creates a new batch file and adds the command to it
	filename := nameIn + ".bat"      //create a valid file name
	file, err := os.Create(filename) //create the batch file
	var content string
	if browser == 0 {
		content = "start Microsoft-Edge:\"" + linkIn + "\"" //create the command
	} else {
		content = "start firefox " + linkIn
	}

	writer := bufio.NewWriter(file) //make a writer
	if err != nil {
		log.Fatal(err) //Error handling
	}

	_, err = writer.WriteString(content) //write the command into the batch file
	if err != nil {
		log.Fatal(err) //error handling
	}

	_ = writer.Flush() //flush the writer

}
