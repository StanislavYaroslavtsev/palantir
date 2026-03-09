package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/StanislavYaroslavtsev/palantir/internal/adapters/sqlite"
	"github.com/StanislavYaroslavtsev/palantir/internal/config"
	"github.com/StanislavYaroslavtsev/palantir/internal/node"
)

const banner = `
██████╗  █████╗ ██╗      █████╗ ███╗   ██╗████████╗██╗██████╗ 
██╔══██╗██╔══██╗██║     ██╔══██╗████╗  ██║╚══██╔══╝██║██╔══██╗
██████╔╝███████║██║     ███████║██╔██╗ ██║   ██║   ██║██████╔╝
██╔═══╝ ██╔══██║██║     ██╔══██║██║╚██╗██║   ██║   ██║██╔══██╗
██║     ██║  ██║███████╗██║  ██║██║ ╚████║   ██║   ██║██║  ██║
╚═╝     ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝   ╚═╝╚═╝  ╚═╝`

func main() {
	fmt.Println(banner)

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config", "error", err)
		os.Exit(1)
	}

	db, err := sqlite.Open(cfg.DataDir)
	if err != nil {
		slog.Error("db", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	identityRepo := sqlite.NewIdentityRepo(db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n, err := node.New(ctx, cfg, identityRepo)
	if err != nil {
		slog.Error("node", "error", err)
		os.Exit(1)
	}
	defer n.Close()
}
