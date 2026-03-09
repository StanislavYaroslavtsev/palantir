package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	// DataDir is where Palantir stores its database, keys, etc.
	DataDir string

	// DeviceName is a human-readable name for this node.
	DeviceName string

	// ListenPort is the port libp2p will listen on (0 = random).
	ListenPort int
}

func Load() (*Config, error) {
	dataDir, err := defaultDataDir()
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(dataDir, 0o700); err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Config{
		DataDir:    dataDir,
		DeviceName: hostname,
		ListenPort: 0, // random port, libp2p will assign
	}, nil
}

func defaultDataDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".palantir"), nil
}
