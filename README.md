## `beepafter`

Play a sound after command execution. Different sounds are played based on command's success or failure (determined by its exit code).

### Usage

```powershell
beepafter echo hello             # will play success jingle
beepafter rm filedoesntexist.txt # will play failure jingle
```

### Install

```powershell
go install github.com/kitsunemikan/beepafter@latest
```
