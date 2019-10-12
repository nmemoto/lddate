package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/nsf/termbox-go"
)

type term struct {
	Width  int
	Height int
}

func newTerm() *term {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	tw, th := termbox.Size()
	return &term{
		Width:  tw,
		Height: th,
	}
}

/*
	type rowArtStr

	ex.

  .oooo.     .oooo.     .o   .ooooo.        88   .o    .oooo.        88   .oooo.       .ooo
.dP""Y88b   d8P'`Y8b  o888  888' `Y88.     .8' o888   d8P'`Y8b      .8'  d8P'`Y8b    .88'
      ]8P' 888    888  888  888    888    .8'   888  888    888    .8'  888    888  d88'
    .d8P'  888    888  888   `Vbood888   .8'    888  888    888   .8'   888    888 d888P"Ybo.
  .dP'     888    888  888        888'  .8'     888  888    888  .8'    888    888 Y88[   ]88
.oP     .o `88b  d88'  888      .88P'  .8'      888  `88b  d88' .8'     `88b  d88' `Y88   88P
8888888888  `Y8bd8P'  o888o   .oP'     88      o888o  `Y8bd8P'  88       `Y8bd8P'   `88bod8'

*/
type rowArtStr []string

func (artStr rowArtStr) MaxWidth() (max int) {
	for _, str := range artStr {
		if len(str) > max {
			max = len(str)
		}
	}
	return
}

/*
	type artStr

	ex.

  .oooo.     .oooo.     .o   .ooooo.        88   .o    .oooo.        88   .oooo.       .ooo
.dP""Y88b   d8P'`Y8b  o888  888' `Y88.     .8' o888   d8P'`Y8b      .8'  d8P'`Y8b    .88'
      ]8P' 888    888  888  888    888    .8'   888  888    888    .8'  888    888  d88'
    .d8P'  888    888  888   `Vbood888   .8'    888  888    888   .8'   888    888 d888P"Ybo.
  .dP'     888    888  888        888'  .8'     888  888    888  .8'    888    888 Y88[   ]88
.oP     .o `88b  d88'  888      .88P'  .8'      888  `88b  d88' .8'     `88b  d88' `Y88   88P
8888888888  `Y8bd8P'  o888o   .oP'     88      o888o  `Y8bd8P'  88       `Y8bd8P'   `88bod8'



                              .oooooo..o
                             d8P'    `Y8
                             Y88bo.      oooo  oooo  ooo. .oo.
                              `"Y8888o.  `888  `888  `888P"Y88b
                                  `"Y88b  888   888   888   888
                             oo     .d8P  888   888   888   888
                             8""88888P'   `V88V"V8P' o888o o888o



              .oooo.     .o        .oooo.     .oooo.         oooooooo   oooooooo
             d8P'`Y8b  o888      .dP""Y88b   d8P'`Y8b       dP"""""""  dP"""""""
            888    888  888            ]8P' 888    888     d88888b.   d88888b.
            888    888  888          <88b.  888    888         `Y88b      `Y88b
            888    888  888  o8o      `88b. 888    888 o8o       ]88        ]88
            `88b  d88'  888  `"' o.   .88P  `88b  d88' `"' o.   .88P  o.   .88P
             `Y8bd8P'  o888o o8o `8bd88P'    `Y8bd8P'  o8o `8bd88P'   `8bd88P'
                             `"'                       `"'
*/
type artStr []rowArtStr

func (str artStr) Print(w io.Writer) {
	for _, rowArtStr := range str {
		for _, row := range rowArtStr {
			fmt.Fprintln(w, row)
		}
	}
}

func (str artStr) setTermPos(updown string, leftright string, term *term) artStr {
	switch leftright {
	case "left":
	case "center":
		str = str.termWidthCenter(term.Width)
	case "right":
		str = str.termWidthRight(term.Width)
	}
	switch updown {
	case "top":
	case "center":
		str = str.termHeightCenter(term.Height)
	case "bottom":
		str = str.termHeightBottom(term.Height)
	}
	return str
}

func (str artStr) termWidthCenter(tmWidth int) artStr {
	diff := (tmWidth - str.Width()) / 2
	for i, rowArtStr := range str {
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	return str
}

func (str artStr) termWidthRight(tmWidth int) artStr {
	diff := (tmWidth - str.Width())
	for i, rowArtStr := range str {
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	return str
}

func (str artStr) termHeightCenter(tmHeight int) artStr {
	diff := (tmHeight - str.Height()) / 2
	var empty rowArtStr
	for i := 0; i < diff; i++ {
		empty = append(empty, "")
	}
	str = append([]rowArtStr{empty}, str...)
	str = append(str, empty)
	return str
}

func (str artStr) termHeightBottom(tmHeight int) artStr {
	diff := (tmHeight - str.Height())
	var empty rowArtStr
	for i := 0; i < diff; i++ {
		empty = append(empty, "")
	}
	str = append([]rowArtStr{empty}, str...)
	return str
}

func (str artStr) setPos(pos string) artStr {
	switch position {
	case "left":
	case "center":
		str = str.Center()
	case "right":
		str = str.Right()
	}
	return str
}

func (str artStr) Center() artStr {
	maxWidth := str.Width()
	for i, rowArtStr := range str {
		if maxWidth == rowArtStr.MaxWidth() {
			continue
		}
		diff := (maxWidth - rowArtStr.MaxWidth()) / 2
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	return str
}

func (str artStr) Right() artStr {
	maxWidth := str.Width()
	for i, rowArtStr := range str {
		if maxWidth == rowArtStr.MaxWidth() {
			continue
		}
		diff := maxWidth - rowArtStr.MaxWidth()
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	return str
}

func (str artStr) Height() (height int) {
	for _, rowArtStr := range str {
		height += len(rowArtStr)
	}
	return
}

func (str artStr) Width() (maxWidth int) {
	for _, rowArtStr := range str {
		if w := rowArtStr.MaxWidth(); w > maxWidth {
			maxWidth = w
		}
	}
	return
}
