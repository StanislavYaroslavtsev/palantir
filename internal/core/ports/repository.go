package ports

import "github.com/StanislavYaroslavtsev/palantir/internal/core/domain"

type IdentityRepository interface {
	Load() (*domain.Identity, error)
	Save(identity domain.Identity) error
}
