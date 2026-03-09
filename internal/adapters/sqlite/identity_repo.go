package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/StanislavYaroslavtsev/palantir/internal/core/domain"
)

type IdentityRepo struct {
	db *DB
}

func NewIdentityRepo(db *DB) *IdentityRepo {
	return &IdentityRepo{db: db}
}

func (r *IdentityRepo) Load() (*domain.Identity, error) {
	var identity domain.Identity

	err := r.db.sql.QueryRow(
		`SELECT peer_id, priv_key FROM identity WHERE id = 1`,
	).Scan(&identity.PeerID, &identity.PrivKey)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("identity load: %w", err)
	}

	return &identity, nil
}

func (r *IdentityRepo) Save(identity domain.Identity) error {
	_, err := r.db.sql.Exec(`
		INSERT INTO identity (id, peer_id, priv_key)
		VALUES (1, ?, ?)
		ON CONFLICT(id) DO UPDATE SET peer_id = excluded.peer_id
	`, identity.PeerID, identity.PrivKey)

	if err != nil {
		return fmt.Errorf("identity save: %w", err)
	}
	return nil
}
