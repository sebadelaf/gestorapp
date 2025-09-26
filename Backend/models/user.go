package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	//el id es un ObjectID que es un tipo especial de MongoDB y si este es vacio entonces mongo lo genera
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

func InsertUser(user User) error {
	collection := mongoClient.Database(DB_NAME).Collection(COLLECTION_NAME) //obtener la coleccion a trabajar
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func InsertManyUsers(users []User) error {
	// convertir la lista en una lista de interfaces vacias para InsertMany porque acepta cualquier tipo
	newUsers := make([]interface{}, len(users))
	for i, user := range users {
		newUsers[i] = user
	}
	collection := mongoClient.Database(DB_NAME).Collection(COLLECTION_NAME) //obtener la coleccion a trabajar
	_, err := collection.InsertMany(context.TODO(), newUsers)
	return err
}

func UpdateUser(userid string, user User) error {
	iduser, err := primitive.ObjectIDFromHex(userid) //convertir el id de string a ObjectID
	if err != nil {                                  //recordar que err siempre sera no nulo si es que hay un error
		return err
	}
	filter := bson.M{"_id": iduser}                                                                     //filtro para encontrar el usuario por su id
	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email, "password": user.Password}} //nuevos datos a actualizar

	collection := mongoClient.Database(DB_NAME).Collection(COLLECTION_NAME) //obtener la coleccion a trabajar
	_, err = collection.UpdateOne(context.TODO(), filter, update)           //actualizar el usuario con el filtro y los nuevos datos
	return err
}

func DeleteUser(userid string) error {
	iduser, err := primitive.ObjectIDFromHex(userid) //convertir el id de string a ObjectID
	if err != nil {                                  //recordar que err siempre sera no nulo si es que hay un error
		return err
	}
	filter := bson.M{"_id": iduser} //filtro para encontrar el usuario por su id
	collection := mongoClient.Database(DB_NAME).Collection(COLLECTION_NAME)
	_, err = collection.DeleteOne(context.TODO(), filter) //eliminar el usuario con el filtro
	return err
}

func GetAllUsers() ([]User, error) {
	collection := mongoClient.Database(DB_NAME).Collection(COLLECTION_NAME)

	var users []User
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(userid string) (User, error) {
	iduser, err := primitive.ObjectIDFromHex(userid) //convertir el id de string a ObjectID
	if err != nil {                                  //recordar que err siempre sera no nulo si es que hay un error
		return User{}, err
	}
	filter := bson.M{"_id": iduser} //filtro para encontrar el usuario por su id
	collection := mongoClient.Database(DB_NAME).Collection(COLLECTION_NAME)
	var user User
	err = collection.FindOne(context.TODO(), filter).Decode(&user) //buscar el usuario con el filtro y decodificarlo en la variable user
	return user, err
}
