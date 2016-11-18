package main

import (
	"text/template"
	"net/http"
	"fmt"
	"strconv"
	"math/rand"
)

//Render gohtml
func renderTemplate(w http.ResponseWriter, tmpl string, node *map[string][]string) {
	t, _ := template.ParseFiles(tmpl + ".gohtml")
	t.Execute(w,node)
}

//Generate Random Perpetrators for initial state
func generateRandomPerpetrators(gameState *map[string][]string) {
	i := 0
	for i < 3 {
		value := rand.Intn(7)
		if (*gameState)["perpetrators"][value] == "0" {
			(*gameState)["perpetrators"][value] = "1"
			i++
		}
	}
}

//Generate Random Criminals
func generateRandomCriminals(gameState *map[string][]string) {
	i := 0
	for i < 7 {
		(*gameState)["randomCriminals"][i] = "0"
		i++
	}
	i = 0
	for i < 3 {
		value := rand.Intn(7)
		if (*gameState)["randomCriminals"][value] == "0" {
			(*gameState)["randomCriminals"][value] = "1"
			i++
		}
	}
}

//Return number of criminals matched with perpetrators
func getMatchingPerpetratorCriminals(gameState *map[string][]string) int {
	i := 0
	match := 0
	for i < 7 {
		if (*gameState)["randomCriminals"][i] == "1" {
			if (*gameState)["perpetrators"][i] == "1" {
				match++
			}
		}
		i++
	}
	return match
}


func gameHandler(gameState *map[string][]string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			//Get values for the form and check for match with perpetrators
			value1 := r.Form["value1"][0]
			value2 := r.Form["value2"][0]
			value3 := r.Form["value3"][0]
			intval1, _ := strconv.Atoi(value1)
			intval2, _ := strconv.Atoi(value2)
			intval3, _ := strconv.Atoi(value3)
			(*gameState)["winStatus"] = []string{"No",}
			fmt.Println((*gameState)["perpetrators"][intval1], intval1)
			if (*gameState)["perpetrators"][intval1] == "1" {
				if (*gameState)["perpetrators"][intval2] == "1" {
					if (*gameState)["perpetrators"][intval3] == "1" {
						//User has guessed it correct
						*gameState = getNewGameState()
						generateRandomPerpetrators(gameState)
						(*gameState)["winStatus"] = []string{"Yes",}
					}
				}
			}
		}
		value := 3
		for value > 2 {
			//Generate Random criminals
			generateRandomCriminals(gameState)
			//Make sure you dont get 3 perpatators matching to random criminals
			value = getMatchingPerpetratorCriminals(gameState)
			if value > 2 {
				continue
			}
			(*gameState)["totalPerpetrators"][0] = strconv.Itoa(value)
		}
		renderTemplate(w, "game", gameState)
	}
}

func passHandler(gameState *map[string][]string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		//For default we say user has not won and then we check using gameHandler
		(*gameState)["winStatus"] = []string{"No",}
		http.Redirect(w,r,"/game/",http.StatusSeeOther)
	}
}

func getNewGameState() map[string][]string {
	//gameState is a dictionary of strings
	gameState := map[string][]string{
		"criminals" : {"a","b","c","d","e","f","g"},
		"perpetrators" : {"0","0","0","0","0","0","0"},
		"randomCriminals" : {"0","0","0","0","0","0","0"},
		"totalPerpetrators" : {"0",},
		"winStatus" : {"No"},
	}
	return gameState
}
func main() {
	gameState := getNewGameState()
	//Seed is kept constant
	rand.Seed(1)
	//Generate Perpertrators
	generateRandomPerpetrators(&gameState)
	//Generate Random criminals 
	generateRandomCriminals(&gameState)

	//Print current game state 
	fmt.Printf("%+v\n", gameState)
	//using link 127.0.0.1:8080/game/ to run the game
	http.HandleFunc("/game/",gameHandler(&gameState))
	///pass/ serves for getting a post request
	http.HandleFunc("/pass/",passHandler(&gameState))
	//Start listening
	http.ListenAndServe(":8080",nil)
}


