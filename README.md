# timelog

[![LICENSE](https://img.shields.io/github/license/qbart/timelog)](https://github.com/qbart/timelog/blob/master/LICENSE)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbart/timelog)](https://goreportcard.com/report/github.com/qbart/timelog)
[![Last commit](https://img.shields.io/github/last-commit/qbart/timelog)](https://github.com/qbart/timelog/commits/master)



Time logging in CLI.

## Install autocompleter (bash + fzf using complete)

```
timelog autocomplete install >> ~/.bash_profile
```

## Usage

### Print current timelog

```
timelog
```


```
3  row(s)
---
2020-02-08 07:36 07:37 hello
2020-02-08 07:37 07:43 world
2020-02-08 08:21 ...   another
---
7m0s
```

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
