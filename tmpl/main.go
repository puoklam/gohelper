package tmpl

import (
	"fmt"
	"strings"
)

func modFile(module, version string, pkgs []string) []byte {
	lines := make([]string, 0, len(pkgs)+4)
	lines = append(lines, "module "+module)
	lines = append(lines, "")
	lines = append(lines, "go "+version)
	lines = append(lines, "")
	lines = append(lines, pkgs...)
	return []byte(strings.Join(lines, "\n"))
}

func Basic(module, version string) map[string][]byte {
	files := make(map[string][]byte)
	files["main.go"] = []byte(`/*
basic template
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("basic")
}
`)
	files["go.mod"] = modFile(module, version, nil)
	return files
}

func Http(module, version string) map[string][]byte {
	files := make(map[string][]byte)
	files["main.go"] = []byte(`/*
http server template
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("http server")
}
`)
	files["go.mod"] = modFile(module, version, nil)
	return files
}

func Cli(module, version string) map[string][]byte {
	files := make(map[string][]byte)
	files["main.go"] = []byte(fmt.Sprintf(`/*
cli template
*/
package main

import (
	"fmt"
	"log"
	"os"

	"%s/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`, module))
	files["cmd/root.go"] = []byte(fmt.Sprintf(`/*
cli template
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{Use: "%s"}

func init() {
	fmt.Println("Initialising command")
	initCmdPrint()
	initCmdTimes()
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
}

func Execute() error {
	return rootCmd.Execute()
}
`, module[strings.LastIndex(module, "/")+1:]))
	files["cmd/print.go"] = []byte(`/*
cli template
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var flag1 bool

var cmdPrint *cobra.Command = &cobra.Command{
	Use:   "print [string to print]",
	Short: "Print anything to the screen",
	Long: "print is for printing anything back to the screen. For many years people have printed back to the screen.",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(flag1)
		fmt.Println("Print: " + strings.Join(args, " "))
		// error will be caught at Execute
		return nil
	},
}

func initCmdPrint() {
	cmdPrint.Flags().BoolVarP(&flag1, "f1", "", false, "usage")
}
`)
	files["cmd/echo.go"] = []byte(`/*
cli template
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: "echo is for echoing anything back. Echo works a lot like print, except it has a child command.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
}
`)
	files["cmd/times.go"] = []byte(`/*
cli template
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var echoTimes int

var cmdTimes = &cobra.Command{
	Use:   "times [string to echo]",
	Short: "Echo anything to the screen more times",
	Long: "echo things multiple times back to the user by providing a count and a string.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},
}

func initCmdTimes() {
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
}
`)
	files["go.mod"] = modFile(module, version, []string{"require github.com/spf13/cobra v1.3.0"})
	return files
}
