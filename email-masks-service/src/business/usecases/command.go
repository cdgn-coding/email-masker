package usecases

type CommandUseCase[T interface{}] interface {
	Execute(input T) error
}
