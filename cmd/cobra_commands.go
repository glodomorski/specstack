package cmd

import "github.com/spf13/cobra"

func noop(*cobra.Command, []string) {}

func WireUpCobraHarness(harness *CobraHarness) *cobra.Command {
	root := &cobra.Command{
		Use: "spec",
	}

	root.SetOutput(harness.stdout)

	root.AddCommand(
		commandConfig(harness),
		commandGitHooks(harness),
		commandMetadata(harness),
		commandPull(harness),
		commandPush(harness),
	)

	root.PersistentPreRunE = harness.PersistentPreRunE

	return root
}

func commandConfig(harness *CobraHarness) *cobra.Command {
	root := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
	}
	list := &cobra.Command{
		Use:  "list",
		Args: cobra.NoArgs,
	}
	get := &cobra.Command{
		Use:     "get <key>",
		Args:    cobra.ExactArgs(1),
		Example: "$ spec config get project.name",
	}
	set := &cobra.Command{
		Use:     "set <key>=<value>",
		Example: "$ spec config set project.name=myProject",
	}

	root.AddCommand(
		list,
		get,
		set,
	)

	list.RunE = harness.ConfigList
	get.RunE = harness.ConfigGet
	set.Args = harness.SetKeyValueArgs
	set.RunE = harness.ConfigSet

	return root
}

func commandGitHooks(harness *CobraHarness) *cobra.Command {
	root := &cobra.Command{
		Use:     "git-hook",
		Aliases: []string{"gh"},
		Short:   "Low-level git hook interactions",
	}
	exec := &cobra.Command{
		Use:     "exec <pre-push|post-merge|post-commit>",
		Args:    cobra.ExactArgs(1),
		Example: "$ spec git-hook exec pre-push",
	}

	root.AddCommand(
		exec,
	)

	exec.RunE = harness.GitHookExec

	return root
}

func commandMetadata(harness *CobraHarness) *cobra.Command {
	root := &cobra.Command{
		Use:     "metadata",
		Aliases: []string{"md"},
		Short:   "Manage low level metadata",
	}
	add := &cobra.Command{
		Use:     "add",
		Short:   "Add metadata to a story or scenario",
		Example: "$ spec metadata add --story my_story key=value",
		PreRunE: harness.SnapshotScenarioMetadata,
	}
	list := &cobra.Command{
		Use:     "list",
		Args:    cobra.ExactArgs(0),
		Aliases: []string{"ls"},
		Short:   "Show metadata for a story or scenario",
		Example: "$ spec metadata list --story my_story",
		PreRunE: harness.SnapshotScenarioMetadata,
	}
	commit := &cobra.Command{
		Use:     "commit",
		Example: "$ spec metadata commit",
		PreRunE: harness.SnapshotScenarioMetadata,
		Run:     noop,
	}
	root.AddCommand(
		add,
		list,
		commit,
	)

	root.PersistentFlags().String("story", "", "")
	root.PersistentFlags().String("scenario", "", "")
	add.RunE = harness.MetadataAdd
	add.Args = harness.SetKeyValueArgs
	list.RunE = harness.MetadataList

	return root
}

func commandPull(harness *CobraHarness) *cobra.Command {
	root := &cobra.Command{
		Use:   "pull",
		Short: "Update local repository's metadata",
	}

	root.RunE = harness.Pull

	return root
}

func commandPush(harness *CobraHarness) *cobra.Command {
	root := &cobra.Command{
		Use:   "push",
		Short: "Update remote repository's metadata",
	}

	root.RunE = harness.Push

	return root
}
