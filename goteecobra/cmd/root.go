/*
Copyright © 2020 Alan Diaz <adiazny@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var canAppend bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goteecobra",
	Short: "Tee utility written in Go",
	Long: `The tee utility copies standard input to standard output, making a copy in zero or more
	files.  The output is unbuffered.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: goTee,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goteecobra.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&canAppend, "Append", "a", false, "Append the output to the files rather than overwriting them.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".goteecobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".goteecobra")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func goTee(cmd *cobra.Command, args []string) {
	multiWriter := io.MultiWriter(os.Stdout)
	writers := make([]io.Writer, 1)
	writers[0] = os.Stdout

	if len(args) > 0 {

		for i := 0; i < len(args); i++ {
			file := CreateFile(args[i], canAppend)
			defer file.Close()
			writers = append(writers, file)
		}

		multiWriter = io.MultiWriter(writers...)
	}

	teeReader := io.TeeReader(os.Stdin, multiWriter)
	ioutil.ReadAll(teeReader)
}

// CreateFile returns a file with append or no append ability
func CreateFile(fileName string, canAppend bool) *os.File {
	fileMode := os.O_TRUNC

	if canAppend {
		fileMode = os.O_APPEND
	}

	file, err := os.OpenFile(fileName, fileMode|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	return file

}
