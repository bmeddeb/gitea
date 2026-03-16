// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repository

import (
	"os"
	"strconv"
	"strings"

	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/setting"
)

// env keys for git hooks need (GITFX_ is the primary prefix, GITEA_ is kept for backward compatibility)
const (
	EnvRepoName     = "GITFX_REPO_NAME"
	EnvRepoUsername = "GITFX_REPO_USER_NAME"
	EnvRepoID       = "GITFX_REPO_ID"
	EnvRepoIsWiki   = "GITFX_REPO_IS_WIKI"
	EnvPusherName   = "GITFX_PUSHER_NAME"
	EnvPusherEmail  = "GITFX_PUSHER_EMAIL"
	EnvPusherID     = "GITFX_PUSHER_ID"
	EnvKeyID        = "GITFX_KEY_ID" // public key ID
	EnvDeployKeyID  = "GITFX_DEPLOY_KEY_ID"
	EnvPRID         = "GITFX_PR_ID"
	EnvPRIndex      = "GITFX_PR_INDEX" // not used by GitFX at the moment, it is for custom git hooks
	EnvPushTrigger  = "GITFX_PUSH_TRIGGER"
	EnvIsInternal   = "GITFX_INTERNAL_PUSH"
	EnvAppURL       = "GITFX_ROOT_URL"
	EnvActionPerm   = "GITFX_ACTION_PERM"
)

// Deprecated GITEA_* env var names, kept for backward compatibility with custom hooks
const (
	EnvRepoNameLegacy     = "GITEA_REPO_NAME"
	EnvRepoUsernameLegacy = "GITEA_REPO_USER_NAME"
	EnvRepoIDLegacy       = "GITEA_REPO_ID"
	EnvRepoIsWikiLegacy   = "GITEA_REPO_IS_WIKI"
	EnvPusherNameLegacy   = "GITEA_PUSHER_NAME"
	EnvPusherEmailLegacy  = "GITEA_PUSHER_EMAIL"
	EnvPusherIDLegacy     = "GITEA_PUSHER_ID"
	EnvKeyIDLegacy        = "GITEA_KEY_ID"
	EnvDeployKeyIDLegacy  = "GITEA_DEPLOY_KEY_ID"
	EnvPRIDLegacy         = "GITEA_PR_ID"
	EnvPRIndexLegacy      = "GITEA_PR_INDEX"
	EnvPushTriggerLegacy  = "GITEA_PUSH_TRIGGER"
	EnvIsInternalLegacy   = "GITEA_INTERNAL_PUSH"
	EnvAppURLLegacy       = "GITEA_ROOT_URL"
	EnvActionPermLegacy   = "GITEA_ACTION_PERM"
)

type PushTrigger string

const (
	PushTriggerPRMergeToBase    PushTrigger = "pr-merge-to-base"
	PushTriggerPRUpdateWithBase PushTrigger = "pr-update-with-base"
)

// InternalPushingEnvironment returns an os environment to switch off hooks on push
// It is recommended to avoid using this unless you are pushing within a transaction
// or if you absolutely are sure that post-receive and pre-receive will do nothing
// We provide the full pushing-environment for other hook providers
func InternalPushingEnvironment(doer *user_model.User, repo *repo_model.Repository) []string {
	return append(PushingEnvironment(doer, repo),
		EnvIsInternal+"=true",
		EnvIsInternalLegacy+"=true",
	)
}

// PushingEnvironment returns an os environment to allow hooks to work on push
func PushingEnvironment(doer *user_model.User, repo *repo_model.Repository) []string {
	return FullPushingEnvironment(doer, doer, repo, repo.Name, 0, 0)
}

// FullPushingEnvironment returns an os environment to allow hooks to work on push
func FullPushingEnvironment(author, committer *user_model.User, repo *repo_model.Repository, repoName string, prID, prIndex int64) []string {
	isWiki := "false"
	if strings.HasSuffix(repoName, ".wiki") {
		isWiki = "true"
	}

	authorSig := author.NewGitSig()
	committerSig := committer.NewGitSig()

	environ := append(os.Environ(),
		"GIT_AUTHOR_NAME="+authorSig.Name,
		"GIT_AUTHOR_EMAIL="+authorSig.Email,
		"GIT_COMMITTER_NAME="+committerSig.Name,
		"GIT_COMMITTER_EMAIL="+committerSig.Email,
		// Set new GITFX_ env vars
		EnvRepoName+"="+repoName,
		EnvRepoUsername+"="+repo.OwnerName,
		EnvRepoIsWiki+"="+isWiki,
		EnvPusherName+"="+committer.Name,
		EnvPusherID+"="+strconv.FormatInt(committer.ID, 10),
		EnvRepoID+"="+strconv.FormatInt(repo.ID, 10),
		EnvPRID+"="+strconv.FormatInt(prID, 10),
		EnvPRIndex+"="+strconv.FormatInt(prIndex, 10),
		EnvAppURL+"="+setting.AppURL,
		// Also set deprecated GITEA_ env vars for backward compatibility with custom hooks
		EnvRepoNameLegacy+"="+repoName,
		EnvRepoUsernameLegacy+"="+repo.OwnerName,
		EnvRepoIsWikiLegacy+"="+isWiki,
		EnvPusherNameLegacy+"="+committer.Name,
		EnvPusherIDLegacy+"="+strconv.FormatInt(committer.ID, 10),
		EnvRepoIDLegacy+"="+strconv.FormatInt(repo.ID, 10),
		EnvPRIDLegacy+"="+strconv.FormatInt(prID, 10),
		EnvPRIndexLegacy+"="+strconv.FormatInt(prIndex, 10),
		EnvAppURLLegacy+"="+setting.AppURL,
		"SSH_ORIGINAL_COMMAND=gitea-internal",
	)

	if !committer.KeepEmailPrivate {
		environ = append(environ,
			EnvPusherEmail+"="+committer.Email,
			EnvPusherEmailLegacy+"="+committer.Email,
		)
	}

	return environ
}
