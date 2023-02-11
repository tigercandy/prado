package system

import (
	"errors"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/pkg/utils"
)

const StatusForbidden = "1" // 禁用状态

type User struct {
	IdentityKey string
	UserName    string
	Role        string
}

type UserName struct {
	Username string `gorm:"type:varchar(128)" json:"username"`
}

type PassWord struct {
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT" json:"userId"`
}

type LoginModel struct {
	UserName
	PassWord
}

type UserModel struct {
	NickName    string `gorm:"type:varchar(128)" json:"nickName"`
	Phone       string `gorm:"type:varchar(11)" json:"phone"`
	Email       string `gorm:"type:varchar(128)" json:"email"`
	RoleId      int    `gorm:"type:int(11)" json:"roleId"`
	Sex         string `gorm:"type:varchar(29)" json:"sex"`
	Avatar      string `gorm:"type:varchar(255)" json:"avatar"`
	Salt        string `gorm:"type:varchar(255)" json:"salt"`
	DeptId      int    `gorm:"type:int(11)" json:"deptId"`
	PostId      int    `gorm:"type:int(11)" json:"postId"`
	Remark      string `gorm:"type:varchar(255)" json:"remark"`
	Status      string `gorm:"type:int(1)" json:"status"`
	CreatedUser string `gorm:"type:varchar(128)" json:"createdUser"`
	UpdatedUser string `gorm:"type:varchar(128)" json:"updatedUser"`
	Params      string `gorm:"-" json:"params"`
	BaseModel
}

type SysUser struct {
	SysUserId
	UserModel
	LoginModel
}

type UserExpose struct {
	SysUserId
	UserModel
	LoginModel
	RoleName string `gorm:"column:role_name" json:"role_name"`
}

type UserPage struct {
	SysUserId
	UserModel
	LoginModel
	DeptName string `gorm:"column:dept_name" json:"deptName"`
}

type UserPwd struct {
	OldPwd string `json:"oldPwd" form:"oldPwd"`
	NewPwd string `json:"newPwd" form:"newPwd"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (u *SysUser) Get() (UserExpose UserExpose, err error) {

	var roleModel = new(SysRole)
	roleTable := roleModel.TableName()
	table := orm.Eloquent.Table(u.TableName()).Select([]string{u.TableName() + ".*", roleTable + ".role_name"})
	table = table.Joins("left join " + roleTable + " on " + u.TableName() + ".role_id = " + roleTable + ".role_id")
	if u.UserId != 0 {
		table = table.Where("user_id = ? ", u.UserId)
	}
	if u.Username != "" {
		table = table.Where("username = ? ", u.Username)
	}
	if u.Password != "" {
		table = table.Where("password = ? ", u.Password)
	}
	if u.RoleId != 0 {
		table = table.Where("role_id = ? ", u.RoleId)
	}
	if u.DeptId != 0 {
		table = table.Where("dept_id = ? ", u.DeptId)
	}
	if u.PostId != 0 {
		table = table.Where("post_id = ? ", u.PostId)
	}

	if err = table.First(&UserExpose).Error; err != nil {
		return
	}
	UserExpose.Password = ""
	return
}

func (u *SysUser) GetPage(page, pageSize int) ([]UserPage, int, error) {
	var (
		doc       []UserPage
		total     int64
		deptModel = new(DeptModel)
	)

	deptTable := deptModel.TableName()
	table := orm.Eloquent.Select(u.TableName() + ".*," + deptTable + ".dept_name").Table(u.TableName())
	table = table.Joins("left join " + deptTable + " on " + deptTable + ".dept_id = " + u.TableName() + ".dept_id")

	if u.Username != "" {
		table = table.Where(u.TableName()+".username like ?", "%"+u.Username+"%")
	}
	if u.NickName != "" {
		table = table.Where(u.TableName()+".nickname like ?", "%"+u.NickName+"%")
	}
	if u.Status != "" {
		table = table.Where(u.TableName()+".status = ?", u.Status)
	}

	if u.Phone != "" {
		table = table.Where(u.TableName()+".phone like ?", "%"+u.Phone+"%")
	}

	if u.DeptId != 0 {
		table = table.Where(u.TableName()+".dept_id in (select dept_id from "+deptTable+" where dept_path like ? )", "%"+utils.Int2String(u.DeptId)+"%")
	}

	if err := table.Offset((page - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where(u.TableName() + ".deleted_at IS NULL").Count(&total)
	return doc, utils.Int642Int(total), nil
}

func (u *SysUser) GetList() (UserExpose []UserExpose, err error) {
	var roleModel SysRole
	table := orm.Eloquent.Table(u.TableName()).Select([]string{u.TableName() + ".*", roleModel.TableName() + ".role_name"})
	table = table.Joins("left join " + roleModel.TableName() + " on " + u.TableName() + ".role_id = " + roleModel.TableName() + ".role_id")
	if u.UserId != 0 {
		table = table.Where("user_id = ? ", u.UserId)
	}
	if u.Username != "" {
		table = table.Where("username = ? ", u.Username)
	}
	if u.Password != "" {
		table = table.Where("password = ? ", u.Password)
	}
	if u.RoleId != 0 {
		table = table.Where("role_id = ? ", u.RoleId)
	}
	if u.DeptId != 0 {
		table = table.Where("dept_id = ? ", u.DeptId)
	}
	if u.PostId != 0 {
		table = table.Where("post_id = ? ", u.PostId)
	}
	if err = table.Find(&UserExpose).Error; err != nil {
		return
	}
	return
}

func (u *SysUser) Insert() (id int, err error) {
	hashPwd := utils.EncryptMake(u.Password)
	var count int64
	orm.Eloquent.Table(u.TableName()).Where("username = ? and `deleted_at IS NULL`", u.Username).Count(&count)
	if count > 0 {
		err = errors.New("用户名已存在!")
		return
	}
	u.Password = hashPwd
	if err = orm.Eloquent.Table(u.TableName()).Create(&u).Error; err != nil {
		return
	}
	id = u.UserId
	return
}

func (u *SysUser) Update(id int) (update SysUser, err error) {
	if u.Password != "" {
		u.Password = utils.EncryptMake(u.Password)
	}
	if err = orm.Eloquent.Table(u.TableName()).First(&update, id).Error; err != nil {
		return
	}
	if u.RoleId == 0 {
		u.RoleId = update.RoleId
	}
	if err = orm.Eloquent.Table(u.TableName()).Model(&update).Updates(&u).Error; err != nil {
		return
	}
	return
}

func (u *SysUser) BatchDelete(id []int) (res bool, err error) {
	if err = orm.Eloquent.Table(u.TableName()).Where("user_id in (?)", id).Delete(&SysUser{}).Error; err != nil {
		return
	}
	res = true
	return
}
