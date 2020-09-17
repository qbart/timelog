# timelog

[![LICENSE](https://img.shields.io/github/license/qbart/timelog)](https://github.com/qbart/timelog/blob/master/LICENSE)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbart/timelog)](https://goreportcard.com/report/github.com/qbart/timelog)
[![Last commit](https://img.shields.io/github/last-commit/qbart/timelog)](https://github.com/qbart/timelog/commits/master)


```
Time logging in CLI

Usage:
  timelog [flags]
  timelog [command]

Available Commands:
  adjust       Adjusts time between entries
  archive      Archive data file
  autocomplete Autocomplete for entries
  clear        Clears all entries
  help         Help about any command
  polybar      Polybar configuration
  qlist        Prints all quicklist entries
  start        Starts a new time entry
  stop         Stops active time entry
  version      Prints software version

Flags:
  -h, --help   help for timelog

Use "timelog [command] --help" for more information about a command.
```

## Install

```
git clone git@github.com:qbart/timelog.git
cd timelog
make build
```

## Usage

### Print current timelog

```
timelog
```

![timelog](./doc/timelog.png)


### Start next task

```
timelog start <comment>
```

### Stop current task

```
timelog stop
```

### Clear

```
timelog clear
```

1. Current timelog will be printed.
2. Once confirmed local database will be cleared :warning:.

### Adjust time

```
timelog adjust
```

1. Console UI will start (use arrows or hjkl).
2. Enter to continue.
3. After changes you will see git-like diff to accept/reject changes.

![timelog](./doc/timelog_adjust_step1.png)
![timelog](./doc/timelog_adjust_step2.png)

### Archive

```
timelog archive
```

1. Current timelog will be printed.
2. Once confirmed file will be moved to archive (`~/.config/timelog/archive/`)

## Configuration

### Install autocompleter (bash + fzf using complete)

Tested only in Ubuntu (PR appreciated for other OSes)
```
timelog autocomplete install >> ~/.bash_profile
```

```
timelog [hit TAB]
```
![timelog](./doc/timelog_autocomplete_cmds.png)

### Quicklist

Quicklist is a data source for autocompleter (fzf). Tasks should not contain whitespaces.

```
vim ~/.config/timelog/config.ini
```

```
[quicklist]
task-1
task-2
task-3
hello
```

```
task start [hit TAB]
```
![timelog](./doc/timelog_autocomplete_qlist.png)

## Polybar (or any other bar) integration

You can integrate timelog with your custom bar:
```
timelog polybar format "FORMAT"
```

`FORMAT` exposes following vars in `go` template:
```go
type polybarItem struct {
	Comment         string // task comment
	Duration        string // last task duration
	Total           string // tasks total duration
	Count           int    // task count
	CountNotZero    bool   //
	TotalGtDuration bool   // true when total > duration
}
```

### Polybar example

```
[module/timelog]
type=custom/script
interval=10
exec=timelog polybar format "{{if .CountNotZero }}%{F#011814}%{B#24f5bf} {{.Comment}} %{B-}%{B#0adba6} {{.Duration}} %{B-}{{ if .TotalGtDuration}}%{B#08aa81} {{.Total}} %{B-}{{ end }}%{F-}{{ end }}"
```

#### When duration equals total
![timelog](./doc/timelog_polybar_same.png)
#### When duration does not equal total
![timelog](./doc/timelog_polybar.png)

## Config

## How to contribute?

Ask first before any implementation.
Possible todos:
- edit comments `timelog edit` (similar to `timelog adjust` editor)
- quicklist management from CLI ie. `timelog qlist.add ENTRY`
- multi-autocomplete i.e `timelog start [TAB]`, then `timelog start autocompleted-1 [TAB]` <- currently this will replace `autocompleted-1` with new qlist entry, goal is to append next one
- autocomplete for other OSes (OSx integration anyone?)
