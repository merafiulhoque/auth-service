package requireuser

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func CreateNewHandler() *handler {
	return newHandler()
}
