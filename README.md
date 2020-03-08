# timelog

[![LICENSE](https://img.shields.io/github/license/qbart/timelog)](https://github.com/qbart/timelog/blob/master/LICENSE)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbart/timelog)](https://goreportcard.com/report/github.com/qbart/timelog)
[![Last commit](https://img.shields.io/github/last-commit/qbart/timelog)](https://github.com/qbart/timelog/commits/master)



Time logging in CLI.

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

## How to contribute?

Ask first before any implementation.
