package routes

import (
	"github.com/TiiQu-Network/claim-verifier-prototype/http/controllers"
	"github.com/gobuffalo/buffalo"
)

func Routes(app *buffalo.App) {
	// Home routes
	homeController := new(controllers.Home)
	app.GET("/", homeController.Home)
	app.GET("/regenerate/", homeController.Regenerate)

	// Institution routes
	institutionController := new(controllers.Institution)
	app.Resource("/institution/", &controllers.InstitutionResource{&buffalo.BaseResource{}})
	app.GET("/institution/{institution}/students/", institutionController.Students)
	app.GET("/institution/{institution}/students/toBlockchain/", institutionController.ToBlockchain)
	app.GET("/institution/{institution}/blockchain/", institutionController.Blockchain)
	// Recipient routes
	studentController := new(controllers.Recipient)
	app.Resource("/student/", &controllers.RecipientResource{&buffalo.BaseResource{}})
	app.GET("/student/{student}/to-blockchain/", studentController.ToBlockchain)

	// Platform Members routes
	memberController := new(controllers.Member)
	app.Resource("/member/", &controllers.MemberResource{&buffalo.BaseResource{}})
	app.GET("/member/{member}/certifications/", memberController.Certifications)
	app.GET("/member/{member}/certifications/add/", memberController.AddCertification)
	app.GET("/member/{member}/blockchain/", memberController.Blockchain)

	// Certification routes
	memberCertificationController := new(controllers.MemberCertification)
	app.Resource("/member-certification/", &controllers.MemberCertificationResource{&buffalo.BaseResource{}})
	app.GET("/member-certification/{memberCertification}/to-blockchain/", memberCertificationController.ToBlockchain)

	// Data-hash routes
	app.Resource("/dataHash/", &controllers.DataHashResource{&buffalo.BaseResource{}})
}
