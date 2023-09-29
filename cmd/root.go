/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/Mrpye/golib/dir"
	golog "github.com/Mrpye/golib/log"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
)

var manifest_file string
var disable_required bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hauler",
	Short: "hauler is a template for creating a CLI tool for GitHub Actions Workflows",
	Long:  `hauler is a template for creating a CLI tool for GitHub Actions Workflows.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Print(err)
	}
}

func GenerateDoc() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "gen_docs",
		Short: "This command will build the documents for the cli",
		Long: `
Description:
This command will build the documents for the cli

Example Command:
- hauler gen_docs [--dir ./documents]
		`,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				if f.Name == "file" {
					cmd.Flags().SetAnnotation(f.Name, cobra.BashCompOneRequiredFlag, []string{"false"})
				}
			})
			rootCmd.MarkPersistentFlagDirname("file")
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			if manifest_file == "" {
				manifest_file = "./documents"
			}
			if !dir.DirExists(manifest_file) {
				//Make the directory if it does not exist
				err := os.MkdirAll(manifest_file, os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}
			}
			//******************************************
			// Create the documents directory cli
			//******************************************
			err := doc.GenMarkdownTree(rootCmd, manifest_file)
			if err != nil {
				log.Fatal(err)
			}
			//******************************************
			//Create documents for the actions
			//******************************************
			err = wf_client.BuildActionDoc(path.Join(manifest_file, "actions.md"))
			if err != nil {
				return err
			}

			golog.ActionLogOK("Documents Generated", '-')
			return nil
		},
	}
	cmd.Flags().StringVarP(&manifest_file, "dir", "", "", "Used to specify the directory to store the documents")
	manifest_Cmd.MarkFlagDirname("dir")
	return cmd
}

func init() {

	rootCmd.PersistentFlags().BoolP("help", "", false, "help for hauler")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(GenerateDoc())
}
