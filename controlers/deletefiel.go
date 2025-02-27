package controlers

import (
	"docs/app/Env"
	"docs/app/mongoconnect"
	"docs/app/returnJwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)



func DeletePatientStory(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {
			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")

			if ids == "" && Path == "" {
				c.JSON(404, "Empty Field")
			} else {

				client, ctx := mongoconnect.DBConnection()
				var createDB = client.Database(env.Data_Name).Collection("PatientStory")
				deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
					"_id": ids,
				})
				if deleteerror != nil {
					fmt.Printf("deleteerror: %v\n", deleteerror)
				}
				if deletrezult.DeletedCount == 1 {
					err := os.RemoveAll("./Statics/" + Path)
					if err != nil {
						fmt.Printf("err: %v\n", err)
					}
					c.JSON(200, "succes")
					fmt.Printf("deletrezult: %v\n", deletrezult)
				} else {
					c.JSON(404, "error")
				}
			}
		}
	}
}





func DeletePatners(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error")
		} else {
			ids := c.Request.URL.Query().Get("id")
			Path := c.Request.URL.Query().Get("Path")

			client, ctx := mongoconnect.DBConnection()
			var createDB = client.Database(env.Data_Name).Collection("Partners")
			deletrezult, deleteerror := createDB.DeleteOne(ctx, bson.M{
				"_id": ids,
			})
			if deleteerror != nil {
				fmt.Printf("deleteerror: %v\n", deleteerror)
			}
			if deletrezult.DeletedCount == 1 {
				err := os.RemoveAll("./Statics/" + Path)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				} else {
					c.JSON(200, "succes")
					fmt.Printf("deletrezult: %v\n", deletrezult)
				}
			} else {
				c.JSON(404, "error1")
			}
		}
	}
}














