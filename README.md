# custom-caddy
A template for creating own caddy instances with custom plugins and commands.

## Additional plugins

- [Caddy Jailbait](https://github.com/Tasudo/caddy-jailbait)
- [caddy-exec](https://github.com/abiosoft/caddy-exec)
- [caddy-git](https://github.com/greenpau/caddy-git)
- [caddy-teapot-modules](https://github.com/hairyhenderson/caddy-teapot-module)
- [caddy-webdav](https://github.com/mholt/caddy-webdav)
- [caddy-ftp](https://github.com/n0trace/caddy-ftp)
- [caddy-maxmind-geolocation](https://github.com/porech/caddy-maxmind-geolocation)
- [caddy-brotli](https://github.com/ueffel/caddy-brotli)

## Additional commands

- "rick"
- "proxy-stats"
- "whstart" (start with webhook notification)

## Installation
### with Github CLI
```bash
gh repo clone Minecodes/custom-caddy
cd custom-caddy
go install
go build main.go
```
### with Git
```bash
git clone https://github.com/Minecodes/custom-caddy
cd custom-caddy
go install
go build main.go
```

## Example Caddyfile
```
localhost {
	templates
	file_server {
		root test/
	}
}

rick.localhost {
	respond "I'M PICKEL RICK!"
}

teapot.localhost {
	route {
		teapot
	}
}

jailbait.localhost {
	route {
		jailbait
	}
}
```