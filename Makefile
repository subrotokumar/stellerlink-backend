clean:
	rm *.exe *.exe~

gen:
	go run github.com/99designs/gqlgen generate

clean:
	rm -f /stellerlink-backend.*

dev:
	CompileDaemon -command="./stellerlink-backend"

tunnel:
	cloudflared tunnel --url localhost:8080

start:
	go build -tags netgo -ldflags '-s -w' -o app
	./app