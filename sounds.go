package main

import _ "embed"

//go:embed ok.wav
var dataSoundOk []byte

//go:embed fail.wav
var dataSoundFail []byte
