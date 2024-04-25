package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var PORT int=6969 // Specify port here

var data []map[string]string

func updateDB(){
    new_data,_:=json.Marshal(data)
    fo,_:=os.Create("db.json")
    fo.WriteString(string(new_data))
    fo.Close()
}

func clearDB(c *gin.Context){
    data=data[:0]
    updateDB()
    c.JSON(http.StatusOK,gin.H{"message":"database cleared"})
}

func remove(arr []map[string]string, i int) []map[string]string {
    arr[i] = arr[len(arr)-1]
    return arr[:len(arr)-1]
}

func createNewRecord(c *gin.Context) {
    var new_entry map[string]string
    c.BindJSON(&new_entry)
    new_entry["_id"]="user"+uuid.NewString()
    data=append(data, new_entry)
    updateDB()
    c.JSON(http.StatusCreated,gin.H{"_id":new_entry["_id"]})
}

func getRecordById(c *gin.Context) {
    id:=c.Param("id")
    for _,entry:=range data{
        if(entry["_id"]==id){ 
            c.JSON(http.StatusOK,entry)
            return
        }
    }
    c.JSON(http.StatusNotFound,gin.H{"result":"Not Found"})
}

func updateRecordById(c *gin.Context){
    id:=c.Param("id")
    var update_values map[string]string
    c.BindJSON(&update_values)
    for i,entry:=range data{
        if(entry["_id"]==id){
            for k,v:=range update_values{
                data[i][k]=v
            }
            updateDB()
            c.JSON(http.StatusOK,gin.H{"message":"updated record successfully"})
            return
        }
    }
    c.JSON(http.StatusNotFound,gin.H{"message":"Not Found"})
}

func deleteRecordById(c *gin.Context){
    id:=c.Param("id")
    for i,entry:=range data{
        if(entry["_id"]==id){
            data=remove(data,i)
            updateDB()
            c.JSON(http.StatusOK,gin.H{"message":"deleted successfully"})
            return
        }
    }
    c.JSON(http.StatusNotFound,gin.H{"result":"Not Found"})
}

func filterRecords(c *gin.Context) {
    var filter map[string]string
    c.BindJSON(&filter)
    filtered_data:=[]map[string]string{}
    for _,entry:=range data{
        is_valid:=true
        for k,v:=range filter{
            if _,exists:=entry[k];exists{
                if(entry[k]!=v){
                    is_valid=false
                    break
                }
            }else{
                is_valid=false
                break
            }
        }
        if is_valid{
            filtered_data=append(filtered_data, entry)
        }
    }
    c.JSON(http.StatusOK,filtered_data)
}

func main() {
    tableBytes,_:=os.ReadFile("db.json")
    json.Unmarshal(tableBytes,&data)

    router:=gin.Default()
    router.GET("/api/record/:id",getRecordById)
    router.POST("/api/new",createNewRecord)
    router.POST("/api/filter",filterRecords)
    router.PATCH("/api/patch/:id",updateRecordById)
    router.DELETE("/api/delete/:id",deleteRecordById)
    router.DELETE("/api/clear",clearDB)

    router.Run("localhost:"+fmt.Sprint(PORT))
}
