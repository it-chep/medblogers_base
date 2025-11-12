package dto

type Doctor struct {
	ID   int64
	Name string
	Slug string

	Speciality string // Строка из основной и дополнительных специальностей
	City       string // Строка из основного и дополнительных городов

	Image string

	MainCityID       int64
	MainSpecialityID int64

	S3Key string
}

type Response struct {
	Doctors []Doctor
}

func NewResponse(doctorsMap map[int64]Doctor, orderedIDs []int64) Response {
	mappedDoctors := make([]Doctor, 0, len(doctorsMap))

	for _, doctorID := range orderedIDs {
		doctorData := doctorsMap[doctorID]

		if doctorData.ID == 0 {
			continue
		}

		mappedDoctors = append(mappedDoctors, Doctor{
			ID:    doctorData.ID,
			Name:  doctorData.Name,
			Slug:  doctorData.Slug,
			Image: doctorData.Image,

			MainCityID:       doctorData.MainCityID,
			MainSpecialityID: doctorData.MainSpecialityID,

			Speciality: doctorData.Speciality,
			City:       doctorData.City,

			S3Key: doctorData.S3Key,
		})
	}

	return Response{Doctors: mappedDoctors}
}
