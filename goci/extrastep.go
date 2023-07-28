package main

import (
	"bytes"
	"os/exec"
)

type extraStep struct {
	step
}

func newextraStep(name, exe, message, proj string, args []string) extrastep {
	s := extraStep{}
	s.step = newStep(name, exe, message, proj, args)
	return s
}

func (s extraStep) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Dir = s.proj
	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}
	if out.Len() > 0 {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: nil,
		}
	}
	return s.message, nil
}
