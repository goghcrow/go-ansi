package ansi

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Color int

const (
	Black  Color = 30
	Red    Color = 31
	Green  Color = 32
	Yellow Color = 33
	Blue   Color = 34
	Purple Color = 35
	Cyan   Color = 36
	White  Color = 37
)

func (c Color) Light() Color        { return Color(-int(math.Abs(float64(c)))) }
func (c Color) Text(s string) *Ansi { return New().Fg(c).Text(s) }
func (c Color) Bg() *Ansi           { return New().Bg(c) }
func (c Color) Fg() *Ansi           { return New().Fg(c) }
func (c Color) Bold() *Ansi         { return New().Fg(c).Bold() }
func (c Color) Underline() *Ansi    { return New().Fg(c).Underline() }

func (a *Ansi) Fg(c Color) *Ansi {
	a.fg = c
	return a
}

// Bg Background Color
func (a *Ansi) Bg(c Color) *Ansi {
	a.bg = c
	return a
}

func (a *Ansi) Bold() *Ansi {
	a.bold = true
	return a
}

func (a *Ansi) Underline() *Ansi {
	a.underline = true
	return a
}

func (a *Ansi) Text(s string) *Ansi {
	a.span.buf.WriteString(s)
	return a
}

// Append Span
func (a *Ansi) Append(o *Ansi) *Ansi {
	a.Reset().buf.WriteString(o.Reset().String())
	return a
}

func New() *Ansi { return &Ansi{} }

type Ansi struct {
	span
	buf strings.Builder
}

type span struct {
	fg        Color
	bg        Color
	bold      bool
	underline bool

	buf strings.Builder
}

const rst = "\033[0m"

func (s span) String() string {
	if s.fg == 0 && (s.bold || s.underline) {
		s.fg = Black
	}

	fg := 0
	if s.fg != 0 {
		if s.fg < 0 {
			fg = int(-s.fg) + 60
		} else {
			fg = int(s.fg)
		}
	}

	bg := 0
	if s.bg != 0 {
		if s.bg < 0 {
			bg = int(-s.bg) + 60
		} else {
			bg = int(s.bg)
		}
		bg += 10
	}

	ex := "0;"
	if s.bold {
		ex += "1;"
	}
	if s.underline {
		ex += "4;"

	}

	b := s.buf.String()
	if fg == 0 && bg == 0 {
		return b
	}
	if bg == 0 {
		return fmt.Sprintf("\033[%s%dm%s"+rst, ex, fg, b)
	}
	if fg == 0 {
		return fmt.Sprintf("\033[%dm%s"+rst, bg, b)
	}
	return fmt.Sprintf("\033[%s%dm\033[%dm%s"+rst, ex, fg, bg, b)
}

func (a *Ansi) Reset() *Ansi {
	a.buf.WriteString(a.span.String())
	a.span = span{}
	return a
}

func (a *Ansi) S() string {
	return a.String()
}

func (a *Ansi) String() string {
	if a.span.buf.Len() > 0 {
		a.Reset()
	}
	return a.buf.String()
}

var pat = regexp.MustCompile("\033\\[([014];)*\\d+m")

func Strip(s string) string {
	return pat.ReplaceAllStringFunc(s, func(s string) string { return "" })
}
