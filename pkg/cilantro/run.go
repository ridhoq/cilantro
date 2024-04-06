package cilantro

import (
	"context"
	"fmt"

	"github.com/ridhoq/cilantro/pkg/guac"
	"github.com/ridhoq/cilantro/pkg/sbom"
	"github.com/ridhoq/cilantro/pkg/sbom/syft"
)

func RunCilantro(ctx context.Context, image string, guacGraphqlEndpoint string, sbomGenerator string) error {
	generator, err := getSBOMGenerator(sbomGenerator)
	if err != nil {
		return fmt.Errorf("failed to get SBOM generator: %w", err)
	}

	sbom, err := generator.GenerateSBOM(ctx, image)
	if err != nil {
		return fmt.Errorf("failed to generate SBOM: %w", err)
	}

	err = guac.CollectSBOM(ctx, guacGraphqlEndpoint, image, sbom)
	if err != nil {
		return fmt.Errorf("failed to collect SBOM: %w", err)
	}

	return nil
}

func getSBOMGenerator(sbomGenerator string) (sbom.Generator, error) {
	switch sbomGenerator {
	case "syft":
		return syft.NewSyftGenerator(), nil
	default:
		return nil, fmt.Errorf("unsupported SBOM generator: %s", sbomGenerator)
	}
}
