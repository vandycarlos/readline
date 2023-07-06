package chzyer

import (
	"errors"

	"github.com/chzyer/readline"
)

type Readline struct {
	impl readline.Instance
}

func New(prompt string, historyFile string) (*Readline, error) {
	rlx, err := readline.NewEx(&readline.Config{
		Prompt:      prompt,
		HistoryFile: historyFile,
	})
	return &Readline{*rlx}, err
}

func (rl *Readline) Readline() (string, error) {
	line, err := rl.impl.Readline()
	for errors.Is(err, readline.ErrInterrupt) {
		line, err = rl.impl.Readline()
	}
	return line, err
}

func (rl *Readline) AppendHistory(_ string) error {
	// return rl.impl.SaveHistory(text)
	return nil
}

func (rl *Readline) Close() error {
	return rl.impl.Close()
}
