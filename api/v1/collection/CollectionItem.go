package collection

import (
	"github.com/gin-gonic/gin"
	service "kingford-backend/service/collection"
	"net/http"
)

// @Description 获取分页列表
// @Summary 分页列表
// @Tags CollectionItem
// @Accept application/json
// @Param pageIndex query int false "当前页"
// @Param pageSize query int false "分页条数"
// @success 200 {object} dto.Response
// @Failure 400 {object}  dto.Response "失败"
// @Router /api/collection-item [get]
func GetItemList(c *gin.Context) {
	var listService service.ListItemService
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.GetList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}


// @Description  通过id获取信息
// @Summary  通过Id获取信息
// @Tags CollectionItem
// @Accept application/x-json-stream
// @Param id path string true "主键id"
// @Success 200 {object} dto.Response "成功"
// @Failure 400 object  dto.Response "失败"
// @Router /api/collection-item/{id} [get]
func GetItem(c *gin.Context) {
	deleteService := service.GetItemService{}
	res := deleteService.Get(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

//@Description 创建信息
// @Tags CollectionItem
// @Summary 创建信息
// @Produce  application/json
// @Param collection body service.CreateItemService true "创建信息"
// @success 200 {object} dto.Response "成功"
// @Failure 400 {object}  dto.Response "失败"
// @Router /api/collection-item [post]
func CreateItem(c *gin.Context) {
	var createService service.CreateItemService
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}

// @Description 通过id修改信息
// @Summary 通过id修改信息
// @Tags CollectionItem
// @Accept application/json
// @Param id path string true "主键id"
// @Param collection body service.UpdateItemService true "修改信息"
// @success 200 {object} dto.Response "成功"
// @Failure 400 {object}  dto.Response "失败"
// @Router /api/collection-item/{id} [put]
func UpdateItem(c *gin.Context) {
	var updateService service.UpdateItemService
	if err := c.ShouldBind(&updateService); err == nil {
		id, _ := c.Params.Get("id")
		res := updateService.Update(id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusMethodNotAllowed, err.Error())
	}
}

// @Description  通过id删除信息
// @Summary  通过id删除信息
// @Tags CollectionItem
// @Accept application/x-json-stream
// @Param id path string true "主键id"
// @Success 200 {object} dto.Response "成功"
// @Failure 400 object  dto.Response "失败"
// @Router /api/collection-item/{id} [delete]
func DeleteItem(c *gin.Context) {
	deleteService := service.DeleteItemService{}
	res := deleteService.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
