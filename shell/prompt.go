/* prompt
promopt purpose
1. show current promopt path
2. log past path
*/

package shell

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/pkg/errors"
)

type promopt struct {
	username string
	hostname string
	homePath string
}

func NewPromopt() (*promopt, error) {
	u, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "could not get current user")
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, errors.Wrap(err, "could not get hostname")
	}
	return &promopt{
		username: u.Name,
		hostname: hostname,
		homePath: u.HomeDir,
	}, nil
}

func (p *promopt) String() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "could not get current work dir")
	}

	if strings.HasPrefix(wd, p.homePath) {
		wd = "~" + strings.TrimPrefix(wd, p.homePath)
	}
	return fmt.Sprintf("[%s@%s %s]$ ", p.username, p.hostname, wd), nil
}
