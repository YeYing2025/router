package group

import "github.com/yeying-community/router/internal/admin/model"

func ListCatalog() ([]model.GroupCatalog, error) {
	return model.ListGroupCatalog()
}

func ListPage(page int, pageSize int, keyword string) ([]model.GroupCatalog, int64, error) {
	return model.ListGroupCatalogPage(page, pageSize, keyword)
}

func Get(id string) (model.GroupCatalog, error) {
	return model.GetGroupCatalogByID(id)
}

func Create(item model.GroupCatalog) (model.GroupCatalog, error) {
	return model.CreateGroupCatalog(item)
}

func CreateWithChannels(item model.GroupCatalog, channelIDs []string) (model.GroupCatalog, error) {
	return model.CreateGroupCatalogWithChannels(item, channelIDs)
}

func CreateWithModels(item model.GroupCatalog, channelIDs []string, models []model.GroupModelBindingItem) (model.GroupCatalog, error) {
	return model.CreateGroupCatalogWithModels(item, channelIDs, models)
}

func Update(item model.GroupCatalog) (model.GroupCatalog, error) {
	return model.UpdateGroupCatalog(item)
}

func UpdateWithChannels(item model.GroupCatalog, channelIDs []string) (model.GroupCatalog, error) {
	return model.UpdateGroupCatalogWithChannels(item, channelIDs)
}

func UpdateWithModels(item model.GroupCatalog, channelIDs []string, models []model.GroupModelBindingItem, updateChannels bool, updateModels bool) (model.GroupCatalog, error) {
	return model.UpdateGroupCatalogWithModels(item, channelIDs, models, updateChannels, updateModels)
}

func Delete(id string) error {
	return model.DeleteGroupCatalog(id)
}

func ListChannels(id string) ([]model.GroupChannelItem, error) {
	return model.ListGroupChannels(id)
}

func ListModels(id string) (model.GroupModelsPayload, error) {
	return model.ListGroupModelsPayload(id)
}

func ReplaceChannels(id string, channelIDs []string) error {
	return model.ReplaceGroupChannels(id, channelIDs)
}

func ReplaceChannelsWithItems(id string, items []model.GroupChannelItem) error {
	return model.ReplaceGroupChannelsWithItems(id, items)
}

func ReplaceModels(id string, channelIDs []string, models []model.GroupModelBindingItem, explicitChannels bool) error {
	return model.ReplaceGroupModels(id, channelIDs, models, explicitChannels)
}

func ReplaceSingleModel(id string, modelName string, models []model.GroupModelBindingItem) error {
	return model.ReplaceSingleGroupModel(id, modelName, models)
}

func DeleteSingleModel(id string, modelName string) error {
	return model.DeleteSingleGroupModel(id, modelName)
}

func GetDailyQuotaSnapshot(id string, userID string, bizDate string) (model.GroupDailyQuotaSnapshot, error) {
	return model.GetGroupDailyQuotaSnapshot(id, userID, bizDate)
}
