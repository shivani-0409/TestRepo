package handlers
import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
)

// func NewFindUser(rt *runtime) users.FindUsersHandler{
func NewFindUser(rt *ShivaniCustomServerExample1.Runtime) users.FindUsersHandler{
	return &findUser{rt:rt}
}
type findUser struct{
	rt *ShivaniCustomServerExample1.Runtime
}
func (f *findUser) Handle(fup users.FindUsersParams) middleware.Responder{
	var limit int32
	var name string
	var us []*domain.User
	if fup.Limit != nil && fup.Name != nil {
		limit = *fup.Limit
		name = *fup.Name
		us,_ = f.rt.GetManager().ListUsers(name, limit)
	}else{
		limit = *fup.Limit
		us,_ = f.rt.GetManager().ListUsers("", limit)
	}
	usResponse:=[]*models.User{}
	for _,usr:=range us{
		usResponse=append(usResponse,asUserResponse(usr))
	}
	res:=users.NewFindUsersOK().WithPayload(usResponse)
	return res
}

