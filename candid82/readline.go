package peterh

import (
	"os"

	"github.com/candid82/liner"
)

type Readline struct {
	impl        *liner.State
	prompt      string
	historyFile string
}

func New(prompt string, historyFile string) (*Readline, error) {
	impl := liner.NewLiner()
	impl.SetCtrlCAborts(true)
	var f *os.File
	var err error
	if historyFile != "" {
		f, err = os.Open(historyFile)
		if err != nil {
			return nil, err
		}
		_, err = impl.ReadHistory(f)
		_ = f.Close()
	}
	rl := &Readline{
		impl:        impl,
		prompt:      prompt,
		historyFile: historyFile,
	}
	return rl, err
}

func (rl *Readline) Readline() (string, error) {
	return rl.impl.Prompt(rl.prompt)
}

func (rl *Readline) AppendHistory(text string) error {
	rl.impl.AppendHistory(text)
	return nil
}

func (rl *Readline) Close() error {
	if rl.historyFile == "" {
		return nil
	}
	f, err := os.Create(rl.historyFile)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = rl.impl.WriteHistory(f)
	if err != nil {
		return err
	}
	return nil
}
