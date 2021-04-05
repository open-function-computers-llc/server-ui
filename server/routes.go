package server

import "github.com/julienschmidt/httprouter"

func (s *Server) bindRoutes() {
	openRoutes := map[string]httprouter.Handle{
		"GET:/login":         s.showLoginForm(),
		"POST:/handle-login": s.handleLogin(),
	}

	protectedRoutes := map[string]httprouter.Handle{
		"GET:/":                     s.handleIndex(),              // show dashboard
		"GET:/stats":                s.handleStats(),              // show JSON
		"GET:/sites":                s.handleSites(),              // show index of all sites
		"GET:/details/:url":         s.handleSiteDetails(),        // show details about a single site
		"GET:/lock/:url":            s.handleLockSite(),           // run the script to lock a site
		"GET:/unlock/:url":          s.handleUnlockSite(),         // run the script to unlock a site
		"GET:/stats/:url/:duration": s.handleSiteStats(),          // show site statistics file if it exists
		"GET:/create-site":          s.handleCreateForm(),         //show form to create a new site
		"POST:/create-site":         s.handleCreate(),             // deploy new site from skeleton
		"GET:/create-skeleton":      s.handleCreateSkeletonForm(), // show form to create new skeleton
		"POST:/create-skeleton":     s.handleCreateSkeleton(),     // create a new skeleton from existing site
		"POST:/disable":             s.handleDisableSite(),        // disable an existing site
		"POST:/enable":              s.handleEnableSite(),         // enable a disabled site
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

	// set up websocket connection for script output
	s.router.GET("/ws", s.handleWebSocket)
	s.Log("WS:   /ws")
}
