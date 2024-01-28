package database

import (
	"backend/model"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	GetUserbyUUID(id string, c *gin.Context)
	LoginUser(User model.Users, c *gin.Context)
	UpdateDailySpends(User *model.Users, c *gin.Context)
	UpdateTotalAmount(User *model.Users, c *gin.Context)
	InsertUser(User model.Users, c *gin.Context)
	GetUsers(c *gin.Context) map[string]interface{}
	Health() map[string]string
}

type service struct {
	db *mongo.Client
}

var Collection_Main *mongo.Collection

var (
	// host        = os.Getenv("DB_HOST")
	// port        = os.Getenv("DB_PORT")
	DB_USER      = os.Getenv("DB_USERNAME")
	DB_PASSWORD  = os.Getenv("DB_ROOT_PASSWORD")
	DB_NAME      = os.Getenv("DB_NAME")
	DB_NEW_USERS = os.Getenv("DB_COLLECTION_NEW_REG")
	DB_MAIN      = os.Getenv("DB_COLLECTION_MAIN")
	//database = os.Getenv("DB_DATABASE")
)

func New() Service {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.ifstf4g.mongodb.net/%s?retryWrites=true&w=majority", DB_USER, DB_PASSWORD, DB_NAME)))

	if err != nil {
		log.Fatal(err)

	}
	Collection_Main = client.Database(DB_NAME).Collection(DB_MAIN)
	return &service{
		db: client,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) LoginUser(User model.Users, c *gin.Context) {
	var temp model.Users
	filter := bson.D{{Key: "username", Value: User.Username}}
	err := Collection_Main.FindOne(context.Background(), filter).Decode(&temp)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found"})
		return
	}

	if User.Password != temp.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Password"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Logged In Successfully",
		"user":    temp,
	})
}

func (s *service) InsertUser(User model.Users, c *gin.Context) {

	var temp model.Users
	filter := bson.D{{Key: "email", Value: User.Email}}
	err := Collection_Main.FindOne(context.TODO(), filter).Decode(&temp)

	if err == nil {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"message": "User already exists",
		})
		return
	}

	filter = bson.D{{Key: "username", Value: User.Username}}
	err = Collection_Main.FindOne(context.TODO(), filter).Decode(&temp)

	if err == nil {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"message": "User already exists, Username already Taken",
		})
		return
	}

	User.UUID = uuid.NewString()
	//insert in mongoDB
	_, err = Collection_Main.InsertOne(context.Background(), User)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User added successfully",
	})
}

func (s *service) GetUsers(c *gin.Context) map[string]interface{} {
	var users []model.Users

	cursor, err := Collection_Main.Find(context.TODO(), bson.D{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in fetching users"})
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.TODO()) {
		var user model.Users

		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return map[string]interface{}{
		"users": users,
	}
}

func (s *service) UpdateTotalAmount(User *model.Users, c *gin.Context) {

	var tempUser model.Users
	err := Collection_Main.FindOne(context.TODO(), bson.M{"uuid": User.UUID}).Decode(&tempUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found"})
	}

	filter := bson.D{{Key: "uuid", Value: User.UUID}}
	newTotalAmount := append(tempUser.Total_Amt, User.Total_Amt...)
	fmt.Printf("new%+v", newTotalAmount)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "total_amt", Value: newTotalAmount}}}}

	_, err = Collection_Main.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in updating total amount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Total amount updated successfully",
	})

}

func (s *service) UpdateDailySpends(User *model.Users, c *gin.Context) {

	var tempUser model.Users
	err := Collection_Main.FindOne(context.TODO(), bson.M{"uuid": User.UUID}).Decode(&tempUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found"})
	}

	filter := bson.D{{Key: "uuid", Value: User.UUID}}
	newSpends := append(tempUser.Spends, User.Spends...)
	fmt.Printf("new%+v", newSpends)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "spends", Value: newSpends}}}}

	_, err = Collection_Main.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in updating daily spends"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daily spends updated successfully",
	})
}

func (s *service) GetUserbyUUID(id string, c *gin.Context) {

	var userFound model.Users

	err := Collection_Main.FindOne(context.TODO(), bson.D{{Key: "uuid", Value: id}}).Decode(&userFound)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found !!!",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userFound,
	})
}
