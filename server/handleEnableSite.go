package server

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleEnableSite() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			s.LogError(err.Error())
			return
		}

		domain := r.Form.Get("domain")

		allDisabledDomains := []string{}

		cmd := exec.Command("ls", "-1", "/etc/httpd/conf.d/")
		reader, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout

		// Make a new channel which will be used to ensure we get all output
		done := make(chan struct{})

		scanner := bufio.NewScanner(reader)
		go func() {
			for scanner.Scan() {
				line := scanner.Text()

				// short filenames be damned
				if len(line) < 10 {
					continue
				}

				if (line[len(line)-9:]) == ".disabled" {
					allDisabledDomains = append(allDisabledDomains, line)
				}
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
		fmt.Println(allDisabledDomains)

		s.Log(domain)
		for _, d := range allDisabledDomains {
			if d == domain+".conf.disabled" {
				// disable me
				s.Log("about to enable " + domain)
				os.Rename("/etc/httpd/conf.d/"+domain+".conf.disabled", "/etc/httpd/conf.d/"+domain+".conf")
				os.Rename("/etc/httpd/conf.d/"+domain+"-le-ssl.conf.disabled", "/etc/httpd/conf.d/"+domain+"-le-ssl.conf")
				exec.Command("apachectl", "graceful").Run()
				fmt.Fprint(w, "Domain "+domain+" was enabled")
				return
			}
		}

		fmt.Fprint(w, "Domain "+domain+" was not found in the list of disabled domains")
	}
}
