package apiserver

//APISERVER ...
type APIServer struct{}

//New ...
func New() *APIServer {
	return &APIServer{}
}

//start
func (s *APIServer) Start() error {
	return nil
}
