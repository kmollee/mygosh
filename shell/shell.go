package shell

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Shell struct {
	r    *bufio.Reader
	text string // welcome text
	p    *promopt
}

func NewShell(r io.Reader, t string) (*Shell, error) {
	p, err := NewPromopt()
	if err != nil {
		return nil, err
	}
	return &Shell{
		r:    bufio.NewReader(r),
		text: t,
		p:    p,
	}, nil
}

func (s *Shell) Init() {
	fmt.Println(s.text)
	// TODO: add configure
}

func (s *Shell) Read() ([]string, error) {

	content, err := s.r.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "could not read command from reader")
	}
	content = strings.TrimSpace(content)
	content = strings.TrimRight(content, "\n")
	return strings.Fields(content), err
}

// show prompt
func (s *Shell) Prompt() {
	c, _ := s.p.String()
	fmt.Print(c)

}
