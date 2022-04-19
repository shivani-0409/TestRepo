package main

import (
    "fmt"
	"flag"
	"log"
	"github.com/go-openapi/loads"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations"
    "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/api/handlers"
    "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
    _ "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db/inmemory"
    _ "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db/mongo"
)

func main() {
	var portFlag = flag.Int("port", 4000, "Port to run this service on")
    // load embedded swagger file
    swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
    if err != nil {
        log.Fatalln(err)
    }

    // create new service API
    api := operations.NewShivaniCustomServerExample1API(swaggerSpec)
    // added by shivani
    v:=ShivaniCustomServerExample1.NewRunTime("Shivani")
    // v:=ShivaniCustomServerExample1.Runtime{AppName:"Shivani"}
    api.UsersFindUsersHandler=handlers.NewFindUser(v)
    fmt.Println("The Application Name is",v.AppName)
    api.UsersAddUserHandler=handlers.NewAddUser(v)
    api.UsersDeleteUserHandler=handlers.NewDeleteUser(v)
    api.UsersViewUserHandler=handlers.NewViewUser(v)
    // added by shivani
    server := restapi.NewServer(api)
    defer server.Shutdown()
    
    // parse flags
    flag.Parse()
    // set the port this service will be run on
    server.Port = *portFlag

    // TODO: Set Handle

    // serve API
    if err := server.Serve(); err != nil {
        log.Fatalln(err)
    }
}