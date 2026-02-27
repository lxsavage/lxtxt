# LXTXT

A modal text editor written in Go and inspired by VIM.

## Installation

WIP

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

In addition to these commands, the following commands can be typed to switch
into other modes:

- `i`: enter `INSERT` mode
- `:`: **NOT YET IMPLEMENTED** enter `COMMAND` mode

> [!NOTE]
> You can return to `NORMAL` mode from any other mode by pressing `esc`

####  `INSERT`

`INSERT` is primarily for directly modifying text. Typing text in this mode will
insert it where the cursor currently sits, and arrow keys can still be used to
move around here similar to `NORMAL` mode.

#### `COMMAND`

> [!IMPORTANT]
> This mode is currently not implemented but will be before the 1.0.0 release.

`COMMAND` is for typing textual commands with arguments to perform actions that
are not possible in `NORMAL` mode.
