package event

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/messaging"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/eventgrid/azsystemevents"
)

func ParseImagePushEventData(data io.ReadCloser) (*azsystemevents.ContainerRegistryImagePushedEventData, error) {
	var cloudEvent messaging.CloudEvent
	if err := json.NewDecoder(data).Decode(&cloudEvent); err != nil {
		return nil, errors.New("not a CloudEvent, only CloudEvent is supported")
	}
	if cloudEvent.Type != string(azsystemevents.TypeContainerRegistryImagePushed) {
		return nil, errors.New("not a ContainerRegistryImagePushed event")
	}
	var imagePushedEventData azsystemevents.ContainerRegistryImagePushedEventData
	if err := json.Unmarshal(cloudEvent.Data.([]byte), &imagePushedEventData); err != nil {
		return nil, errors.New("failed to unmarshal ContainerRegistryImagePushedEventData")
	}

	return &imagePushedEventData, nil
}
