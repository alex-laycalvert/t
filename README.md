# t

Yet another time-tracking tool I wanted to build because I got bored. There's other tools that do what I want, I just wanted to make my own.

## Installation

### Manual (Only way right now)

I made this on Arch linux, so it might not work everywhere.

```bash
# Clone the repo
git clone https://github.com/alex-laycalvert/t

# Navigate
cd ./t

# Install
make install
```

## Usage

### Start a new project timer

```bash
t start <project>
```

### See project timer details

```bash
t show <project>
```

### Stop a project timer

```bash
t stop <project>
```

### Show projects

```bash
t projects
```

## What I Want

- Start a timer for a project

```bash
t start <project>
```

- Let multiple projects run at the same time
- Autocomplete `<project>` for commands
- Stop a project timer and see the elapsed time, pauses, etc.

```bash
t stop <project>
```

- See how long I've worked on a project

```bash
t show <project>
```

## TODO

- [ ] Better error handling when unqiue constraint fails
- [x] Better DB creation / config handling
- [x] Configuration file
- [ ] The `show` command with no sub args should probably show all projects instead of a separate `projects` command
