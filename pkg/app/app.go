package app

import (
	"github.com/Abhishek-More/customapi-backend/pkg/utils"
	"github.com/Abhishek-More/customapi-backend/pkg/data"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-chi/chi/v5"

)

type App struct {
	Client *mongo.Client
	Router      *chi.Mux
}

func (a* App) InitializeApp() {
	utils.LoadEnvironmentVariables()

	a.Client = db.Connect()

}