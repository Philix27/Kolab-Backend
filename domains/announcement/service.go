package announcement

import (
	"harmony/libs/helper"

	"github.com/go-playground/validator"
)

type Service struct {
	repository iRepository
	validate   *validator.Validate
}

func NewService(repository iRepository, validate *validator.Validate) iService {
	return &Service{
		repository: repository,
		validate:   validate,
	}
}

// Create implements iService.
func (s *Service) Create(data createAnnouncementDto) {
	err := s.validate.Struct(data)
	helper.ErrorPanic(err, "Create announcement service")

	model := AnnouncementM{
		Title:    data.Title,
		Subtitle: data.Subtitle,
	}

	s.repository.Create(model)
}

// Delete implements iService.
func (s *Service) Delete(dataId int) {
	s.repository.Delete(dataId)
}

// FindAll implements iService.
func (s *Service) FindAll() (list []announcementResponseDto) {
	result := s.repository.FindAll()

	var announceList []announcementResponseDto

	for _, val := range result {
		an := announcementResponseDto{
			Id:    int(val.ID),
			Title: val.Title,
		}
		announceList = append(announceList, an)
	}

	return announceList
}

// FindById implements iService.
func (s *Service) FindById(dataId int) (data announcementResponseDto, err error) {
	announceData, err := s.repository.FindById(dataId)
	helper.ErrorPanic(err, "FindById announcement service")
	response := announcementResponseDto{
		Id:       int(announceData.ID),
		Title:    announceData.Title,
		Subtitle: announceData.Subtitle,
	}

	return response, err
}

// Update implements iService.
func (s *Service) Update(data updateAnnouncementDto) {
	announceData, err := s.repository.FindById(data.Id)
	helper.ErrorPanic(err, "Update announcement service")
	announceData.Title = data.Title
	announceData.Subtitle = data.Subtitle

	s.repository.Update(announceData)
}
