package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOneStudent(ctx context.Context, coll *mongo.Collection, newStudent Student) {
	newId, err := coll.InsertOne(ctx, newStudent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID baru : %v \n", (*newId).InsertedID)
}

func FindAllStudent(ctx context.Context, coll *mongo.Collection) {
	var results []bson.M //map[string]interface{}
	allDocumentStudentCursor, err := coll.Find(ctx, bson.D{})

	if err != nil {
		log.Fatalln(err)
	}

	defer allDocumentStudentCursor.Close(ctx)

	err = allDocumentStudentCursor.All(ctx, &results)

	if err != nil {
		log.Fatalln(err)
	}

	for _, doc := range results {
		fmt.Printf("id:%v, name:%v, age:%v", doc["_id"], doc["name"], doc["age"])
	}
}

func FindStudentByGenderAndAge(ctx context.Context, coll *mongo.Collection, gender string, age int) {
	filterGenderAndAge := bson.D{
		{
			"$and", bson.A{
				bson.D{
					{"gender", gender},
					{"age", age},
				},
			},
		},
	}

	projection := bson.D{
		{"_id", 0},
		{"name", 1},
	}

	findOpts := options.Find().SetProjection(projection)

	var results []bson.M

	resultCursor, err := coll.Find(ctx, filterGenderAndAge, findOpts)
	if err != nil {
		log.Fatalln(err)
	}

	defer resultCursor.Close(ctx)
	err = resultCursor.All(ctx, &results)

	if err != nil {
		log.Fatalln(err)
	}

	for _, doc := range results {
		fmt.Printf("name: %v \n", doc["name"])
	}

}

func FindStudentStructByGenderAndAge(ctx context.Context, coll *mongo.Collection, gender string, age int) {
	filterGenderAndAge := bson.D{
		{
			"$and", bson.A{
				bson.D{
					{"gender", gender},
					{"age", age},
				},
			},
		},
	}

	projection := bson.D{
		{"_id", 0},
		{"name", 1},
	}

	findOpts := options.Find().SetProjection(projection)
	results := make([]*Student, 0)

	resultCursor, err := coll.Find(ctx, filterGenderAndAge, findOpts)
	if err != nil {
		log.Fatalln(err)
	}

	defer resultCursor.Close(ctx)

	for resultCursor.Next(ctx) {
		var row Student
		err := resultCursor.Decode(&row)
		if err != nil {
			log.Fatalln(err)
		}
		results = append(results, &row)
	}

	for _, doc := range results {
		fmt.Printf("name:%v \n", doc.Name)
	}
}

func CountDocument(ctx context.Context, coll *mongo.Collection) {
	itemCount, err := coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Jumlah", itemCount)

}

func CountDocumentByAge(ctx context.Context, coll *mongo.Collection, age int) {
	itemCount, err := coll.CountDocuments(ctx, bson.M{"age": age})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Jumlah", itemCount)

}

func CountProductByCategory(ctx context.Context, productColl *mongo.Collection, category string) {
	matchStage := bson.D{{"$match", bson.D{{"category", category}}}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$category"},
		{"total", bson.D{{"$sum", 1}}},
	}}}

	aggCursor, err := productColl.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage}) //pipeline buat ngebungkus match sama group
	if err != nil {
		log.Fatalln(err)
	}

	defer aggCursor.Close(ctx)

	var results []bson.M
	if err = aggCursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	for _, doc := range results {
		fmt.Println()
		fmt.Printf("Group: %v, Total: %v \n", doc["_id"], doc["total"])
	}
}
