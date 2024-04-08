package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anchore/clio"
	"github.com/anchore/fangs"
	"github.com/ridhoq/cilantro/pkg/cilantro"
	"github.com/ridhoq/cilantro/pkg/handler/webhook"
	"github.com/spf13/cobra"
)

type CilantroConfig struct {
	GUACGraphQLEndpoint string `mapstructure:"guac-graphql-endpoint"`
	SBOMGenerator       string `mapstructure:"sbom-generator"`
}

func (c *CilantroConfig) AddFlags(flags fangs.FlagSet) {
	flags.StringVarP(
		&c.GUACGraphQLEndpoint, "guac-graphql-server", "g",
		"GUAC GraphQL server endpoint",
	)
	flags.StringVarP(
		&c.SBOMGenerator, "sbom-generator", "s",
		"SBOM generator to use",
	)
}

func RunCommand(app clio.Application) *cobra.Command {
	cfg := &CilantroConfig{
		GUACGraphQLEndpoint: "http://localhost:8080/query",
		SBOMGenerator:       "syft",
	}

	return app.SetupCommand(&cobra.Command{
		Use:   "run <image>",
		Short: "Generate and collect SBOM for an image",
		Long:  "Generate a Software Bill of Materials (SBOM) for a container image and collect it into Graph for Understanding Artifact Composition (GUAC)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cilantro.RunCilantro(cmd.Context(), args[0], cfg.GUACGraphQLEndpoint, cfg.SBOMGenerator)
		},
	}, cfg)
}

func WebhookCommand(app clio.Application) *cobra.Command {
	cfg := &CilantroConfig{
		GUACGraphQLEndpoint: "http://localhost:8080/query",
		SBOMGenerator:       "syft",
	}

	return app.SetupCommand(&cobra.Command{
		Use: "webhook",
		Run: func(cmd *cobra.Command, args []string) {
			mux := http.NewServeMux()

			mux.HandleFunc("POST /webhook", webhook.ImagePushWebhookHandler)

			fmt.Println("Server is running on port 5050")
			http.ListenAndServe(":5050", mux)
		},
	}, cfg)
}

func main() {
	cfg := clio.NewSetupConfig(clio.Identification{
		Name:    "cilantro",
		Version: "v0.1.0",
	})

	app := clio.New(*cfg)

	root := app.SetupRootCommand(&cobra.Command{})

	root.AddCommand(RunCommand(app))
	root.AddCommand(WebhookCommand(app))

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
