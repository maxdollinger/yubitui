package clipboard

import (
	"errors"
	"os/exec"
)

type Clipboard struct {
	cmd string
}

type ClipbardNotFoundError struct{}

func (ce ClipbardNotFoundError) Error() string {
	return "no Clipboard found"
}

var clipboards = []string{"wl-copy", "pb-copy"}

func InitClipboard() (*Clipboard, error) {
	cmd := ""
	for _, clip := range clipboards {
		which := exec.Command("which", clip)
		err := which.Run()
		if err == nil {
			cmd = clip
			break
		}
	}

	if cmd == "" {
		return nil, errors.New("no clipboard found")
	}

	return &Clipboard{
		cmd: cmd,
	}, nil
}

func (c *Clipboard) Copy(str string) error {
	copyCmd := exec.Command(c.cmd, str)

	err := copyCmd.Run()
	if err != nil {
		return err
	}

	return nil
}
