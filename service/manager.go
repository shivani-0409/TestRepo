package service
import(
	"fmt"
	"time"
	// "github.com/go-openapi/swag"	 
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db"
	"github.com/segmentio/ksuid"
	// "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
	// "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
)

type manager struct{
	// ds map[int]*domain.User
	ds db.DataStore
}

func (m *manager) CreateUser(input *domain.User) (string,error){
	input.CreatedAt=time.Now().UTC()
	input.ID=ksuid.New().String()
		if len(input.Name) <=2 {
			return "",&domain.Error{Code:400,Message:"Name should be atleast 3 chars long"}
		}
		return m.ds.AddUser(input)
}

func (m *manager) ListUsers(queryName string, queryLimit int32) ([]*domain.User,error){
	return m.ds.ListUsers(queryName,queryLimit)
}

func (m *manager) DeleteUser(i string) error{
	_,err:=m.ds.ViewUser(i)
	if err != nil{
		return &domain.Error{Code:404,Message:"User doesn't exist"}
	}
	if err := m.ds.DeleteUser(fmt.Sprint(i)); err != nil{
		return &domain.Error{Code:404,Message:"User doesn't exist"}
	}
	return nil
  }

  func (m *manager) ViewUser(i string) (*domain.User,error){
	fmt.Printf("The type of id is %T",fmt.Sprint(i))
	return m.ds.ViewUser(i)
  }
   
   