package middlewaresHandlers

import _pkgMiddlewaresMiddlewaresUsecases "github.com/MarkTBSS/056_Middlewares_Module/modules/middlewares/middlewaresUsecases"

type IMiddlewaresHandler interface {
}

type middlewaresHandler struct {
	middlewaresUsecase _pkgMiddlewaresMiddlewaresUsecases.IMiddlewaresUsecase
}

func MiddlewaresHandler(middlewaresUsecase _pkgMiddlewaresMiddlewaresUsecases.IMiddlewaresUsecase) IMiddlewaresHandler {
	return &middlewaresHandler{
		middlewaresUsecase: middlewaresUsecase,
	}
}
