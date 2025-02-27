package controlers

import (
	"docs/app/Env"
	"docs/app/mongoconnect"
	"docs/app/structs"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetStatistic(c *gin.Context) {
	var Forlist = []structs.AddStatistic{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Statistic")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddStatistic
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}



func GetStatisticforproject(c *gin.Context) {
	var Forlist = []structs.AddStatisticForCenter{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Statistic")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.AddStatisticForCenter
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}



func Get_ChangedNumber_for_project(c *gin.Context) {
	var Forlist = []structs.ChangNumber{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("change_number_for_center")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.ChangNumber
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}



func GetPatientStory(c *gin.Context) {
	var Forlist = []structs.Patient_story{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("PatientStory")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Patient_story
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}

func GetPatners(c *gin.Context) {
	var Forlist = []structs.Partner{}

	connect, ctx := mongoconnect.DBConnection()
	var createDB = connect.Database(env.Data_Name).Collection("Partners")

	var singlerezult, singerror = createDB.Find(ctx, bson.M{})
	if singerror != nil {
		fmt.Printf("singerror: %v\n", singerror)
	}

	for singlerezult.Next(ctx) {
		var datafromdb structs.Partner
		fmt.Printf("datafromdb: %v\n", datafromdb)
		singlerezult.Decode(&datafromdb)

		Forlist = append(Forlist, datafromdb)
	}
	c.JSON(200, Forlist)
}


