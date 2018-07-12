[![Build Status](https://travis-ci.org/rickb777/ical.svg?branch=master)](https://travis-ci.org/rickb777/ical)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickb777/ical)](https://goreportcard.com/report/github.com/rickb777/ical)

# ical

Simple ical (https://tools.ietf.org/html/rfc5545) encoder for golang.

This repo is forked from github.com/soh335/ical via github.com/ajcollins/ical.

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
