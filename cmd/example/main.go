package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	// liner "github.com/vandycarlos/readline/chzyer"
	// liner "github.com/vandycarlos/readline/simple"
	// liner "github.com/vandycarlos/readline/peterh"
	liner "github.com/vandycarlos/readline/candid82"
)

func main() {
	rl, err := liner.New(">>> ", "~/.liner-history")
	if err != nil {
		log.Fatal("error creating readline functionality - " + err.Error())
	}
	fmt.Println("liner v0.0.1")
	for {
		line, err := rl.Readline()
		if errors.Is(err, io.EOF) {
			break
		}
		_ = rl.AppendHistory(line)
		fmt.Println(line)
	}
	_ = rl.Close()
	os.Exit(0)
}
