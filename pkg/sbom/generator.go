package sbom

import "context"

type Generator interface {
	// GenerateSBOM generates a Software Bill of Materials (SBOM) for the given image.
	GenerateSBOM(ctx context.Context, image string) ([]byte, error)
	// Type() returns the type of the SBOM generator.
	Type() string
}
