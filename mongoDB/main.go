package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URI  = "mongodb://localhost"
	PORT = "27017"
)

type Client struct {
	conn *mongo.Client
}

func createClient() Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s:%s", URI, PORT)))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	return Client{conn: client}
}

func listDatabases(client Client) []string {
	allDatabases, err := client.conn.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	return allDatabases
}

func listCollections(client Client, database string) []string {
	allCollections, err := client.conn.Database(database).ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	return allCollections
}

func showAllDocuments(client Client, database string, collection string) {
	var result []bson.M
	cursor, _ := client.conn.Database(database).Collection(collection).Find(context.TODO(), bson.M{}, options.Find())

	if err := cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total Documents In Collection : ", len(result))

	for _, v := range result {
		fmt.Print("{\n")
		for key, each := range v {
			fmt.Printf("%s :  %s\n", key, each)
		}
		fmt.Print("}\n")
	}
}

func findOneDocument(client Client, database string, collection string, filter bson.D) {
	var result bson.M

	if err := client.conn.Database(database).Collection(collection).FindOne(context.TODO(), filter, options.FindOne()).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatal("No Documents Founded")
		} else {
			log.Fatal("Unknown Error Occurred")
		}
	}

	fmt.Println(result)
}

func deleteOneDocument(client Client, database string, collection string, filter bson.D) {
	deleted, err := client.conn.Database(database).Collection(collection).DeleteOne(context.TODO(), filter, options.Delete())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatal("No Documents Founded")
		} else {
			log.Fatal("Unknown Error Occurred")
		}
	}

	fmt.Println("Document Deleted : ", deleted)
}

func createNewCollection(client Client, database string, collection string, document bson.D) {
	inserted, err := client.conn.Database(database).Collection(collection).InsertOne(context.TODO(), document, options.InsertOne())
	if err != nil {
		log.Fatal("Error occurred while creating a new document")
	}

	fmt.Println("Document Created : ", inserted)
}

func updateCollection(client Client, database string, collection string, filter bson.D, document bson.D) {
	updateDocument := bson.M{
		"$set": document,
	}

	updated, err := client.conn.Database(database).Collection(collection).UpdateOne(context.TODO(), filter, updateDocument, options.MergeUpdateOptions())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatal("No Documents Founded To Update")
		}
		log.Fatal("Some Error Occur While Updating")
	}

	fmt.Println("Document Updated : ", updated)
}

func getAllKeysOfCollection(client Client, database string, collection string) []string {
	var result []bson.M
	var keys []string

	cursor, _ := client.conn.Database(database).Collection(collection).Find(context.TODO(), bson.M{}, options.Find())

	if err := cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	for _, v := range result {
		for key, _ := range v {
			keys = append(keys, key)
		}
		break
	}

	return keys
}

func handleFilter(scanner bufio.Scanner, client Client, database string, collection string) bson.D {
	allKeys := getAllKeysOfCollection(client, database, collection)

	fmt.Println("Select Key For Searching")

	for i, v := range allKeys {
		fmt.Printf("%d: %s\n", i, v)
	}

	fmt.Print("Select Key : ")
	scanner.Scan()
	userKeyInput, _ := strconv.Atoi(scanner.Text())
	selectedKey := allKeys[userKeyInput]
	fmt.Println("You Selected : ", selectedKey)
	fmt.Printf("Filter %s By : ", selectedKey)
	scanner.Scan()
	userFilterInput := scanner.Text()
	filter := bson.D{{selectedKey, userFilterInput}}

	return filter
}

func getDocumentForInsertion(scanner bufio.Scanner, client Client, database string, collection string) bson.D {
	allKeys := getAllKeysOfCollection(client, database, collection)
	var document bson.D

	fmt.Println("Total Key Founded : ", len(allKeys))

	for _, v := range allKeys {
		fmt.Printf("Enter Value Of %s : ", v)
		scanner.Scan()
		userValue := scanner.Text()
		toInsert := bson.E{Key: v, Value: userValue}
		document = append(document, toInsert)
	}

	return document
}

func main() {
	cli := createClient()
	fmt.Println("Connected With Local Mongo Database")
	databases := listDatabases(cli)
	fmt.Println("Available Databases Are : ", len(databases))
	for i, v := range databases {
		fmt.Printf("%d: %s\n", i, v)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Database Number To Select : ")
	scanner.Scan()
	userDatabaseInput, _ := strconv.Atoi(scanner.Text())
	if userDatabaseInput > len(databases)-1 || userDatabaseInput < 0 {
		fmt.Println("choose from available databases (wrong number detected)")
		os.Exit(1)
	}

	selectedDatabase := databases[userDatabaseInput]
	fmt.Println("You Selected : ", selectedDatabase)
	fmt.Println("Getting All Collections For : ", selectedDatabase)
	collections := listCollections(cli, selectedDatabase)
	fmt.Println("-------------------------------------------------")
	fmt.Println("Available Collections Are : ", len(collections))
	for i, v := range collections {
		fmt.Printf("%d: %s\n", i, v)
	}
	fmt.Print("Enter Collection Number To Select : ")
	scanner.Scan()
	userCollectionInput, _ := strconv.Atoi(scanner.Text())
	if userCollectionInput > len(collections)-1 || userCollectionInput < 0 {
		fmt.Println("choose one the available collection")
		os.Exit(1)
	}

	selectedCollection := collections[userCollectionInput]
	fmt.Println("You Selected : ", selectedCollection)
	fmt.Println("-------------------------------------------------")
	fmt.Println("----- Choose Options -----")
	fmt.Println("1. Show All Documents")
	fmt.Println("2. Find A Specific Document")
	fmt.Println("3. Delete A Document")
	fmt.Println("4. Create A New Document")
	fmt.Println("5. Update A Collection")
	fmt.Print("Choose Any Of The Above : ")
	scanner.Scan()
	userChoice, _ := strconv.Atoi(scanner.Text())
	switch userChoice {
	case 1:
		showAllDocuments(cli, selectedDatabase, selectedCollection)
	case 2:
		filter := handleFilter(*scanner, cli, selectedDatabase, selectedCollection)
		findOneDocument(cli, selectedDatabase, selectedCollection, filter)
	case 3:
		filter := handleFilter(*scanner, cli, selectedDatabase, selectedCollection)
		deleteOneDocument(cli, selectedDatabase, selectedCollection, filter)
	case 4:
		document := getDocumentForInsertion(*scanner, cli, selectedDatabase, selectedCollection)
		createNewCollection(cli, selectedDatabase, selectedCollection, document)
	case 5:
		filter := handleFilter(*scanner, cli, selectedDatabase, selectedCollection)
		document := getDocumentForInsertion(*scanner, cli, selectedDatabase, selectedCollection)
		updateCollection(cli, selectedDatabase, selectedCollection, filter, document)
	}
}
