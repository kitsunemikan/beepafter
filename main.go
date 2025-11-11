package main

import (
    "fmt"
    "os"
    "time"
    "bytes"
    "os/exec"

    "github.com/gopxl/beep"
    "github.com/gopxl/beep/wav"
    "github.com/gopxl/beep/speaker"
)

func playAndWait(streamer beep.Streamer) {
    done := make(chan struct{})
    speaker.Play(beep.Seq(streamer, beep.Callback(func() {
        done <- struct{}{}
    })))

    <-done
}

func main() {
    if len(os.Args) == 1 {
        fmt.Printf("Usage:\n    beep-on-done APP ARGS...\n")
        return
    }

    soundOk, format, err := wav.Decode(bytes.NewReader(dataSoundOk))
    if err != nil {
        fmt.Printf("error: decode ok sound: %v\n", err)
        os.Exit(1)
    }
    defer soundOk.Close()

    soundFail, format, err := wav.Decode(bytes.NewReader(dataSoundFail))
    if err != nil {
        fmt.Printf("error: decode ok sound: %v\n", err)
        os.Exit(1)
    }
    defer soundFail.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second / 10))

    cmd := exec.Command(os.Args[1], os.Args[2:]...)
    if cmd.Err != nil {
        fmt.Printf("error: %v\n", cmd.Err)
        os.Exit(1)
    }

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err = cmd.Run()
    if err != nil {
        playAndWait(soundFail)

        errExit, ok := err.(*exec.ExitError)
        if ok {
            fmt.Printf("Process exited with error code %d\n", errExit.ExitCode())
        } else {
            fmt.Printf("error: run command %v: %v\n", os.Args[1:], err)
            os.Exit(1)
        }
    } else {
        playAndWait(soundOk)
    }
}

