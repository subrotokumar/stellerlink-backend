clean:
	rm *.exe *.exe~

gen:
	go run github.com/99designs/gqlgen generate

dev:
	rm -f /stellerlink-backend.*
	CompileDaemon -command="./stellerlink-backend"

tunnel:
	cloudflared tunnel --url localhost:8080