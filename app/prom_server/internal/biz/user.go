package biz

import (
	"context"
	"strconv"

	query "github.com/aide-cloud/gorm-normalize"
	"github.com/go-kratos/kratos/v2/log"

	"prometheus-manager/api/perrors"
	"prometheus-manager/pkg/after"
	"prometheus-manager/pkg/helper/consts"
	"prometheus-manager/pkg/helper/middler"
	"prometheus-manager/pkg/helper/model"
	"prometheus-manager/pkg/helper/model/systemscopes"
	"prometheus-manager/pkg/helper/valueobj"
	"prometheus-manager/pkg/util/password"
	"prometheus-manager/pkg/util/slices"

	"prometheus-manager/app/prom_server/internal/biz/bo"
	"prometheus-manager/app/prom_server/internal/biz/repository"
)

type (
	UserBiz struct {
		log *log.Helper

		userRepo repository.UserRepo
		roleRepo repository.RoleRepo
		dataRepo repository.DataRepo
	}
)

func NewUserBiz(
	userRepo repository.UserRepo,
	dataRepo repository.DataRepo,
	roleRepo repository.RoleRepo,
	logger log.Logger,
) *UserBiz {
	return &UserBiz{
		log:      log.NewHelper(logger),
		userRepo: userRepo,
		dataRepo: dataRepo,
		roleRepo: roleRepo,
	}
}

// GetUserInfoById 获取用户信息
func (b *UserBiz) GetUserInfoById(ctx context.Context, id uint32) (*bo.UserBO, error) {
	userBo, err := b.userRepo.Get(ctx, systemscopes.UserInIds(id), systemscopes.UserPreloadRoles())
	if err != nil {
		return nil, err
	}
	return userBo, nil
}

// CheckNewUser 检查新用户信息
func (b *UserBiz) CheckNewUser(ctx context.Context, userBo *bo.UserBO) error {
	if userBo == nil {
		return perrors.ErrorInvalidParams("用户信息不能为空")
	}

	wheres := []query.ScopeMethod{
		systemscopes.UserEqName(userBo.Username),
		systemscopes.UserEqEmail(userBo.Email),
		systemscopes.UserEqPhone(userBo.Phone),
	}
	list, err := b.userRepo.Find(ctx, wheres...)
	if err != nil {
		return err
	}

	for _, v := range list {
		if v.Id == userBo.Id {
			continue
		}

		if v.Username == userBo.Username {
			return perrors.ErrorInvalidParams("用户名已存在")
		}
		if v.Email == userBo.Email {
			return perrors.ErrorInvalidParams("邮箱已存在")
		}
		if v.Phone == userBo.Phone {
			return perrors.ErrorInvalidParams("手机号已存在")
		}
	}
	return nil
}

// CreateUser 创建用户
func (b *UserBiz) CreateUser(ctx context.Context, userBo *bo.UserBO) (*bo.UserBO, error) {
	var err error
	userBo.Salt = password.GenerateSalt()
	userBo.Password, err = password.GeneratePassword(userBo.Password, userBo.Salt)
	if err != nil {
		return nil, err
	}

	userBo, err = b.userRepo.Create(ctx, userBo)
	if err != nil {
		return nil, err
	}

	return userBo, nil
}

// UpdateUserById 更新用户信息
func (b *UserBiz) UpdateUserById(ctx context.Context, id uint32, userBo *bo.UserBO) (*bo.UserBO, error) {
	userBo, err := b.userRepo.Update(ctx, userBo, systemscopes.RoleInIds(id))
	if err != nil {
		return nil, err
	}
	return userBo, nil
}

// UpdateUserStatusById 更新用户状态
func (b *UserBiz) UpdateUserStatusById(ctx context.Context, status valueobj.Status, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	userBo := &bo.UserBO{Status: status}
	_, err := b.userRepo.Update(ctx, userBo, systemscopes.RoleInIds(ids...))
	return err
}

// DeleteUserByIds 删除用户
func (b *UserBiz) DeleteUserByIds(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return b.userRepo.Delete(ctx, systemscopes.RoleInIds(ids...))
}

// GetUserList 获取用户列表
func (b *UserBiz) GetUserList(ctx context.Context, pgInfo query.Pagination, scopes ...query.ScopeMethod) ([]*bo.UserBO, error) {
	userBos, err := b.userRepo.List(ctx, pgInfo, scopes...)
	if err != nil {
		return nil, err
	}
	return userBos, nil
}

// LoginByUsernameAndPassword 登录
func (b *UserBiz) LoginByUsernameAndPassword(ctx context.Context, username, pwd string) (userBO *bo.UserBO, token string, err error) {
	userBO, err = b.userRepo.Get(ctx, systemscopes.UserEqName(username), systemscopes.UserPreloadRoles())
	if err != nil {
		return
	}

	if err = password.ValidatePasswordErr(pwd, userBO.Password, userBO.Salt); err != nil {
		return
	}

	// 没有角色
	if len(userBO.Roles) == 0 {
		token, err = middler.IssueToken(userBO.Id, "")
		return
	}

	// 获取上次默认角色
	key := consts.UserRoleKey.KeyUint32(userBO.Id).String()
	client, err := b.dataRepo.Client()
	if err != nil {
		b.log.Error(err)
		err = perrors.ErrorUnknown("系统错误")
		return
	}

	cacheRoleIdStr := client.Get(ctx, key).String()
	searchRole := func(roleInfo *bo.RoleBO) bool {
		cacheRoleId, _ := strconv.Atoi(cacheRoleIdStr)
		return roleInfo.Id == uint32(cacheRoleId)
	}
	// 如果上次默认角色还在角色列表中
	if slices.ContainsOf(userBO.Roles, searchRole) {
		token, err = middler.IssueToken(userBO.Id, cacheRoleIdStr)
		return
	}

	roleId := userBO.Roles[0].Id
	roleIdStr := strconv.Itoa(int(roleId))
	token, err = middler.IssueToken(userBO.Id, roleIdStr)
	return
}

