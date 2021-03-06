// Copyright (c) 2017 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"

	"github.com/mattermost/platform/app"
	"github.com/mattermost/platform/model"
	"github.com/mattermost/platform/utils"
)

func InitStatus() {
	l4g.Debug(utils.T("api.status.init.debug"))

	BaseRoutes.User.Handle("/status", ApiHandler(getUserStatus)).Methods("GET")

}

func getUserStatus(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireUserId()
	if c.Err != nil {
		return
	}

	if statusMap, err := app.GetUserStatusesByIds([]string{c.Params.UserId}); err != nil {
		c.Err = err
		return
	} else {
		if len(statusMap) == 0 {
			c.Err = model.NewAppError("UserStatus", "api.status.user_not_found.app_error", nil, "", http.StatusNotFound)
			return
		} else {
			w.Write([]byte(statusMap[0].ToJson()))
		}
	}
}
