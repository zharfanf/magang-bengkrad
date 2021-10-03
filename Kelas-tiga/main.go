// +build linux,amd64,go1.16.5,!cgo

package main

import (
	"encoding/json"
	"github.com/aeramu/mongolib"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var db *mongolib.Collection

func main(){
	//client, err := mongolib.NewSingletonClient(context.Background(), "mongodb+srv://admin:admin@qiup-wrbox.mongodb.net")
	//if err != nil {
	//	log.Fatalln("[Init DB Client]", err)
	//}
	//db = mongolib.NewDatabase(client, "bengkrad").Coll("Faishal")

	r := mux.NewRouter()
	r.HandleFunc("/log", Log).Methods("POST")
	log.Println("Server started at :9000")
	log.Fatalln(http.ListenAndServe(":9000", r))
}


func Log(w http.ResponseWriter, r *http.Request){
	input,_ := ioutil.ReadAll(r.Body)
	var hasil Message
	json.Unmarshal(input, &hasil)
	log.Println(hasil)
}


type Message struct{
	Message string
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ini root"))

}
//
////var latestArticle Article
//
//
//func getArticleByID(w http.ResponseWriter, r *http.Request){
//	vars := mux.Vars(r)
//	//var articles []Article
//	//var hehe1 Article
//	var hehe Article
//	tes,_ :=  primitive.ObjectIDFromHex(vars["id"])
//	db.Query().Equal("id",tes).FindOne(context.Background()).Consume(&hehe)
//	//db.Query().Find(context.Background()).Consume(&articles)
//	//for _,tes := range articles{
//	//	string_objectID := tes.ID.Hex()
//	//	if vars["id"] == string_objectID{
//	//		hehe1 = tes
//	//	}
//	//}
//	//hehe2,_ := json.Marshal(hehe1)
//	//w.Write(hehe2)
//	hehe3,_ := json.Marshal(hehe)
//	w.Write(hehe3)
//}
//
//func getArticle(w http.ResponseWriter, r *http.Request){
//	//latestArticle = Article{
//	//	Title: "Hello World guys",
//	//	Body: "HUHU",
//	//	Author: "Gig",
//	//}
//	//vars := mux.Vars(r)
//	var articles []Article
//	db.Query().Find(context.Background()).Consume(&articles)
//	hehe,_ := json.Marshal(articles)
//	w.Write(hehe)
//}
//
//func createArticle(w http.ResponseWriter, r *http.Request){
//	var article Article
//	v,_ := ioutil.ReadAll(r.Body)
//	json.Unmarshal(v, &article)
//	article.ID = mongolib.NewObjectID()
//	//log.Println(articles.Title)
//	db.Save(context.Background(), article.ID, article)
//}
//
//func Products(w http.ResponseWriter, r *http.Request) {
//	product := Product{
//		ID : "1234",
//		Name : "Magnum",
//		cat : "Ice Cream",
//		Price : 10000,
//	}
//	jsonString, err := json.Marshal(product)
//	if(err != nil){
//		log.Println(err)
//		return
//	}
//	w.Write(jsonString)
//}
//
//func Users(w http.ResponseWriter, r *http.Request) {
//	user := User{
//		ID : "id1234",
//		Name : "Faishal",
//		Address : "Bandung",
//	}
//	jsonString, err := json.Marshal(user)
//	if(err != nil){
//		log.Println(err)
//		return
//	}
//	w.Write(jsonString)
//
//}
//
//type User struct{
//	ID string
//	Name string
//	Address string
//
//}
//
//type Product struct{
//	ID string
//	Name string
//	cat string
//	Price int
//}
//
//type Article struct{
//	Title string
//	Body string
//	Author string `json: "author"`
//	ID primitive.ObjectID `bson:”id”`
//}


