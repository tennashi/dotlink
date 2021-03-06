package got

import (
	"context"
	"errors"
	"flag"
	"fmt"
)

type versionCmd struct {
	rootCmd *got
}

func newVersionCmd(rootCmd *got) *versionCmd {
	return &versionCmd{rootCmd: rootCmd}
}

func (v *versionCmd) parse(args []string) error {
	return nil
}

func (v *versionCmd) run(ctx context.Context) error {
	printVersion(v.rootCmd.IOStream.Out)
	return nil
}

type helpCmd struct {
	rootCmd *got
	fs      *flag.FlagSet
}

var help = `got - dotfiles and packages manager

Usage:
  got command [arguments]

Commands:
  version
    print got command version

  clone [repo_name]
    clone your dotfiles repository

  link
    create symbolic links

  clean
    remove all symbolic links

  push
    push dotfiles update to your dotfiles repository

`

func newHelpCmd(rootCmd *got) *helpCmd {
	fs := flag.NewFlagSet("got-help", flag.ContinueOnError)
	fs.SetOutput(rootCmd.IOStream.Err)
	fs.Usage = func() {
		fmt.Fprint(rootCmd.IOStream.Out, help)
		fs.PrintDefaults()
	}
	return &helpCmd{rootCmd: rootCmd, fs: fs}
}

func (c *helpCmd) parse(args []string) error {
	return c.fs.Parse(args)
}

func (c *helpCmd) run(ctx context.Context) error {
	c.fs.Usage()
	if c.fs.NArg() < 2 {
		return fmt.Errorf("must specify sub-command")
	}
	subCmd := c.fs.Arg(1)
	if subCmd != "help" {
		return fmt.Errorf("no such command: %s", subCmd)
	}
	return nil
}

type cloneCmd struct {
	rootCmd *got

	fs *flag.FlagSet
}

func newCloneCmd(rootCmd *got) *cloneCmd {
	fs := flag.NewFlagSet("got-clone", flag.ContinueOnError)
	return &cloneCmd{rootCmd: rootCmd, fs: fs}
}

func (c *cloneCmd) parse(args []string) error {
	return c.fs.Parse(args)
}

func (c *cloneCmd) run(ctx context.Context) error {
	if c.fs.NArg() < 3 {
		return errors.New("must specify the dotfiles repository URL")
	}
	repoName := c.fs.Arg(2)
	return clone(c.rootCmd.IOStream, repoName)
}

type linkCmd struct {
	rootCmd *got

	fs *flag.FlagSet
}

func newLinkCmd(rootCmd *got) *linkCmd {
	fs := flag.NewFlagSet("got-link", flag.ContinueOnError)
	return &linkCmd{
		rootCmd: rootCmd,
		fs:      fs,
	}
}

func (c *linkCmd) parse(args []string) error {
	return c.fs.Parse(args)
}

func (c *linkCmd) run(ctx context.Context) error {
	return link()
}
