package createadmin

import (
	"context"
	"math/rand"

	env "docs/app/Env"
	"docs/app/baner"
	"docs/app/emptyfieldcheker"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	returnjwt "docs/app/returnJwt"
	"docs/app/structs"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Createadmin() {
	client, ctx := mongoconnect.DBConnection()
	Connections := client.Database(env.Data_Name).Collection("Users")
	// if admin exist not insert
	found := Connections.FindOne(ctx, bson.M{
		"permission": "Admin",
	})
	var shablon structs.UserStruct
	found.Decode(&shablon)

	if shablon.Permission == "Admin" {
		fmt.Println("admin is exsist")
	} else {
		ID := primitive.NewObjectID().Hex()
		Password := "Admin"
		Hashed, _ := hashedpasswod.HashPassword(Password)
		Connections.InsertOne(context.Background(), structs.UserStruct{
			Id:         ID,
			Photo:      "",
			Phone:      "123",
			Password:   Hashed,
			Permission: "Admin",
			Email:      "murtazobroimshoevm4@gmail.com",
			Ru: structs.RussianUser{
				Name: "Admin",
			},
			En: structs.EnglishUser{
				Name: "Admin",
			},
		})
	}
}

func UpdateAdmin(c *gin.Context) {
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
			var Update_Admin structs.UserStruct
			c.ShouldBindJSON(&Update_Admin)
			fmt.Printf("Update_statistic: %v\n", Update_Admin)
			Emptyfield, err := emptyfieldcheker.EmptyField(Update_Admin, "Updatedemail", "Id", "Audience", "Issuer", "Subject")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				deleteeror := os.RemoveAll("./Statics/" + "AdminPhoto")
				if deleteeror != nil {
					fmt.Printf("err: %v\n", deleteeror)
				} else {

					folder_Name := "AdminPhoto"
					err := os.Mkdir("Statics"+"/"+folder_Name, os.ModePerm)
					if err != nil {
						if os.IsExist(err) {
							fmt.Println(" Папка уже существует.")
						} else {
							fmt.Println("Ошибка при создании папки")
							return
						}
					} else {
						fmt.Println("Папка успешно создана.")
					}
					rndName := rand.Intn(10000)
					ForImage := fmt.Sprintf("image_%v.png", rndName)
					Update_Admin.Photo = baner.ImageFunc(Update_Admin.Photo, ForImage, folder_Name)
					client, ctx := mongoconnect.DBConnection()
					Connections := client.Database(env.Data_Name).Collection("Users")
					Hashed, _ := hashedpasswod.HashPassword(Update_Admin.Password)
					_, err2 := Connections.UpdateOne(ctx,
						bson.M{
							"phone": Update_Admin.Phone,
						},
						bson.D{
							{Key: "$set", Value: bson.M{
								"photo":      folder_Name + "/" + Update_Admin.Photo,
								"phone":      Update_Admin.Phone,
								"password":   Hashed,
								"permission": "Admin",
								"email":      Update_Admin.Email,
								"ru": structs.RussianUser{
									Name: Update_Admin.Ru.Name,
								},
								"en": structs.EnglishUser{
									Name: Update_Admin.En.Name,
								},
							},
							},
						},
					)
					if err2 != nil {
						fmt.Printf("err: %v\n", err)
					}
					c.JSON(200, "succes")
				}
			}
		}
	}
}
