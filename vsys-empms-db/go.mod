module vsys.empms.dbhelper

go 1.22.5

require (
	github.com/gorilla/mux v1.8.1
	vsys.empms.commons v0.0.0

)

require github.com/lib/pq v1.10.9

// require github.com/golang-jwt/jwt/v5 v5.2.1 // indirect

replace vsys.empms.commons => ../vsys-empms-commons

replace vsys.empms.rest => ../vsys-empms-rest
