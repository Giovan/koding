package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroupChannel(t *testing.T) {
	Convey("while testing group channel", t, func() {

		Convey("channel should be there", nil)

		Convey("owner should be able to update it", nil)

		Convey("normal user should not be able to update it", nil)

		Convey("owner cant delete it", nil)

		Convey("normal user cant delete it", nil)

		Convey("owner cant delete it", func() {
			account := models.NewAccount()
			account.OldId = AccountOldId.Hex()
			account, err := createAccount(account)
			So(err, ShouldBeNil)
			So(account, ShouldNotBeNil)

			channel1, err := createChannelByGroupNameAndType(account.Id, "testgroup", models.Channel_TYPE_GROUP)
			So(err, ShouldBeNil)
			So(channel1, ShouldNotBeNil)

			err = deleteChannel(account.Id, channel1.Id)
			So(err, ShouldNotBeNil)
		})

		Convey("normal user cant delete it", func() {
			account := models.NewAccount()
			account.OldId = AccountOldId.Hex()
			account, err := createAccount(account)
			So(err, ShouldBeNil)
			So(account, ShouldNotBeNil)

			channel1, err := createChannelByGroupNameAndType(account.Id, "testgroup", models.Channel_TYPE_GROUP)
			So(err, ShouldBeNil)
			So(channel1, ShouldNotBeNil)

			err = deleteChannel(rand.Int63(), channel1.Id)
			So(err, ShouldNotBeNil)
		})

		Convey("member can post status update", nil)

		Convey("non-member can not post status update", nil)

		Convey("member can post status update", nil)

		Convey("non-member can not post status update", nil)

	})
}
