package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"replicate.ai/cli/pkg/config"
	"replicate.ai/cli/pkg/global"
	"replicate.ai/cli/pkg/list"
	"replicate.ai/cli/pkg/slices"
	"replicate.ai/cli/pkg/storage"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls [project-storage-url]",
		Short:   "List experiments in this project",
		Aliases: []string{"list"},
		RunE:    listExperiments,
		Args:    cobra.MaximumNArgs(1),
	}

	cmd.Flags().StringP("format", "f", "table", "Output format (table/json)")

	return cmd
}

func listExperiments(cmd *cobra.Command, args []string) error {
	var storageURL string
	// FIXME (bfirsh): perhaps better as a flag? (just putting this here to remind ourselves to have this discussion)
	if len(args) == 1 {
		storageURL = args[0]
	} else {
		conf, _, err := config.FindConfigInWorkingDir(global.SourceDirectory)
		if err != nil {
			return err
		}
		storageURL = conf.Storage
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}
	validFormats := []string{list.FormatJSON, list.FormatTable}
	if !slices.ContainsString(validFormats, format) {
		return fmt.Errorf("%s is not a valid format. Valid formats are: %s", format, strings.Join(validFormats, ", "))
	}

	store, err := storage.ForURL(storageURL)
	if err != nil {
		return err
	}

	return list.Experiments(store, format)
}
