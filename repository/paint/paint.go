package paint

import "painteer/model"

type PaintRepository interface {
	FindPostIDsByPrefecture(groupId model.GroupId) ([]model.PostsByPrefecture, error)
}
