package reporter

import (
	"net/http"
	netUrl "net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// open opens the specified URL in the default browser of the user.
func open(url string) error {

	url, err := sanitizeURL(url)
	if err != nil {
		return errors.Wrap(err, "invalid url")
	}

	for {
		time.Sleep(time.Second)

		log.Debug("Checking if started...")

		resp, err := http.Get(url) // nolint
		if err != nil {
			log.Warn("Failed:", err)
			continue
		}
		err = resp.Body.Close()
		if err != nil {
			log.Errorf("failed to close body: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Warn("Not OK:", resp.StatusCode)
			continue
		}
		// Reached this point: server is up and running!
		break
	}
	log.Debug("SERVER UP AND RUNNING!")
	log.Debug("Opening browser")

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	log.Debugf("try to execute: %s %s", cmd, args)
	return exec.Command(cmd, args...).Start()
}

// sanitizeURL validates url and add protocol prefix
func sanitizeURL(url string) (string, error) {
	// workaround for this issue: https://github.com/golang/go/issues/18824
	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// validate url
	u, err := netUrl.Parse(url)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse url")
	}
	return u.String(), nil
}

func processFormValue(data string) []string {
	var res []string
	arr := strings.Split(data, "\r\n")
	// remove empty strings
	for _, a := range arr {
		if a != "" {
			res = append(res, a)
		}
	}
	return res
}
