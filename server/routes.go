package server

import "github.com/julienschmidt/httprouter"

func (s *Server) bindRoutes() {
	openRoutes := map[string]httprouter.Handle{
		"GET:/login":         s.showLoginForm(),
		"POST:/handle-login": s.handleLogin(),
	}

	protectedRoutes := map[string]httprouter.Handle{
		"GET:/":                 s.handleIndex(),              // show dashboard
		"GET:/sites":            s.handleSites(),              // show index of all sites
		"GET:/create-site":      s.handleCreateForm(),         //show form to create a new site
		"POST:/create-site":     s.handleCreate(),             // deploy new site from skeleton
		"GET:/create-skeleton":  s.handleCreateSkeletonForm(), // show form to create new skeleton
		"POST:/create-skeleton": s.handleCreateSkeleton(),     // create a new skeleton from existing site
		"POST:/disable":         s.handleDisableSite(),        // disable an existing site
		"POST:/enable":          s.handleEnableSite(),         // enable a disabled site
	}

	s.Log("Available routes:")
	// bind the routes without authentication middleware
	for signature, handler := range openRoutes {
		if signature[0:3] == "GET" {
			s.router.GET(s.routePrefix+signature[4:], handler)
			s.Log("GET:  " + s.routePrefix + signature[4:])
		}
		if signature[0:4] == "POST" {
			s.router.POST(s.routePrefix+signature[5:], handler)
			s.Log("POST: " + s.routePrefix + signature[5:])
		}
	}

	// bind protected routes with the authentication middleware
	for signature, handler := range protectedRoutes {
		if signature[0:3] == "GET" {
			s.router.GET(s.routePrefix+signature[4:], s.checkSession(handler))
			s.Log("GET:  " + s.routePrefix + signature[4:])
		}
		if signature[0:4] == "POST" {
			s.router.POST(s.routePrefix+signature[5:], s.checkSession(handler))
			s.Log("POST: " + s.routePrefix + signature[5:])
		}
	}
}
