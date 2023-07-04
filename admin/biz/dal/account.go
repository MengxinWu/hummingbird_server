package dal

import "admin/biz/model/login"

func GetAccount(name string) (*login.Account, error) {
	var (
		account *login.Account
		err     error
	)
	db := DB.Model(login.Account{})
	if err = db.Where("name = ?", name).Find(&account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
