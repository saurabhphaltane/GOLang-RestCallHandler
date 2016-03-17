package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
	"encoding/json"

)

//Global Data Structure to store the user Profiles 
var hashmap map[string]interface{}

//POST Profile to Application

func postprofile(w http.ResponseWriter, r *http.Request) {
   
        var f interface{}
        json.NewDecoder(r.Body).Decode(&f)
        m := f.(map[string]interface{})
//Check if the user Profile already exists for the given Id:
//if not create a Hashmap for the given id

      _, present := hashmap[string(m["email"].(string))]
       
       if (!present){
            if (len(hashmap)==0){

            hashmap = make(map[string]interface{})

            }
           
//Assign the Value to Glabal Hashmap with Email Id as the key and entire Json as Value.
       	    hashmap[string(m["email"].(string))]=m
            w.WriteHeader(201)

       }else{ 
 
            fmt.Println("Profile Exists for Given MailId: ",string(m["email"].(string)))
            w.WriteHeader(200)
       
       }
   
}



//Retrive the Posted Profile with Email Id as Input key 
func getprofile(w http.ResponseWriter, r *http.Request) {
      
            params := r.URL.Query()
            emailId  := params.Get(":emailId")

//Check the Presence of Record for the corresponding email id in Global hashMap
            result, present := hashmap[string(emailId)]

       if (present){

             json.NewEncoder(w).Encode(result)
            
       }else{ 
 
       	    fmt.Println("%s %s","Profile does not Exists ",string(emailId))
            w.WriteHeader(200)
            fmt.Fprintf(w, "%s\n %s ","No Profile Found : the Email Id",string(emailId))
        }
 }


//Handle the Delete Profile Utility of the HashMap

func deleteprofile(w http.ResponseWriter, r *http.Request){

//Fetch the Email Id and check the presence of Record in the Global HashMap
        params := r.URL.Query()
        emailId  := params.Get(":emailId")

        _, present := hashmap[string(emailId)]

        if (present){

//Execte the Delete if Record Found  
        fmt.Println("this is Hashmap  : ",hashmap)

        delete(hashmap,string(emailId))  

        fmt.Println("%s %s","Profile Deleted",string(emailId))
        w.WriteHeader(204)
   
    }else {
     	  fmt.Println("%s %s","No Profile found to Delete :",string(emailId))
        w.WriteHeader(200)
        fmt.Fprintf(w, "%s\n %s","No Profile found to Delete :",string(emailId))

    }

}


func putprofile(w http.ResponseWriter, r *http.Request){


    params := r.URL.Query()
    emailId  := params.Get(":emailId")

//   Check the Presence of Record Corresponding to the given Email Id

    _, present := hashmap[string(emailId)]

    if (present){

      //create a Map temporary update the map and save the map the said id and move on
               
     //Fetch the HashMap to be updated from the Global Map

    m := hashmap[string(emailId)]
     //Receive the JSON to be Updated and conver the same to MAP Structure
    var f interface{}

    json.NewDecoder(r.Body).Decode(&f)
    recvdmap := f.(map[string]interface{})


        for recvdkey, recvdvalue := range recvdmap {
          mmap:=m.(map[string]interface{}) 

        mmap[recvdkey]=recvdvalue

        for dkey, dvalue := range mmap {
        fmt.Println("Key:", dkey, "Value:", dvalue)
        }

        hashmap[string(emailId)]=mmap
      
        w.WriteHeader(204)
    
      }



    }else{
  	//Handle the Record Not Found Exception 
        w.WriteHeader(200)
     
       //	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
       fmt.Fprintf(w, "%s\n %s","Please Check Email Id No Profile Found ",string(emailId))

    }

}

func main() {
	
    //Create A new Router to Handle the Request and Handle the Requests for corresponding API's.
	mux := routes.New()

	mux.Post("/profile", postprofile)
	mux.Get("/profile/:emailId", getprofile)
	mux.Put("/profile/:emailId",putprofile)
	mux.Del("/profile/:emailId",deleteprofile)
	http.Handle("/", mux)
  fmt.Println("Server Listening on port : 3000")
	http.ListenAndServe(":3000", nil)
  

}
