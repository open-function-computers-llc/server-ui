package server

import "os"

func (s *Server) bindRoutes() {
	prefix := "/" + os.Getenv("ROUTE_PREFIX")

	// admin homepage
	s.router.GET(prefix+"/", s.handleIndex())

	// deploy new site from skeleton
	s.router.GET(prefix+"/create-site", s.handleCreateForm())
	s.router.POST(prefix+"/create-site", s.handleCreate())

	// create a new skeleton from existing site
	s.router.GET(prefix+"/create-skeleton", s.handleCreateSkeletonForm())
	s.router.POST(prefix+"/create-skeleton", s.handleCreateSkeleton())

	// disable an existing site
	s.router.POST(prefix+"/disable", s.handleDisableSite())

	// enable a disabled site
	s.router.POST(prefix+"/enable", s.handleEnableSite())
}
