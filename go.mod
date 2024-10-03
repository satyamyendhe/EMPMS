module vsys.empms.web

go 1.22.5

require (
	github.com/gorilla/mux v1.8.1
	vsys.empms.commons v0.0.0

)

require (
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/rs/cors v1.11.0
)

replace vsys.empms.commons => ../vsys-empms-commons
