package dbconnection

import (
	"context"
	"errors"
	"snipetz/auth/models"

	"github.com/f7ed0/golog/lg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connector struct {
	cli *mongo.Client
	db  *mongo.Database
}

func NewConnector() Connector {
	return Connector{}
}

func (con *Connector) Init(uri string) (err error) {
	con.cli, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	con.db = con.cli.Database("auth")
	if con.db == nil {
		err = errors.New("No database")
	}

	return
}

func (con Connector) UserCreationValid(user models.User) (bool, error) {
	col := con.db.Collection("users")
	lg.Debug.Println(col)
	res := col.FindOne(
		context.TODO(),
		bson.D{
			bson.E{
				Key: "$or", Value: bson.A{
					bson.D{{Key: "username", Value: user.Username}},
					bson.D{{Key: "email", Value: user.Email}},
					bson.D{{Key: "uid", Value: user.Uid}},
				},
			},
		},
	)

	if res.Err() == mongo.ErrNoDocuments {
		return true, nil
	} else if res.Err() != nil {
		return false, res.Err()
	}
	return false, nil
}

var Cntr Connector = NewConnector()
