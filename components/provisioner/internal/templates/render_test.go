package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const renderTemplate = `Shoot: {{ .ShootName }}
Gardener Project: {{ .ProjectName }}
Gardener Secret: {{ .GardenerSecretName }}
Region: {{ .Region }}
Gardener Domain: {{ .GardenerDomain }}

OIDC Issuer URL: {{ .OIDC.IssuerURL }}
OIDC Client Id: {{ .OIDC.ClientId }}
OIDC Client Secret: {{ .OIDC.ClientSecret }}
OIDC Admin Group: {{ .OIDC.AdminGroup }}
OIDC Admin Group Namespace: {{ .OIDC.AdminGroupNamespace }}
OIDC Developer Group: {{ .OIDC.DeveloperGroup }}`

func TestRenderTemplates(t *testing.T) {

	expectedRender := `Shoot: my-shoot
Gardener Project: my-project
Gardener Secret: my-secret
Region: eu-west
Gardener Domain: live.ondemand.com

OIDC Issuer URL: https://issuer.com
OIDC Client Id: abcd
OIDC Client Secret: efgh
OIDC Admin Group: admin
OIDC Admin Group Namespace: admins
OIDC Developer Group: devs`

	values := Values{
		ShootName:          "my-shoot",
		ProjectName:        "my-project",
		GardenerSecretName: "my-secret",
		Region:             "eu-west",
		GardenerDomain:     "live.ondemand.com",
		OIDC: OIDCConfig{
			IssuerURL:           "https://issuer.com",
			ClientId:            "abcd",
			ClientSecret:        "efgh",
			AdminGroup:          "admin",
			AdminGroupNamespace: "admins",
			DeveloperGroup:      "devs",
		},
	}

	result, err := RenderTemplate(renderTemplate, values)
	require.NoError(t, err)

	assert.Equal(t, expectedRender, string(result))
}
