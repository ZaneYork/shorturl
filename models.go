package main

import (
	"time"
)

// URL ...
type URL struct {
	ID        string    `storm:"id"`
	URL       string    `storm:"index"`
	Name      string    `storm:"index"`
	CreatedAt time.Time `storm:"index"`
	UpdatedAt time.Time `storm:"index"`
}

func GenerateID(length int) string {
	for {
		id := RandomString(length)
		err := db.One("ID", id, nil)
		if err != nil {
			return id
		}
	}
}

func NewURL(target string, urlLength int) (url *URL, err error) {
	url = &URL{ID: GenerateID(urlLength), URL: target, CreatedAt: time.Now()}
	err = db.Save(url)
	return
}

// SetName ...
func (u *URL) SetName(name string) error {
	u.Name = name
	u.UpdatedAt = time.Now()
	return db.Save(&u)
}
