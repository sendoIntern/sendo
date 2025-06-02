package usecase

// func (r rewardAnalyticService) CreateRewardAnalytic(ctx context.Context, data entity.MissionRewardCommission) error {
// 	logContext := logger.EnhanceWith(ctx)
// 	logContext.Infof("CreateRewardAnalytic start %+v", data)

// 	if r.rewardAnalyticsRepo == nil {
// 		logContext.Error("CreateRewardAnalytic rewardAnalyticsRepo is nil")
// 		return errors.New("CreateRewardAnalytic rewardAnalyticsRepo is nil")
// 	}

// 	if err := r.rewardAnalyticsRepo.CreateReward(ctx, data); err != nil {
// 		logContext.Errorw("CreateRewardAnalytic error", "error", err, "data", data)
// 		return err
// 	}

// 	logContext.Infof("CreateRewardAnalytic success: %s", utils.DumpJson(data))
// 	return nil
// }
