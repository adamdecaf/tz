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

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <input>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nFormat aliases:\n")
		fmt.Fprintf(os.Stderr, "  rfc3339, iso, iso8601 -> RFC3339\n")
		fmt.Fprintf(os.Stderr, "  ansic -> ANSIC\n")
		fmt.Fprintf(os.Stderr, "  rfc822, rfc822z -> RFC822 variants\n")
		fmt.Fprintf(os.Stderr, "  rfc1123, rfc1123z -> RFC1123 variants\n")
		fmt.Fprintf(os.Stderr, "  rfc3339nano -> RFC3339Nano\n")
		fmt.Fprintf(os.Stderr, "  kitchen -> Kitchen\n")
		fmt.Fprintf(os.Stderr, "  stamp, stampmilli, stampmicro, stampnano -> Stamp variants\n")
		fmt.Fprintf(os.Stderr, "  unixdate -> UnixDate\n")
		fmt.Fprintf(os.Stderr, "  rubydate -> RubyDate\n")
	}
}

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
	// check format against known aliases
	switch strings.ToLower(format) {
	case "rfc3339", "iso", "iso8601":
		format = time.RFC3339
	case "ansic":
		format = time.ANSIC
	case "rfc822":
		format = time.RFC822
	case "rfc822z":
		format = time.RFC822Z
	case "rfc1123":
		format = time.RFC1123
	case "rfc1123z":
		format = time.RFC1123Z
	case "rfc3339nano":
		format = time.RFC3339Nano
	case "kitchen":
		format = time.Kitchen
	case "stamp":
		format = time.Stamp
	case "stampmilli":
		format = time.StampMilli
	case "stampmicro":
		format = time.StampMicro
	case "stampnano":
		format = time.StampNano
	case "unixdate":
		format = time.UnixDate
	case "rubydate":
		format = time.RubyDate
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
