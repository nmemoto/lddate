# lddate - Large Display Date on the terminal

## How to use

### Build

```
$ git clone git@github.com:nmemoto/lddate.git
$ cd lddate
$ go build
```

### Run

```
$ ./lddate
  .oooo.     .oooo.     .o   .ooooo.        88   .o    .oooo.        88   .oooo.         .o       .oooo.     .oooo.         .oooo.     .oooo.         .oooo.       .ooo
.dP""Y88b   d8P'`Y8b  o888  888' `Y88.     .8' o888   d8P'`Y8b      .8'  d8P'`Y8b      .d88      d8P'`Y8b  .dP""Y88b      .dP""Y88b  .dP""Y88b      .dP""Y88b    .88'
      ]8P' 888    888  888  888    888    .8'   888  888    888    .8'  888    888   .d'888     888    888       ]8P'           ]8P'       ]8P'           ]8P'  d88'
    .d8P'  888    888  888   `Vbood888   .8'    888  888    888   .8'   888    888 .d'  888     888    888     <88b.          <88b.      <88b.          .d8P'  d888P"Ybo.
  .dP'     888    888  888        888'  .8'     888  888    888  .8'    888    888 88ooo888oo   888    888      `88b. o8o      `88b.      `88b. o8o   .dP'     Y88[   ]88
.oP     .o `88b  d88'  888      .88P'  .8'      888  `88b  d88' .8'     `88b  d88'      888     `88b  d88' o.   .88P  `"' o.   .88P  o.   .88P  `"' .oP     .o `Y88   88P
8888888888  `Y8bd8P'  o888o   .oP'     88      o888o  `Y8bd8P'  88       `Y8bd8P'      o888o     `Y8bd8P'  `8bd88P'   o8o `8bd88P'   `8bd88P'   o8o 8888888888  `88bod8'
```

### Options

```
$ ./lddate -h
Usage of ./lddate:
  -d int
        update date duration (default 1)
  -duration int
        update date duration (default 1)
  -f string
        date format (default "%Y/%m/%d %H:%M:%S")
  -font string
        font (default "roman")
  -format string
        date format (default "%Y/%m/%d %H:%M:%S")
```

You can choose from the following fonts

https://github.com/common-nighthawk/go-figure#supported-fonts

## License

MIT