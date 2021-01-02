package router

import (
	"github.com/gin-gonic/gin"
	"gover-server/controller"
	"gover-server/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("tokenAuth", middleware.TokenAuthHandler)

	r.POST("login", controller.LoginCheck)
	r.POST("loginAdmin", controller.LoginAdminCheck)

	r.POST("department/add", controller.AddDepartment)
	r.GET("department/all", controller.SendAllDepartments)
	r.GET("department/get/:deptId", controller.SendDepartment)
	r.POST("department/delete", controller.DepartmentDelete)
	r.POST("department/update", controller.UpdateDepartment)

	r.GET("gover/all", controller.SendAllGovernors)
	r.POST("gover/add", controller.AddGovernor)
	r.POST("gover/delete", controller.DeleteGovernor)

	r.GET("business/all", controller.SendAllBusiness)
	r.POST("business/add/:deptId", controller.AddBusiness)
	r.POST("business/delete/:busId", controller.DeleteBusiness)
	r.POST("businesses/get", controller.SendAllBusinessOfDeptByDeptId)
	r.GET("businesses/getHot", controller.SendHotBusiness)
	r.POST("bus/get", controller.SendBusiness)
	r.POST("bus/update", controller.UpdateBusiness)
	r.POST("bus/addTemplate", controller.AddBusTemplate)

	r.POST("material/get", controller.SendMaterials)
	r.POST("material/add", controller.AddMaterial)
	r.POST("material/update", controller.UpdateMaterial)
	r.POST("material/delete", controller.DeleteMaterial)

	r.POST("process/add", controller.AddProcessForBus)
	r.POST("process/addMaterial/:processId", controller.AddProcessMaterial)
	r.GET("process/all/:busId", controller.SendAllProcessOfBus)
	r.POST("process/delete", controller.DeleteProcess)
	r.POST("process/deleteMaterial", controller.DeleteProcessMaterial)

	r.GET("chat/getHot", controller.GetHotChats)
	r.GET("chat/get/:chatId", controller.SendChat)
	r.GET("chat/getSubChat/:chatId", controller.SendSubChat)
	r.POST("chat/create", controller.CreateChat)
	r.POST("chat/createSubChat", controller.CreateSubChat)
	r.POST("chat/deleteSubChat", controller.DeleteSubChat)
	r.POST("chat/like", controller.LikeChat)
	r.POST("chat/getByTitle", controller.SendChatByTitle)

	r.GET("comment/getDept/:deptId", controller.SendDeptComment)
	r.POST("comment/feedBack", controller.FeedBack)
	return r
}
