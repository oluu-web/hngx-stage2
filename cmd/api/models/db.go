package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var dbName = "Peoplestore"

// ConnectToDB establishes a connection to the Mongo Database
func ConnectToDB() error {
	errr := godotenv.Load()
	if errr != nil {
		log.Fatal("Error loading .env file: ", errr)
	}
	mongoURI := os.Getenv("MONGOURI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	log.Println("connected successfully")
	return nil
}

// GetDBCollection returns a reference to the collection in the Mongo Database
func GetDBCollection(collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func ValidateFields(person Person) error {
	if !validName(person.Name) {
		return fmt.Errorf("name must contain only letters")
	}

	if !validAge(person.Age) {
		return fmt.Errorf("age must be an integer")
	}

	if !validGender(person.Gender) {
		return fmt.Errorf("Gender should be male or female")
	}

	return nil
}

// validation functions
func validName(name string) bool {
	// Check if the name is a valid string without numbers
	return !containsNumbers(name)
}

func validAge(age int) bool {
	// Check if the age is a positive integer
	return age >= 0
}

func validGender(gender string) bool {
	// Check if the gender is either "male" or "female"
	return strings.EqualFold(gender, "male") || strings.EqualFold(gender, "female")
}

func containsNumbers(s string) bool {
	// Check if the string contains any digit (0-9)
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

// this checks if a person with a given name alresdy exists in the database.
func CheckDuplicate(name string) (bool, error) {
	collection := GetDBCollection("People")

	filter := bson.M{"name": name}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// creates a new person in the database, and return the result
func CreateNewPerson(person Person) (string, error) {
	collection := GetDBCollection("People")

	result, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(string); ok {
		return oid, nil
	}

	return "Created Successfully", nil
}

// retrieves a person by name from the database
func GetPerson(name string) (Person, error) {
	collection := GetDBCollection("People")

	filter := bson.M{
		"name": name,
	}

	var person Person
	err := collection.FindOne(context.Background(), filter).Decode(&person)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Person{}, fmt.Errorf("Person not found")
		}
		return Person{}, err
	}
	return person, nil
}

// updates an existing person in the database
func UpdatePerson(name string, updatedPerson Person) error {
	collection := GetDBCollection("People")

	filter := bson.M{"name": name}
	update := bson.M{
		"$set": bson.M{
			"name":   updatedPerson.Name,
			"age":    updatedPerson.Age,
			"gender": updatedPerson.Gender,
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// delete person by name from the database
func DeletePerson(name string) error {
	collection := GetDBCollection("People")

	// check if person exists before deletio
	_, err := GetPerson(name)
	if err != nil {
		return err
	}

	filter := bson.M{"name": name}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
