package manager

import (
	"sync"

	"authentication/infra"
	"authentication/service"
)

type ServiceManager interface {
	AuthService() service.AuthService
}

type serviceManager struct {
	infra infra.Infra
	repo  RepoManager
}

func NewServiceManager(infra infra.Infra) ServiceManager {
	return &serviceManager{
		infra: infra,
		repo:  NewRepoManager(infra),
	}
}

var (
	authServiceOnce sync.Once
	authService     service.AuthService
)

func (sm *serviceManager) AuthService() service.AuthService {
	authServiceOnce.Do(func() {
		authService = sm.repo.AuthRepo()
	})

	return authService
}
