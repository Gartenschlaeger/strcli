package commands

import (
	"flag"

	"github.com/Gartenschlaeger/strcli/pkg/strutilities"
)

type fieldCommandsFlagValues struct {
	index     *int
	separator *string
}

func FieldCommand(input string, args []string) (string, error) {
	fv := fieldCommandsFlagValues{}

	fs := flag.NewFlagSet("field", flag.ExitOnError)
	fv.index = fs.Int("index", 0, "Zero based index of the field.")
	fv.separator = fs.String("separator", " ", "Separator character between fields.")

	err := fs.Parse(args)
	if err != nil {
		panic(err)
	}

	field, err := strutilities.GetField(input, *fv.index, *fv.separator)
	if err != nil {
		return "", err
	}

	return field, nil
}
