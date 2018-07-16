package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
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
	paris, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	utcJanNoon := time.Date(2014, time.Month(1), 1, 12, 0, 0, 0, time.UTC)
	utcJulyNoon := time.Date(2014, time.Month(7), 1, 12, 0, 0, 0, time.UTC)
	parisJanNoon := time.Date(2014, time.Month(1), 2, 12, 0, 0, 0, paris)
	parisJulyNoon := time.Date(2014, time.Month(7), 2, 12, 0, 0, 0, paris)

	cases := []struct {
		dt  DateTimeValue
		exp string
	}{
		{DateTime(utcJanNoon), "DT;VALUE=DATE-TIME:20140101T120000Z\n"},
		{DateTime(utcJanNoon).With(parameter.TZid("UTC")), "DT;VALUE=DATE-TIME:20140101T120000Z\n"},
		{DateTime(parisJanNoon), "DT;VALUE=DATE-TIME:20140102T120000\n"},
		{DateTime(parisJanNoon).With(parameter.TZid("Europe/Paris")), "DT;VALUE=DATE-TIME;TZID=Europe/Paris:20140102T120000\n"},
		//{DateTime(parisJanNoon).UTC(), "DT;VALUE=DATE-TIME:20140102T120000Z\n"},

		{DateTime(utcJulyNoon), "DT;VALUE=DATE-TIME:20140701T120000Z\n"},
		{DateTime(parisJulyNoon), "DT;VALUE=DATE-TIME:20140702T120000\n"},
		//{DateTime(parisJulyNoon).UTC(), "DT;VALUE=DATE-TIME:20140702T120000Z\n"},
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
			t.Errorf("%d: expected %s, got %s", i, c.exp, s)
		}
	}
}
