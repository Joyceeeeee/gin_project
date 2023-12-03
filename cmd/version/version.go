package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"youdangzhe/internal/global"
)

var (
	Cmd = &cobra.Command{
		Use:     "version",
		Short:   "GetUserInfoByEmail version info",
		Example: "go-layout version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(global.Version)
		},
	}
)
