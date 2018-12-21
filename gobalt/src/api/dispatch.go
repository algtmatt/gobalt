package api

import (
	"encoding/json"
	"fmt"
	"gobalt/src/api/endpoint"
	"log"
)

func (api *API) Login(username string, password string, authMethod string) (string, error) {

	var token string

	log.Println("Logging into Salt")

	mappedCredentials := mapCredentials(username, password, authMethod)
	verifyCredentials(mappedCredentials)

	body, err := api.Transport.Fetch("/login", mappedCredentials)

	if err != nil {
		handleFetchErr(err)

	} else {
		objmap := unmarshallLogin(body)
		token = objmap.Return[0].Token
	}
	fmt.Printf("token: %s", token)

	if err != nil || token == "" {
		log.Fatal("Error attempting to login")
	}

	log.Printf("Current auth stored value: %s\n", api.Session.Token)
	log.Printf("Setting auth token on %s\n", api.Session.Token)

	// Set the token
	api.Session.Token = token
	fmt.Printf("Current auth stored value: %s\n", api.Session.Token)
	return "", nil
}

func unmarshallLogin(body []byte) endpoint.SaltAPIListReturn {
	var objmap = new(endpoint.SaltAPIListReturn)
	err := json.Unmarshal(body, &objmap)
	if err != nil {
		panic(err)
	}
	return *objmap
}

func handleFetchErr(err error) {
	panic(err) // TODO implement overall handling
}

func mapCredentials(username string, password string, authMethod string) map[string]string {
	return map[string]string{
		"username": username,
		"password": password,
		"eauth":    authMethod,
	}
}

func verifyCredentials(credentials map[string]string) {
	if credentials["username"] == "" || credentials["password"] == "" {
		log.Fatal("Must specify username and password in Gobalt configuration file")
	}
}

func (api *API) Keys() (endpoint.SaltAPIKeyReturn, error) {
	fmt.Println("Retreiving keys from Salt")
	ret, err := api.Transport.Fetch("/keys")
	if err != nil {
		panic(err)
	}
	print(ret)
	var objmap = new(endpoint.SaltAPIKeyReturn)
	err = json.Unmarshal(ret, &objmap)

	if err != nil {
		panic(err)
	}
	fmt.Printf(objmap.Return.Minions[0])
	return *objmap, err
}

//FIXME incomplete
func (api *API) Key(key string) (endpoint.SaltAPIKeyReturn, error) {
	ret, err := api.Transport.Fetch("/keys/" + key)
	print(ret)
	return endpoint.SaltAPIKeyReturn{}, err
}
