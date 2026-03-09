package node

import (
	"crypto/sha256"
	"fmt"

	"github.com/StanislavYaroslavtsev/palantir/internal/core/domain"
	"github.com/StanislavYaroslavtsev/palantir/internal/core/ports"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

func LoadOrCreateIdentity(repo ports.IdentityRepository) (crypto.PrivKey, error) {
	stored, err := repo.Load()
	if err != nil {
		return nil, err
	}

	if stored != nil {
		privKey, err := crypto.UnmarshalPrivateKey(stored.PrivKey)
		if err != nil {
			return nil, fmt.Errorf("unmarshal privkey: %w", err)
		}
		return privKey, nil
	}

	privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
	if err != nil {
		return nil, fmt.Errorf("generate keypair: %w", err)
	}

	privBytes, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return nil, fmt.Errorf("marshal privkey: %w", err)
	}

	peerID, err := peer.IDFromPrivateKey(privKey)
	if err != nil {
		return nil, fmt.Errorf("peer id: %w", err)
	}

	if err := repo.Save(domain.Identity{
		PeerID:  peerID.String(),
		PrivKey: privBytes,
	}); err != nil {
		return nil, fmt.Errorf("save identity: %w", err)
	}

	fmt.Println("  ✨ new identity generated and saved")
	return privKey, nil
}

func BuildFingerprint(pub crypto.PubKey) string {
	raw, err := pub.Raw()
	if err != nil {
		return "unknown"
	}
	sum := sha256.Sum256(raw)
	return fmt.Sprintf(
		"%02x%02x:%02x%02x:%02x%02x:%02x%02x",
		sum[0], sum[1], sum[2], sum[3],
		sum[4], sum[5], sum[6], sum[7],
	)
}
