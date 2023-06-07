package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	validate *validator.Validate
}
 
func NewTagsServiceImpl(tagRepository repository.TagsRepository,validate *validator.Validate) TagsService{
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		validate: validate,
	}
}


//create implements TagsService 
func(t *TagsServiceImpl) Create (tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

//Delete implements TagsService
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

//Find all implements TagsService
func(t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagsResponse
	for _, value := range result{
		tag := response.TagsResponse{
				Id : value.Id,
				Name: value.Name,
		}
		tags=append(tags, tag)
	}
	return tags
}

//Find by ID implements TagsService
func(t *TagsServiceImpl) FindById(tagsId int ) response.TagsResponse{
	tagData,err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id: tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}


//update implements TagsService
func(t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData,err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}