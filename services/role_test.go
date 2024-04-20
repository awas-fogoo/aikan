package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"one/common"
	"one/model"
	"testing"
)

func CreateRole(db *gorm.DB, name string, permissions []model.Permission) (*model.Role, error) {
	role := model.Role{Name: name, Permissions: permissions}
	if err := db.Create(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
func AssignRoleToUser(db *gorm.DB, userID uint, roleID uint) error {
	var user model.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}
	var role model.Role
	if err := db.First(&role, roleID).Error; err != nil {
		return err
	}

	// 将角色添加到用户的角色列表中
	return db.Model(&user).Association("Roles").Append(&role)
}
func InitializeRoles(db *gorm.DB) {
	// 创建无权限的基本角色
	role, err := CreateRole(db, "Guest", nil)
	if err != nil {
		return
	}
	fmt.Println(role)
	// 创建管理员角色，并分配所有权限
	allPermissions := []model.Permission{{Name: "manage_users"}, {Name: "manage_roles"}, {Name: "edit_content"}}
	createRole, err := CreateRole(db, "Administrator", allPermissions)
	if err != nil {
		return
	}
	fmt.Println(createRole)
}
func AddRole(c *gin.Context) {
	var newRoleData struct {
		Name        string `json:"name"`
		Permissions []uint `json:"permissions"` // 权限ID列表
	}
	if err := c.ShouldBindJSON(&newRoleData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	permissions := make([]model.Permission, len(newRoleData.Permissions))
	for i, permID := range newRoleData.Permissions {
		permissions[i] = model.Permission{Model: gorm.Model{ID: permID}}
	}

	role, err := CreateRole(common.DB, newRoleData.Name, permissions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role created", "role": role})
}

func TestAddRole(t *testing.T) {

}
