package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/mdp/qrterminal"
)

// Participant represents a tech workshop participant
type Participant struct {
	Initials string `csv:"Initialen"`
	Username string `csv:"Username"`
	Name     string `csv:"Name"`
	Unit     string `csv:"Abteilung"`
}

func main() {
	helloworld := "Hello World!"

	fmt.Println(helloworld)

	file, err := os.OpenFile("participants.csv", os.O_RDONLY, os.ModePerm)

	if err != nil {
		panic(err)
	}

	var participants []Participant

	gocsv.UnmarshalFile(file, &participants)

	qrcode := ""

	for index, participant := range participants {
		fmt.Printf("%02d %-30s %-10s %-15s (%s)\n", index+1, participant.Name, participant.Unit, participant.Username, participant.Initials)
		qrcode += fmt.Sprintln(participant.Name)
	}

	printQRcode(qrcode)

}

func printQRcode(s string) {
	qrterminal.Generate(s, qrterminal.L, os.Stdout)
}
