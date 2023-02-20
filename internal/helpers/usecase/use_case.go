package usecase

import "context"

type UseCase[In, Out any] interface {
	Execute(ctx context.Context, input *In) (output *Out, err error)
}
