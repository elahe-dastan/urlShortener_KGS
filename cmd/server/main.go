package server

import (
	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/elahe-dastan/urlShortener_KGS/db"
	"github.com/elahe-dastan/urlShortener_KGS/service"
	"github.com/elahe-dastan/urlShortener_KGS/store"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				d := db.New(cfg.Database)
				api := service.API{
					Map:      *store.NewMap(d),
					ShortURL: *store.NewShortURL(d)}
				api.Run(cfg.Log)
			},
		},
	)
}
