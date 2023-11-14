package authority

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/utils/entx"

	"github.com/kubeTasker/kubeTaskerServer/rpc/ent"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/utils/errorhandler"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		err := tx.Role.UpdateOneID(in.RoleId).ClearMenus().Exec(l.ctx)
		if err != nil {
			return err
		}

		err = tx.Role.UpdateOneID(in.RoleId).AddMenuIDs(in.MenuId...).Exec(l.ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
