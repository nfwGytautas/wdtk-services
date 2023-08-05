package context

func (ad *AuthData) GetUser(identifier string) (User, error) {
	u := User{}

	if err := ad.verifyConnection(); err != nil {
		return u, err
	}

	err := ad.dbConn.Model(User{}).Where("identifier = ?", identifier).Take(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (ad *AuthData) GetUserByID(identifier uint) (User, error) {
	u := User{}

	if err := ad.verifyConnection(); err != nil {
		return u, err
	}

	err := ad.dbConn.Model(User{}).First(&u, identifier).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (ad *AuthData) CreateUser(user *User) error {
	if err := ad.verifyConnection(); err != nil {
		return err
	}

	err := ad.dbConn.Create(user).Error
	return err
}
