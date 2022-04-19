package handlers
import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
)
func NewAddUser(rt *ShivaniCustomServerExample1.Runtime) users.AddUserHandler{
	return &addUser{rt:rt}
}
type addUser struct{
	rt *ShivaniCustomServerExample1.Runtime
}
func (au *addUser) Handle(aup users.AddUserParams) middleware.Responder{
	// n:="Sharma Shivani"
	// fmt.Println("The name is",*aup.Body.Name)
	// fmt.Println("The name is",aup.Body.Address)
	// fmt.Println("The name is",aup.Body.ID)
	// us:= &models.User{Address:"PQR",ID:1,Name:&n}
	// res:=users.NewAddUserCreated().WithPayload(us)
	// usr:=&domain.User{ID:aup.Body.ID,Name:*aup.Body.Name,Address:aup.Body.Address}
	// au.rt.GetManager().CreateUser(aup.Body)
	usr:=&domain.User{Name:*aup.Body.Name,Address:aup.Body.Address}
	fmt.Println("The application name(from POST) is ",au.rt.AppName)
	fmt.Println("The len of name is",len(*aup.Body.Name))
	if _,err :=au.rt.GetManager().CreateUser(usr); err != nil{
		return users.NewAddUserBadRequest().WithPayload(asErrorResponse(err.(*domain.Error)))
	}
	
	return users.NewAddUserCreated()
}

