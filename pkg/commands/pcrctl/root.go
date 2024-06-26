package pcrctl

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yhlooo/podmig/pkg/commands/pcrctl/options"
	"github.com/yhlooo/podmig/pkg/utils/cmdutil"
)

// NewRootCommand 创建一个 pcrctl 命令
func NewRootCommand() *cobra.Command {
	return NewRootCommandWithOptions(options.NewDefaultOptions())
}

// NewRootCommandWithOptions  使用指定选项创建一个 pcrctl 命令
func NewRootCommandWithOptions(opts options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "pcrctl",
		Short:        "Checkpoint/Restore kubernetes pod on node",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// 校验全局选项
			if err := opts.Global.Validate(); err != nil {
				return err
			}
			// 设置日志
			logger := cmdutil.SetLogger(cmd, opts.Global.Verbosity)

			logger.V(1).Info(fmt.Sprintf("command: %q, args: %#v, options: %#v", cmd.Name(), args, opts))
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// 绑定选项到命令行参数
	opts.Global.AddPFlags(cmd.PersistentFlags())

	// 添加子命令
	cmd.AddCommand(
		NewCheckpointCommandWithOptions(&opts.Checkpoint),
		NewRestoreCommandWithOptions(&opts.Restore),
	)

	return cmd
}
