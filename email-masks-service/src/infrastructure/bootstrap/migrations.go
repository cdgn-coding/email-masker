package bootstrap

type Migrations interface {
	Apply()
}
