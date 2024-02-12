package templates

type MetaData struct {
	Hostname string
	UUID     string
}

const MetaDataTemplate = `instance_id: {{.Hostname}}-{{.UUID}}
local-hostname: {{.Hostname}}
`
