package guac

import (
	"context"
	"fmt"

	"github.com/guacsec/guac/pkg/handler/processor"
	"github.com/guacsec/guac/pkg/ingestor"
)

func CollectSBOM(ctx context.Context, guacGraphqlEndpoint string, imageReference string, sbom []byte) error {
	doc := &processor.Document{
		Blob:   sbom,
		Type:   processor.DocumentSPDX,
		Format: processor.FormatJSON,
		SourceInformation: processor.SourceInformation{
			Collector: "cilantro",
			Source:    imageReference,
		},
	}

	err := ingestor.Ingest(ctx, doc, guacGraphqlEndpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to ingest SBOM into GUAC: %w", err)
	}

	return nil
}
