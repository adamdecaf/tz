# tz

[![GoDoc](https://godoc.org/github.com/adamdecaf/tz?status.svg)](https://godoc.org/github.com/adamdecaf/tz)
[![Build Status](https://github.com/adamdecaf/tz/workflows/Go/badge.svg)](https://github.com/adamdecaf/tz/actions)
[![Coverage Status](https://codecov.io/gh/adamdecaf/tz/branch/master/graph/badge.svg)](https://codecov.io/gh/adamdecaf/tz)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdecaf/tz)](https://goreportcard.com/report/github.com/adamdecaf/tz)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/adamdecaf/tz/master/LICENSE)

Timezone conversion CLI

## Install

Download the [latest release for your architecture](https://github.com/adamdecaf/tz/releases/latest).

## Usage

```
$ tz $(date)
America/Chicago  Tue Feb 13 13:09:57 CST 2024
UTC              Tue Feb 13 19:09:57 UTC 2024
```

tz reads the `TZ=America/Chicago` environmental variable too and supports multiple output timezones.

```
$ tz -to America/New_York,America/Belize $(date)
America/Belize    Tue Feb 13 11:26:21 CST 2024
America/Chicago   Tue Feb 13 11:26:21 CST 2024
America/New_York  Tue Feb 13 12:26:21 EST 2024
UTC               Tue Feb 13 17:26:21 UTC 2024
```

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
