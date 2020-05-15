package server

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleDisableSite() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			s.LogError(err.Error())
			return
		}

		domain := r.Form.Get("domain")

		allDomains := []string{}

		cmd := exec.Command("ls", "-1", "/var/www/prod")
		reader, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout

		// Make a new channel which will be used to ensure we get all output
		done := make(chan struct{})

		scanner := bufio.NewScanner(reader)
		go func() {
			for scanner.Scan() {
				line := scanner.Text()
				allDomains = append(allDomains, line)
			}

			// We're all done, unblock the channel
			done <- struct{}{}
		}()

		// Start the command and check for errors
		cmd.Start()

		// Wait for all output to be processed
		<-done

		// Wait for the command to finish
		cmd.Wait()
		fmt.Println(allDomains)

		s.Log(domain)
		for _, d := range allDomains {
			if d == domain {
				// disable me
				s.Log("about to disable " + domain)
				os.Rename("/etc/httpd/conf.d/"+domain+".conf", "/etc/httpd/conf.d/"+domain+".conf.disabled")
				os.Rename("/etc/httpd/conf.d/"+domain+"-le-ssl.conf", "/etc/httpd/conf.d/"+domain+"-le-ssl.conf.disabled")
				exec.Command("apachectl", "graceful").Run()
				fmt.Fprint(w, "Domain "+domain+" was disabled")
				return
			}
		}

		fmt.Fprint(w, "Domain "+domain+" was not found in the list of available domains")
	}
}
