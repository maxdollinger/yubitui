package clipboard

import (
	"errors"
	"os/exec"
)

type Clipboard struct {
	cp    string
	paste string
}

type ClipbardNotFoundError struct{}

func (ce ClipbardNotFoundError) Error() string {
	return "no Clipboard found"
}

var (
	cpCmds    = []string{"wl-copy", "pb-copy"}
	pasteCmds = []string{"wl-paste", "pb-paste"}
)

func InitClipboard() (*Clipboard, error) {
	cmdIdx := -1
	for i, clip := range cpCmds {
		which := exec.Command("which", clip)
		err := which.Run()
		if err == nil {
			cmdIdx = i
			break
		}
	}

	if cmdIdx < 0 {
		return nil, errors.New("no clipboard found")
	}

	return &Clipboard{
		cp:    cpCmds[cmdIdx],
		paste: pasteCmds[cmdIdx],
	}, nil
}

func (c *Clipboard) Copy(str string) error {
	cmd := exec.Command(c.cp, str)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (c *Clipboard) Paste() string {
	cmd := exec.Command(c.paste)

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(output)
}
