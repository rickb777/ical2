[![Build Status](https://travis-ci.org/ajcollins/ical.svg?branch=master)](https://travis-ci.org/ajcollins/ical)
[![Go Report Card](https://goreportcard.com/badge/github.com/ajcollins/ical)](https://goreportcard.com/report/github.com/ajcollins/ical)

# ical

Simple ical encoder for golang. From github.com/soh335/ical.

## Support Extensions

* X-WR-CALNAME
* X-WR-CALDESC
* X-WR-TIMEZONE

## Support Components

* [x] Event Component
* [ ] To-do Component
* [ ] Journal Component
* [ ] Free/Busy Component
* [ ] Time Zone Component
* [ ] Alarm Component

## VEVENT additions

* Free / Busy using time transparency (TRANSP)
* Location (from github.com/colm2/ical)

## TODO

* 75 octets folding (https://tools.ietf.org/html/rfc5545#section-3.1)
