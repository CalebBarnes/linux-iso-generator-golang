package templates

import (
	"bytes"
	"text/template"
)

type UserData struct {
	Hostname string
	Username string
	Password string
	SSHKeys  []string // Assume this is a slice of SSH public keys.
}

const UserDataTemplate = `#cloud-config
autoinstall:
  version: 1
  keyboard:
    layout: us
  identity:
    hostname: {{.Hostname}}
    username: {{.Username}}
    password: "{{.Password}}"
  ssh:
    allow-pw: true
    {{- if .SSHKeys}}
    authorized-keys:
    {{- range .SSHKeys}}
      - {{.}}
    {{- end}}
    {{- end}}
    install-server: true
`

func GetUserDataConfig(userData UserData) string {
	tmpl, err := template.New("user-data").Parse(UserDataTemplate)
	if err != nil {
		panic(err)
	}

	var userDataBuffer bytes.Buffer
	if err := tmpl.Execute(&userDataBuffer, userData); err != nil {
		panic(err)
	}

	return userDataBuffer.String()
}
