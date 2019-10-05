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
  -p string
        left (default "left")
  -position string
        left (default "left")
```

You can choose from the following fonts

https://github.com/common-nighthawk/go-figure#supported-fonts

## example

```
$ ./lddate -f %Y/%m/%d$'\n'%H:%M:%S
  .oooo.     .oooo.     .o   .ooooo.        88   .o    .oooo.        88   .oooo.         .o
.dP""Y88b   d8P'`Y8b  o888  888' `Y88.     .8' o888   d8P'`Y8b      .8'  d8P'`Y8b      .d88
      ]8P' 888    888  888  888    888    .8'   888  888    888    .8'  888    888   .d'888
    .d8P'  888    888  888   `Vbood888   .8'    888  888    888   .8'   888    888 .d'  888
  .dP'     888    888  888        888'  .8'     888  888    888  .8'    888    888 88ooo888oo
.oP     .o `88b  d88'  888      .88P'  .8'      888  `88b  d88' .8'     `88b  d88'      888
8888888888  `Y8bd8P'  o888o   .oP'     88      o888o  `Y8bd8P'  88       `Y8bd8P'      o888o



  .oooo.     .oooo.         .oooo.     .oooo.         oooooooo  .ooooo.
.dP""Y88b  .dP""Y88b      .dP""Y88b   d8P'`Y8b       dP""""""" d88'   `8.
      ]8P'       ]8P'           ]8P' 888    888     d88888b.   Y88..  .8'
    .d8P'      <88b.          <88b.  888    888         `Y88b   `88888b.
  .dP'          `88b. o8o      `88b. 888    888 o8o       ]88  .8'  ``88b
.oP     .o o.   .88P  `"' o.   .88P  `88b  d88' `"' o.   .88P  `8.   .88P
8888888888 `8bd88P'   o8o `8bd88P'    `Y8bd8P'  o8o `8bd88P'    `boood8'

```

```
$ ./lddate -f %Y/%m/%d$'\n'%H:%M:%S -font standard
 ____     ___    _    ___       __  _    ___       __   ___    _  _
 |___ \   / _ \  / |  / _ \     / / / |  / _ \     / /  / _ \  | || |
   __) | | | | | | | | (_) |   / /  | | | | | |   / /  | | | | | || |_
  / __/  | |_| | | |  \__, |  / /   | | | |_| |  / /   | |_| | |__   _|
 |_____|  \___/  |_|    /_/  /_/    |_|  \___/  /_/     \___/     |_|
  ____    _____       _____   ____        ____    _____
 |___ \  |___ /   _  |___ /  |___ \   _  | ___|  |___ /
   __) |   |_ \  (_)   |_ \    __) | (_) |___ \    |_ \
  / __/   ___) |  _   ___) |  / __/   _   ___) |  ___) |
 |_____| |____/  (_) |____/  |_____| (_) |____/  |____/
```

```
$ ./lddate -format %Y/%m/%d$'\n'%H:%M:%S -p center
  .oooo.     .oooo.     .o   .ooooo.        88   .o    .oooo.        88   .oooo.       .ooo
.dP""Y88b   d8P'`Y8b  o888  888' `Y88.     .8' o888   d8P'`Y8b      .8'  d8P'`Y8b    .88'
      ]8P' 888    888  888  888    888    .8'   888  888    888    .8'  888    888  d88'
    .d8P'  888    888  888   `Vbood888   .8'    888  888    888   .8'   888    888 d888P"Ybo.
  .dP'     888    888  888        888'  .8'     888  888    888  .8'    888    888 Y88[   ]88
.oP     .o `88b  d88'  888      .88P'  .8'      888  `88b  d88' .8'     `88b  d88' `Y88   88P
8888888888  `Y8bd8P'  o888o   .oP'     88      o888o  `Y8bd8P'  88       `Y8bd8P'   `88bod8'



           .oooo.     .oooo.         .oooo.     .oooo.         .oooo.     oooooooo
          d8P'`Y8b   d8P'`Y8b      .dP""Y88b  .dP""Y88b       d8P'`Y8b   dP"""""""
         888    888 888    888           ]8P'       ]8P'     888    888 d88888b.
         888    888 888    888         <88b.      <88b.      888    888     `Y88b
         888    888 888    888 o8o      `88b.      `88b. o8o 888    888       ]88
         `88b  d88' `88b  d88' `"' o.   .88P  o.   .88P  `"' `88b  d88' o.   .88P
          `Y8bd8P'   `Y8bd8P'  o8o `8bd88P'   `8bd88P'   o8o  `Y8bd8P'  `8bd88P'
                               `"'                       `"'
```

```
$ ./lddate -format %Y/%m/%d$'\n'%a$'\n'%H:%M:%S -p center
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

```

## License

MIT
