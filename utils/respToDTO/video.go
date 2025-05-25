package respToDTO

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/dto"
)

func GetVideoDTO(video *model.Video) *dto.VideoDTO {
	return &dto.VideoDTO{
		ID:            video.ID,
		UserID:        video.User.ID,
		User:          *GetUserDTO(&video.User),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}

func GetVideoListDTO(videos []model.Video) *[]dto.VideoDTO {
	videoInfo := make([]dto.VideoDTO, len(videos))
	for i := 0; i < len(videos); i++ {
		video := videos[i]
		videoInfo[i] = *GetVideoDTO(&video)
	}
	return &videoInfo
}
