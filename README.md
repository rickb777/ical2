[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/rickb777/ical2)
[![Build Status](https://travis-ci.org/rickb777/ical2.svg?branch=master)](https://travis-ci.org/rickb777/ical2)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickb777/ical2)](https://goreportcard.com/report/github.com/rickb777/ical2)

# ical2

Simple ical (https://tools.ietf.org/html/rfc5545) encoder for golang.

This repo is a rewritten fork from github.com/ajcollins/ical, which was orignally from github.com/soh335/ical.

## Installation

    go get -u github.com/rickb777/ical2

or

    dep ensure -add github.com/rickb777/ical2

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



