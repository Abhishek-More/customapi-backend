package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Abhishek-More/customapi-backend/pkg/data/models"
	"github.com/Abhishek-More/customapi-backend/pkg/utils"
)

func (a *App) InitializeRouter() {
	a.Router = mux.NewRouter()
	
	//TODO: Show API Info
	a.Router.HandleFunc("/", utils.RouteNotImplemented)

	//TODO: Show Posts
	a.Router.HandleFunc("/sample/posts", a.PostsHandler)

	//TODO: Show Users
	a.Router.HandleFunc("/sample/users", utils.RouteNotImplemented)

	http.ListenAndServe(":8000", a.Router)
}

func (a *App) PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	collection := a.Client.Database("sample-data").Collection("posts")
	
	var posts []models.Post

	cur, err := collection.Find(context.TODO(), bson.M{})
	utils.CheckError(err)
	
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var post models.Post
		err := cur.Decode(&post)
		utils.CheckError(err)
		
		posts = append(posts, post)
	}

	err = cur.Err()
	utils.CheckError(err)

	json.NewEncoder(w).Encode(posts)
}

