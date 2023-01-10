package cmd

import (
	"fmt"
	"os"

	"github.com/ForeverSRC/kaeya-ctl/pkg/service"
	"github.com/spf13/cobra"
)

const (
	ctlVersion = "1.0.0"
)

const (
	flagAddress     = "address"
	flagInteractive = "interactive"
)

var (
	address         string
	interactiveMode bool
)

var rootCmd = &cobra.Command{
	Use:   "kaeyactl",
	Short: "kaeyactl is a command-line tool for Kaeya",
	Long: `kaeyactl is a command-line tool for Kaeya, a database.
See https://github.com/ForeverSRC/kaeya for more information.`,
	Version: ctlVersion,
	Run:     RootCmdRun,
}

func init() {
	rootCmd.Flags().BoolVarP(&interactiveMode, flagInteractive, "i", false, "interactive mode or not")
	rootCmd.Flags().StringVarP(&address, flagAddress, "a", "", "address of kaeya-server")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func RootCmdRun(cmd *cobra.Command, args []string) {
	if interactiveMode {
		if address == "" {
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "server address is needed")
			return
		}
		var interaction service.DBInteraction
		interaction = service.NewDefaultDBInteraction(cmd.InOrStdin(), cmd.OutOrStdout())

		_ = interaction.Interactive(cmd.Context(), address)
	}

}
