package model

import "context"

type GroupModelChannelRepository struct {
	GetRandomSatisfiedChannel                func(group string, model string, ignoreFirstPriority bool) (*Channel, error)
	ListSatisfiedChannels                    func(group string, model string) ([]*Channel, error)
	AddGroupModelChannels                    func(channel *Channel) error
	DeleteGroupModelChannels                 func(channel *Channel) error
	UpdateGroupModelChannels                 func(channel *Channel) error
	RefreshGroupModelChannelsByChannelStatus func(channelId string, status bool) error
	GetTopChannelByModel                     func(group string, model string) (*Channel, error)
	GetGroupModels                           func(ctx context.Context, group string) ([]string, error)
}

var groupModelChannelRepo GroupModelChannelRepository

func BindGroupModelChannelRepository(repo GroupModelChannelRepository) {
	groupModelChannelRepo = repo
}

func mustGroupModelChannelRepo() GroupModelChannelRepository {
	if groupModelChannelRepo.GetRandomSatisfiedChannel == nil {
		panic("group model channel repository not initialized")
	}
	if groupModelChannelRepo.ListSatisfiedChannels == nil {
		panic("group model channel repository not initialized")
	}
	return groupModelChannelRepo
}
