package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
)

const (
	ExitCodeOK = iota
	ExitCodeParserFlagError
)

const (
	defaultFormat   = "%Y/%m/%d %H:%M:%S"
	defaultDuration = 1 * time.Second
	defaultFont     = "roman"
	defaultPosition = "left"
	defaultColor    = "white"
)

var (
	format, font, position, strColor string
	duration                         time.Duration
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {

	//https://github.com/common-nighthawk/go-figure#supported-fonts
	fonts := []string{"3-d", "3x5", "5lineoblique", "acrobatic", "alligator", "alligator2", "alphabet", "avatar", "banner",
		"banner3-D", "banner3", "banner4", "barbwire", "basic", "bell", "big", "bigchief", "binary", "block", "bubble", "bulbhead",
		"calgphy2", "caligraphy", "catwalk", "chunky", "coinstak", "colossal", "computer", "contessa", "contrast", "cosmic", "cosmike",
		"cricket", "cursive", "cyberlarge", "cybermedium", "cybersmall", "diamond", "digital", "doh", "doom", "dotmatrix", "drpepper",
		"eftichess", "eftifont", "eftipiti", "eftirobot", "eftitalic", "eftiwall", "eftiwater", "epic", "fender", "fourtops", "fuzzy", "goofy",
		"gothic", "graffiti", "hollywood", "invita", "isometric1", "isometric2", "isometric3", "isometric4", "italic", "ivrit", "jazmine",
		"jerusalem", "katakana", "kban", "larry3d", "lcd", "lean", "letters", "linux", "lockergnome", "madrid", "marquee", "maxfour", "mike",
		"mini", "mirror", "mnemonic", "morse", "moscow", "nancyj-fancy", "nancyj-underlined", "nancyj", "nipples", "ntgreek", "o8", "ogre",
		"pawp", "peaks", "pebbles", "pepper", "poison", "puffy", "pyramid", "rectangles", "relief", "relief2", "rev", "roman", "rot13", "rounded",
		"rowancap", "rozzo", "runic", "runyc", "sblood", "script", "serifcap", "shadow", "short", "slant", "slide", "slscript", "small", "smisome1",
		"smkeyboard", "smscript", "smshadow", "smslant", "smtengwar", "speed", "stampatello", "standard", "starwars", "stellar", "stop", "straight",
		"tanja", "tengwar", "term", "thick", "thin", "threepoint", "ticks", "ticksslant", "tinker-toy", "tombstone", "trek", "tsalagi", "twopoint",
		"univers", "usaflag", "wavy", "weird",
	}
	// The following is relatively high visibility
	/*
		fonts := []string{"3-d", "3x5", "5lineoblique", "alligator2", "alphabet", "avatar", "banner",
			"banner3-D", "banner3", "banner4", "basic", "bell", "big", "bigchief", "block", "bubble", "bulbhead",
			"catwalk", "chunky", "colossal", "computer", "contessa", "contrast", "cosmic", "cosmike",
			"cricket", "cursive", "digital", "doom", "drpepper",
			"eftifont", "eftirobot", "eftitalic", "epic", "fender", "fuzzy",
			"gothic", "graffiti", "hollywood", "invita", "italic", "jazmine",
			"katakana", "kban", "larry3d", "lcd", "letters", "linux", "marquee", "maxfour",
			"mini", "moscow", "nancyj-fancy", "nancyj-underlined", "nancyj", "nipples", "ntgreek", "o8", "ogre",
			"pawp", "peaks", "pebbles", "pepper", "poison", "puffy", "pyramid", "rev", "roman", "rounded",
			"rozzo", "sblood", "script", "serifcap", "shadow", "slant", "slide", "slscript", "small",
			"smkeyboard", "smscript", "smshadow", "smslant", "speed", "stampatello", "standard", "starwars", "stop", "straight",
			"tanja", "thick", "thin", "tinker-toy", "tombstone", "trek",
			"univers", "usaflag", "wavy", "weird",
		}
	*/
	positions := []string{"left", "center", "right"}
	flags := flag.NewFlagSet("lddate", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	flags.StringVar(&format, "f", defaultFormat, "date format")
	flags.StringVar(&format, "format", defaultFormat, "date format")
	flags.DurationVar(&duration, "d", defaultDuration, "update date duration")
	flags.DurationVar(&duration, "duration", defaultDuration, "update date duration")
	flags.StringVar(&font, "font", defaultFont, "font")
	flags.StringVar(&position, "p", defaultPosition, "left")
	flags.StringVar(&position, "position", defaultPosition, "left")
	flags.StringVar(&strColor, "c", defaultColor, "text color")
	flags.StringVar(&strColor, "color", defaultColor, "text color")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParserFlagError
	}
	if !contains(fonts, font) {
		font = defaultFont
	}
	if !contains(positions, position) {
		position = defaultPosition
	}

	switch strColor {
	case "red":
		color.Set(color.FgRed)
	case "blue":
		color.Set(color.FgBlue)
	case "green":
		color.Set(color.FgGreen)
	case "yellow":
		color.Set(color.FgYellow)
	case "magenta":
		color.Set(color.FgMagenta)
	case "cyan":
		color.Set(color.FgCyan)
	case "black":
		color.Set(color.FgBlack)
	case "white":
		color.Set(color.FgWhite)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	canceled := false
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		<-sc
		cancel()
		canceled = true
	}()
	goFmtStr := formatDate(format)

	cnt := 0
	for now := range now(ctx, duration) {
		print("\033[H\033[2J")
		nowStr := now.Format(goFmtStr)
		var artStr artStr
		for _, str := range strings.Split(nowStr, "\n") {
			artStr = append(artStr, figure.NewFigure(str, font, true).Slicify())
		}
		maxWidth := maxArtStrWidth(artStr)

		if err := termbox.Init(); err != nil {
			panic(err)
		}
		w, _ := termbox.Size()
		termbox.Close()
		if w > maxWidth {
			switch position {
			case "left":
				artStr.Print(c.outStream)
			case "center":
				artStr.CenterPrint(c.outStream)
			case "right":
				artStr.RightPrint(c.outStream)
			}
		} else {
			fmt.Println("Increase the width of the terminal")
		}

		cnt++
	}
	if canceled {
		fmt.Fprintln(c.outStream, "canceled!")
	}
	return ExitCodeOK
}

func now(ctx context.Context, duration time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func(ch chan<- time.Time) {
		defer close(ch)
	loop:
		for now := range time.Tick(duration) {
			select {
			case <-ctx.Done():
				break loop
			default:
			}
			ch <- now
		}
	}(ch)
	return ch
}

func formatDate(format string) string {
	var matchTypes = []struct {
		dateCmdFmt string
		goFmt      string
	}{
		{
			dateCmdFmt: "%Y",
			goFmt:      "2006",
		},
		{
			dateCmdFmt: "%m",
			goFmt:      "01",
		},
		{
			dateCmdFmt: "%d",
			goFmt:      "02",
		},
		{
			dateCmdFmt: "%H",
			goFmt:      "15",
		},
		{
			dateCmdFmt: "%M",
			goFmt:      "04",
		},
		{
			dateCmdFmt: "%S",
			goFmt:      "05",
		},
		{
			dateCmdFmt: ".%3N",
			goFmt:      ".000",
		},
		{
			dateCmdFmt: ".%6N",
			goFmt:      ".000000",
		},
		{
			dateCmdFmt: ".%N",
			goFmt:      ".000000000",
		},
		{
			dateCmdFmt: "%A",
			goFmt:      "Monday",
		},
		{
			dateCmdFmt: "%a",
			goFmt:      "Mon",
		},
	}
	goFmtStr := format
	for _, m := range matchTypes {
		goFmtStr = strings.Replace(goFmtStr, m.dateCmdFmt, m.goFmt, -1)
	}
	goFmtStr = strings.Replace(goFmtStr, "\\n", "\n", -1)
	return goFmtStr
}

type artStr [][]string

func (str artStr) Print(w io.Writer) {
	for _, rowArtStr := range str {
		for _, row := range rowArtStr {
			fmt.Fprintln(w, row)
		}
	}
}

func (str artStr) CenterPrint(w io.Writer) {
	maxWidth := maxArtStrWidth(str)
	for i, rowArtStr := range str {
		if maxWidth == maxRowArtStrWidth(rowArtStr) {
			continue
		}
		diff := (maxWidth - maxRowArtStrWidth(rowArtStr)) / 2
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	str.Print(w)
}

func (str artStr) RightPrint(w io.Writer) {
	maxWidth := maxArtStrWidth(str)
	for i, rowArtStr := range str {
		if maxWidth == maxRowArtStrWidth(rowArtStr) {
			continue
		}
		diff := maxWidth - maxRowArtStrWidth(rowArtStr)
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	str.Print(w)
}

func maxArtStrWidth(str artStr) (maxWidth int) {
	for _, rowArtStr := range str {
		if maxRowArtStrWidth(rowArtStr) > maxWidth {
			maxWidth = maxRowArtStrWidth(rowArtStr)
		}
	}
	return
}

func maxRowArtStrWidth(rowArtStr []string) int {
	max := 0
	for _, str := range rowArtStr {
		if len(str) > max {
			max = len(str)
		}
	}
	return max
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
