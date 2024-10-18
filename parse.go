package main

import (
	"fmt"
	"time"
)

// TODO(adam): need // 2024-02-12 15:31 UTC‑06:00 with fancy hyphen?

var (
	// formats is a list of various date and time formats found in the wild
	// Sources:
	//   - https://github.com/mjibson/goread/blob/master/utils.go#L162
	//   -
	//   - https://github.com/dolthub/go-mysql-server/blob/0a8f91740cdcc1309e01548ce1db37f9a0b2c8f2/sql/types/datetime.go#L62
	formats = []string{
		"01-02-2006",
		"01/02/2006 - 15:04",
		"01/02/2006 15:04:05 MST",
		"01/02/2006 3:04 PM",
		"01/02/2006",
		"02 Jan 2006 15:04 MST",
		"02 Jan 2006 15:04:05 -0700",
		"02 Jan 2006 15:04:05 MST",
		"02 Jan 2006 15:04:05 UT",
		"02 Jan 2006 15:04:05",
		"02 Jan 2006",
		"02 Monday, Jan 2006 15:04",
		"02-01-2006 15:04:05 MST",
		"02-01-2006",
		"02.01.2006 -0700",
		"02.01.2006 15:04",
		"02.01.2006 15:04:05",
		"02/01/2006 - 15:04",
		"02/01/2006 15:04 MST",
		"02/01/2006 15:04:05",
		"02/01/2006",
		"06-1-2 15:04",
		"06/1/2 15:04",
		"1/2/2006 15:04:05 MST",
		"1/2/2006 3:04:05 PM MST",
		"1/2/2006 3:04:05 PM",
		"1/2/2006",
		"15:04 02.01.2006 -0700",
		"15:04 MST",
		"2 Jan 2006 15:04:05 -0700",
		"2 Jan 2006 15:04:05 MST",
		"2 Jan 2006 15:04:05 Z",
		"2 Jan 2006",
		"2 January 2006 15:04:05 -0700",
		"2 January 2006 15:04:05 MST",
		"2 January 2006",
		"2-1-2006",
		"2.1.2006 15:04:05",
		"2/1/2006",
		"2006 January 02",
		"2006-01-02 00:00:00.0 15:04:05.0 -0700",
		"2006-01-02 15:04 MST-05:00",
		"2006-01-02 15:04 MST‑05:00",
		"2006-01-02 15:04",
		"2006-01-02 15:04:",
		"2006-01-02 15:04:.",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04 MST",
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05-0700",
		"2006-01-02 15:04:05-07:00",
		"2006-01-02 15:04:05.",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.999999999 -0700 MST",
		"2006-01-02 15:04:05Z",
		"2006-01-02 15:4",
		"2006-01-02 at 15:04:05",
		"2006-01-02",
		"2006-01-02T15:04-07:00",
		"2006-01-02T15:04:05 -0700",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05-07:00:00",
		"2006-01-02T15:04:05:-0700",
		"2006-01-02T15:04:05:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04Z",
		"2006-1-02T15:04:05Z",
		"2006-1-2 15:04:05",
		"2006-1-2 15:4:5.999999",
		"2006-1-2",
		"2006-1-2T15:04:05Z",
		"2006/01/02",
		"20060102",
		"20060102150405",
		"6-1-2 15:04",
		"6/1/2 15:04",
		"Jan 02 2006 03:04:05PM",
		"Jan 02, 2006",
		"Jan 2, 2006 15:04:05 MST",
		"Jan 2, 2006 3:04:05 PM MST",
		"Jan 2, 2006 3:04:05 PM",
		"Jan 2, 2006",
		"January 02, 2006 03:04 PM",
		"January 02, 2006 15:04",
		"January 02, 2006 15:04:05 MST",
		"January 02, 2006",
		"January 2, 2006 03:04 PM",
		"January 2, 2006 15:04:05 MST",
		"January 2, 2006 15:04:05",
		"January 2, 2006 3:04 PM",
		"January 2, 2006",
		"January 2, 2006, 3:04 p.m.",
		"Mon , 02 Jan 2006 15:04:05 MST",
		"Mon 02 Jan 2006 15:04:05 -0700",
		"Mon 2 Jan 2006 15:04:05 MST",
		"Mon Jan 02 2006 15:04:05 -0700",
		"Mon Jan 02 2006 15:04:05 GMT-0700 (MST)",
		"Mon Jan 02, 2006 3:04 pm",
		"Mon Jan 2 15:04 2006",
		"Mon Jan 2 15:04:05 2006 MST",
		"Mon, 02 Jan 06 15:04:05 MST",
		"Mon, 02 Jan 2006 15 -0700",
		"Mon, 02 Jan 2006 15:04 -0700",
		"Mon, 02 Jan 2006 15:04 MST",
		"Mon, 02 Jan 2006 15:04:05 --0700",
		"Mon, 02 Jan 2006 15:04:05 -07",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"Mon, 02 Jan 2006 15:04:05 -07:00",
		"Mon, 02 Jan 2006 15:04:05 00",
		"Mon, 02 Jan 2006 15:04:05 GMT-0700",
		"Mon, 02 Jan 2006 15:04:05 MST -0700",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 MST-07:00",
		"Mon, 02 Jan 2006 15:04:05 UT",
		"Mon, 02 Jan 2006 15:04:05 Z",
		"Mon, 02 Jan 2006 15:04:05",
		"Mon, 02 Jan 2006 15:04:05MST",
		"Mon, 02 Jan 2006 15:4:5 Z",
		"Mon, 02 Jan 2006 3:04:05 PM MST",
		"Mon, 02 Jan 2006",
		"Mon, 02 Jan 2006, 15:04:05 MST",
		"Mon, 02 January 2006",
		"Mon, 2 Jan 06 15:04:05 -0700",
		"Mon, 2 Jan 06 15:04:05 MST",
		"Mon, 2 Jan 15:04:05 MST",
		"Mon, 2 Jan 2006 15:04 -0700",
		"Mon, 2 Jan 2006 15:04 MST",
		"Mon, 2 Jan 2006 15:04",
		"Mon, 2 Jan 2006 15:04:05 -0700 MST",
		"Mon, 2 Jan 2006 15:04:05 -0700",
		"Mon, 2 Jan 2006 15:04:05 MST",
		"Mon, 2 Jan 2006 15:04:05 UT",
		"Mon, 2 Jan 2006 15:04:05",
		"Mon, 2 Jan 2006 15:04:05-0700",
		"Mon, 2 Jan 2006 15:04:05-07:00",
		"Mon, 2 Jan 2006 15:04:05MST",
		"Mon, 2 Jan 2006 15:4:5 -0700 GMT",
		"Mon, 2 Jan 2006 15:4:5 MST",
		"Mon, 2 Jan 2006 3:04:05 PM -0700",
		"Mon, 2 Jan 2006",
		"Mon, 2 Jan 2006, 15:04 -0700",
		"Mon, 2 January 2006 15:04 MST",
		"Mon, 2 January 2006 15:04:05 -0700",
		"Mon, 2 January 2006 15:04:05 MST",
		"Mon, 2 January 2006",
		"Mon, 2 January 2006, 15:04 -0700",
		"Mon, 2 January 2006, 15:04:05 MST",
		"Mon, 2, Jan 2006 15:4",
		"Mon, 2006-01-02 15:04",
		"Mon, Jan 02,2006 15:04:05 MST",
		"Mon, Jan 2 2006 15:04 MST",
		"Mon, Jan 2 2006 15:04:05 -0700",
		"Mon, Jan 2 2006 15:04:05 -700",
		"Mon, Jan 2, 2006 15:04 MST",
		"Mon, Jan 2, 2006 15:04:05 MST",
		"Mon, January 02, 2006 15:04:05 MST",
		"Mon, January 02, 2006, 15:04:05 MST",
		"Mon, January 2 2006 15:04:05 -0700",
		"Mon,02 Jan 2006 15:04 MST",
		"Mon,02 Jan 2006 15:04:05 -0700",
		"Mon,02 January 2006 14:04:05 MST",
		"Mon,2 Jan 2006",
		"Monday, 02 January 2006 15:04:05 -0700",
		"Monday, 02 January 2006 15:04:05 MST",
		"Monday, 02 January 2006 15:04:05",
		"Monday, 2 Jan 2006 15:04:05 -0700",
		"Monday, 2 Jan 2006 15:04:05 MST",
		"Monday, 2 January 2006 15:04:05 -0700",
		"Monday, 2 January 2006 15:04:05 MST",
		"Monday, January 02, 2006",
		"Monday, January 2, 2006 03:04 PM",
		"Monday, January 2, 2006 15:04:05 MST",
		"Monday, January 2, 2006",
		time.ANSIC,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RubyDate,
		time.UnixDate,
	}
)

func parse(input string) (time.Time, string, error) {
	for i := range formats {
		ts, err := time.Parse(formats[i], input)
		if err != nil {
			continue
		}
		return ts, formats[i], nil
	}
	return time.Time{}, "", fmt.Errorf("unsupported format %s", input)
}
