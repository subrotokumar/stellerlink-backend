clean:
	rm *.exe *.exe~

gen:
	go run github.com/99designs/gqlgen generate

dev:
	CompileDaemon -command="./stellerlink-backend"