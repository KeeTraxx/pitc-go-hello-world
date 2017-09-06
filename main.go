package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/mdp/qrterminal"
)

func main() {
	// Variable assingment
	// Note: type is implicitely set
	helloworld := "Hello World!"

	fmt.Println(helloworld)

	// Open a file for read operations
	file, err := os.OpenFile("participants.csv", os.O_RDONLY, os.ModePerm)

	// Error handling, the Go idiomatic way
	if err != nil {
		panic(err)
	}

	// Declare a slice
	var participants []Participant

	gocsv.UnmarshalFile(file, &participants)

	qrcode := ""

	// loop through slice
	for index, participant := range participants {
		// Call a struct method
		participant.PrettyPrint(index)

		// Append to string
		qrcode += fmt.Sprintln(participant.Name)
	}

	// Call a package method
	printQRcode(qrcode)

}

// Non-exported package method
func printQRcode(s string) {
	qrterminal.Generate(s, qrterminal.L, os.Stdout)
}

// Participant represents a tech workshop participant
type Participant struct {
	Initials string `csv:"Initialen"`
	Username string `csv:"Username"`
	Name     string `csv:"Name"`
	Unit     string `csv:"Abteilung"`
}

// PrettyPrint prints the Participant prettily to stdOut
func (p *Participant) PrettyPrint(index int) {
	fmt.Printf(
		"%02d %-30s %-10s %-15s (%s)\n",
		index+1,
		p.Name,
		p.Unit,
		p.Username,
		p.Initials,
	)
}
