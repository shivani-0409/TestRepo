package ShivaniCustomServerExample1
import(
	// "fmt"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/service"
)

func NewRunTime(an string) (*Runtime){
	return &Runtime{
		AppName:an,
		svc: service.NewManager("mongo"),
	}
}

type Runtime struct{
	AppName string
	svc service.Manager
}

func (rt *Runtime) GetApplicationName() string{
	return rt.AppName
}

func (rt *Runtime) SetApplicationName(s string){
	rt.AppName = s
}

// func (rt *Runtime) InitiateManager() *service.Manager{
// 	v:=service.NewManager()
// 	fmt.Println("Hey I am Shivani Sharma",v)
// 	return v
// }

func (rt *Runtime)GetManager() service.Manager{
	return rt.svc
}

