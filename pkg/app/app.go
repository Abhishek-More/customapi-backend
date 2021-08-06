package app

import (
	"github.com/Abhishek-More/customapi-backend/pkg/utils"
	"github.com/Abhishek-More/customapi-backend/pkg/data/db"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gorilla/mux"
)

type App struct {
	Client *mongo.Client
	Router *mux.Router
}

func (a* App) InitializeApp() {
	utils.LoadEnvironmentVariables()

	a.Client = db.Connect()
	a.InitializeRouter()


	defer db.Disconnect(a.Client)

}