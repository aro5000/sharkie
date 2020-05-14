package main

type targetdata struct {
	Url            string
	Host           string
	Path           string
	Proto          string
	Port           string
	Expected       int
	DisplaySuccess bool
	Sleep          float64
}

var TDATA targetdata

type tracking struct {
	Twohundreds     int
	Threehundreds   int
	Fourhundreds    int
	Fivehundreds    int
	Failed          int
	Total           int
	Server          string
}

var TRACKINGLIST []tracking
