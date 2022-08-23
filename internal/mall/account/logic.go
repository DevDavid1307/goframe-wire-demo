package account

type logic struct {
	repo Dao
}

func (l *logic) GetUser() string {
	return l.repo.GetUser()
}

func (l *logic) Greet() string {
	return l.repo.Greeting()
}
