/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package model

type VDESrcMember struct {
	ID                   int64  `gorm:"primary_key; not null" json:"id"`
	VDEPubKey            string `gorm:"not null" json:"vde_pub_key"`
	VDEComment           string `gorm:"not null" json:"vde_comment"`
	VDEName              string `gorm:"not null" json:"vde_name"`
	VDEIp                string `gorm:"not null" json:"vde_ip"`
	VDEType              int64  `gorm:"not null" json:"vde_type"`
	ContractRunHttp      string `gorm:"not null" json:"contract_run_http"`
	ContractRunEcosystem string `gorm:"not null" json:"contract_run_ecosystem"`

	UpdateTime int64 `gorm:"not null" json:"update_time"`
	CreateTime int64 `gorm:"not null" json:"create_time"`

func (m *VDESrcMember) Delete() error {
	return DBConn.Delete(m).Error
}

func (m *VDESrcMember) GetAll() ([]VDESrcMember, error) {
	var result []VDESrcMember
	err := DBConn.Find(&result).Error
	return result, err
}
func (m *VDESrcMember) GetOneByID() (*VDESrcMember, error) {
	err := DBConn.Where("id=?", m.ID).First(&m).Error
	return m, err
}

func (m *VDESrcMember) GetOneByPubKey(VDEPubKey string) (*VDESrcMember, error) {
	err := DBConn.Where("vde_pub_key=?", VDEPubKey).First(&m).Error
	return m, err
}

func (m *VDESrcMember) GetAllByType(Type int64) ([]VDESrcMember, error) {
	result := make([]VDESrcMember, 0)
	err := DBConn.Table("vde_src_member").Where("vde_type = ?", Type).Find(&result).Error
	return result, err
}