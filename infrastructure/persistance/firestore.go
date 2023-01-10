package persistance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
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

	doc, err := Client.Doc(path).Get(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data = doc.Data()
	data["id"] = doc.Ref.ID

	return data, err
}

// function to delete a firestore document in a given path
func Delete(path string, id string) error {
	docRef := Client.Collection(path).Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// function to create a doc in firestore in a given path
func Create(path string, data interface{}, id string) error {
	docRef := Client.Collection(path).Doc(id)
	_, err := docRef.Set(ctx, data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// functions to update a existing doc with its id and path
func Update(path string, data interface{}, id string) error {
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
	value, ok := dataMap["warnings_record"]

	docRef := Client.Collection(path).Doc(id)

	if ok {
		warningsArr, ok := value.([]interface{})
		if ok {
			for _, item := range(warningsArr) {
				_, err = docRef.Update(ctx, []firestore.Update{
					{
						FieldPath: []string{"warnings_record"},
						Value:     firestore.ArrayUnion(item),
					},
				})
				if err != nil {
					return err
				}
			}
		}
		delete(dataMap, "warnings_record")
	}

	_, err = docRef.Set(ctx, dataMap, firestore.MergeAll)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// function to list all documents of a Collection
func List(path string) ([]map[string]interface{}, error) {
	ctx := context.Background()
	var docArr []map[string]interface{}

	collRef := Client.Collection(path)

	iter := collRef.Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		docPath := doc.Ref.Path
		parts := strings.Split(docPath, "/")

		// Get the portion of the path after the "servers" collection
		subpath := strings.Join(parts[5:], "/")

		docMap, err := Get(subpath)
		docArr = append(docArr, docMap)
	}

	return docArr, nil

}
