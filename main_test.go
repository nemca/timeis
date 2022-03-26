package main

import (
	"bytes"
	"testing"
)

func TestGetTime(t *testing.T) {
	cases := []struct {
		name     string
		timezone string
		want     string
		wantErr  string
	}{
		{
			name:     "test UTC",
			timezone: utc,
			want:     "UTC",
		},
		{
			name:     "bad timezone",
			timezone: "Foo/Bar",
			want:     "UTC",
			wantErr:  "unknown time zone Foo/Bar",
		},
	}

	for _, c := range cases {
		utcTime, err := getTime(c.timezone)
		if err != nil && err.Error() != c.wantErr {
			t.Fatalf("%s: got error: %v; want error: %v", c.name, err.Error(), c.wantErr)
		}
		got := utcTime.Format("MST")
		if got != c.want {
			t.Fatalf("%s: got: %v; want: %v", c.name, got, c.want)
		}
	}
}

func TestFormatTime(t *testing.T) {
	cases := []struct {
		name      string
		zone      string
		localTime string
		localDate string
		delta     string
		want      string
	}{
		{
			name:      "test UTC",
			zone:      "UTC",
			localTime: "11:12:45",
			localDate: "26-03-2022",
			delta:     "+0000",
			want:      "UTC\t11:12:45\t26-03-2022\t+0000\n",
		},
	}

	for _, c := range cases {
		var b bytes.Buffer
		formatTime(&b, c.zone, c.localTime, c.localDate, c.delta)
		got := b.String()
		if got != c.want {
			t.Fatalf("%s: got: %v; want: %v", c.name, got, c.want)
		}
	}
}
