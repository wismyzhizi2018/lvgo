package User

import (
	"fmt"
	"order/app/Http/Models/Kit"

	"gorm.io/gorm"
)

type Order_staff_list struct {
	StaffId        string `orm:"staff_id" json:"staff_id"`               // 员工id
	StaffName      string `orm:"staff_name" json:"staff_name"`           // 员工姓名
	UnionId        string `orm:"union_id" json:"union_id"`               // 员工在当前开发者企业账号范围内的唯一标识，系统生成，固定值，不会改变
	Mobile         string `orm:"mobile" json:"mobile"`                   // 手机号
	JobNumber      string `orm:"job_number" json:"job_number"`           // 工号
	WorkPlace      string `orm:"work_place" json:"work_place"`           // 工作地点
	Avatar         string `orm:"avatar" json:"avatar"`                   // 头像
	Position       string `orm:"position" json:"position"`               // 职位
	IsDimission    int    `orm:"is_dimission" json:"is_dimission"`       // 是否离职(1=未离职，2=离职)
	IsAdmin        int    `orm:"is_admin" json:"is_admin"`               // 是否是企业的管理员，1表示是，0表示不是
	IsBoss         int    `orm:"is_boss" json:"is_boss"`                 // 是否为企业的老板，1表示是，0表示不是
	IsHide         int    `orm:"is_hide" json:"is_hide"`                 // 是否隐藏号码，1表示是，0表示不是
	IsLeader       int    `orm:"is_leader" json:"is_leader"`             // 是否是部门的主管，1表示是，0表示不是
	IsSenior       int    `orm:"is_senior" json:"is_senior"`             // 是否是高管，1表示是，0表示不是
	Active         int    `orm:"active" json:"active"`                   // 该用户是否激活了钉钉，1表示是，0表示不是
	EmployeeType   int    `orm:"employee_type" json:"employee_type"`     // 员工类型(1=全职，2=兼职，3=实习)
	EmployeeStatus int    `orm:"employee_status" json:"employee_status"` // 在职员工状态（2=试用期，3=正式，4=离职，5=待离职）
	Department     string `orm:"department" json:"department"`           // 成员所属部门id列表
	Email          string `orm:"email" json:"email"`                     // 员工的邮箱
	HiredDate      string `orm:"hired_date" json:"hired_date"`           // 入职时间 日期格式
	StateCode      string `orm:"state_code" json:"state_code"`           // 国家地区码
	Remark         string `orm:"remark" json:"remark"`                   // 备注
	UpdatedAt      string `orm:"updated_at" json:"updated_at"`           // 更新时间
	CreatedAt      string `orm:"created_at" json:"created_at"`           // 创建时间
}

func (*Order_staff_list) TableName() string {
	return "order_staff_list"
}

func GetUserInfo(mobile string) (orderUser *Order_staff_list, err error) {
	// 查询数据库

	// WhereMap := map[string]interface{}{}
	// WhereMap["order_code"] = orderCode

	var result Order_staff_list
	fmt.Println(Kit.DB)
	if Kit.DB != nil {
	}
	if err := Kit.DB.Where("mobile = ?", mobile).First(&result).Error; err == nil {
		return &result, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}

type result struct {
	Data  *[]Order_staff_list `json:"data"`
	Count int64               `json:"total"`
}

// GetUserList 获取多条记录
// Usage:
//  user,_ := User.GetUserList(s.CurrentPage, s.PageSize)
func GetUserList(pageSize int, currentPage int) (res *result, err error) {
	// 查询数据库

	// WhereMap := map[string]interface{}{}
	// WhereMap["order_code"] = orderCode
	pageSize = (currentPage - 1) * pageSize

	// fmt.Println(Kit.DB)
	if Kit.DB != nil {
	}
	var UserList []Order_staff_list
	var Count int64

	// db.Limit(3).Find(&users)
	// db.Offset(3).Find(&users)
	if err := Kit.DB.Find(&UserList).Count(&Count).Error; err != nil {
		return nil, err
	}

	if err := Kit.DB.Offset(pageSize).Limit(currentPage).Find(&UserList).Error; err == nil {
		return &result{&UserList, Count}, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}
