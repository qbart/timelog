# timelog

Time logging in CLI.

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
