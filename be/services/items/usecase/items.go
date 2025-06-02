package usecase

import (
	"be/pkg/db"
	"be/services/items/model/entity"
)

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

func GetAllItems()([]entity.Item, error) {
	var items []entity.Item
	result := db.DB.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func GetItemById(itemId string) (*entity.Item, error) {
	var item entity.Item
	result := db.DB.First(&item, "id = ?", itemId)
	if result.Error != nil {
		return nil, result.Error
	}
	db.DB.Model(&item).Update("view", item.View+1)
	return &item, nil
}