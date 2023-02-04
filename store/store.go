package store

import (
	"context"
	"errors"

	interfaces "github.com/ibrahimfarhan/voting-app/voting-app-server/store/interfaces"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/store/mongostore"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
)

type Store struct {
	connector interfaces.DBConnector
	interfaces.UserStore
	interfaces.TeamStore
	interfaces.TokenStore
}

func InitStore(ctx context.Context, dbType string) Store {
	var connector interfaces.DBConnector
	var UserStore interfaces.UserStore
	var TeamStore interfaces.TeamStore
	var TokenStore interfaces.TokenStore

	switch dbType {
	case "mongo":
		connector = mongostore.NewMongoConnector()
		c := connector.(*mongostore.MongoConnector)
		c.DB.CreateCollection(ctx, "users")
		c.DB.CreateCollection(ctx, "teams")
		c.DB.CreateCollection(ctx, "tokens")
		UserStore = mongostore.NewUserStore(c.DB.Collection("users"))
		TeamStore = mongostore.NewTeamStore(c.DB.Collection("teams"))
		TokenStore = mongostore.NewTokenStore(c.DB.Collection("tokens"))
	default:
		panic(errors.New("Unsupported db type"))
	}

	utils.PanicOnErr(connector.Open(ctx))
	utils.PanicOnErr(UserStore.RegisterIndexes())
	utils.PanicOnErr(TeamStore.RegisterIndexes())
	utils.PanicOnErr(TokenStore.RegisterIndexes())

	return Store{
		connector:  connector,
		UserStore:  UserStore,
		TeamStore:  TeamStore,
		TokenStore: TokenStore,
	}
}

func (s *Store) Close(ctx context.Context) error {
	return s.connector.Close(ctx)
}
