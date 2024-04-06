package event

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseImagePushEventData(t *testing.T) {
	// Create a mock io.ReadCloser with the desired input data
	inputData := `{
		"id": "831e1650-001e-001b-66ab-eeb76e069631",
		"source": "/subscriptions/<subscription-id>/resourceGroups/<resource-group-name>/providers/Microsoft.ContainerRegistry/registries/<name>",
		"subject": "aci-helloworld:v1",
		"type": "Microsoft.ContainerRegistry.ImagePushed",
		"time": "2018-04-25T21:39:47.6549614Z",
		"data": {
			"id": "31c51664-e5bd-416a-a5df-e5206bc47ed0",
			"timestamp": "2018-04-25T21:39:47.276585742Z",
			"action": "push",
			"target": {
				"mediaType": "application/vnd.docker.distribution.manifest.v2+json",
				"size": 3023,
				"digest": "sha256:213bbc182920ab41e18edc2001e06abcca6735d87782d9cef68abd83941cf0e5",
				"length": 3023,
				"repository": "aci-helloworld",
				"tag": "v1"
			},
			"request": {
				"id": "7c66f28b-de19-40a4-821c-6f5f6c0003a4",
				"host": "demo.azurecr.io",
				"method": "PUT",
				"useragent": "docker/18.03.0-ce go/go1.9.4 git-commit/0520e24 os/windows arch/amd64 UpstreamClient(Docker-Client/18.03.0-ce \\\\(windows\\\\))"
			}
		},
		"specversion": "1.0"
	}`

	mockData := io.NopCloser(strings.NewReader(inputData))

	// Call the function under test
	eventData, err := ParseImagePushEventData(mockData)

	// Assert that no error occurred
	assert.NoError(t, err)

	assert.Equal(t, "sha256:213bbc182920ab41e18edc2001e06abcca6735d87782d9cef68abd83941cf0e5", *eventData.Target.Digest)
}
