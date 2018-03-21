package routes

import (
	"github.com/TiiQu-Network/claim-verifier-prototype/http/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controllers.Handler)

	institutionController := new(controllers.Institution)
	//TODO below should not belong to this controller
	http.HandleFunc("/institution/regenerate/", institutionController.Regenerate)

	http.HandleFunc("/institution/", controllers.Handler)
	http.HandleFunc("/institution/{institution}/students/", institutionController.Students)
	http.HandleFunc("/institution/{institution}/students/toBlockchain/", institutionController.ToBlockchain)
	http.HandleFunc("/institution/{institution}/blockchain/", institutionController.Blockchain)

	studentController := new(controllers.Student)
	http.HandleFunc("/student/", controllers.Handler)
	http.HandleFunc("/student/{student}/to-blockchain/", studentController.ToBlockchain)

	memberController := new(controllers.Member)
	http.HandleFunc("/member/", controllers.Handler)
	http.HandleFunc("/member/{member}/certifications/", memberController.Certifications)
	http.HandleFunc("/member/{member}/certifications/add/", memberController.AddCertification)
	http.HandleFunc("/member/{member}/blockchain/", memberController.Blockchain)

	memberCertificationController := new(controllers.MemberCertification)
	http.HandleFunc("/member-certification/", controllers.Handler)
	http.HandleFunc("/member-certification/{memberCertification}/to-blockchain/", memberCertificationController.ToBlockchain)

	dataHashController := new(controllers.DataHash)
	http.HandleFunc("/dataHash/", dataHashController.Resource)
}
