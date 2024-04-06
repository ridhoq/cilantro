package syft

import (
	"context"
	"fmt"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/format"
	"github.com/anchore/syft/syft/format/spdxjson"
	"github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
)

type syftGenerator struct {
}

func NewSyftGenerator() *syftGenerator {
	return &syftGenerator{}
}

func (s *syftGenerator) Type() string {
	return "syft"
}

func (s *syftGenerator) GenerateSBOM(ctx context.Context, image string) ([]byte, error) {
	src, err := getSource(ctx, image)
	if err != nil {
		return nil, fmt.Errorf("failed to get source: %w", err)
	}

	sbom, err := getSBOM(ctx, src)
	if err != nil {
		return nil, fmt.Errorf("failed to get SBOM: %w", err)
	}

	return formatSBOM(sbom)
}

func getSource(ctx context.Context, input string) (source.Source, error) {
	src, err := syft.GetSource(ctx, input, nil)
	if err != nil {
		return nil, err
	}

	return src, nil
}

func getSBOM(ctx context.Context, src source.Source) (*sbom.SBOM, error) {
	sbom, err := syft.CreateSBOM(ctx, src, nil)
	if err != nil {
		return nil, err
	}

	return sbom, nil
}

func formatSBOM(sbom *sbom.SBOM) ([]byte, error) {
	spdxEncoder, err := spdxjson.NewFormatEncoderWithConfig(spdxjson.EncoderConfig{Version: "2.3", Pretty: true})
	if err != nil {
		return nil, fmt.Errorf("failed to create SPDX JSON encoder: %w", err)
	}
	bytes, err := format.Encode(*sbom, spdxEncoder)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
