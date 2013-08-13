package commands

import (
	"fmt"
	"github.com/gedex/ginsta/clients"
	"github.com/gedex/ginsta/utils"
	"os"
	"text/template"
)

var cmdHelp = &Command{
	Usage: "help [command]",
	Short: "Show help",
	Long:  `Shows usage for a command.`,
}

func init() {
	cmdHelp.Callback = runHelp
}

func runHelp(r *Runner, cmd *Command, args []string) {
	if len(args) == 0 {
		printUsage()
		os.Exit(0)
	}

	if len(args) > 1 {
		utils.Check(fmt.Errorf("too many arguments"))
	}

	for _, cmd := range All() {
		if cmd.Name() == args[0] {
			cmd.PrintUsage()
			os.Exit(0)
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic: %q. Run '%s help'.\n", args[0], clients.Name)
	os.Exit(2)
}

var usageTemplate = template.Must(template.New("usage").Parse(
	`Usage: {{.Name}} [command] [options] [arguments]

Users Commands:{{range .UsersCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-18s"}}  {{.Short}}{{end}}{{end}}{{end}}

Relationships Commands:{{range .RelationshipsCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-18s"}}  {{.Short}}{{end}}{{end}}{{end}}

Media Commands:{{range .MediaCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-18s"}}  {{.Short}}{{end}}{{end}}{{end}}

Comments Commands:{{range .CommentsCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-18s"}}  {{.Short}}{{end}}{{end}}{{end}}

Basic Commands:{{range .BasicCommands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-18s"}}  {{.Short}}{{end}}{{end}}{{end}}
`))

func printUsage() {
	usageTemplate.Execute(os.Stdout, struct {
		Name                  string
		UsersCommands         []*Command
		RelationshipsCommands []*Command
		MediaCommands         []*Command
		CommentsCommands      []*Command
		BasicCommands         []*Command
	}{
		clients.Name,
		UsersCommands,
		RelationshipsCommands,
		MediaCommands,
		CommentsCommands,
		BasicCommands,
	})
}

func usage() {
	printUsage()
	os.Exit(2)
}
