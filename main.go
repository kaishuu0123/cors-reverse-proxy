package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

var (
	Version  = "unset"
	Revision = "unset"
)

var (
	targetURL     string
	host          string
	port          uint64
	origin        string
	serverLogging bool
	version       bool
)

func init() {
	envTargetURL := os.Getenv("CORS_REVERSE_PROXY_TARGET_URL")

	envHost := os.Getenv("CORS_REVERSE_PROXY_HOST")
	if envHost == "" {
		envHost = "localhost"
	}

	envPort, _ := strconv.ParseUint(os.Getenv("CORS_REVERSE_PROXY_PORT"), 10, 16)
	if envPort == 0 {
		envPort = 8081
	}

	envOrigin := os.Getenv("CORS_REVERSE_PROXY_ORIGIN")
	if envOrigin == "" {
		envOrigin = "*"
	}
	envServerLogging, _ := strconv.ParseBool(os.Getenv("CORS_REVERSE_PROXY_SERVER_LOGGING"))

	pflag.StringVarP(&targetURL, "target-url", "t", envTargetURL, "")
	pflag.StringVarP(&host, "host", "h", envHost, "")
	pflag.Uint64VarP(&port, "port", "p", envPort, "")
	pflag.StringVarP(&origin, "origin", "o", envOrigin, "")
	pflag.BoolVarP(&serverLogging, "server-logging", "l", envServerLogging, "")
	pflag.BoolVarP(&version, "version", "v", false, "")
}

func run(targetURL string) error {
	target, err := url.Parse(targetURL)
	if err != nil {
		return err
	}

	modifyCORSResponse := func(res *http.Response) error {
		res.Header.Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
		res.Header.Set("Access-Control-Allow-Credentials", "true")
		res.Header.Set("Access-Control-Allow-Origin", origin)
		return nil
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(target)
	reverseProxy.ModifyResponse = modifyCORSResponse
	http.Handle("/", reverseProxy)

	addr := host + ":" + strconv.FormatUint(port, 10)
	return http.ListenAndServe(addr, nil)
}

func printHeader() {
	log.Println("Start CORS Reverse Proxy")
	log.Println("")
	log.Printf("Target URL: %s\n", targetURL)
	log.Printf("Host: %s\n", host)
	log.Printf("Port: %d\n", port)
	log.Println("")
	log.Printf("Please access to http://%s:%d/\n", host, port)
	log.Println("")
}

func main() {
	pflag.Parse()

	if version {
		log.Printf("%s version: %s revision: %s", os.Args[0], Version, Revision)
		os.Exit(0)
	}

	if targetURL == "" {
		log.Fatal("Target URL(--target-url or -t) option is required.")
	}

	printHeader()

	if err := run(targetURL); err != nil {
		log.Fatal(err)
	}
}
