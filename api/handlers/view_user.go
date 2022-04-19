package handlers
import (
	// "fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
)
type viewUser struct{
	rt *ShivaniCustomServerExample1.Runtime
}
func NewViewUser(rt *ShivaniCustomServerExample1.Runtime) users.ViewUserHandler{
	return &viewUser{rt:rt}
}

func (d *viewUser) Handle(vu users.ViewUserParams) middleware.Responder{
	us,err :=d.rt.GetManager().ViewUser(vu.ID)
	if err != nil{
		derr,ok:=err.(domain.Err)
		
		if ok {
			switch derr.StatusCode(){
				case 404:
					 return users.NewViewUserNotFound().WithPayload(asErrorResponse(err.(*domain.Error)))
			}
		} else{
			return users.NewViewUserDefault(500).WithPayload(asErrorResponse(&domain.Error{Message:"Internal Server Error"}))
		}
    }
	return users.NewViewUserOK().WithPayload(asUserResponse(us))
} 