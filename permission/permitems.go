// AUTOMATICALLY GENERATED FILE - DO NOT EDIT!
// Please run 'go generate' to update this file.
//
// Copyright 2015 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package permission

var (
	PermAll                         = PermissionRegistry.get("")
	PermApp                         = PermissionRegistry.get("app")
	PermAppCreate                   = PermissionRegistry.get("app.create")
	PermAppDelete                   = PermissionRegistry.get("app.delete")
	PermAppDeploy                   = PermissionRegistry.get("app.deploy")
	PermAppDeployRollback           = PermissionRegistry.get("app.deploy.rollback")
	PermAppRead                     = PermissionRegistry.get("app.read")
	PermAppReadDeploy               = PermissionRegistry.get("app.read.deploy")
	PermAppReadEnv                  = PermissionRegistry.get("app.read.env")
	PermAppReadLog                  = PermissionRegistry.get("app.read.log")
	PermAppReadMetric               = PermissionRegistry.get("app.read.metric")
	PermAppRun                      = PermissionRegistry.get("app.run")
	PermAppUpdate                   = PermissionRegistry.get("app.update")
	PermAppUpdateCname              = PermissionRegistry.get("app.update.cname")
	PermAppUpdateCnameAdd           = PermissionRegistry.get("app.update.cname.add")
	PermAppUpdateCnameRemove        = PermissionRegistry.get("app.update.cname.remove")
	PermAppUpdateEnv                = PermissionRegistry.get("app.update.env")
	PermAppUpdateEnvSet             = PermissionRegistry.get("app.update.env.set")
	PermAppUpdateEnvUnset           = PermissionRegistry.get("app.update.env.unset")
	PermAppUpdateGrant              = PermissionRegistry.get("app.update.grant")
	PermAppUpdateLog                = PermissionRegistry.get("app.update.log")
	PermAppUpdatePool               = PermissionRegistry.get("app.update.pool")
	PermAppUpdateRestart            = PermissionRegistry.get("app.update.restart")
	PermAppUpdateRevoke             = PermissionRegistry.get("app.update.revoke")
	PermAppUpdateTeamowner          = PermissionRegistry.get("app.update.teamowner")
	PermAppUpdateUnit               = PermissionRegistry.get("app.update.unit")
	PermAppUpdateUnitAdd            = PermissionRegistry.get("app.update.unit.add")
	PermAppUpdateUnitRemove         = PermissionRegistry.get("app.update.unit.remove")
	PermIaas                        = PermissionRegistry.get("iaas")
	PermIaasRead                    = PermissionRegistry.get("iaas.read")
	PermNode                        = PermissionRegistry.get("node")
	PermNodeCreate                  = PermissionRegistry.get("node.create")
	PermNodeDelete                  = PermissionRegistry.get("node.delete")
	PermNodeRead                    = PermissionRegistry.get("node.read")
	PermNodeUpdate                  = PermissionRegistry.get("node.update")
	PermPlatform                    = PermissionRegistry.get("platform")
	PermPlatformCreate              = PermissionRegistry.get("platform.create")
	PermPlatformDelete              = PermissionRegistry.get("platform.delete")
	PermPlatformUpdate              = PermissionRegistry.get("platform.update")
	PermRole                        = PermissionRegistry.get("role")
	PermRoleAssign                  = PermissionRegistry.get("role.assign")
	PermRoleCreate                  = PermissionRegistry.get("role.create")
	PermRoleDelete                  = PermissionRegistry.get("role.delete")
	PermRoleUpdate                  = PermissionRegistry.get("role.update")
	PermServiceInstance             = PermissionRegistry.get("service-instance")
	PermServiceInstanceCreate       = PermissionRegistry.get("service-instance.create")
	PermServiceInstanceDelete       = PermissionRegistry.get("service-instance.delete")
	PermServiceInstanceRead         = PermissionRegistry.get("service-instance.read")
	PermServiceInstanceUpdate       = PermissionRegistry.get("service-instance.update")
	PermServiceInstanceUpdateBind   = PermissionRegistry.get("service-instance.update.bind")
	PermServiceInstanceUpdateGrant  = PermissionRegistry.get("service-instance.update.grant")
	PermServiceInstanceUpdateRevoke = PermissionRegistry.get("service-instance.update.revoke")
	PermServiceInstanceUpdateUnbind = PermissionRegistry.get("service-instance.update.unbind")
	PermTeam                        = PermissionRegistry.get("team")
	PermTeamCreate                  = PermissionRegistry.get("team.create")
	PermTeamDelete                  = PermissionRegistry.get("team.delete")
	PermTeamUpdate                  = PermissionRegistry.get("team.update")
	PermTeamUpdateAddMember         = PermissionRegistry.get("team.update.add-member")
	PermTeamUpdateRemoveMember      = PermissionRegistry.get("team.update.remove-member")
	PermUser                        = PermissionRegistry.get("user")
	PermUserCreate                  = PermissionRegistry.get("user.create")
	PermUserDelete                  = PermissionRegistry.get("user.delete")
	PermUserList                    = PermissionRegistry.get("user.list")
	PermUserUpdate                  = PermissionRegistry.get("user.update")
)