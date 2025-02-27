package handlers

import (
	"docs/app/controlers"
	"docs/app/createadmin"

	"github.com/gin-gonic/gin"
)

func Handlers() {
	createadmin.Createadmin()
	r := gin.Default()
	r.Use(controlers.Cors)
	r.POST("/check/email",controlers.CheckEmail)
	r.POST("/check/code",controlers.CheckSecretCode)

	r.POST("/login",controlers.Login)
	r.POST("/updateAdmin",createadmin.UpdateAdmin)


	r.POST("/add/statistic",controlers.AddStatistic)
	r.POST("/add_pationt_story",controlers.AddPatientStory)
	r.POST("/add/partner",controlers.AddPartner)
	r.POST("/Add_statistic_for_project",controlers.Add_statistic_for_center)
	r.POST("/Add_Mumber",controlers.Change_Number_in_Project)
	
	
	r.POST("/update/statistic",controlers.UpdateStatistic)
	r.POST("/updateServiceNumber",controlers.UpdateServiceNumber)
	r.POST("/updatecenterstatistic",controlers.UpdateStatisticForCenter)



	
	r.DELETE("/delete/pationt/story",controlers.DeletePatientStory)
	r.DELETE("/delete/partner",controlers.DeletePatners)




	
	r.GET("/get/statistic",controlers.GetStatistic)
	r.GET("/get/pationt/story",controlers.GetPatientStory)
	r.GET("/get/patners",controlers.GetPatners)
	r.GET("/getCenter",controlers.GetStatisticforproject)
	r.GET("/get_project_number",controlers.Get_ChangedNumber_for_project)





	

	r.Run(":2020")
}


