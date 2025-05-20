package intermediate

import (
	"flag"
	"fmt"
	"os"
)


func main() {

	var stringFlag string
	flag.StringVar(&stringFlag,"user","John","Name of user")

	subcomman1 := flag.NewFlagSet("FirstSub", flag.ExitOnError)
	subcomman2 := flag.NewFlagSet("SecondSub", flag.ExitOnError)

	firstFlag := subcomman1.Bool("processing",false,"Command processing status")
	secondFlag := subcomman1.Int("bytes", 1024,"Byte length of result")

	flagsc2 := subcomman2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcomman1.Parse(os.Args[2:])
		fmt.Println("subcommand1")
		fmt.Println("processing:", firstFlag)
		fmt.Println("bytes:", secondFlag)
	case "secondSub":
		subcomman2.Parse(os.Args[2:])
		fmt.Println("subcommand2:")
		fmt.Println("language:", flagsc2)
	default:
		fmt.Println("no subcommand entered.")
		os.Exit(1)

	}


}