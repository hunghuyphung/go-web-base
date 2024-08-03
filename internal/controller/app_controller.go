package controller

type AppController interface {
	Health() string
}

type AppControllerImpl struct {
}

func NewAppController() AppController {
	return &AppControllerImpl{}
}

func (controller AppControllerImpl) Health() string {
	return "Service up!"
}
