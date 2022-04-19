package handlers
import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
	// "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
)
type deleteUser struct{
	rt *ShivaniCustomServerExample1.Runtime
}
func NewDeleteUser(rt *ShivaniCustomServerExample1.Runtime) users.DeleteUserHandler{
	return &deleteUser{rt:rt}
}

// func (d *deleteUser) Handle(del users.DeleteUserParams) middleware.Responder{
// 	//if the val returing from DeleteUser = nil then return NewDestroyOneNoContent else not nil then return 404
// 	isError:=d.rt.GetManager().DeleteUser(del.ID)
// 	if isError == nil{
// 		res:=users.NewDeleteUserNoContent()
// 		return res
// 	}
// 	res:=users.NewDeleteUserNotFound().WithPayload(isError)
// 	return res
// }

func (d *deleteUser) Handle(del users.DeleteUserParams) middleware.Responder{
	err :=d.rt.GetManager().DeleteUser(del.ID)
	fmt.Println("The Delete status from handler",err)
	if err != nil{
		derr,ok:=err.(domain.Err)
		fmt.Println("The typecasted Delete status from handler",derr.StatusCode())
		if ok {
			switch derr.StatusCode(){
				case 404:
					return users.NewDeleteUserNotFound().WithPayload(asErrorResponse(err.(*domain.Error)))
			}
		} else{
			return users.NewDeleteUserDefault(500).WithPayload(asErrorResponse(&domain.Error{Message:"Internal Server Error"}))
		}
    }
	
	return users.NewDeleteUserNoContent()
} 