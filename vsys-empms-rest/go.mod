module vsys.empms.rest

go 1.22.5

require (
	github.com/gorilla/mux v1.8.1
	vsys.empms.commons v0.0.0
)

replace vsys.empms.commons => ../vsys-empms-commons
