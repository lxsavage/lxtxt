# LXTXT

A modal text editor written in Go and inspired by VIM.

## Installation

Installation scripts are coming soon, but for now the project needs to manually
build for your system.

Dependencies
- Go CLI >=1.26.0

To install:

```bash
make install
```

See the makefile for more commands.

## Usage

```sh
lxtxt <filepath>
```

### Modes

With being a modal editor, LXTXT has multiple modes that allow the user to do
different types of actions. These are detailed below.

#### `NORMAL`

This is the default mode. You can use navigational commands to move around the
buffer as well as save/quit from here. The following commands are available by
directly typing them:

- `arrow keys/hjkl`: move around by one character
- `_`: move to beginning of the current line
- `$`: move to end of the current line
- `D`: delete the current line
- `Q`: quit the editor without saving
- `W`: write the buffer to the file
- `!`: discard changes and revert to initial state when opening file
- `?`: show the manpage (if installed)

> [!NOTE]
> In most actions, you can type a number before the action command character to
> repeat that command multiple times!

In addition to these commands, the following commands can be typed to switch
into other modes:

- `i`: enter `INSERT` mode

> [!NOTE]
> You can return to `NORMAL` mode from any other mode by pressing `esc`.

####  `INSERT`

`INSERT` is primarily for directly modifying text. Typing text in this mode will
insert it where the cursor currently sits, and arrow keys can still be used to
move around here similar to `NORMAL` mode.
