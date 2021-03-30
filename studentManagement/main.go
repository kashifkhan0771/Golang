package main

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
	"time"
)

type Client struct {
	conn *mongo.Client
}

type Student struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func createClient() Client {
	cli, err := mongo.NewClient(options.Client())
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err := cli.Connect(ctx); err != nil {
		log.Fatal("Cannot connect with database : ", err)
	}

	return Client{conn: cli}
}

func getStudentDetails(scanner bufio.Scanner) Student {
	fmt.Println("Press Enter If You Dont Want To Add Any Value")
	stu := &Student{}
	fmt.Print("Enter Student Id: ")
	scanner.Scan()
	stu.ID = scanner.Text()

	fmt.Print("Enter Student Name : ")
	scanner.Scan()
	stu.Name = scanner.Text()

	fmt.Print("Enter Student Age : ")
	scanner.Scan()
	stu.Age, _ = strconv.Atoi(scanner.Text())

	if stu.ID == "" {
		log.Fatal("Id cannot be empty")
	}

	return *stu
}

func registerNewStudent(client Client, stu *Student) {
	if _, err := client.conn.Database("student").Collection("records").InsertOne(context.TODO(), stu); err != nil {
		_ = fmt.Errorf("failed to register new student : %s", err)
	}

	fmt.Println("Registered New Student : ", stu)
}

func deleteStudent(client Client, id string) {
	if _, err := client.conn.Database("student").Collection("records").DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		_ = fmt.Errorf("failed to delete student : %s", err)
	}

	fmt.Println("Student Deleted : ", id)
}

func updateStudent(client Client, id string, updateStu *Student) {
	if _, err := client.conn.Database("student").Collection("records").UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updateStu}); err != nil {
		_ = fmt.Errorf("failed to update student : %s", err)
	}

	fmt.Println("Student Updated : ", id)
}

func getStudent(client Client, id string) {
	var student *Student
	if err := client.conn.Database("student").Collection("records").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&student); err != nil {
		_ = fmt.Errorf("failed to get student : %s", err)
	}

	fmt.Println("Id : ", student.ID)
	fmt.Println("Name : ", student.Name)
	fmt.Println("Age : ", student.Age)

}

func getStudentID(scanner bufio.Scanner) string{
	fmt.Print("Please Enter Student Id : ")
	scanner.Scan()
	studentId := scanner.Text()

	return studentId
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cli := createClient()
	fmt.Println("Press 1 For Registering New Student")
	fmt.Println("Press 2 For Deleting Student")
	fmt.Println("Press 3 For Updating Student Record")
	fmt.Println("Press 4 To Get A Student Record")
	fmt.Print("Choose : ")
	scanner.Scan()
	choice := scanner.Text()
	switch choice {
	case "1":
		student := getStudentDetails(*scanner)
		registerNewStudent(cli, &student)
	case "2":
		studentId := getStudentID(*scanner)
		deleteStudent(cli, studentId)
	case "3":
		fmt.Println("***** Which Student To Update *****")
		studentId := getStudentID(*scanner)
		fmt.Println("***** New Updated Values *****")
		updated := getStudentDetails(*scanner)
		updateStudent(cli, studentId, &updated)
	case "4":
		studentId := getStudentID(*scanner)
		getStudent(cli, studentId)
	}
}
