package command

import (
	"bufio"
	"fmt"
	"os"

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
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
