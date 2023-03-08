package main

import (
	"flag"
	"fmt"

	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	"github.com/go-resty/resty/v2"

	// plug in Caddy modules here
	_ "github.com/Tasudo/caddy-jailbait/v2"
	_ "github.com/abiosoft/caddy-exec"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/greenpau/caddy-git"
	_ "github.com/hairyhenderson/caddy-teapot-module"
	_ "github.com/mholt/caddy-webdav"
	_ "github.com/n0trace/caddy-ftp"
	_ "github.com/porech/caddy-maxmind-geolocation"
	_ "github.com/ueffel/caddy-brotli"
)

var (
	wh_discord bool   = true
	wh_url     string = "https://discord.com/api/webhooks/id/token"
	wh_post    bool   = true
	wh_message string = "Caddy has been started"
)

func main() {
	caddycmd.RegisterCommand(caddycmd.Command{
		Name: "rick",
		Func: caddycmd.CommandFunc(func(f caddycmd.Flags) (int, error) {
			fmt.Println("I'M PICKLE RICK!\nhttps://youtu.be/_gRnvDRFYN4")
			return 0, nil
		}),
		Usage: "",
		Short: "Find it out yourself",
		Long:  "Find it out yourself",
		Flags: &flag.FlagSet{},
	})

	caddycmd.RegisterCommand(caddycmd.Command{
		Name: "proxy-stats",
		Func: caddycmd.CommandFunc(func(f caddycmd.Flags) (int, error) {
			req, err := caddycmd.AdminAPIRequest("http://localhost:2019", "GET", "/reverse_proxy/upstreams", nil, nil)
			if err != nil {
				return 1, err
			}

			fmt.Println(req.Body)
			return 0, nil
		}),
		Usage: "",
		Short: "Shows the reverse proxy stats",
		Long:  "Shows the reverse proxy stats",
		Flags: &flag.FlagSet{},
	})

	caddycmd.RegisterCommand(caddycmd.Command{
		Name: "whstart",
		Func: caddycmd.CommandFunc(func(f caddycmd.Flags) (int, error) {
			client := resty.New()
			if wh_discord {
				_, err := client.R().
					SetHeader("Content-Type", "application/json").
					SetBody(`{"content": "` + wh_message + `"}`).
					Post(wh_url)

				if err != nil {
					return 1, err
				}
			} else {
				if wh_post {
					_, err := client.R().
						SetHeader("Content-Type", "application/json").
						SetBody(wh_message).
						Post(wh_url)

					if err != nil {
						return 1, err
					}
				} else {
					_, err := client.R().
						SetHeader("Content-Type", "application/json").
						Get(wh_url)

					if err != nil {
						return 1, err
					}
				}
			}
			caddycmd.Commands()["start"].Func(f)
			return 0, nil
		}),
		Usage: "",
		Short: "Starts the Caddy process in the background, sends a message to your webhook and then returns",
		Long:  "Starts the Caddy process in the background, sends a message to your webhook and then returns",
		Flags: &flag.FlagSet{},
	})

	caddycmd.Main()
}
