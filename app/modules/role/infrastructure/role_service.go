package infrastructure

type RoleService struct{}

func (us *RoleService) find() []Role {
	return rolesDummy
}
