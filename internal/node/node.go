package node

import (
	"context"
	"fmt"

	"github.com/StanislavYaroslavtsev/palantir/internal/config"
	"github.com/StanislavYaroslavtsev/palantir/internal/core/ports"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

type Node struct {
	host        host.Host
	cfg         *config.Config
	fingerprint string
}

func New(ctx context.Context, cfg *config.Config, identityRepo ports.IdentityRepository) (*Node, error) {
	privKey, err := LoadOrCreateIdentity(identityRepo)
	if err != nil {
		return nil, fmt.Errorf("identity: %w", err)
	}

	listenAddr, err := multiaddr.NewMultiaddr(
		fmt.Sprintf("/ip4/0.0.0.0/udp/%d/quic-v1", cfg.ListenPort),
	)
	if err != nil {
		return nil, fmt.Errorf("multiaddr: %w", err)
	}

	h, err := libp2p.New(
		libp2p.Identity(privKey),
		libp2p.ListenAddrs(listenAddr),
	)
	if err != nil {
		return nil, fmt.Errorf("libp2p.New: %w", err)
	}

	n := &Node{
		host:        h,
		cfg:         cfg,
		fingerprint: BuildFingerprint(privKey.GetPublic()),
	}
	n.printInfo()

	return n, nil
}

func (n *Node) PeerID() peer.ID {
	return n.host.ID()
}

func (n *Node) Fingerprint() string {
	return n.fingerprint
}

func (n *Node) Host() host.Host {
	return n.host
}

func (n *Node) Close() error {
	return n.host.Close()
}

func (n *Node) printInfo() {
	fmt.Printf("PeerID : %s\n", n.host.ID().String())
	fmt.Printf("Fingerprint : %s\n", n.fingerprint)
	fmt.Printf("Addrs  : ")
	for _, a := range n.host.Addrs() {
		fmt.Printf("%s ", a.String())
	}
	fmt.Println()
}
