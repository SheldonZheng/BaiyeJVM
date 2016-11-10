package main
import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage()
	flag.BoolVar(&cmd.helpFlag, "help", false,"print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
}
