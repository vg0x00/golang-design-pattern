package proxy

import (
	"math/rand"
	"testing"
)

func TestProxy(t *testing.T) {
	unrealDatabase := UserList{} // faked user database

	t.Run("inti database", func(t *testing.T) {
		rand.Seed(32321)
		var id int32
		for i := 0; i < 5; i++ {
			id = rand.Int31()
			unrealDatabase = append(unrealDatabase, User{Id: id})
		}
	})

	knownIds := []int32{unrealDatabase[0].Id, unrealDatabase[1].Id, unrealDatabase[2].Id}
	var proxy_obj UserFinderProxy

	t.Run("proxy", func(t *testing.T) {
		proxy_obj = UserFinderProxy{
			UserCache:      UserList{},
			UserCacheSize:  2,
			UnrealDatabase: &unrealDatabase,
		}
	})

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy_obj.FindUser(knownIds[0])
		if err != nil {
			t.Error(err.Error())
		}
		if user.Id != knownIds[0] {
			t.Errorf("user id should be the same but not:\n got: %d\n expect: %d\n",
				user.Id, knownIds[0])
		}
		if len(proxy_obj.UserCache) != 1 {
			t.Errorf("cache length should be 1 after first query, but got: %d\n",
				len(proxy_obj.UserCache))
		}
		if len(proxy_obj.UserCache) == 1 && proxy_obj.UserCache[0].Id != knownIds[0] {
			t.Errorf("cache id not matched:\n got: %d\n, expected: %d\n",
				proxy_obj.UserCache[0].Id, knownIds[0])
		}
		if proxy_obj.LastQueryFromCache {
			t.Errorf("after the first time query, data should not from cache\n")
		}
	})

	t.Run("FindUser - find user the same as before", func(t *testing.T) {
		user, err := proxy_obj.FindUser(knownIds[0])
		if err != nil {
			t.Error(err.Error())
		}
		if user.Id != knownIds[0] {
			t.Errorf("user id should be the same but not:\n got: %d\n expect: %d\n",
				user.Id, knownIds[0])
		}
		if len(proxy_obj.UserCache) != 1 {
			t.Errorf("cache length should be 1 after first query, but got: %d\n",
				len(proxy_obj.UserCache))
		}
		if len(proxy_obj.UserCache) == 1 && proxy_obj.UserCache[0].Id != knownIds[0] {
			t.Errorf("cache id not matched:\n got: %d\n, expected: %d\n",
				proxy_obj.UserCache[0].Id, knownIds[0])
		}
		if !proxy_obj.LastQueryFromCache {
			t.Errorf("user data should return from cache\n")
		}
	})

	t.Run("FindUser - find user the 3rd user", func(t *testing.T) {
		user1, err := proxy_obj.FindUser(knownIds[0])
		if err != nil {
			t.Error(err.Error())
		}

		user2, _ := proxy_obj.FindUser(knownIds[1])
		if proxy_obj.LastQueryFromCache {
			t.Errorf("the user does not stored in cache, should not return from cache\n")
		}

		user3, _ := proxy_obj.FindUser(knownIds[2])
		if proxy_obj.LastQueryFromCache {
			t.Errorf("the user does not stored in cache, should not return from cache\n")
		}

		if len(proxy_obj.UserCache) != 2 {
			t.Errorf("user cache size should be 2 after 3rd user query: %d\n",
				len(proxy_obj.UserCache))
		}

		for i := 0; i < 2; i++ {
			if len(proxy_obj.UserCache) == 2 && proxy_obj.UserCache[i].Id == user1.Id {
				t.Errorf("user 1 should not in cache now")
			}
		}

		for _, u := range proxy_obj.UserCache {
			if u != user2 && u != user3 {
				t.Errorf("unexpected user in cache\n")
			}
		}
	})

}
