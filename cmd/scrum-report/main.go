package main

import (
	"log"
	"net"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/scrum-report/internal/config"
	"github.com/oleg-balunenko/scrum-report/internal/logger"
	"github.com/oleg-balunenko/scrum-report/internal/reporter"
)

func main() {
	printVersion()

	cfg := config.Load()
	if cfg.Debug {
		ip, err := getIP()
		if err != nil {
			log.Fatalf("failed to get ip: %v", err)
		}

		cfg.Host = ip
	}

	logger.SetUp(cfg)

	r := reporter.New(cfg)

	log.Fatal(r.Run())
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", errors.Wrap(err, "failed to get addresses")
	}

	var ip string

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}

	return ip, nil
}
