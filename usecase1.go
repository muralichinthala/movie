package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
	"log"
)


type Movies struct {
    XMLName xml.Name `xml:"movies"`
    Movies   []Movie   `xml:"movie"`
}

type Movie struct {
    XMLName xml.Name `xml:"movie"`
    Title   string   `xml:"title,attr"`
    Genre   string   `xml:"genre,attr"`
    ReleaseDate  string   `xml:"releaseDate,attr"`
}



func main() {

    // Open our xmlFile
    xmlFile, err := os.Open("movies.xml")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened movies.xml")
    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(xmlFile)

	//fmt.Println(string(byteValue))

    movies := Movies{}

    xml.Unmarshal(byteValue, &movies)
    if err != nil {
		log.Fatal(err)
	}

    for i := 0; i < len(movies.Movies); i++ {
        fmt.Println("Moive Title: " + movies.Movies[i].Title)
        fmt.Println("Movie Genre: " + movies.Movies[i].Genre)
        fmt.Println("Movie Release Date : " + movies.Movies[i].ReleaseDate)
    }

    listOfMovies := make(map[string]interface{})

	for i , _ := range movies.Movies {
		listOfMovies[movies.Movies[i].Genre] = movies.Movies[i].Title
	}

	fmt.Println(listOfMovies)

	movieLits_genre := getMovieListBasedOnGenre("drama" , listOfMovies)

	fmt.Printf("movie list based on genre : %v ", movieLits_genre)
}


func getMovieListBasedOnGenre(s string , data map[string]interface{}) interface{} {

    return data[s]
}
