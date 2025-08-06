# NerdFont-Cheat-Sheet

`nerdfont.csv` is a generated CSV that includes all of the Nerd Font icons and their description, as found on [Nerd Font's Cheat Sheet website](https://www.nerdfonts.com/cheat-sheet).

`nerdfont.txt` is a more readable version of that file

`main.go` and `go.mod` are files needed for nerdfonts-cs cli tool

## Usage

You will need to run `nerdfonts-cs` inside an terminal that has a Nerd Font installed

## Installation

A simple CLI tool not much better than `cat nerdfont.csv | grep debian` can be installed with golang
```bash
go install
```

And test with 
```bash
nerdfonts-cs debian
```

## Updating the files 

A javascript file is included `updateScript.js` with instructions on how to generate `nerdfont.csv` and `nerdfont.txt`
