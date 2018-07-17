package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"strings"
	"testing"
	"time"
)

func TestDateTimeZero(t *testing.T) {
	defined := TStamp(time.Time{}).IsDefined()
	if defined {
		t.Error("zero value should be undefined")
	}
}

func TestDateTimeRender(t *testing.T) {
	// choose one that's not UTC
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	utcJanNoon := time.Date(2014, time.Month(1), 1, 12, 0, 0, 0, time.UTC)
	utcJulyNoon := time.Date(2014, time.Month(7), 1, 12, 0, 0, 0, time.UTC)
	parisJanNoon := time.Date(2014, time.Month(1), 2, 12, 0, 0, 0, berlin)
	parisJulyNoon := time.Date(2014, time.Month(7), 2, 12, 0, 0, 0, berlin)

	cases := []struct {
		dt  DateTimeValue
		exp string
	}{
		{DateTime(utcJanNoon), "DT;VALUE=DATE-TIME:20140101T120000Z\n"},
		{DateTime(utcJanNoon).With(parameter.TZid("UTC")), "DT;VALUE=DATE-TIME:20140101T120000Z\n"},
		{DateTime(parisJanNoon), "DT;VALUE=DATE-TIME:20140102T120000\n"},
		{DateTime(parisJanNoon).With(parameter.TZid("Europe/Berlin")), "DT;VALUE=DATE-TIME;TZID=Europe/Berlin:20140102T120000\n"},

		{DateTime(utcJulyNoon), "DT;VALUE=DATE-TIME:20140701T120000Z\n"},
		{DateTime(parisJulyNoon), "DT;VALUE=DATE-TIME:20140702T120000\n"},

		{DateTime(utcJanNoon).AsDate(), "DT;VALUE=DATE:20140101\n"},
		{DateTime(utcJanNoon).With(parameter.TZid("UTC")).AsDate(), "DT;TZID=UTC;VALUE=DATE:20140101\n"},
		{DateTime(parisJanNoon).AsDate(), "DT;VALUE=DATE:20140102\n"},
		{DateTime(parisJanNoon).With(parameter.TZid("Europe/Berlin")).AsDate(), "DT;TZID=Europe/Berlin;VALUE=DATE:20140102\n"},
	}

	for i, c := range cases {
		b := &bytes.Buffer{}
		x := ics.NewBuffer(b, "\n")
		x.WriteValuerLine(true, "DT", c.dt)
		err := x.Flush()
		if err != nil {
			t.Errorf("%d: unexpected error %v", i, err)
		}
		s := b.String()
		if s != c.exp {
			t.Errorf("%d: expected %s, got %s", i, strings.TrimSpace(c.exp), s)
		}
	}
}

func TestNonTextValuesShouldIncludeType(t *testing.T) {
	now := time.Now()

	cases := []struct {
		v   ics.Valuer
		exp string
	}{
		{DateTime(now), ";VALUE=DATE-TIME;"},
		{Date(now), ";VALUE=DATE;"},
		{Duration("PT1H"), ";VALUE=DURATION;"},
		{Geo(1, 2), ";VALUE=FLOAT;"},
		{Integer(1), ";VALUE=INTEGER;"},
		{URI("a:b:c"), ";VALUE=URI;"},
	}

	for i, c := range cases {
		b := &bytes.Buffer{}
		x := ics.NewBuffer(b, "\n")
		x.WriteValuerLine(true, "X", c.v)
		err := x.Flush()
		if err != nil {
			t.Errorf("%d: unexpected error %v", i, err)
		}
		s := b.String()
		// ignore differences between COLON and SEMICOLON
		s = strings.Replace(s, ":", ";", -1)
		if !strings.Contains(s, c.exp) {
			t.Errorf("%d: expected %q within %q", i, c.exp, s)
		}
	}
}

func TestDefaultValuesShouldNotBeDefined(t *testing.T) {
	now := time.Now()

	cases := []struct {
		v   ics.Valuer
		exp bool
	}{
		{DateTime(now), true},
		{Date(now), true},
		{Duration("PT1H"), true},
		{Geo(1, 2), true},
		{Integer(1), true},
		{URI("a:b:c"), true},
		{Text("abc"), true},

		{DateTimeValue{}, false},
		{GeoValue{}, false},
		{IntegerValue{}, false},
		{URI(""), false},
		{Text(""), false},
	}

	for i, c := range cases {
		if c.exp {
			if !c.v.IsDefined() {
				t.Errorf("%d: expected defined %#q", i, c.v)
			}
		} else {
			if c.v.IsDefined() {
				t.Errorf("%d: expected undefined %#q", i, c.v)
			}
		}
	}
}
