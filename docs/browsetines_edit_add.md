## browsetines edit add

Add an activity

### Synopsis

Add an activity with provided name to routine
with given ID.
If position is not specified or is -1, activity will be
inserted at last position.
If you provide a time, a stopwatch will run next
activity in routine when it reaches zero.
Flag prompt overrides delay.

```
browsetines edit add <id> <name> [<position>] [<time>] [flags]
```

### Options

```
  -d, --delay duration   delay before starting activity
  -h, --help             help for add
  -p, --prompt           prompt before starting activity
  -u, --urls strings     urls to open when activity starts
```

### SEE ALSO

* [browsetines edit](browsetines_edit.md)	 - Edit routine

###### Auto generated by spf13/cobra on 17-Jan-2021
