package controlers

import (
	"docs/app/Env"
	"docs/app/baner"
	"docs/app/emptyfieldcheker"
	"docs/app/hashedpasswod"
	"docs/app/mongoconnect"
	"docs/app/returnJwt"
	"docs/app/structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func Login(c *gin.Context) {
	var LoginTemp structs.UserStruct
	c.ShouldBindJSON(&LoginTemp)
	EmptyField, err := emptyfieldcheker.EmptyField(LoginTemp,"Photo", "Name", "Surname", "Email", "Id", "Permission", "Ru", "En", "Audience", "Issuer", "Subject")
	if EmptyField {
		c.JSON(404, err)
	} else {
		client, ctx := mongoconnect.DBConnection()

		DBConnect := client.Database(env.Data_Name).Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"phone": LoginTemp.Phone,
		})

		var userdata structs.UserStruct
		result.Decode(&userdata)
		isValidPass := hashedpasswod.CompareHashPasswords(userdata.Password, LoginTemp.Password)
		fmt.Println(isValidPass)
		key := returnjwt.GenerateToken(userdata.Phone, userdata.Permission, userdata.Id)
		if userdata.Phone == LoginTemp.Phone {
			if isValidPass {
				http.SetCookie(c.Writer, &http.Cookie{
					Name:     env.Data_Name,
					Value:    key,
					Expires:  time.Now().Add(60 * time.Hour),
					Domain:   "",
					Path:     "/",
					Secure:   false,
					HttpOnly: false,
					SameSite: http.SameSiteLaxMode,
				})
				c.JSON(200, "success")
			} else {
				c.JSON(404, "Not valid pass")
			}
		} else {
			c.JSON(400, "Wrong phone")
		}
	}

}





func AddStatistic(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error:only admin have ecses to add")
		} else {
			var statistic_shablon structs.AddStatistic

			c.ShouldBindJSON(&statistic_shablon)

			Emptyfield, err := emptyfieldcheker.EmptyField(statistic_shablon, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("Statistic")

				ID := primitive.NewObjectID().Hex()
				insertrezult, inserterror := createDB.InsertOne(ctx, bson.M{
					"_id":      ID,
					"Quantity": statistic_shablon.Quantity,
					"ru": structs.LangForStatistic{
						Description: statistic_shablon.Ru.Description,
					},
					"en": structs.LangForStatistic{
						Description: statistic_shablon.En.Description,
					},
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(201, "succes")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}
		}

	}
}






func AddPatientStory(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error:only admin have ecses to add")
		} else {

			var Patient_data structs.Patient_story
			c.ShouldBindJSON(&Patient_data)

			Emptyfield, err := emptyfieldcheker.EmptyField(Patient_data, "Id", "Audience", "Issuer", "Subject")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				folderName := "PatientData"

				// Создание папки
				err := os.Mkdir("Statics"+"/"+folderName, os.ModePerm)
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
				Patient_data.Photo = baner.ImageFunc(Patient_data.Photo, ForImage, folderName)
				fmt.Printf("Imagstruct: %v\n", Patient_data)
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("PatientStory")
				fmt.Printf("Patient_data: %v\n", Patient_data)
				ID := primitive.NewObjectID().Hex()
				insertrezult, inserterror := createDB.InsertOne(ctx, bson.M{

					"_id":   ID,
					"photo": folderName + "/" + Patient_data.Photo,
					"ru": structs.LangForPatient{
						Full_Name:        Patient_data.Ru.Full_Name,
						Description: Patient_data.Ru.Description,
						Quot:      Patient_data.Ru.Quot,
					},
					"en": structs.LangForPatient{
						Full_Name:        Patient_data.En.Full_Name,
						Description: Patient_data.En.Description,
						Quot:      Patient_data.En.Quot,
					},
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "succes")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}
		}

	}
}







func AddPartner(c *gin.Context) {
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

			var PartnerData structs.Partner
			c.ShouldBindJSON(&PartnerData)
			Emptyfield, err := emptyfieldcheker.EmptyField(PartnerData, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {

				folderName := "Partners"

				// Создание папки
				err := os.Mkdir("Statics"+"/"+folderName, os.ModePerm)
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
				PartnerData.Logo = baner.ImageFunc(PartnerData.Logo, ForImage, folderName)
				fmt.Printf("Imagstruct: %v\n", PartnerData)
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database("OpenHearts").Collection("Partners")
				fmt.Printf("Patient_data: %v\n", PartnerData)
				ID := primitive.NewObjectID().Hex()
				insertrezult, inserterror := createDB.InsertOne(ctx, bson.M{

					"_id":  ID,
					"logo": folderName + "/" + PartnerData.Logo,
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(200, "succes")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}
		}
	}
}






func Add_statistic_for_center(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error:only admin have ecses to add")
		} else {
			var statistic_shablon structs.AddStatisticForCenter

			c.ShouldBindJSON(&statistic_shablon)

			Emptyfield, err := emptyfieldcheker.EmptyField(statistic_shablon, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("statistic_for_center")

				ID := primitive.NewObjectID().Hex()
				insertrezult, inserterror := createDB.InsertOne(ctx, bson.M{
					"_id":      ID,
					"Quantity": statistic_shablon.Quantity,
					"ru": structs.LangForProjectStatistic{
						Description: statistic_shablon.Ru.Description,
						Name: statistic_shablon.Ru.Name,
					},
					"en": structs.LangForProjectStatistic{
						Description: statistic_shablon.En.Description,
						Name: statistic_shablon.En.Name,
					},
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(201, "succes")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}
		}

	}
}


func Change_Number_in_Project(c *gin.Context) {
	var cookidata, cookieerror = c.Request.Cookie(env.Data_Name)
	if cookieerror != nil {
		fmt.Printf("cookieerror: %v\n", cookieerror)
		c.JSON(404, "error Not Cookie found")
		fmt.Printf("cookidata: %v\n", cookidata)
	} else {
		SecretKeyData, isvalid := returnjwt.Validate(cookidata.Value)
		if SecretKeyData.Permission != "Admin" && isvalid {
			c.JSON(404, "error:only admin have ecses to add")
		} else {
			var statistic_shablon structs.ChangNumber

			c.ShouldBindJSON(&statistic_shablon)

			Emptyfield, err := emptyfieldcheker.EmptyField(statistic_shablon, "Id")
			if Emptyfield {
				c.JSON(404, err)
			} else {
				client, ctx := mongoconnect.DBConnection()

				var createDB = client.Database(env.Data_Name).Collection("change_number_for_center")

				ID := primitive.NewObjectID().Hex()
				insertrezult, inserterror := createDB.InsertOne(ctx, bson.M{
					"_id":      ID,
					"Quantity": statistic_shablon.Quantity,
				})
				if inserterror != nil {
					fmt.Printf("inserterror: %v\n", inserterror)
				} else {
					c.JSON(201, "succes")
					fmt.Printf("insertrezult: %v\n", insertrezult)
				}
			}
		}

	}
}



// func Cors(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.6.237:5173")
// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
// 	if c.Request.Method == "OPTIONS" {
// 		c.AbortWithStatus(200)
// 		return
// 	}

// 	c.Next()
// }

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5502")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}

	c.Next()
}
