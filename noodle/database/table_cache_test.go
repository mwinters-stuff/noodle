package database_test

import (
	"testing"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TableCacheTestSuite struct {
	suite.Suite
}

func (suite *TabTableTestSuite) TestAdd() {
	cache := database.NewTableCache[models.User]()

	u1 := models.User{
		DN:          "Dn1",
		DisplayName: "User1",
		GivenName:   "User",
		ID:          1,
		Surname:     "One",
		UIDNumber:   1001,
		Username:    "userone",
	}

	u2 := models.User{
		DN:          "Dn2",
		DisplayName: "User2",
		GivenName:   "User",
		ID:          2,
		Surname:     "Two",
		UIDNumber:   1002,
		Username:    "usertwo",
	}

	cache.Add(1, u1)
	cache.Add(2, u2)

	found, user := cache.GetID(2)
	require.True(suite.T(), found)
	require.Equal(suite.T(), u2, user)

	found, user = cache.GetID(1)
	require.True(suite.T(), found)
	require.Equal(suite.T(), u1, user)

	found, user = cache.GetID(3)
	require.False(suite.T(), found)
	require.NotEqual(suite.T(), u1, user)
	require.NotEqual(suite.T(), u2, user)

	u1x := u1

	u1x.DN = "DNCHANGED"
	u1x.DisplayName = "New DisplayName"
	require.NotEqual(suite.T(), u1, u1x)

	found, user = cache.GetID(1)
	require.True(suite.T(), found)
	require.NotEqual(suite.T(), u1x, user)
	require.Equal(suite.T(), u1, user)

	cache.Update(1, u1x)

	found, user = cache.GetID(1)
	require.True(suite.T(), found)
	require.NotEqual(suite.T(), u1, user)
	require.Equal(suite.T(), u1x, user)

	found, user = cache.GetID(22)
	require.False(suite.T(), found)
	require.Equal(suite.T(), models.User{}, user)

	found, userfound := cache.Find(func(index int64, value models.User) bool {
		return value.Username == "usertwo"
	})

	require.True(suite.T(), found)
	require.Equal(suite.T(), u2, *userfound)

	found, userfound = cache.Find(func(index int64, value models.User) bool {
		return value.Username == "userthree"
	})

	require.False(suite.T(), found)
	require.Nil(suite.T(), userfound)

	var count = 0
	cache.ForEach(func(index int64, value models.User) bool {
		count++
		return true
	})
	require.Equal(suite.T(), 2, count)

	count = 0
	cache.ForEach(func(index int64, value models.User) bool {
		if index == 2 {
			return false
		}
		count++
		return true
	})
	require.Equal(suite.T(), 1, count)

	cache.DeleteIndex(2)
	found, _ = cache.GetID(2)
	require.False(suite.T(), found)

	cache.DeleteValue(u1)
	found, _ = cache.GetID(1)
	require.True(suite.T(), found)

	cache.DeleteValue(u1x)
	found, _ = cache.GetID(1)
	require.False(suite.T(), found)
}

func (suite *TabTableTestSuite) TestFindAll() {
	cache := database.NewTableCache[models.UserGroup]()

	ug1 := models.UserGroup{
		ID:        1,
		GroupID:   1,
		UserID:    1,
		GroupName: "Group 1",
	}
	ug2 := models.UserGroup{
		ID:        2,
		GroupID:   1,
		UserID:    2,
		GroupName: "Group 1",
	}
	ug3 := models.UserGroup{
		ID:        3,
		GroupID:   2,
		UserID:    2,
		GroupName: "Group 2",
	}
	ug4 := models.UserGroup{
		ID:        4,
		GroupID:   2,
		UserID:    3,
		GroupName: "Group 2",
	}

	cache.Add(1, ug1)
	cache.Add(2, ug2)
	cache.Add(3, ug3)
	cache.Add(4, ug4)

	found, ugs := cache.FindAll(func(index int64, value models.UserGroup) bool { return value.GroupID == 1 })
	require.True(suite.T(), found)
	require.Len(suite.T(), ugs, 2)
	require.ElementsMatch(suite.T(), []models.UserGroup{
		{
			ID:        1,
			GroupID:   1,
			UserID:    1,
			GroupName: "Group 1",
		},
		{
			ID:        2,
			GroupID:   1,
			UserID:    2,
			GroupName: "Group 1",
		},
	}, ugs)

	found, ugs = cache.FindAll(func(index int64, value models.UserGroup) bool { return value.UserID == 2 })
	require.True(suite.T(), found)
	require.Len(suite.T(), ugs, 2)
	require.ElementsMatch(suite.T(), []models.UserGroup{
		{
			ID:        3,
			GroupID:   2,
			UserID:    2,
			GroupName: "Group 2",
		},
		{
			ID:        2,
			GroupID:   1,
			UserID:    2,
			GroupName: "Group 1",
		},
	}, ugs)

	found, ugs = cache.FindAll(func(index int64, value models.UserGroup) bool { return value.UserID == 44 })
	require.False(suite.T(), found)
	require.Len(suite.T(), ugs, 0)
}

func TestTableCacheSuite(t *testing.T) {
	suite.Run(t, new(TabTableTestSuite))
}
