package simple

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Readline struct {
	prompt string
	reader bufio.Reader
	writer io.Writer
}

func New(prompt string, historyFile string) (*Readline, error) {
	_ = historyFile
	return &Readline{
		prompt: prompt,
		reader: *bufio.NewReader(os.Stdin),
		writer: os.Stdout,
	}, nil
}

func (rl *Readline) Readline() (string, error) {
	_, _ = fmt.Fprintf(rl.writer, rl.prompt)
	return rl.reader.ReadString('\n')
}

func (rl *Readline) AppendHistory(_ string) error {
	return nil
}

func (rl *Readline) Close() error {
	return nil
}
