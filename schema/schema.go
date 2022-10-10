package schema

type Sample struct {
	Schema string `json:"$schema"`
	Global Global `json:"global"`
}

type Global struct {
	FabricVersion string     `json:"fabricVersion"`
	TLS           bool       `json:"tls"`
	Monitoring    Monitoring `json:"monitoring"`
}

type Monitoring struct {
	Loglevel string `json:"loglevel"`
}
