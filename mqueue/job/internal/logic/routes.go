package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"zero-shop/mqueue/job/internal/svc"
	"zero-shop/mqueue/job/jobtype"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	// 延迟任务队列
	mux.Handle(jobtype.DeferCloseProductOrder, NewCloseProductOrderHandler(l.svcCtx))
	mux.Handle(jobtype.DeferCloseSeckillOrder, NewCloseSeckillOrderHandler(l.svcCtx))

	return mux
}
