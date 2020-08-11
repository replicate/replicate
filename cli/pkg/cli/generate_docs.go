package cli

import (
	"os"
	"path"

	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

func newGenerateDocsCommand(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "generate-docs",
		Short:  "",
		RunE:   generateDocs(rootCmd),
		Args:   cobra.ExactArgs(0),
		Hidden: true,
	}

	cmd.Flags().StringP("docs-folder", "f", "", "Path to replicate-docs")
	if err := cmd.MarkFlagRequired("docs-folder"); err != nil {
		panic(err)
	}

	return cmd
}

func generateDocs(rootCmd *cobra.Command) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		docsFolder, err := cmd.Flags().GetString("docs-folder")
		if err != nil {
			return err
		}

		docsPath := path.Join(docsFolder, "cli.md")
		err = genMarkdownSingleFile(rootCmd, docsPath)
		if err != nil {
			return err
		}
		return nil
	}
}

func genMarkdownSingleFile(cmd *cobra.Command, path string) error {
	// TODO: support more than two levels of commands?

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, `---
id: cli
title: CLI reference
---

This is the reference for the Replicate CLI commands. You can also see this in the terminal by running `+"`replicate --help` or `replicate command --help`"+`.

`)

	cmd.DisableAutoGenTag = true

	fmt.Fprintln(f, "## Commands")
	fmt.Fprintln(f)

	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}

		fmt.Fprintf(f, "* [`%s %s`](#replicate-%s) – %s\n", cmd.Name(), c.Name(), c.Name(), c.Short)
	}

	fmt.Fprintln(f, "")

	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		c.DisableAutoGenTag = true
		if err := genMarkdown(c, f); err != nil {
			return err
		}
	}

	return nil
}

func genMarkdown(cmd *cobra.Command, f *os.File) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	buf := new(bytes.Buffer)
	name := cmd.CommandPath()

	short := cmd.Short
	long := cmd.Long

	buf.WriteString(fmt.Sprintf("## `%s`\n\n", name))
	if long != "" {
		buf.WriteString(long + "\n\n")
	} else {
		buf.WriteString(short + "\n\n")
	}

	if cmd.Runnable() {
		buf.WriteString("### Usage\n\n")
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.UseLine()))

		if len(cmd.Example) > 0 {
			buf.WriteString("### Examples\n\n")
			buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.Example))
		}

		if err := printOptions(buf, cmd, name); err != nil {
			return err
		}
	}
	_, err := buf.WriteTo(f)
	return err
}

func printOptions(buf *bytes.Buffer, cmd *cobra.Command, name string) error {
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(buf)
	parentFlags := cmd.InheritedFlags()
	hasFlags := flags.HasAvailableFlags() || parentFlags.HasAvailableFlags()

	if !hasFlags {
		return nil
	}

	buf.WriteString("### Flags\n\n```\n")

	if flags.HasAvailableFlags() {
		flags.PrintDefaults()
		buf.WriteString("\n")
	}

	parentFlags.SetOutput(buf)
	if parentFlags.HasAvailableFlags() {
		parentFlags.PrintDefaults()
	}

	buf.WriteString("```\n")

	return nil
}
