module myapp

go 1.17

replace github.com/ghibbo/speedy => ../speedy

require (
	github.com/ghibbo/speedy v0.0.0-20211020185947-fdb7bd24a9b8
	github.com/go-chi/chi/v5 v5.0.4
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.1.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
