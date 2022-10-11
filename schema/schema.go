package schema

type Sample struct {
	Schema   string    `json:"$schema"`
	Global   Global    `json:"global"`
	Orgs     []Org     `json:"orgs"`
	Channels []Channel `json:"channels"`
}

type Global struct {
	FabricVersion string     `json:"fabricVersion"`
	TLS           bool       `json:"tls"`
	Monitoring    Monitoring `json:"monitoring"`
}

type Monitoring struct {
	Loglevel string `json:"loglevel"`
}

type Org struct {
	Organization OrgOrganization `json:"organization"`
	Orderers     []OrgOrderer    `json:"orderers,omitempty"`
	CA           *OrgCA          `json:"ca,omitempty"`
	Peer         *OrgPeer        `json:"peer,omitempty"`
}

type OrgOrganization struct {
	Name    string `json:"name"`
	Domain  string `json:"domain"`
	MspName string `json:"mspName,omitempty"`
}

type OrgOrderer struct {
	GroupName string `json:"groupName"`
	Prefix    string `json:"prefix"`
	Type      string `json:"type"`
	Instances int    `json:"instances"`
}

type OrgCA struct {
	Prefix string `json:"prefix"`
}

type OrgPeer struct {
	Prefix    string `json:"prefix"`
	Instances int    `json:"instances"`
	DB        string `json:"db"`
}

type Channel struct {
	Name string       `json:"name"`
	Orgs []ChannelOrg `json:"orgs"`
}

type ChannelOrg struct {
	Name  string   `json:"name"`
	Peers []string `json:"peers"`
}
