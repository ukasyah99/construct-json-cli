package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukasyah99/construct-json-cli/lib"
	"github.com/ukasyah99/construct-json-cli/schema"
)

type OrgInput struct {
	Name    string
	Domain  string
	MspName string
}

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		sample := schema.Sample{
			Schema: "https://github.com/hyperledger/releases/download/1.1.0/schema.json",
			Global: schema.Global{
				TLS: true,
			},
			Orgs: []schema.Org{},
		}

		err := lib.SelectItem(&sample.Global.FabricVersion, "Fabric version", []string{"1.4.6", "2.2.4"})
		if err != nil {
			return
		}

		err = lib.SelectItem(&sample.Global.Monitoring.Loglevel, "Monitoring log level", []string{"off", "debug"})
		if err != nil {
			return
		}

		var doneInputOrgs bool
		for !doneInputOrgs {
			fmt.Println("\nCreate new organization:")

			var org schema.Org

			if err := lib.Input(&org.Organization.Name, "Name"); err != nil {
				break
			}

			if err := lib.Input(&org.Organization.Domain, "Domain"); err != nil {
				break
			}

			if err := lib.Input(&org.Organization.MspName, "MspName"); err != nil {
				break
			}

			sample.Orgs = append(sample.Orgs, org)

			var createAnotherOrg string
			err = lib.SelectItem(&createAnotherOrg, "Want to create another one", []string{"yes", "no"})
			if err != nil {
				break
			}

			if createAnotherOrg == "no" {
				break
			}
		}

		s, _ := json.MarshalIndent(sample, "", "\t")
		fmt.Println(string(s))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
