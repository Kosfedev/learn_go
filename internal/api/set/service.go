package set

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/set_v1"
)

type Implementation struct {
	desc.UnimplementedSetV1Server
	setService service.SetService
}

func NewImplementation(setService service.SetService) *Implementation {
	return &Implementation{
		setService: setService,
	}
}
