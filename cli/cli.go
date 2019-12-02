package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AppProvider ...
type AppProvider interface {
	Run()
	Command() *cobra.Command
}

type app struct {
	cmd        *RootCommand
	configFile string
}

// New ...
func New(root *RootCommand) AppProvider {
	return &app{cmd: root}
}

func (a *app) Command() *cobra.Command {
	return a.cmd.Command
}

func (a *app) Run() {
	a.cmd.PersistentFlags().StringVar(&a.configFile, "config", "", "set config file")
	cobra.OnInitialize(a.initConfig)

	a.cmd.SetOutput(os.Stdout)
	if err := a.cmd.Execute(); err != nil {
		a.cmd.SetOutput(os.Stderr)
		a.cmd.Println(err)
		os.Exit(1)
	}
}

func (a *app) initConfig() {
	if a.configFile != "" {
		viper.SetConfigFile(a.configFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
