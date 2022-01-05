package main

import (
	"context"
	"fmt"
	"log"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)



const uri = "mongodb://localhost:27017"

func main(){
	credential := options.Credential{
		Username: "jack",
		Password: "test",
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(uri).SetAuth(credential)

	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()

	client,err := mongo.Connect(ctx,clientOptions)
	if err !=nil{
		log.Fatal(err)
	}

	defer func ()  {
		if err=client.Disconnect(ctx); err!=nil{
			panic(err)
		}
	}()

	err = client.Ping(ctx,readpref.Primary())
	if err !=nil{
		panic(err)
	}
	fmt.Println("Successfully to connect and pinged")



	// coll := client.Database("db_enigma").Collection("students") //Define collection yang mana
	productColl := client.Database("db_enigma").Collection("products")
	// const layout = "2006-01-02"

	// dt,_ :=time.Parse(layout,"2022-01-05")
	// // newId, err := coll.InsertOne(ctx,bson.D{
	// // 	{"name","Jojon"},
	// // 	{"gender","M"},
	// // 	{"age",35},
	// // 	{"joinDate",primitive.NewDateTimeFromTime(dt)},
	// // 	{"idCard", "301"},
	// // 	{"senior",true},
	// // })

	// newStudent := Student{
	// 	Id: primitive.NewObjectID(),
	// 	Name: "Raya",
	// 	Gender: "F",
	// 	Age: 20,
	// 	JoinDate: dt,
	// 	IdCard: "304",
	// 	Senior: false,
	// }

	// newId,err := coll.InsertOne(ctx,newStudent)

	// if err !=nil{
	// 	log.Fatal(err)
	// }

	// fmt.Printf("ID baru : %v \n", (*newId).InsertedID)

	// InsertOneStudent(ctx,coll,newStudent)

	// FindAllStudent(ctx,coll)
	// CountDocument(ctx,coll)
	// CountDocumentByAge(ctx,coll,23)

	// fmt.Println("With bson.D")
	// FindStudentByGenderAndAge(ctx,coll,"F",23)

	// fmt.Println("With Student Struct")
	// FindStudentStructByGenderAndAge(ctx,coll,"F",23)

	CountProductByCategory(ctx,productColl,"food")

}