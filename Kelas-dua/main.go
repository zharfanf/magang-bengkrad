// +build linux,amd64,go1.16.5,!cgo

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aeramu/mongolib"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	//"go.mongodb.org/mongo-driver/"
)

var db *mongolib.Collection

func main(){
	client, err := mongolib.NewSingletonClient(context.Background(), "mongodb+srv://admin:admin@qiup-wrbox.mongodb.net")
	if err != nil {
		log.Fatalln("[Init DB Client]", err)
	}
	db = mongolib.NewDatabase(client, "bengkrad").Coll("Faishal")

	r := mux.NewRouter()
	r.HandleFunc("/", Root)
	r.HandleFunc("/articles", createArticle).Methods("POST")
	r.HandleFunc("/users", Users)
	r.HandleFunc("/articles/{id}", getArticleByID).Methods("GET")
	r.HandleFunc("/articles", getArticle).Methods("GET")
	http.Handle("/", r)
	log.Println("Server started at :8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
	//r.Host("www.example.com")
	//r.Host("{subdomain:[a-z]+}.example.com")
	//r.PathPrefix("/products/")
	//sendToLog("log ini")
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ini root"))

}

//var latestArticle Article

func sendToLog(message string){
	n := Message{Message: message}
	b,_ := json.Marshal(n)
	req,_ := http.NewRequest(http.MethodPost, "http://localhost:9000/log", bytes.NewReader(b))
	http.DefaultClient.Do(req)
	//log.Println(message)
}

type Message struct{
	Message string
}

func getArticleByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	//var articles []Article
	//var hehe1 Article
	var hehe Article
	tes,_ :=  primitive.ObjectIDFromHex(vars["id"])
	db.Query().Equal("id",tes).FindOne(context.Background()).Consume(&hehe)
	//db.Query().Find(context.Background()).Consume(&articles)
	//for _,tes := range articles{
	//	string_objectID := tes.ID.Hex()
	//	if vars["id"] == string_objectID{
	//		hehe1 = tes
	//	}
	//}
	//hehe2,_ := json.Marshal(hehe1)
	//w.Write(hehe2)
	hehe3,_ := json.Marshal(hehe)
	w.Write(hehe3)
	sendToLog("get article by id")
}

func getArticle(w http.ResponseWriter, r *http.Request){
	//latestArticle = Article{
	//	Title: "Hello World guys",
	//	Body: "HUHU",
	//	Author: "Gig",
	//}
	//vars := mux.Vars(r)
	var articles []Article
	db.Query().Find(context.Background()).Consume(&articles)
	hehe,_ := json.Marshal(articles)
	w.Write(hehe)
	sendToLog("get article list")
}

func createArticle(w http.ResponseWriter, r *http.Request){
	var article Article
	v,_ := ioutil.ReadAll(r.Body)
	json.Unmarshal(v, &article)
	article.ID = mongolib.NewObjectID()
	//log.Println(articles.Title)
	db.Save(context.Background(), article.ID, article)
	sendToLog("insert article")
}

func Products(w http.ResponseWriter, r *http.Request) {
	product := Product{
		ID : "1234",
		Name : "Magnum",
		cat : "Ice Cream",
		Price : 10000,
	}
	jsonString, err := json.Marshal(product)
	if(err != nil){
		log.Println(err)
		return
	}
	w.Write(jsonString)
}

func Users(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID : "id1234",
		Name : "Faishal",
		Address : "Bandung",
	}
	jsonString, err := json.Marshal(user)
	if(err != nil){
		log.Println(err)
		return
	}
	w.Write(jsonString)

}

type User struct{
	ID string
	Name string
	Address string

}

type Product struct{
	ID string
	Name string
	cat string
	Price int
}

type Article struct{
	Title string
	Body string
	Author string `json: "author"`
	ID primitive.ObjectID `bson:”id”`
}


