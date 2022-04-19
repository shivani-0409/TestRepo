package inmemory
import(
	"fmt"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db"
	// "github.com/segmentio/ksuid"
)

func init(){
	db.RegisterDataStore("inmemory",NewClient)
	fmt.Println("Call is in Init")
}

func NewClient() (db.DataStore,error){
	return &client{ds:make(map[string]*domain.User)},nil
}

type client struct{
	ds map[string]*domain.User
}

func (c *client) AddUser(user *domain.User) (string, error){
	// user.ID= ksuid.New().String()
	c.ds[user.ID]=user
	return user.ID,nil
}
func (c *client) ListUsers(queryName string, queryLimit int32) ([]*domain.User,error){
	var userInfo =[]*domain.User{}
	if queryName != "" && queryLimit !=0{
		fmt.Println("The queryName and queryLimit are",queryName,queryLimit)
	}
	var lim int32 = 1
	for _,value := range c.ds{
		if queryName !=""{
			if queryName == value.Name{
				userInfo = append(userInfo, value)
			}
		} else{
			userInfo = append(userInfo, value)
		}
		if queryLimit != 0{
			if lim == queryLimit{
				break
			}
			lim++
		} 
	}
	return userInfo,nil
}
func (c *client) DeleteUser(id string) error{
	fmt.Println("The DeleteUser ID param is",id)
	if _, ok := c.ds[id]; ok{
		fmt.Println("The DeleteUser status param is",ok)
		delete(c.ds,id)
		return nil
  }
return &domain.Error{Code:404,Message:"User doesn't exist"}
}

func (c *client) ViewUser(id string) (*domain.User,error){
	// var userInfo =[]*domain.User{}
	var userInfo *domain.User
	var ok bool
	if userInfo, ok = c.ds[id]; !ok{
		 return nil,&domain.Error{Code:404,Message:"User doesn't exist"}
  	}
	// for _,value := range c.ds{
	// 	fmt.Println("The value, id and value in view user",value,id,c.ds[id])
	// 	if value == c.ds[id]{
	// 		// userInfo = append(userInfo, value)
	// 		userInfo=value
	// 	}
	// }
	// fmt.Println("The value of userInfo is",userInfo.Name)
	// if userInfo.Name == ""{
	// 	return nil,nil
	// }
	return userInfo,nil
}

