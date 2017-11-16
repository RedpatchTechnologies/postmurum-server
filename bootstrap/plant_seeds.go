package main

import (
	"encoding/json"
	"fmt"
	"github.com/RedPatchTechnologies/postmurum-server/models"
	//"github.com/caarlos0/env"
	"github.com/markbates/pop"

	//"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"os"
	//"time"
)

/*
type Organization struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Domain    string    `json:"domain" db:"domain"`
}*/

func main() {

	db, dberr := pop.Connect("development")
	if dberr != nil {
		log.Panic(dberr)
	}

	fmt.Printf("db is %+v\n", db)

	file, e := ioutil.ReadFile("./seeds/organization.seed")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//var jsontype jsonobject
	//json.Unmarshal(file, []models.Organization{})
	//fmt.Printf("Results: %v\n", jsontype)

	orgsFromJson := []models.Organization{}

	//keys := make(users, 0)
	json.Unmarshal(file, &orgsFromJson)
	fmt.Printf("%#v", orgsFromJson)

	for _, element := range orgsFromJson {
		models.DB.Transaction(func(tx *pop.Connection) error {
			org := &models.Organization{Name: element.Name}
			err := tx.Create(org)
			if err != nil {
				panic(err)
			}
			return err
		})
	}

	orgQuery := models.DB
	orgsfromDB := []models.Organization{}
	err := orgQuery.All(&orgsfromDB)
	fmt.Printf("all orgs is %+v\n", orgsfromDB)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n Numbrer of all orgs is %+v\n", len(orgsfromDB))

	/*
		query := models.DB
		users := []models.Organization{}
		err := query.All(&users)
		fmt.Printf("users is %+v\n", users)

		// seed a few records
		models.DB.Transaction(func(tx *pop.Connection) error {
			person := &models.Person{Name: "Mark"}
			err := tx.Create(person)
			if err != nil {
				return errors.WithStack(err)
			}
			pet := &models.Pet{Name: "Ringo", PersonID: person.ID}
			return tx.Create(pet)
		})
	*/

}
