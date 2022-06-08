package cli

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

// AreYouSure displays user prompt for confirmation.
func AreYouSure(msg string, yes func(), no func()) {
	red := color.New(color.FgRed)
	red.Print(msg, " Y/n: ")

	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')

	confirmed := s == "\n"
	s = strings.TrimSpace(s)
	confirmed = confirmed || (len(s) == 1 && (s[0] == 'y' || s[0] == 'Y'))
	if confirmed {
		yes()
	} else {
		no()
	}
}
