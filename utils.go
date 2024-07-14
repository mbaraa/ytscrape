package main

import (
	"errors"
	"strings"
	"time"
	"unicode"
)

// getDuration returns a time.Duration from the given string duration, and an occurring error.
//
// Examples:
// 10:12 => 601000000000 or "10m12s"
//
// 12:34:56 => 45245000000000 or "12h34m56s"
var getDuration = durationer()

func durationer() func(strDuration string) (time.Duration, error) {
	durationSeparators := [3]rune{'s', 'm', 'h'}
	return func(strDuration string) (time.Duration, error) {
		startIdx := 0
		colonsCount := 0
		for _, chr := range strDuration {
			if chr == ':' {
				colonsCount++
			}
		}
		if colonsCount > 2 {
			return 0, errors.New("invalid iso duration")
		}
		for i, chr := range strDuration {
			if unicode.IsDigit(chr) {
				startIdx = i
				break
			}
		}
		refinedStrDuration := strings.Builder{}
		for i, chr := range strDuration[startIdx:] {
			if chr == ':' || i == len(strDuration)-1 {
				refinedStrDuration.WriteRune(durationSeparators[colonsCount])
				colonsCount--
				continue
			}
			refinedStrDuration.WriteRune(chr)
		}

		duration, err := time.ParseDuration(refinedStrDuration.String())
		if err != nil {
			return 0, err
		}

		return duration, nil
	}
}
