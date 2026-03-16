// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_26

import (
	"xorm.io/xorm"
)

func RenameActionsUser(x *xorm.Engine) error {
	// Update the actions system user (id=-2) to use the new GitFX branding.
	// This user is normally a virtual/in-memory user, but some installations
	// may have persisted it in the database (e.g., via older migrations or
	// direct inserts). Update it if it exists.
	_, err := x.Exec("UPDATE `user` SET `name` = 'gitfx-actions', `lower_name` = 'gitfx-actions', `full_name` = 'GitFX Actions', `email` = 'actions@gitfx.com' WHERE `id` = -2")
	return err
}
