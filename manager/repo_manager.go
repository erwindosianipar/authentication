package manager

import (
	"sync"

	"authentication/infra"
	"authentication/repo"
)

type RepoManager interface {
	AuthRepo() repo.AuthRepo
}

type repoManager struct {
	infra infra.Infra
}

func NewRepoManager(infra infra.Infra) RepoManager {
	return &repoManager{infra: infra}
}

var (
	authRepoOnce sync.Once
	authRepo     repo.AuthRepo
)

func (rm *repoManager) AuthRepo() repo.AuthRepo {
	authRepoOnce.Do(func() {
		authRepo = repo.NewAuthRepo(rm.infra.GormDB())
	})

	return authRepo
}
