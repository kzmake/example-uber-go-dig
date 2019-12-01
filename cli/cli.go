package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var b Builder

func init() {
	r := NewRootCommand()
	os := NewOperationCommands()

	b = New().Root(r).Operations(os...)
}

// Builder defines ...
type Builder interface {
	Root(cmd *cobra.Command) Builder
	Operations(cmds ...*cobra.Command) Builder
	Targets(cmds ...map[string]*cobra.Command) Builder
	Build() App
}

// App defines ...
type App interface {
	Pre(func(*cobra.Command, []string) error)
	Run()
	Post(func(*cobra.Command, []string) error)
}

type builder struct {
	root       *cobra.Command
	operations map[string]*cobra.Command
	targets    map[string][]*cobra.Command
}

// New builds cli builder.
func New() Builder {
	return &builder{
		operations: map[string]*cobra.Command{},
		targets:    map[string][]*cobra.Command{},
	}
}

func (b *builder) Root(cmd *cobra.Command) Builder {
	b.root = cmd
	return b
}

func (b *builder) Operations(cmds ...*cobra.Command) Builder {
	operations := map[string]*cobra.Command{}
	for _, c := range cmds {
		operations[c.Name()] = c
	}
	b.operations = update(b.operations, operations)

	return b
}

// Targets ...
func Targets(targets ...map[string]*cobra.Command) Builder { return b.Targets(targets...) }
func (b *builder) Targets(targets ...map[string]*cobra.Command) Builder {
	b.targets = merge(b.targets, list(targets...))

	return b
}

// Build ...
func Build() App { return b.Build() }
func (b *builder) Build() App {
	// add target commands
	for k, v := range b.targets {
		b.operations[k].AddCommand(v...)
	}

	// add operation commands
	operations := []*cobra.Command{}
	for _, op := range b.operations {
		operations = append(operations, op)
	}
	b.root.AddCommand(operations...)

	return &app{cmd: b.root}
}

type app struct {
	cmd        *cobra.Command
	configFile string
}

func (a *app) Pre(f func(*cobra.Command, []string) error) {
	a.cmd.PersistentPreRunE = f
}

func (a *app) Post(f func(*cobra.Command, []string) error) {
	a.cmd.PersistentPostRunE = f
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
