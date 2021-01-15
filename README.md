# browsetines

Command-line tool to control how much time you invest in certain activities. You can create routines, which are collections of activities, and set an amount of time to invest to each one.

## Commands

### run

Runs a routine, if routine was already started, continues from where it was left.
If you leave in the middle of an activity, you will continue from where you left.

```
Usage:
    browsetines run [options] <ID>
    browsetines run -h | --help

Options:
    -h, --help          Show this screen.
    -b, --beginning     Start from first activity.
```

### skip

Skips a number of activities in a routine. By default skips one activity.

```
Usage:
    browsetines skip [options] <ID>
    browsetines skip -h | --help

Options:
    -h, --help                      Show this screen.
    -n <number>, --number <number>  Number of activities to skip.
```

### restart

Restarts a routine, so if you use `run`, you will start from the first activity.

```
Usage:
    browsetines restart <ID>
    browsetines restart -h | --help

Options:
    -h, --help  Show this screen.
```

### remove

Removes a routine.

```
Usage:
    browsetines remove <ID>
    browsetines remove -h | --help

Options:
    -h, --help  Show this screen.
```

### edit

Selects a routine to edit it using subcommands.

```
Usage:
    browsetines edit -h | --help
Options:
    -h, --help  Show this screen.
```

#### add

Edits routine by adding a new activity with given name.

If you specify a time, activity will start a stopwatch and will run next activity in routine when stopwatch reaches zero.

If you don't specify a time, it will display a timer instead.

```
Usage:
    browsetines edit <ID> add [-pdu] <time> <name>
    browsetines edit <ID> add [-p] <name>
    browsetines edit <ID> add -h | --help

Options:
    -h, --help                          Show this screen.
    -p, --prompt                        Prompt before starting timer.
    -d <timespan>, --delay <timespan>   Delay before starting stopwatch.
    -u <URL>..., --url <URL>...        URLs to open when active.
```

#### remove

Removes an activity.

```
Usage:
    browsetines edit <ID> remove <ID>
    browsetines edit <ID> remove -h | --help

Options:
    -h, --help  Show this screen.
```
