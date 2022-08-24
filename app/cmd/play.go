package cmd

import (
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "testing playground",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {

}
