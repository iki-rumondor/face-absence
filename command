docker run --name postgres -e POSTGRES_PASSWORD=pgroot -p 5432:5432 -d postgres

docker build -t go-absence .

docker container create --name backend -e PGHOST=pg-absence -e PORT=8080 -e PGPORT=5432 -e PGUSER=postgres -e PGPASSWORD=pgroot -e PGNAME=face_absence -e SSLMODE="disable" -p 8080:8080 go-absence

docker container start backend

docker container run --name absence --rm -it -e PGHOST=postgres -e PGPORT=5432 -e PGUSER=postgres -e PGPASSWORD=pgroot -e PGNAME=face_absence -e SSLMODE="disable" -p 8080:8080 go-absence

docker network create face-absence

docker network connect face-absence pg-absence
docker network connect face-absence backend