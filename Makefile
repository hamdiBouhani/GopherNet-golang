postgres:
	docker run --name postgis  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgis/postgis
createdb:
	docker exec -it postgis createdb --username=postgres --owner=postgres gopher_net