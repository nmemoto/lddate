package main

import "flag"

const (
	defaultFormat   = "%Y/%m/%d %H:%M:%S"
	defaultDuration = 1
	defaultFont     = "roman"
)

func parseArgs() (format string, duration int, font string) {

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

	flag.StringVar(&format, "f", defaultFormat, "date format")
	flag.StringVar(&format, "format", defaultFormat, "date format")
	flag.IntVar(&duration, "d", defaultDuration, "update date duration")
	flag.IntVar(&duration, "duration", defaultDuration, "update date duration")
	flag.StringVar(&font, "font", defaultFont, "font")
	flag.Parse()
	if !contains(fonts, font) {
		font = defaultFont
	}
	return
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
