package controlers

import (
	env "docs/app/Env"
	"docs/app/emptyfieldcheker"
	"docs/app/hallpers"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	"docs/app/structs"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func random(min, max int) int {
	if min > max {
		return min // Return min if it’s the only valid output
	}
	return rand.Intn(max-min+1) + min // +1 to include max in the range
}

func CheckEmail(c *gin.Context) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	GetEmail := c.Request.URL.Query().Get("email")
	// var EmailShTempmt structs.UserStruct

	client, ctx := mongoconnect.DBConnection()
	DBConnect := client.Database("OpenHearts").Collection("Users")

	result := DBConnect.FindOne(ctx, bson.M{
		"email": GetEmail,
	})
	var DBdat structs.UserStruct
	result.Decode(&DBdat)
	fmt.Printf("DBdat: %v\n", DBdat)

	if DBdat.Email != "" {
		secretCode := random(100000, 900000)
		fmt.Printf("secretCode: %v\n", secretCode)

		data := structs.Verify{
			Id:      string(primitive.NewObjectID().Hex()),
			Code:    secretCode ,
			Email:   DBdat.Email,
			User_Id: DBdat.Id,
		}
		DBConnect2 := client.Database("OpenHearts").Collection("Code")

		Insert, err := DBConnect2.InsertOne(ctx, bson.M{
			"_id":     string(primitive.NewObjectID().Hex()),
			"code":    secretCode,
			"email":   DBdat.Email,
			"user_Id": DBdat.Id,
		})
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			endPath := hallpers.CreateHTML(data, "./Email/mail.html", "./Email/codetwo.html")
			hallpers.SendEmail(DBdat.Email, "Verify Code", endPath)
			c.JSON(200, "succes")
			fmt.Printf("Insert: %v\n", Insert)
		}
	} else {
		c.JSON(404, "error email not found")
	}

}

func CheckSecretCode(c *gin.Context) {
	var Update_Pass structs.UpdatePassword
	c.ShouldBindJSON(&Update_Pass)
	Emptyfield, err := emptyfieldcheker.EmptyField(Update_Pass,"Email")
	if Emptyfield {
		c.JSON(404, err)
	} else {
		client, ctx := mongoconnect.DBConnection()
		DBConnect := client.Database(env.Data_Name).Collection("Code")
		result := DBConnect.FindOne(ctx, bson.M{
			"code":  Update_Pass.Code,
		})
		var DBdat structs.UpdatePassword
		result.Decode(&DBdat)
		fmt.Printf("DBdat: %v\n", DBdat)

		if DBdat.Code != 0 {
			digits := "0123456789"
			specials := "~%//!@#$?|"
			all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
				"abcdefghijklmnopqrstuvwxyz" +
				digits + specials
			length := 9
			buf := make([]byte, length)
			buf[0] = digits[rand.Intn(len(digits))]
			buf[1] = specials[rand.Intn(len(specials))]
			for i := 2; i < length; i++ {
				buf[i] = all[rand.Intn(len(all))]
			}
			rand.Shuffle(len(buf), func(i, j int) {
				buf[i], buf[j] = buf[j], buf[i]
			})
			str := string(buf) // Например "3i[g0|)z"

			data := structs.CheckPassword{
				Id:       string(primitive.NewObjectID().Hex()),
				Email:    DBdat.Email,
				Password: str,
			}
			fmt.Printf("str: %v\n", str)
			endPath := hallpers.CreateHTML(data, "./Email/resetpass.html", "./Email/reserTwo.html")
			hallpers.SendEmail(DBdat.Email, "Verify Code", endPath)
			

			hashpass, _ := hashedpasswod.HashPassword(str)
			DBConnect := client.Database("OpenHearts").Collection("Users")

			result := DBConnect.FindOne(ctx, bson.M{
				"code": Update_Pass.Code,
			})
			var UpdatePass structs.UserStruct
			result.Decode(&UpdatePass)
			fmt.Printf("UpdatePass: %v\n", UpdatePass)

			if DBdat.Email != "" {
				_, err := DBConnect.UpdateOne(ctx,
					bson.M{
						"email": DBdat.Email,
					},
					bson.D{
						{Key: "$set", Value: bson.M{
							"password": hashpass,
						},
						},
					},
				)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				c.JSON(200, "succes")
			} else {
				c.JSON(404, "error email not found")
			}
		}else {
      c.JSON(404,"error email not found22")
    }

	}

}
