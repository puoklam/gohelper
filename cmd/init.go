package cmd

import (
	"github.com/puoklam/gohelper/cp"
	"github.com/puoklam/gohelper/tmpl"
	"github.com/spf13/cobra"
)

var httpTmpl, cliTmpl bool
var version string

var cmdInit *cobra.Command = &cobra.Command{
	Use:   "init [project name]",
	Short: "Init prohect",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, module := range args {
			var tpl map[string][]byte
			if httpTmpl {
				tpl = tmpl.Http(module, version)
			} else if cliTmpl {
				tpl = tmpl.Cli(module, version)
			} else {
				tpl = tmpl.Basic(module, version)
			}
			for dst, src := range tpl {
				if err := cp.File(dst, src); err != nil {
					return err
				}
			}
		}
		// fmt.Println("Initialising: " + strings.Join(args, " "))
		// error will be caught at Execute
		return nil
	},
}

func initCmdInit() {
	cmdInit.Flags().BoolVarP(&httpTmpl, "http", "", false, "Use http server application template")
	cmdInit.Flags().BoolVarP(&cliTmpl, "cli", "", false, "Use command line application template")
	cmdInit.Flags().StringVarP(&version, "version", "v", "1.16", "Project Go version")
}
