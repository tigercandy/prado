package system

import (
	"errors"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/pkg/utils"
)

type DeptModel struct {
	DeptId      int         `gorm:"primary_key:AUTO_INCREMENT" json:"deptId"`
	ParentId    int         `gorm:"type:int(11)" json:"parentId"`
	DeptPath    string      `gorm:"type:varchar(255)" json:"deptPath"`
	DeptName    string      `gorm:"type:varchar(128)" json:"deptName"`
	Sort        int         `gorm:"type:int(4)" json:"sort"`
	Leader      int         `gorm:"type:int(11)" json:"leader"`
	Phone       string      `gorm:"type:varchar(11)" json:"phone"`
	Email       string      `gorm:"type:varchar(64)" json:"email"`
	Status      string      `gorm:"type:int(1)" json:"status"`
	CreatedUser string      `gorm:"type:varchar(64)" json:"createdUser"`
	UpdatedUser string      `gorm:"type:varchar(64)" json:"updatedUser"`
	Params      string      `gorm:"-" json:"params"`
	Children    []DeptModel `gorm:"-" json:"children"`
	BaseModel
}

type DeptLabel struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []DeptModel `gorm:"-" json:"children"`
}

func (DeptModel) TableName() string {
	return "sys_dept"
}

func (d *DeptModel) Get() (DeptModel, error) {
	var dept DeptModel
	table := orm.Eloquent.Table(d.TableName())
	if d.DeptId != 0 {
		table = table.Where("dept_id = ? ", d.DeptId)
	}
	if d.DeptName != "" {
		table = table.Where("dept_name = ? ", d.DeptName)
	}
	if err := table.First(&dept).Error; err != nil {
		return dept, err
	}

	return dept, nil
}

func (d *DeptModel) GetPage() ([]DeptModel, error) {
	var depts []DeptModel

	table := orm.Eloquent.Select("*").Table(d.TableName())
	if d.DeptId != 0 {
		table = table.Where("dept_id = ? ", d.DeptId)
	}
	if d.DeptName != "" {
		table = table.Where("dept_name = ? ", d.DeptName)
	}
	if d.Status != "" {
		table = table.Where("status = ? ", d.Status)
	}
	if d.DeptPath != "" {
		table = table.Where("dept_path like %?% ", d.DeptPath)
	}

	if err := table.Order("sort").Find(&depts).Error; err != nil {
		return nil, err
	}

	return depts, nil
}

func (d *DeptModel) Create() (DeptModel, error) {
	var dept DeptModel
	if err := orm.Eloquent.Table(d.TableName()).Create(d).Error; err != nil {
		return dept, err
	}
	deptPath := "/" + utils.Int2String(d.DeptId)
	if int(d.ParentId) != 0 {
		var cp DeptModel
		orm.Eloquent.Table(d.TableName()).Where("dept_id = ? ", d.ParentId).First(&cp)
		deptPath = cp.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	var mp = map[string]string{}
	mp["deptPath"] = deptPath
	if err := orm.Eloquent.Table(d.TableName()).Where("dept_id = ? ", d.DeptId).Updates(mp).Error; err != nil {
		return dept, err
	}
	dept = *d
	d.DeptPath = deptPath
	return dept, nil
}

func (d *DeptModel) Update(id int) (update DeptModel, err error) {
	if err = orm.Eloquent.Table(d.TableName()).Where("dept_id = ? ", id).First(&update).Error; err != nil {
		return
	}

	deptPath := "/" + utils.Int2String(d.DeptId)
	if int(d.ParentId) != 0 {
		var cp DeptModel
		orm.Eloquent.Table(d.TableName()).Where("dept_id = ? ", d.ParentId).First(&cp)
		deptPath = cp.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	d.DeptPath = deptPath
	if d.DeptPath != "" && d.DeptPath != update.DeptPath {
		return update, errors.New("不允许修改上级部门!")
	}

	if err = orm.Eloquent.Table(d.TableName()).Model(&update).Updates(&d).Error; err != nil {
		return
	}
	return
}

func (d *DeptModel) Delete(id int) (success bool, err error) {
	user := SysUser{}
	user.DeptId = id
	userList, err := user.GetList()
	if len(userList) > 0 {
		return false, errors.New("当前部门存在用户，不允许删除!")
	}
	if err = orm.Eloquent.Table(d.TableName()).Where("dept_id = ? ", id).Delete(&DeptModel{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
