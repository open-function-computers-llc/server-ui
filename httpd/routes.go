package main

func (s *server) bindRoutes() {
	s.router.GET("/", s.handleIndex())
}
