package authlogin

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/hoysoft/authlogin/models"
)

// oprations for NodeRoles
type NodeRolesController struct {
	AdminController
}

func (this *NodeRolesController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create NodeRoles
// @Param	body		body 	models.NodeRoles	true		"body for NodeRoles content"
// @Success 200 {int} models.NodeRoles.Id
// @Failure 403 body is empty
// @router / [post]
func (this *NodeRolesController) Post() {
	var v models.NodeRoles
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if id, err := models.AddNodeRoles(&v); err == nil {
		this.Data["json"] = map[string]int64{"id": id}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Get
// @Description get NodeRoles by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.NodeRoles
// @Failure 403 :id is empty
// @router /:id [get]
func (this *NodeRolesController) GetOne() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNodeRolesById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJson()
}

// @Title Get All
// @Description get NodeRoles
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.NodeRoles
// @Failure 403
// @router / [get]
func (this *NodeRolesController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := this.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := this.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				this.Data["json"] = errors.New("Error: invalid query key/value pair")
				this.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllNodeRoles(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJson()
}

// @Title Update
// @Description update the NodeRoles
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.NodeRoles	true		"body for NodeRoles content"
// @Success 200 {object} models.NodeRoles
// @Failure 403 :id is not int
// @router /:id [put]
func (this *NodeRolesController) Put() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.NodeRoles{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdateNodeRolesById(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Delete
// @Description delete the NodeRoles
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *NodeRolesController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteNodeRoles(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}
