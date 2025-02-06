package paint

import "painteer/model"

type PaintRepository interface {
	CountPostsByPrefecture(groupId model.GroupId) (*model.Count, error)
}
