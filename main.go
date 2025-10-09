package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/adamdecaf/tz/pkg/parse"
)

var (
	flagTo     = flag.String("to", "", "Timezone(s) used to convert input into")
	flagFormat = flag.String("format", "", "Format used to display output")
)

// tz accepts a date/timestamp as input and outputs the time in the user's
// local timezone and UTC.
func main() {
	flag.Parse()

	input := strings.Join(flag.Args(), " ")
	if input == "" {
		fmt.Println("No input provided") //nolint:forbidigo
		os.Exit(1)
	}

	in, format, err := parse.Time(input)
	if err != nil {
		fmt.Printf("ERROR %s\n", err) //nolint:forbidigo
		os.Exit(1)
	}
	if *flagFormat != "" {
		format = *flagFormat
	}

	now := time.Now()
	utc := in.In(time.UTC).Format(format)
	out := in.In(now.Location()).Format(format)
	tz := now.Location().String()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	lines := []string{
		fmt.Sprintf("UTC\t%s", utc),
		fmt.Sprintf("%s\t%s", tz, out),
	}

	if *flagTo != "" {
		tzs := strings.Split(*flagTo, ",")
		for i := range tzs {
			loc, err := time.LoadLocation(tzs[i])
			if err != nil {
				fmt.Printf("ERROR parsing -to=%s failed: %v\n", tzs[i], err) //nolint:forbidigo
				os.Exit(1)
			}
			to := fmt.Sprintf("%s\t%s", loc.String(), in.In(loc).Format(format))
			lines = append(lines, to)
		}
	}

	slices.Sort(lines)
	lines = slices.Compact(lines)

	for i := range lines {
		fmt.Fprintln(w, lines[i]) //nolint:forbidigo
	}
}
