package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

// CmdList prints list
func CmdList(c *cli.Context) {

	var fp *os.File
	var err error

	fp, err = os.Open(c.Args()[0])

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "/* Begin") {
			fs := strings.Fields(t)
			fmt.Println(fs[2])
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
