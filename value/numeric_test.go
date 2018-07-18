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
	berlin, _ := time.LoadLocation("Europe/Berlin")

	utcJanNoon := time.Date(2014, time.Month(1), 1, 12, 0, 0, 0, time.UTC)
	utcJulyNoon := time.Date(2014, time.Month(7), 1, 12, 0, 0, 0, time.UTC)
	parisJanNoon := time.Date(2014, time.Month(1), 2, 12, 0, 0, 0, berlin)
	parisJulyNoon := time.Date(2014, time.Month(7), 2, 12, 0, 0, 0, berlin)

	cases := []struct {
		dt  DateTimeValue
		exp string
	}{
		{TStamp(utcJanNoon), ":20140101T120000Z\n"},
		{DateTime(utcJanNoon), ";VALUE=DATE-TIME:20140101T120000Z\n"},
		{DateTime(utcJanNoon).With(parameter.TZid("UTC")), ";VALUE=DATE-TIME:20140101T120000Z\n"},
		{DateTime(parisJanNoon), ";VALUE=DATE-TIME:20140102T120000\n"},
		{DateTime(parisJanNoon).With(parameter.TZid("Europe/Berlin")), ";VALUE=DATE-TIME;TZID=Europe/Berlin:20140102T120000\n"},

		{DateTime(utcJulyNoon), ";VALUE=DATE-TIME:20140701T120000Z\n"},
		{DateTime(parisJulyNoon), ";VALUE=DATE-TIME:20140702T120000\n"},

		{DateTime(utcJanNoon).AsDate(), ";VALUE=DATE:20140101\n"},
		{DateTime(utcJanNoon).AsDate().With(parameter.TZid("UTC")), ";VALUE=DATE;TZID=UTC:20140101\n"},
		{DateTime(parisJanNoon).AsDate(), ";VALUE=DATE:20140102\n"},
		{DateTime(parisJanNoon).AsDate().With(parameter.TZid("Europe/Berlin")), ";VALUE=DATE;TZID=Europe/Berlin:20140102\n"},
	}

	for i, c := range cases {
		b := &bytes.Buffer{}
		x := ics.NewBuffer(b, "\n")
		x.WriteValuerLine(true, "X", c.dt)
		err := x.Flush()
		if err != nil {
			t.Errorf("%d: unexpected error %v", i, err)
		}
		s := b.String()
		// ignore the "X" label
		s = s[1:]
		if s != c.exp {
			t.Errorf("%d: expected %s, got %s", i, strings.TrimSpace(c.exp), s)
		}
	}
}

func TestFreeBusyRender(t *testing.T) {
	utcJanNoon := time.Date(2014, time.Month(2), 3, 12, 4, 5, 0, time.UTC)

	cases := []struct {
		dt  FreeBusyValue
		exp string
	}{
		{FreeBusyOf(utcJanNoon, time.Hour), ";VALUE=PERIOD:20140203T120405Z/PT1H\n"},
	}

	for i, c := range cases {
		b := &bytes.Buffer{}
		x := ics.NewBuffer(b, "\n")
		x.WriteValuerLine(true, "X", c.dt)
		err := x.Flush()
		if err != nil {
			t.Errorf("%d: unexpected error %v", i, err)
		}
		s := b.String()
		// ignore the "X" label
		s = s[1:]
		if s != c.exp {
			t.Errorf("%d: expected %s, got %s", i, strings.TrimSpace(c.exp), s)
		}
	}
}

func TestBinaryRender(t *testing.T) {

	cases := []struct {
		dt  ics.Valuer
		exp string
	}{
		{Binary([]byte("ABC")), ";VALUE=BINARY;ENCODING=BASE64:QUJD\n"},
		{Binary([]byte("A}~B")), ";VALUE=BINARY;ENCODING=BASE64:QX1+Qg==\n"},
	}

	for i, c := range cases {
		b := &bytes.Buffer{}
		x := ics.NewBuffer(b, "\n")
		x.WriteValuerLine(true, "X", c.dt)
		err := x.Flush()
		if err != nil {
			t.Errorf("%d: unexpected error %v", i, err)
		}
		s := b.String()
		// ignore the "X" label
		s = s[1:]
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
		{FreeBusyOf(now, 1), ";VALUE=PERIOD;"},
		{Geo(1, 2), ";VALUE=FLOAT;"},
		{Integer(1), ";VALUE=INTEGER;"},
		{URI("a:b:c"), ";VALUE=URI;"},
		{Binary([]byte("x")), ";VALUE=BINARY;"},
		{Binary([]byte("x")), ";ENCODING=BASE64;"},
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
		{Binary([]byte("x")), true},

		{DateTimeValue{}, false},
		{GeoValue{}, false},
		{IntegerValue{}, false},
		{URI(""), false},
		{Text(""), false},
		{Binary(nil), false},
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
