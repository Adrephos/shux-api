package persistance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var Client *firestore.Client
var ctx = context.Background()
var app *firebase.App

// fuction to intilize the firestore client
func init() {
	var err error
	sapath, _ := os.LookupEnv("SHUX_API_SA")

	sa := option.WithCredentialsFile(sapath)
	app, err = firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	Client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

// function that returns a map with the data of a firestore doc
func Get(path string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var err error

	doc, err := Client.Doc(path).Get(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data = doc.Data()
	id := "id"
	data[id] = doc.Ref.ID

	return data, err
}

// function to delete a firestore document in a given path
func Delete(path string, id string) error {
	docRef := Client.Collection(path).Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil{
		fmt.Println(err)
		return err
	}

	return nil
}

// function to create a doc in firestore in a given path
func Create(path string, data interface{}, id string) error{
	docRef := Client.Collection(path).Doc(id)
	_, err := docRef.Set(ctx, data)
    if err != nil {
		fmt.Println(err)
        return err
    }

    return nil
}

// functions to update a existing doc with its id and path
func Update(path string, data interface{}, id string) error{
	dataJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
        return err
	}

	var dataMap map[string]interface{}
    err = json.Unmarshal(dataJson, &dataMap)
    if err != nil {
        return err
    }

	docRef := Client.Collection(path).Doc(id)
	_, err = docRef.Set(ctx, dataMap, firestore.MergeAll)
    if err != nil {
		fmt.Println(err)
        return err
    }

    return nil
}
