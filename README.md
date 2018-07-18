[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/rickb777/ical2)
[![Build Status](https://travis-ci.org/rickb777/ical2.svg?branch=master)](https://travis-ci.org/rickb777/ical2)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickb777/ical2)](https://goreportcard.com/report/github.com/rickb777/ical2)

# ical2

Simple iCalendar encoder for Go. See https://tools.ietf.org/html/rfc5545

There is no parsing (unmarshalling) implementation yet, although the design will support this.

This repo is a rewritten fork from github.com/ajcollins/ical, which was orignally from github.com/soh335/ical.

## Installation

    go get -u github.com/rickb777/ical2

or

    dep ensure -add github.com/rickb777/ical2

## Supported Components

* [x] Event Component (except for recurrence rules)
* [ ] Event recurrence rules (RECURRENCE-ID EXDATE RDATE RRULE)
* [ ] To-do Component
* [ ] Journal Component
* [x] Free/Busy Component
* [ ] Time Zone Component
* [ ] Alarm Component
* [ ] Parameter Value Encoding https://tools.ietf.org/html/rfc6868,
* [ ] Non-Gregorian Recurrence Rules https://tools.ietf.org/html/rfc7529
* [ ] Calendar Availability https://tools.ietf.org/html/rfc7953
* [x] New Properties https://tools.ietf.org/html/rfc7986
