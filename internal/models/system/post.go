package system

import (
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/pkg/utils"
)

type PostModel struct {
	PostId      int    `gorm:"primary_key;AUTO_INCREMENT" json:"postId"`
	PostName    string `gorm:"type:varchar(128)" json:"postName"`
	PostCode    string `gorm:"type:varchar(128)" json:"postCode"`
	Sort        int    `gorm:"type:int(4)" json:"sort"`
	Status      string `gorm:"type:int(1)" json:"status"`
	CreatedUser string `gorm:"type:varchar(128)" json:"createdUser"`
	UpdatedUser string `gorm:"type:varchar(128)" json:"updatedUser"`
	Remark      string `gorm:"type:varchar(255)" json:"remark"`
	Params      string `gorm:"-" json:"params"`
	BaseModel
}

func (PostModel) TableName() string {
	return "sys_post"
}

func (p *PostModel) Create() (PostModel, error) {
	var doc PostModel
	result := orm.Eloquent.Table(p.TableName()).Create(&p)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *p
	return doc, nil
}

func (p *PostModel) Get() (PostModel, error) {
	var doc PostModel

	table := orm.Eloquent.Table(p.TableName())
	if p.PostId != 0 {
		table = table.Where("post_id = ? ", p.PostId)
	}
	if p.PostName != "" {
		table = table.Where("post_name like ? ", "%"+p.PostName+"%")
	}
	if p.PostCode != "" {
		table = table.Where("post_code = ? ", p.PostCode)
	}
	if p.Status != "" {
		table = table.Where("status = ? ", p.Status)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (p *PostModel) GetList() ([]PostModel, error) {
	var doc []PostModel

	table := orm.Eloquent.Table(p.TableName())
	if p.PostId != 0 {
		table = table.Where("post_id = ? ", p.PostId)
	}
	if p.PostName != "" {
		table = table.Where("post_name like ? ", "%"+p.PostName+"%")
	}
	if p.PostCode != "" {
		table = table.Where("post_code = ? ", p.PostCode)
	}
	if p.Status != "" {
		table = table.Where("status = ? ", p.Status)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (p *PostModel) GetPage(page, pageSize int) ([]PostModel, int, error) {
	var (
		total int64
		doc   []PostModel
	)

	table := orm.Eloquent.Select("*").Table(p.TableName())
	if p.PostId != 0 {
		table = table.Where("post_id = ? ", p.PostId)
	}
	if p.PostName != "" {
		table = table.Where("post_name like ? ", "%"+p.PostName+"%")
	}
	if p.PostCode != "" {
		table = table.Where("post_code = ? ", p.PostCode)
	}
	if p.Status != "" {
		table = table.Where("status = ? ", p.Status)
	}

	if err := table.Order("sort").Offset((page - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}

	table.Where("`deleted_at` IS NULL").Count(&total)

	return doc, utils.Int642Int(total), nil
}

func (p *PostModel) Update(id int) (update PostModel, err error) {
	if err = orm.Eloquent.Table(p.TableName()).First(&update, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table(p.TableName()).Model(&update).Updates(&p).Error; err != nil {
		return
	}
	return
}

func (p *PostModel) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(p.TableName()).Where("post_id = ? ", id).Delete(&PostModel{}).Error; err != nil {
		success = false
		return
	}

	success = true
	return
}

func (p *PostModel) BatchDelete(id []int) (success bool, err error) {
	if err = orm.Eloquent.Table(p.TableName()).Where("post_id in (?)", id).Delete(&PostModel{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