// Logout 退出登录
func (b *UserBiz) Logout(ctx context.Context, authClaims *middler.AuthClaims) error {
	client, err := b.dataRepo.Client()
	if err != nil {
		return err
	}
	return middler.Expire(ctx, client, authClaims)
}

// RefreshToken 刷新token
func (b *UserBiz) RefreshToken(ctx context.Context, authClaims *middler.AuthClaims, roleId uint32) (userBO *bo.UserBO, token string, err error) {
	roleIdStr := strconv.Itoa(int(roleId))
	defer func() {
		key := consts.UserRoleKey.KeyUint32(authClaims.ID).String()
		client, err := b.dataRepo.Client()
		if err != nil {
			b.log.Error(err)
			return
		}
		if err = client.Set(ctx, key, roleIdStr, 0).Err(); err != nil {
			b.log.Errorf("cache user role err: %v", err)
			return
		}
	}()

	userBO, err = b.userRepo.Get(context.Background(), systemscopes.UserInIds(authClaims.ID), systemscopes.UserPreloadRoles())
	if err != nil {
		err = perrors.ErrorUnknown("系统错误")
		return
	}

	// 如果用户没有可用角色, 则直接置空处理
	if len(userBO.Roles) == 0 {
		roleIdStr = ""
		token, err = middler.IssueToken(authClaims.ID, roleIdStr)
		return
	}

	// 更改角色成功
	compareFun := func(do *bo.RoleBO) bool {
		return do.Id == roleId
	}

	// 切换的角色不存在, 则检查已有角色和token内角色
	if !slices.ContainsOf(userBO.Roles, compareFun) {
		compareFunCurrRoleId := func(do *bo.RoleBO) bool {
			currRoleId, _ := strconv.Atoi(authClaims.Role)
			return do.Id == uint32(currRoleId)
		}
		// 先默认为token内的角色
		roleIdStr = authClaims.Role
		// 如果token的角色不在已有的角色列表中, 则默认第一个角色
		if !slices.ContainsOf(userBO.Roles, compareFunCurrRoleId) {
			roleIdStr = strconv.Itoa(int(userBO.Roles[0].Id))
		}
	}

	token, err = middler.IssueToken(authClaims.ID, roleIdStr)
	return
}

// EditUserPassword 修改密码
func (b *UserBiz) EditUserPassword(ctx context.Context, authClaims *middler.AuthClaims, oldPassword, newPassword string) (*bo.UserBO, error) {
	userDo, err := b.userRepo.Get(ctx, systemscopes.UserInIds(authClaims.ID))
	if err != nil {
		return nil, err
	}
	// 验证旧密码
	if err = password.ValidatePasswordErr(oldPassword, userDo.Password, userDo.Salt); err != nil {
		return nil, err
	}

	// 加密新密码
	if userDo.Password, err = password.GeneratePassword(newPassword, userDo.Salt); err != nil {
		return nil, err
	}

	newUserDo := &bo.UserBO{
		Id:       userDo.Id,
		Password: userDo.Password,
	}

	// 更新密码
	if newUserDo, err = b.userRepo.Update(ctx, newUserDo, systemscopes.UserInIds(userDo.Id)); err != nil {
		return nil, err
	}

	return newUserDo, nil
}

// UpdateUserRoleRelation 更新用户角色关系
func (b *UserBiz) UpdateUserRoleRelation(userId uint32) {
	go func() {
		defer after.Recover(b.log)
		db, err := b.dataRepo.DB()
		if err != nil {
			b.log.Errorf("cache user role err: %v", err)
			return
		}
		client, err := b.dataRepo.Client()
		if err != nil {
			b.log.Errorf("cache user role err: %v", err)
			return
		}
		if err = model.CacheUserRole(db, client, userId); err != nil {
			b.log.Errorf("cache user role err: %v", err)
			return
		}
	}()
}

// RelateRoles 关联角色
func (b *UserBiz) RelateRoles(ctx context.Context, userId uint32, roleIds []uint32) error {
	defer b.UpdateUserRoleRelation(userId)
	userDo, err := b.userRepo.Get(ctx, systemscopes.UserInIds(userId))
	if err != nil {
		return err
	}

	// 查询角色
	if len(roleIds) > 0 {
		roleDos, err := b.roleRepo.Find(ctx, systemscopes.RoleInIds(roleIds...))
		if err != nil {
			return err
		}
		userDo.Roles = roleDos
	}

	// 关联角色
	if err = b.userRepo.RelateRoles(ctx, userDo, userDo.Roles); err != nil {
		return err
	}

	return nil
}
