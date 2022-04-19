package db
import(
	"fmt"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	// "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db/inmemory"
)

type DataStore interface{
	AddUser(*domain.User) (string, error)
	ListUsers(string,int32) ([]*domain.User,error)
	DeleteUser(string) error
	ViewUser(string) (*domain.User,error)
}
type DatastoreFactory func() (DataStore, error)
var factories map[string]DatastoreFactory
func RegisterDataStore(key string,value DatastoreFactory){
	if factories == nil{
		factories=make(map[string]DatastoreFactory)
	}
	factories[key]=value
}
func NewDataStore(dbType string) (DataStore,error){
	
	// switch dbType {
	// case "inmemory":
	// 	//inmemo client obj
	// 	return factories[dbType]()
	// case "mongo":
	// 	//inmemo client obj
	// default :
	// 	// return nil
	// }
	fmt.Println("The length is",len(factories))
	return factories[dbType]()
}