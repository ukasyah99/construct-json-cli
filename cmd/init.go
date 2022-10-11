package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/elliotchance/pie/v2"
	"github.com/spf13/cobra"
	"github.com/ukasyah99/construct-json-cli/lib"
	"github.com/ukasyah99/construct-json-cli/schema"
	"golang.org/x/exp/slices"
)

type OrgInput struct {
	Name    string
	Domain  string
	MspName string
}

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		// Default sample configuration
		sample := schema.Sample{
			Schema: "https://github.com/hyperledger/releases/download/1.1.0/schema.json",
			Global: schema.Global{
				FabricVersion: "1.4.6",
				TLS:           true,
				Monitoring:    schema.Monitoring{Loglevel: "debug"},
			},
			Orgs:     []schema.Org{},
			Channels: []schema.Channel{},
		}

		err := lib.SelectItem(&sample.Global.FabricVersion, "Fabric version", []string{"1.4.6", "2.2.4"})
		if err != nil {
			return
		}

		err = lib.SelectItem(&sample.Global.Monitoring.Loglevel, "Monitoring log level", []string{"off", "debug"})
		if err != nil {
			return
		}

		// Start creating orgs
		for true {
			fmt.Println("\nCreate new org:")

			var org schema.Org

			if err := lib.Input(&org.Organization.Name, "Org Name"); err != nil {
				break
			}

			if err := lib.Input(&org.Organization.Domain, "Domain"); err != nil {
				break
			}

			if err := lib.Input(&org.Organization.MspName, "MspName"); err != nil {
				break
			}

			// Start creating orderers
			var hasOrderers string
			if err := lib.SelectItem(&hasOrderers, "Has Orderers", []string{"yes", "no"}); err != nil {
				break
			}
			if hasOrderers == "yes" {
				for true {
					fmt.Println("\nCreate new org orderer:")

					var orderer schema.OrgOrderer

					if err := lib.Input(&orderer.GroupName, "Group name"); err != nil {
						break
					}

					if err := lib.Input(&orderer.Prefix, "Prefix"); err != nil {
						break
					}

					if err := lib.Input(&orderer.Type, "Type"); err != nil {
						break
					}

					if err := lib.InputNumber(&orderer.Instances, "Instances"); err != nil {
						break
					}

					org.Orderers = append(org.Orderers, orderer)

					var createAnotherOrderer string
					err = lib.SelectItem(&createAnotherOrderer, "Want to create another orderer", []string{"yes", "no"})
					if err != nil {
						break
					}

					if createAnotherOrderer == "no" {
						break
					}
				}
			}
			// Done creating orderers

			var hasCA string
			if err := lib.SelectItem(&hasCA, "Has CA", []string{"yes", "no"}); err != nil {
				break
			}
			if hasCA == "yes" {
				// CA values are defined automatically for simplicity
				ca := schema.OrgCA{Prefix: "ca"}
				org.CA = &ca
			}

			var hasPeer string
			if err := lib.SelectItem(&hasPeer, "Has Peer", []string{"yes", "no"}); err != nil {
				break
			}
			if hasPeer == "yes" {
				// Peer values are defined automatically for simplicity
				peer := schema.OrgPeer{
					Prefix:    "peer",
					Instances: 2,
					DB:        "LevelDb",
				}
				org.Peer = &peer
			}

			sample.Orgs = append(sample.Orgs, org)

			var createAnotherOrg string
			err = lib.SelectItem(&createAnotherOrg, "Want to create another org", []string{"yes", "no"})
			if err != nil {
				break
			}

			if createAnotherOrg == "no" {
				break
			}
		}
		// Done creating orgs

		// Start creating channels
		for true {
			fmt.Println("\nCreate new channel:")

			var channel schema.Channel

			if err := lib.Input(&channel.Name, "Channel Name"); err != nil {
				break
			}

			// Filter orgs that have peers
			orgs := pie.Of(sample.Orgs).
				Filter(func(o schema.Org) bool {
					return o.Peer != nil
				}).Result

			// Get org names and peers
			var orgNames []string
			var orgPeers [][]string
			for _, org := range orgs {
				orgNames = append(orgNames, org.Organization.Name)

				// Automatically generate peer names based on number of instances
				var peers []string
				for i := 0; i < org.Peer.Instances; i++ {
					peers = append(peers, fmt.Sprintf("peer%d", i))
				}
				orgPeers = append(orgPeers, peers)
			}

			// Start adding orgs into channel
			for true {
				fmt.Println("\nAdd org into channel:")

				var channelOrg schema.ChannelOrg

				if err := lib.SelectItem(&channelOrg.Name, "Org name", orgNames); err != nil {
					break
				}

				orgNameIndex := slices.Index(orgNames, channelOrg.Name)
				channelOrg.Peers = orgPeers[orgNameIndex]

				channel.Orgs = append(channel.Orgs, channelOrg)

				var addAnotherOrg string
				err = lib.SelectItem(&addAnotherOrg, "Want to add another org", []string{"yes", "no"})
				if err != nil {
					break
				}

				if addAnotherOrg == "no" {
					break
				}
			}
			// Done adding orgs into channel

			sample.Channels = append(sample.Channels, channel)
			var createAnotherChannel string
			err = lib.SelectItem(&createAnotherChannel, "Want to create another channel", []string{"yes", "no"})
			if err != nil {
				break
			}

			if createAnotherChannel == "no" {
				break
			}
		}
		// Done creating channels

		s, _ := json.MarshalIndent(sample, "", "\t")
		// fmt.Println(string(s))

		_ = ioutil.WriteFile("sample.json", s, 0644)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
