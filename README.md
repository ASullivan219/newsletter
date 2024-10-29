# Newsletter

Go, templ, tailwind, SQLite

## TODO:

-[X] Implement Confirmation Email on subscribe  
-[X] Validate Email addresses with email package on sign up  
-[X] Create a successful sign up component to return  
-[ ] Implement relationship feature in sign-up  
-[X] Generate random codes  
-[X] Delete components folder  
-[X] Create representative docker compose file  
-[ ] Create 0.0.1 release and pull image to vps  
-[ ] Install tailwind and build styles locally to move away from cdn  
-[ ] Host and serve htmx javascript from static folder to move away from cdn  

## Run with Go locally

`go run ./cmd/main.go`

## Run with docker

- `sudo docker build -t newsletter .`
- `docker-compose up -d`


## Develop

- Install go templ  
    - `go install github.com/a-h/templ/cmd/templ@latest`
    - `https://templ.guide/quick-start/installation`  
- Install air for live re load  
    - `go install github.com/air-verse/air@latest`  
    - `https://github.com/air-verse/air`
- Make a resources folder in project root, and `subscribers.db` file within.
    - `mkdir resources`  
    - `touch resources/subscribers.db`  
- Run command `air` from the project root.  
- relevant files will be watched by air, start devloping locally and project will be rebuilt automatically on save.  


