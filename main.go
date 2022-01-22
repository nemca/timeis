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
	"os"
	"text/tabwriter"
	"time"
)

const (
	moscow     string = "Europe/Moscow"
	utc        string = "UTC"
	timeFormat string = "15:04:05"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	utcLocation, err := time.LoadLocation(utc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parse timezone: %s\n", err.Error())
		return
	}
	utcTime := time.Now().In(utcLocation)
	fmt.Fprintf(w, "UTC:\t%s\n", utcTime.Format(timeFormat))

	moscowLocation, err := time.LoadLocation(moscow)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parse timezone: %s\n", err.Error())
		return
	}
	moscowTime := time.Now().In(moscowLocation)
	fmt.Fprintf(w, "Moscow:\t%s\n", moscowTime.Format(timeFormat))

	w.Flush()
}
