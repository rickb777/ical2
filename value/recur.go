package value

import (
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"github.com/rickb777/ical2/parameter/valuetype"
	"strconv"
	"time"
)

type Weekday time.Weekday

const (
	Undefined Weekday = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// String returns abbreviated English name of the day ("SU", "MO", ...).
func (d Weekday) String() string { return days[d] }

var days = [...]string{
	"XX",
	"SU",
	"MO",
	"TU",
	"WE",
	"TH",
	"FR",
	"SA",
}

//-------------------------------------------------------------------------------------------------

type WeekDayNum struct {
	OrdWk   int
	WeekDay Weekday
}

func (wdn WeekDayNum) InRange() bool {
	return -53 <= wdn.OrdWk && wdn.OrdWk <= 53
}

func (wdn WeekDayNum) String() string {
	if wdn.OrdWk == 0 {
		return wdn.WeekDay.String()
	}
	return fmt.Sprintf("%d%s", wdn.OrdWk, wdn.WeekDay)
}

var SU = WeekDayNum{WeekDay: Sunday}
var MO = WeekDayNum{WeekDay: Monday}
var TU = WeekDayNum{WeekDay: Tuesday}
var WE = WeekDayNum{WeekDay: Wednesday}
var TH = WeekDayNum{WeekDay: Thursday}
var FR = WeekDayNum{WeekDay: Friday}
var SA = WeekDayNum{WeekDay: Saturday}

//-------------------------------------------------------------------------------------------------

const (
	MINUTELY = "MINUTELY"
	HOURLY   = "HOURLY"
	DAILY    = "DAILY"
	WEEKLY   = "WEEKLY"
	MONTHLY  = "MONTHLY"
	YEARLY   = "YEARLY"
)

//-------------------------------------------------------------------------------------------------

// RecurrenceValue holds an integer.
type RecurrenceValue struct {
	Parameters parameter.Parameters
	Freq       string
	Interval   uint
	Count      uint
	Until      time.Time
	ByWeekNo   []int
	ByMonth    []uint
	ByHour     []uint
	ByMinute   []uint
	BySecond   []uint
	ByDay      []WeekDayNum
	ByMonthDay []int
	ByYearDay  []int
	BySetPos   []int
	WeekStart  Weekday
}

// Recurrence returns a new RecurrenceValue. It has VALUE=INTEGER.
func Recurrence(freq string) RecurrenceValue {
	return RecurrenceValue{
		Parameters: parameter.Parameters{valuetype.Type(valuetype.RECUR)},
		Freq:       freq,
	}
}

// IsDefined tests whether the value has been explicitly defined or is default.
func (v RecurrenceValue) IsDefined() bool {
	return v.Freq != ""
}

// With appends parameters to the value.
func (v RecurrenceValue) With(params ...parameter.Parameter) RecurrenceValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v RecurrenceValue) WriteTo(w ics.StringWriter) (err error) {
	err = v.Validate()
	if err != nil {
		return err
	}

	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, err = w.WriteString("FREQ")
	w.WriteByte('=')
	_, err = w.WriteString(v.Freq)
	writeParam(v.Interval > 0, w, "INTERVAL", strconv.Itoa(int(v.Interval)))
	writeParam(v.Count > 0, w, "COUNT", strconv.Itoa(int(v.Count)))
	writeParam(!v.Until.IsZero(), w, "UNTIL", v.Until.Format(dateTimeLayoutZ))
	writeIntList(len(v.ByWeekNo) > 0, w, "BYWEEKNO", v.ByWeekNo)
	writeUintList(len(v.ByMonth) > 0, w, "BYMONTH", v.ByMonth)
	writeUintList(len(v.ByHour) > 0, w, "BYHOUR", v.ByHour)
	writeUintList(len(v.ByMinute) > 0, w, "BYMINUTE", v.ByMinute)
	writeUintList(len(v.BySecond) > 0, w, "BYSECOND", v.BySecond)
	writeWeekDayNumList(len(v.ByDay) > 0, w, "BYDAY", v.ByDay)
	writeIntList(len(v.ByMonthDay) > 0, w, "BYMONTHDAY", v.ByMonthDay)
	writeIntList(len(v.ByYearDay) > 0, w, "BYYEARDAY", v.ByYearDay)
	writeIntList(len(v.BySetPos) > 0, w, "BYSETPOS", v.BySetPos)
	writeParam(v.WeekStart > 0, w, "WKST", v.WeekStart.String())
	return err
}

func writeParam(predicate bool, w ics.StringWriter, key, value string) {
	if predicate {
		w.WriteByte(';')
		w.WriteString(key)
		w.WriteByte('=')
		w.WriteString(value)
	}
}

func writeIntList(predicate bool, w ics.StringWriter, key string, value []int) {
	if predicate {
		w.WriteByte(';')
		w.WriteString(key)
		w.WriteByte('=')
		comma := ""
		for _, v := range value {
			w.WriteString(comma)
			w.WriteString(strconv.Itoa(v))
			comma = ","
		}
	}
}

func writeUintList(predicate bool, w ics.StringWriter, key string, value []uint) {
	if predicate {
		w.WriteByte(';')
		w.WriteString(key)
		w.WriteByte('=')
		comma := ""
		for _, v := range value {
			w.WriteString(comma)
			w.WriteString(strconv.Itoa(int(v)))
			comma = ","
		}
	}
}

func writeWeekDayNumList(predicate bool, w ics.StringWriter, key string, value []WeekDayNum) {
	if predicate {
		w.WriteByte(';')
		w.WriteString(key)
		w.WriteByte('=')
		comma := ""
		for _, v := range value {
			w.WriteString(comma)
			w.WriteString(v.String())
			comma = ","
		}
	}
}

// Validate confirms that the recurrence parameters are within valid ranges.
func (v RecurrenceValue) Validate() (err error) {
	err = validPositiveList(err, "BySecond", 0, 60, v.BySecond)
	err = validPositiveList(err, "ByMinute", 0, 59, v.ByMinute)
	err = validPositiveList(err, "ByHour", 0, 23, v.ByHour)
	err = validPositiveList(err, "ByMonth", 1, 12, v.ByMonth)
	err = validPlusMinusList(err, "ByMonthDay", 1, 31, v.ByMonthDay)
	err = validPlusMinusList(err, "ByYearDay", 1, 366, v.ByYearDay)
	err = validPlusMinusList(err, "ByWeekNo", 0, 53, v.ByWeekNo)
	err = validPlusMinusList(err, "BySetPos", 1, 366, v.BySetPos)
	err = validWeekDayList(err, "ByDay", v.ByDay)
	return err
}

func validPositiveList(previous error, parameter string, min, max uint, value []uint) error {
	if previous != nil {
		return previous
	}
	for _, v := range value {
		if v < min || v > max {
			return fmt.Errorf("%s value is out of the range %d to %d %v", parameter, min, max, value)
		}
	}
	return nil
}

func validPlusMinusList(previous error, parameter string, min, max int, value []int) error {
	if previous != nil {
		return previous
	}
	for _, v := range value {
		if v < -max || (v > -min && v < min) || v > max {
			return fmt.Errorf("%s value is out of the range %d to %d %v", parameter, min, max, value)
		}
	}
	return nil
}

func validWeekDayList(previous error, parameter string, value []WeekDayNum) error {
	if previous != nil {
		return previous
	}
	for _, v := range value {
		if !v.InRange() {
			return fmt.Errorf("%s value is out of the range -53 to 53 %v", parameter, value)
		}
	}
	return nil
}
