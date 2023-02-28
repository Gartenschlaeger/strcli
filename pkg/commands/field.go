package commands

import (
	"flag"
)

type fieldCommandsFlagValues struct {
	index     *int
	separator *string
}

func FieldCommand(args []string) error {
	fv := fieldCommandsFlagValues{}

	fs := flag.NewFlagSet("field", flag.ExitOnError)
	fv.index = fs.Int("index", 0, "Zero based index of the field.")
	fv.separator = fs.String("separator", " ", "Separator character between fields.")

	err := fs.Parse(args)
	if err != nil {
		panic(err)
	}

	fs.PrintDefaults()

	return nil
}

// func readStdin() (*string, error) {
// 	reader := bufio.NewReader(os.Stdin)

// 	var output string
// 	for {
// 		input, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			return nil, err
// 		}

// 		output += input
// 	}

// 	return &output, nil
// }

// //
// // Returns a field by index and separator
// // e.g. handleField("Hello world", 2, " ") -> "world"
// //
// func handleField(input string, fieldIndex int, separator rune) {

// 	fmt.Println(separator)

// 	startIndex := -1
// 	length := 0

// 	fieldCounter := 0
// 	for i, c := range input {
// 		if c == separator {
// 			fieldCounter += 1
// 			if fieldCounter == fieldIndex {
// 				startIndex = i + 1
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(startIndex, length)

// }
