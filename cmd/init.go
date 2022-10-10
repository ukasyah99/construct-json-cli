package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukasyah99/construct-json-cli/lib"
	"github.com/ukasyah99/construct-json-cli/schema"
)

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		sample := schema.Sample{
			Schema: "https://github.com/hyperledger/releases/download/1.1.0/schema.json",
			Global: schema.Global{
				TLS: true,
			},
		}

		err := lib.SelectItem(&sample.Global.FabricVersion, "Fabric version", []string{"1.4.6", "2.2.4"})
		if err != nil {
			return
		}

		err = lib.SelectItem(&sample.Global.Monitoring.Loglevel, "Monitoring log level", []string{"off", "debug"})
		if err != nil {
			return
		}

		s, _ := json.MarshalIndent(sample, "", "\t")
		fmt.Println(string(s))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
