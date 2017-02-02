package parsetime

import (
	"testing"
	"time"
)

func Test_ParseTime(t *testing.T) {
	n := time.Date(2007, 12, 10, 9, 44, 7, 123123, time.UTC)
	t.Logf("On time: %s", n)
	tf := map[string]time.Time{
		"2006-01-02T15:04:05Z07:00":         time.Date(2007, 12, 10, 9, 44, 7, 0, time.UTC),
		"2006-01-02":                        time.Date(2007, 12, 10, 0, 0, 0, 0, time.UTC),
		"2006-01-02 15:04:05":               time.Date(2007, 12, 10, 9, 44, 7, 0, time.UTC),
		"2006-01-02 15:04":                  time.Date(2007, 12, 10, 9, 44, 0, 0, time.UTC),
		"2006-01-02 15":                     time.Date(2007, 12, 10, 9, 0, 0, 0, time.UTC),
		"Mon Jan 2 15:04:05 -0700 MST 2006": time.Date(2007, 12, 10, 9, 44, 7, 0, time.UTC),
	}

	for f, expectedTime := range tf {
		should := n.Format(f)
		t.Logf("Check Format: parsed time: %s == should: %s", f, should)
		pt, err := ParseTime(should)
		if err != nil {
			t.Fatalf("Failed Format: %s == time:%s == parsed: %s == error: %v", f, n, pt, err)
		}

		if !pt.Equal(expectedTime) {
			t.Fatalf("Failed Format Match: `%s` == RealTime: %s == Parsed: %v == Expected: %v", f, n, pt, expectedTime)
		}
	}

	auxT := map[string]time.Duration{
		"1s":        time.Duration(time.Second),
		"1sec":      time.Duration(time.Second),
		"1m":        time.Duration(time.Minute),
		"1min":      time.Duration(time.Minute),
		"1h":        time.Duration(time.Hour),
		"1hour":     time.Duration(time.Hour),
		"1d":        time.Duration(time.Hour * 24),
		"1day":      time.Duration(time.Hour * 24),
		"now":       time.Duration(0),
		"today":     time.Duration(0),
		"yesterday": time.Duration(-time.Hour * 24),
		"-1d":       time.Duration(-time.Hour * 24),
		"-1h":       time.Duration(-time.Hour),
		"-1m":       time.Duration(-time.Minute),
		"-1s":       time.Duration(-time.Second),
	}

	for str, dur := range auxT {
		t.Logf("Check Format: %s", str)

		// this one is a little harder as "now" changes it can fail if things are real slow
		rt := time.Now().Add(dur)
		p, err := ParseTime(str)

		if rt.Unix() != p.Unix() || err != nil {
			t.Fatalf("Failed `%s` Match %s ||| wanted: %d == got %d", str, err, rt.Unix(), p.Unix())
		}
	}

	intT := map[string]time.Time{
		"1486000961":           time.Unix(1486000961, 0),
		"1486000961.123":       time.Unix(1486000961, 123000000),
		"1486000961.123123":    time.Unix(1486000961, 123123000),
		"1486000961.123123123": time.Unix(1486000961, 123123123),
	}

	for str, rtime := range intT {
		t.Logf("Check Format: %s", str)

		// this one is a little harder as "now" changes it can fail if things are real slow
		p, err := ParseTime(str)
		if rtime.UnixNano() != p.UnixNano() || err != nil {
			t.Fatalf("Failed Number: %s ||| Should: %d == got: %d  (error: %v)", str, rtime.UnixNano(), p.UnixNano(), err)
		}
	}

}

func Test_ParseDuration(t *testing.T) {

	auxT := map[string]time.Duration{
		"1s":    time.Duration(time.Second),
		"1sec":  time.Duration(time.Second),
		"1m":    time.Duration(time.Minute),
		"1min":  time.Duration(time.Minute),
		"1h":    time.Duration(time.Hour),
		"1hour": time.Duration(time.Hour),
		"1d":    time.Duration(time.Hour * 24),
		"1day":  time.Duration(time.Hour * 24),
		"now":   time.Duration(0),
		"today": time.Duration(0),
		"-1d":   time.Duration(-time.Hour * 24),
		"-1h":   time.Duration(-time.Hour),
		"-1m":   time.Duration(-time.Minute),
		"-1s":   time.Duration(-time.Second),
	}

	for str, dur := range auxT {
		t.Logf("Check Format: %s", str)

		// this one is a little harder as "now" changes it can fail if things are real slow
		p, err := ParseDuration(str)
		if p != dur {
			t.Fatalf("Failed %s Match %s", str, err)
		}
	}
}
