package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	copydir      string
	exe          string
	builddir     string
	excludeTests bool

	rootCmd = &cobra.Command{
		Use:   "buildexecute",
		Short: "Build and execute Go projects",
		Long:  "buildexecute builds go projects based on optional flags",
		Run: func(cmd *cobra.Command, args []string) {
			if copydir != "" {
				dest := "./"

				if builddir != "" {
					dest = builddir
				}

				_, err := exec.Command("cp", "-r", copydir, dest).Output()
				if err != nil {
					log.Fatal(err)
				}

				if exe != "" {
					cmd := exec.Command("go", "build")
					cmd.Dir = dest
					_, err := cmd.Output()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&copydir, "copydir", "c", "", "the directory which is to be copied")
	rootCmd.PersistentFlags().StringVarP(&builddir, "builddir", "b", "", "the directory where the files are to be pasted")
	rootCmd.PersistentFlags().StringVarP(&exe, "exe", "e", "", "build executable")
}
