package dto

type SubscribersInfo struct {
	Doctors      map[int64]DoctorSubscribersInfoDTO
	DoctorsCount int64
	SubsCount    string
	OrderedIDs   []int64
}

type DoctorSubscribersInfoDTO struct {
	InstSubsCount        string
	InstSubsCountText    string
	TgSubsCount          string
	TgSubsCountText      string
	YouTubeSubsCount     string
	YouTubeSubsCountText string
}
