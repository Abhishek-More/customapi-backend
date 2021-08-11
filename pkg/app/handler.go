package app

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Abhishek-More/customapi-backend/pkg/data/models"
	"github.com/Abhishek-More/customapi-backend/pkg/utils"
)

func (a *App) InitializeRouter() {
    port := utils.GetPort()
	a.Router = mux.NewRouter()
	
	//TODO: Show API Info
	a.Router.HandleFunc("/", utils.RouteNotImplemented)

	a.Router.HandleFunc("/sample/posts", a.PostsHandler)
	a.Router.HandleFunc("/sample/posts/", a.PostsHandler)
	a.Router.HandleFunc("/sample/posts/{id}", a.PostHandler)
	a.Router.HandleFunc("/sample/posts/{id}/", a.PostHandler)

	a.Router.HandleFunc("/sample/users", a.UsersHandler)
	a.Router.HandleFunc("/sample/users/", a.UsersHandler)
	a.Router.HandleFunc("/sample/users/{id}", a.UserHandler)
	a.Router.HandleFunc("/sample/users/{id}/", a.UserHandler)

    http.ListenAndServe(":"+port, a.Router)
}

func (a *App) PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection := a.Client.Database("sample-data").Collection("posts")
	
	var posts []models.Post

	cur, err := collection.Find(context.TODO(), bson.M{})
	utils.CheckFatalError(err)
	
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var post models.Post
		err := cur.Decode(&post)
		utils.CheckFatalError(err)
		
		posts = append(posts, post)
	}

	err = cur.Err()
	utils.CheckFatalError(err)

	json.NewEncoder(w).Encode(posts)
}

func (a *App) PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection := a.Client.Database("sample-data").Collection("posts")

	var post models.Post
	var params = mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.GetError(err, w)
	}

	filter := bson.M{"_id": id}
	err = collection.FindOne(context.TODO(), filter).Decode(&post)
	utils.CheckError(err, w, 404)
	
	if err == nil {
		json.NewEncoder(w).Encode(post)
	}
}

func (a *App) UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	collection := a.Client.Database("sample-data").Collection("users")
	
	var users []models.User

	cur, err := collection.Find(context.TODO(), bson.M{})
	utils.CheckFatalError(err)
	
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		utils.CheckFatalError(err)
		users = append(users, user)
	}

	err = cur.Err()
	utils.CheckFatalError(err)

	json.NewEncoder(w).Encode(users)
}

func (a *App) UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection := a.Client.Database("sample-data").Collection("users")

	var user models.User
	var params = mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.GetError(err, w)
	}

	filter := bson.M{"_id": id}
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	utils.CheckError(err, w, 404)
	
	if err == nil {
		json.NewEncoder(w).Encode(user)
	}
}

