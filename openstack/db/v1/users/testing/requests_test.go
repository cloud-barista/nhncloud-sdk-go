package testing

import (
	"testing"

	db "github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/db/v1/databases"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/db/v1/users"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
	fake "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreate(t)

	opts := users.BatchCreateOpts{
		{
			Databases: db.BatchCreateOpts{
				db.CreateOpts{Name: "databaseA"},
			},
			Name:     "dbuser3",
			Password: "secretsecret",
		},
		{
			Databases: db.BatchCreateOpts{
				{Name: "databaseB"},
				{Name: "databaseC"},
			},
			Name:     "dbuser4",
			Password: "secretsecret",
		},
	}

	res := users.Create(fake.ServiceClient(), instanceID, opts)
	th.AssertNoErr(t, res.Err)
}

func TestUserList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleList(t)

	expectedUsers := []users.User{
		{
			Databases: []db.Database{
				{Name: "databaseA"},
			},
			Name: "dbuser3",
		},
		{
			Databases: []db.Database{
				{Name: "databaseB"},
				{Name: "databaseC"},
			},
			Name: "dbuser4",
		},
	}

	pages := 0
	err := users.List(fake.ServiceClient(), instanceID).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := users.ExtractUsers(page)
		if err != nil {
			return false, err
		}

		th.CheckDeepEquals(t, expectedUsers, actual)
		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, pages)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDelete(t)

	res := users.Delete(fake.ServiceClient(), instanceID, "{userName}")
	th.AssertNoErr(t, res.Err)
}
