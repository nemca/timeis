/*
Copyright © 2022 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	utc         string = "UTC"
	timeFormat  string = "15:04:05"
	dateFormat  string = "02-01-2006"
	deltaFormat string = "-0700"
	timeString  string = "\t%s\t%s\t%s\n"
)

func main() {
	// CLI flags
	pflag.StringP("utc", "u", "", "Parse argument as an UTC time.")
	pflag.Parse()

	// Name of config file (without extension)
	viper.SetConfigName(".timeis")
	// REQUIRED if the config file does not have the extension in the name
	// path to look for the config file in call multiple times to add many
	// search paths
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/")
	// optionally look for config in the working directory
	viper.AddConfigPath(".")
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			fmt.Fprintln(os.Stderr, "config file not found")
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "failed read config file: %v\n", err)
		os.Exit(1)
	}

	// Tabwriter for formatted output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	// Output header
	fmt.Fprintf(w, "Timezone\tLocal time\tLocal date\tDelta\n")
	fmt.Fprintf(w, "----------\t----------\t----------\t----------\n")

	var userTime time.Time

	utcFlagValue, err := pflag.CommandLine.GetString("utc")
	if err != nil {
		return
	}
	if len(utcFlagValue) > 0 {
		userTime, err = time.Parse(timeFormat, utcFlagValue)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed parse UTC time %s: %v\n", utcFlagValue, err)
			os.Exit(1)
		}
	}
	// We always show UTC time
	utcTime, err := getTime(utc, userTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parse timezone: %v\n", err)
		os.Exit(2)
	}
	formatTime(w, "UTC", utcTime.Format(timeFormat), utcTime.Format(dateFormat), utcTime.Format(deltaFormat))

	// Get timezones from config file and calculate local time
	timeZones := viper.GetStringSlice("timezones")
	for _, zone := range timeZones {
		t, err := getTime(zone, userTime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed parse timezone %s: %v\n", zone, err)
			os.Exit(2)
		}
		formatTime(w, zone, t.Format(timeFormat), t.Format(dateFormat), t.Format(deltaFormat))
	}

	// Flush to output
	w.Flush()
}

func getTime(timezone string, t time.Time) (time.Time, error) {
	now := time.Now()

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	if !t.IsZero() {
		year, month, day := now.Date()
		now = time.Date(year, month, day, t.Hour(), t.Minute(), 0, 0, t.Location())
	}

	return now.In(location), nil
}

func formatTime(w io.Writer, zone string, times ...any) {
	format := fmt.Sprintf("%s%s", zone, timeString)
	fmt.Fprintf(w, format, times...)
}
