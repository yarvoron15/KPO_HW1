package di

import (
	"go.uber.org/dig"
	"vyshkat/internal/domain"
	"vyshkat/internal/service"
	"vyshkat/internal/ui"
)

func BuildContainer() (*dig.Container, error) {
	c := dig.New()
	if err := c.Provide(service.NewStandardServiceCenter, dig.As(new(domain.ServiceCenter))); err != nil {
		return nil, err
	}
	if err := c.Provide(service.NewFleetService); err != nil {
		return nil, err
	}
	if err := c.Provide(ui.NewCLI); err != nil {
		return nil, err
	}
	return c, nil
}
