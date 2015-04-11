package commands

import (
	"bytes"
	"fmt"
	"unicode"
	"github.com/mitchellh/cli"
)

type ExpandIdCommand struct {}

func NewExpandIdCommand() (cli.Command, error) {
	return &ExpandIdCommand{}, nil
}

func (*ExpandIdCommand) Help() string {
	return "Converts the passed in 15 character ID to an 18 character ID"
}

func (*ExpandIdCommand) Run(args []string) int {
	fmt.Printf(Convert15CharIdTo18CharId(args[0]))
	return 0
}

func (*ExpandIdCommand) Synopsis() string {
	return "Converts the passed in 15 character ID to an 18 character ID"
}

// Converts the passed in string from a 15 character id
// to an 18 character ID.  
func Convert15CharIdTo18CharId(sfdcId string) string {
	triplet := make([]string, 3)
	triplet[0] = ReverseChunk(sfdcId[0:5])
	triplet[1] = ReverseChunk(sfdcId[5:10])
	triplet[2] = ReverseChunk(sfdcId[10:15])

	var res bytes.Buffer
	res.WriteString(sfdcId)

	for i, _ := range triplet {
		mask := SFDCBitMaskForChunk(triplet[i])
		res.WriteString(Lookup[mask])
	}

	return res.String()
}

// Reverses the passed in string and returns the result.
func ReverseChunk(chunk string) string {
	runes := make([]rune, len(chunk))
	index := len(chunk) -1

	for _, rune := range chunk {
		runes[index] = rune
		index--		
	}

	return string(runes)
}

// Returns an integer that matches the bitmask of the passed in string
// as described by the algorithm that generates the remaining three characters
// in a Salesforce ID.
func SFDCBitMaskForChunk(chunk string) uint {
	var res uint = 0x00000 
	index := len(chunk) - 1

	for i, rune := range chunk {
		var tmp uint = 0x00000

		if unicode.IsUpper(rune) {
			tmp = 1 << ((uint)(index - i))
		}

		res |= tmp 
	}

	return res
}

// A lookup table that is used by the salesforce ID conversion algorithm
// to determine what the final three characters of the 18 character ID are.
var Lookup = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
}
