package controlers

import (
	// "docs/app/baner"
	env "docs/app/Env"
	"docs/app/emptyfieldcheker"
	// "docs/app/hashedpasswod"

	"docs/app/mongoconnect"
	"docs/app/returnJwt"
	"docs/app/structs"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func UpdateStatistic(c *gin.Context) {
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
			var Update_statistic structs.AddStatistic
			ids := c.Request.URL.Query().Get("id")
			fmt.Printf("ids: %v\n", ids)
			c.ShouldBindJSON(&Update_statistic)
			fmt.Printf("Update_statistic: %v\n", Update_statistic)
			Emptyfield, err := emptyfieldcheker.EmptyField(Update_statistic, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()
				Connections := client.Database(env.Data_Name).Collection("statistic")
				_, err := Connections.UpdateOne(ctx,
					bson.M{
						"_id": ids,
					},
					bson.D{
						{Key:"$set", Value:bson.M {
							"Ru": Update_statistic.Ru,
							"En": Update_statistic.En,
						},
					},
				},	
				)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				c.JSON(200, "succes")
			}
		}
	}
}





func UpdateServiceNumber(c *gin.Context) {
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
			var NewQuontity structs.ChangNumber
			ids := c.Request.URL.Query().Get("id")
			fmt.Printf("ids: %v\n", ids)
			c.ShouldBindJSON(&NewQuontity)
			Emptyfield, err := emptyfieldcheker.EmptyField(NewQuontity, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()
				Connections := client.Database(env.Data_Name).Collection("change_number_for_center")
				_, err := Connections.UpdateOne(ctx,
					bson.M{
						"_id": ids,
					},
					bson.D{
						{Key:"$set", Value:bson.M {
							"quantity":NewQuontity.Quantity,
						},
					},
				},	
				)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				c.JSON(200, "succes")
			}
		}
	}
}




func UpdateStatisticForCenter(c *gin.Context) {
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
			var Update_statistic_Center structs.AddStatisticForCenter
			ids := c.Request.URL.Query().Get("id")
			fmt.Printf("ids: %v\n", ids)
			c.ShouldBindJSON(&Update_statistic_Center)
			Emptyfield, err := emptyfieldcheker.EmptyField(Update_statistic_Center, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()
				Connections := client.Database(env.Data_Name).Collection("statistic_for_center")
				_, err := Connections.UpdateOne(ctx,
					bson.M{
						"_id": ids,
					},
					bson.D{
						{Key:"$set", Value:bson.M {
							"ru": structs.LangForProjectStatistic{
								Description: Update_statistic_Center.Ru.Description,
								Name: Update_statistic_Center.Ru.Name,
							},
							"en": structs.LangForProjectStatistic{
								Description: Update_statistic_Center.En.Description,
								Name: Update_statistic_Center.En.Name,
							},
						},
					},
				},	
				)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				c.JSON(200, "succes")
			}
		}
	}
}




