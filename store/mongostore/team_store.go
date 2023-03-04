package mongostore

import (
	"context"
	"errors"
	"time"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamStore struct {
	entityStore[models.Team]
}

func NewTeamStore(collection *mongo.Collection) *TeamStore {
	return &TeamStore{
		entityStore: entityStore[models.Team]{
			collection: collection,
		},
	}
}

func (s *TeamStore) RegisterIndexes() error {
	unique := true
	_, err := s.collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.M{"name": 1},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})
	return err
}

func (s *TeamStore) Create(t *models.Team) (*models.Team, error) {
	t.ID = utils.NewID()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return s.entityStore.Create(t)
}

func (s *TeamStore) Update(t *models.Team) (*models.Team, error) {
	updatedAttrs := &models.Team{}
	err := utils.UnmarshalFrom(t, updatedAttrs)
	if err != nil {
		return nil, err
	}

	updatedAttrs.Members = nil
	updatedAttrs.Admins = nil
	updatedAttrs.UpdatedAt = time.Now()
	t.UpdatedAt = updatedAttrs.UpdatedAt

	return s.entityStore.Update(updatedAttrs)
}

func (s *TeamStore) GetPublicNonJoined(userID string) ([]*models.Team, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"adminIds":  bson.M{"$ne": userID},
				"memberIds": bson.M{"$ne": userID},
				"isPublic":  true,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "memberIds",
				"foreignField": "_id",
				"as":           "members",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "adminIds",
				"foreignField": "_id",
				"as":           "admins",
			},
		},
		{
			"$project": bson.M{
				"members.createdAt": 0,
				"members.updatedAt": 0,
				"admins.createdAt":  0,
				"admins.updatedAt":  0,
				"members.email":     0,
				"admins.email":      0,
			},
		},
	}

	c, err := s.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	teams := []*models.Team{}

	err = c.All(context.TODO(), &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *TeamStore) GetByIDWithUsers(id string) (*models.Team, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id": id,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "memberIds",
				"foreignField": "_id",
				"as":           "members",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "adminIds",
				"foreignField": "_id",
				"as":           "admins",
			},
		},
		{
			"$project": bson.M{
				"members.createdAt": 0,
				"members.updatedAt": 0,
				"admins.createdAt":  0,
				"admins.updatedAt":  0,
			},
		},
	}

	c, err := s.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	teams := []*models.Team{}

	err = c.All(context.TODO(), &teams)
	if err != nil {
		return nil, err
	}

	if len(teams) != 1 {
		return nil, errors.New(models.ResourceDoesNotExist)
	}

	return teams[0], nil
}

func (s *TeamStore) GetByUserID(userID string, includeUsers bool) ([]*models.Team, error) {
	teams := []*models.Team{}

	matchQuery := bson.M{
		"$or": []bson.M{
			{
				"adminIds": userID,
			},
			{
				"memberIds": userID,
			},
		},
	}

	var c *mongo.Cursor
	var err error

	if includeUsers {
		pipeline := []bson.M{
			{"$match": matchQuery},
			{
				"$lookup": bson.M{
					"from":         "users",
					"localField":   "memberIds",
					"foreignField": "_id",
					"as":           "members",
				},
			},
			{
				"$lookup": bson.M{
					"from":         "users",
					"localField":   "adminIds",
					"foreignField": "_id",
					"as":           "admins",
				},
			},
			{
				"$project": bson.M{
					"members.createdAt": 0,
					"members.updatedAt": 0,
					"admins.createdAt":  0,
					"admins.updatedAt":  0,
					"members.email":     0,
					"admins.email":      0,
					"createdAt":         0,
					"updatedAt":         0,
				},
			},
		}

		c, err = s.collection.Aggregate(context.TODO(), pipeline)
		if err != nil {
			return nil, err
		}

	} else {
		c, err = s.collection.Find(context.TODO(), matchQuery, &options.FindOptions{Projection: bson.M{"adminIds": 0, "users": 0, "createdAt": 0, "updatedAt": 0}})
		if err != nil {
			return nil, err
		}
	}

	err = c.All(context.TODO(), &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *TeamStore) GetCountByUserID(userID string) (int64, error) {
	matchQuery := bson.M{
		"$or": []bson.M{
			{
				"adminIds": userID,
			},
			{
				"memberIds": userID,
			},
		},
	}

	c, err := s.collection.CountDocuments(context.TODO(), matchQuery)
	if err != nil {
		return 0, err
	}

	return c, nil
}

func (s *TeamStore) AddMemberToTeam(teamID, memberID string) error {
	update := bson.M{
		"$push": bson.M{
			"memberIds": memberID,
		},
	}

	r, err := s.collection.UpdateOne(context.TODO(), bson.M{"_id": teamID}, update)
	if err != nil {
		return err
	}

	if r.ModifiedCount != 1 {
		return errors.New(models.ResourceUnmodified)
	}

	return nil
}

func (s *TeamStore) AddAdminToTeam(team *models.Team, adminID string) error {
	updateQuery := bson.M{
		"$push": bson.M{
			"adminIds": adminID,
		},
	}

	for _, m := range team.Members {
		if m.ID == adminID {
			updateQuery["$pull"] = bson.M{"memberIds": adminID}
		}
	}

	r, err := s.collection.UpdateOne(context.TODO(), bson.M{"_id": team.ID}, updateQuery)
	if err != nil {
		return err
	}

	if r.ModifiedCount != 1 {
		return errors.New(models.ResourceUnmodified)
	}

	return nil
}

func (s *TeamStore) RemoveMemberFromTeam(teamID, memberID string) error {
	update := bson.M{
		"$pull": bson.M{
			"memberIds": memberID,
		},
	}

	r, err := s.collection.UpdateOne(context.TODO(), bson.M{"_id": teamID}, update)
	if err != nil {
		return err
	}

	if r.ModifiedCount != 1 {
		return errors.New(models.ResourceUnmodified)
	}

	return nil
}

func (s *TeamStore) RemoveAdminFromTeam(team *models.Team, adminID string) error {
	updateQuery := bson.M{
		"$pull": bson.M{
			"adminIds": adminID,
		},
	}

	var ops *options.UpdateOptions

	if len(team.Admins) == 1 {
		updateQuery = bson.M{
			"$set": bson.M{
				"admins.$[id]": team.Members[0].ID,
			},
		}
		ops = &options.UpdateOptions{ArrayFilters: &options.ArrayFilters{
			Filters: []interface{}{bson.M{"id": adminID}},
		}}
		updateQuery["$pull"] = bson.M{"memberIds": team.Members[0].ID}
	}

	r, err := s.collection.UpdateOne(context.TODO(), bson.M{"_id": team.ID}, updateQuery, ops)
	if err != nil {
		return err
	}

	if r.ModifiedCount != 1 {
		return errors.New(models.ResourceUnmodified)
	}

	return nil
}

func (s *TeamStore) ToggleRole(team *models.Team, userID string, isAdmin bool) error {
	updateQuery := bson.M{
		"$pull": bson.M{},
		"$push": bson.M{},
	}

	if isAdmin {
		updateQuery["$pull"].(bson.M)["adminIds"] = userID
		updateQuery["$push"].(bson.M)["memberIds"] = userID
	} else {
		updateQuery["$pull"].(bson.M)["memberIds"] = userID
		updateQuery["$push"].(bson.M)["adminIds"] = userID
	}

	r, err := s.collection.UpdateOne(context.TODO(), bson.M{"_id": team.ID}, updateQuery)
	if err != nil {
		return err
	}

	if r.ModifiedCount != 1 {
		return errors.New(models.ResourceUnmodified)
	}

	return nil
}
