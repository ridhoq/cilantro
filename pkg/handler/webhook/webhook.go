package webhook

import (
	"fmt"
	"net/http"

	"github.com/ridhoq/cilantro/pkg/event"
)

// ImagePushWebhookHandler is a handler for pushing images to the registry
func ImagePushWebhookHandler(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request body
	_, err := event.ParseImagePushEventData(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}

}
