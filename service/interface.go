package service
import(
	"fmt"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db"

	// "context"
    // "time"
	// "log"
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/bson"
	// // "go.mongodb.org/mongo-driver/bson/primitive"
)
type Manager interface{
	CreateUser(*domain.User) (string,error)
	ListUsers(string,int32) ([]*domain.User,error)
	DeleteUser(string) error
	ViewUser(string) (*domain.User,error)
}

func NewManager(dbType string) Manager{
	//connection to mongodb should be made here
	
	// ctx,cancel:=context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil{
	// 	fmt.Println("The error is",err)
	// }
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil{
	// 	fmt.Println("The error is",err)
	// 	return &manager{ds: make(map[int]*domain.User)}
	// }
	// collection := client.Database("user_app").Collection("users")
	// ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// //Insert
	// res,err:=collection.InsertOne(ctx,bson.D{
	// 	{Key:"_id",Value:"1234"},
	// 	{Key:"name", Value:"Shivani"},
	// 	{Key:"address", Value:"ABC Park Colony"}, 
	// })
	// if err != nil{
	// 	fmt.Println("The error is",err)
	// }
	// id := res.InsertedID
	// fmt.Println("The ID is",id)
	// //Find
	// ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// cur, err := collection.Find(ctx, bson.D{})

	// if err != nil { log.Fatal(err) }
	// defer cur.Close(ctx)
	// for cur.Next(ctx) {
	// 	var result bson.D
	// 	err := cur.Decode(&result)
	// 	if err != nil { log.Fatal(err) }
	// 	fmt.Println("The result is",result)
	// }
	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	//Write for update

	//call factory function here
	// return &manager{ds: make(map[int]*domain.User)}
	ds,err:=db.NewDataStore(dbType)
	if err != nil{
		fmt.Println("The err is",err)
		return nil
	}
	return &manager{ds:ds}
}