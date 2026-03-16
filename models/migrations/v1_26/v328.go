// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_26

import (
	"xorm.io/xorm"
)

func RenameUserThemesGiteaToGitfx(x *xorm.Engine) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if _, err := sess.Exec("UPDATE `user` SET `theme` = REPLACE(`theme`, 'gitea-', 'gitfx-') WHERE `theme` LIKE 'gitea-%'"); err != nil {
		return err
	}

	return sess.Commit()
}
