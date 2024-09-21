package dbconnection

import (
	"context"
	"errors"
	"snipetz/auth/models"
	snipetzerror "snipetz/commons/errors"

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

func (con Connector) UserCreationValid(user models.User) (bool, string, error) {
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
		return true, "", nil
	} else if res.Err() != nil {
		return false, "error", res.Err()
	}
	t := "unknown"
	var usr models.User
	err := res.Decode(&usr)
	if err != nil {
		return false, "error", err
	}
	if usr.Uid == user.Uid {
		return false, "uid", nil
	}
	if usr.Email == user.Email {
		return false, "email", nil
	}
	if usr.Username == user.Username {
		return false, "username", nil
	}
	return false, t, err
}

func (con Connector) CreateUser(user models.User) error {
	col := con.db.Collection("users")
	res, err := col.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	lg.Debug.Println(res.InsertedID)
	return nil
}

func (con Connector) GetUser(username_or_email string) error {
	return snipetzerror.ErrorNotImplemented
}

var Cntr Connector = NewConnector()
