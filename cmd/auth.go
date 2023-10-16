package cmd

import (
	"github.com/sottey/ultodo/ultodo"
	"github.com/spf13/cobra"
)

func init() {
	var (
		webAuthCmdDesc     = "Authenticates you against ultodo.io"
		webAuthCmdLongDesc = `
This will authenticate your local ultodo with ultodo.io.

  Syncing with ultodo.io conveys many benefits:
    - Real-time sync with other ultodo binaries on other computers
    - Manage your lists via the web at app.ultodo.io
    - Use ultodo on your mobile phone
    - Any many others.

  ultodo.io is a paid service.  For more information, See https://ultodo.io/docs/cli/pro_integration`
	)

	var webAuthCmd = &cobra.Command{
		Use:   "auth",
		Long:  webAuthCmdLongDesc,
		Short: webAuthCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().AuthWorkflow()
		},
	}

	var (
		webAuthCheckCmdDesc     = "Checks your authentication status against ultodo.io"
		webAuthCheckCmdLongDesc = "\nCheck your login status to ultodo.io using \"ultodo auth check\"."
	)

	var webAuthCheckCmd = &cobra.Command{
		Use:   "check",
		Long:  webAuthCheckCmdLongDesc,
		Short: webAuthCheckCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			ultodo.NewApp().CheckAuth()
		},
	}

	rootCmd.AddCommand(webAuthCmd)
	webAuthCmd.AddCommand(webAuthCheckCmd)
}
