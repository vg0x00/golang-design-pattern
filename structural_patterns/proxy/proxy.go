package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int) (User, error)
}

type User struct {
	Id int32
}

type UserList []User // using type alais for implement interface

func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].Id == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("user with id: %d not found\n", id)
}

func (u *UserList) AddUser(user User) {
	*u = append(*u, user)
}

type UserFinderProxy struct {
	UserCache          UserList
	UserCacheSize      int
	LastQueryFromCache bool
	UnrealDatabase     *UserList
}

func (p *UserFinderProxy) FindUser(id int32) (User, error) {
	// first find in cache, then database
	user, err := p.UserCache.FindUser(id)
	if err == nil {
		fmt.Println("fetch user from cache")
		p.LastQueryFromCache = true
		return user, nil
	}

	user, err = p.UnrealDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	p.AddUserToCache(user)
	p.LastQueryFromCache = false
	return user, nil
}

func (p *UserFinderProxy) AddUserToCache(user User) {
	if len(p.UserCache) >= p.UserCacheSize {
		// move pointer one step after
		// same size as UserCacheSize defined
		p.UserCache = append(p.UserCache[1:], user)
	} else {
		p.UserCache.AddUser(user)
	}
}
