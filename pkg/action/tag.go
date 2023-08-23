package action

import (
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	getTagsAction   = "getTags"
	addTagAction    = "addTag"
	editTagAction   = "editTag"
	deleteTagAction = "deleteTag"
)

type getTagsResponse struct {
	Ok   bool        `mapstructure:"ok"`
	Msg  *string     `mapstructure:"msg"`
	Tags []state.Tag `mapstructure:"tags"`
}

type addTagRequest struct {
	New   bool   `json:"new"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type addTagResponse struct {
	Ok  bool       `mapstructure:"ok"`
	Msg *string    `mapstructure:"msg"`
	Tag *state.Tag `mapstructure:"tag"`
}

type editTagRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type editTagResponse struct {
	Ok  bool       `mapstructure:"ok"`
	Msg *string    `mapstructure:"msg"`
	Tag *state.Tag `mapstructure:"tag"`
}

type deleteTagResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

func GetTags(c StatefulEmiter) ([]state.Tag, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return nil, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	response, err := c.Emit(getTagsAction, defaultEmitTimeout)
	if err != nil {
		return nil, NewErrActionFailed(getTagsAction, err.Error())
	}

	// unmarshal data
	data := &getTagsResponse{}
	if err := decode(response, data); err != nil {
		return nil, NewErrActionFailed(getTagsAction, err.Error())
	}

	// check if the response is ok
	if !data.Ok {
		return nil, NewErrActionFailed(getTagsAction, *data.Msg)
	}

	// set tags in state
	err = c.State().SetTags(data.Tags)
	if err != nil {
		return nil, NewErrActionFailed(getTagsAction, err.Error())
	}

	return data.Tags, nil
}

func AddTag(c StatefulEmiter, name string, color string) (*state.Tag, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return nil, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(addTagAction, defaultEmitTimeout, &addTagRequest{
		New:   true,
		Name:  name,
		Color: color,
	})
	if err != nil {
		return nil, NewErrActionFailed(addTagAction, err.Error())
	}

	// unmarshal data
	data := &addTagResponse{}
	if err := decode(response, data); err != nil {
		return nil, NewErrActionFailed(addTagAction, err.Error())
	}

	// check if the response is ok
	if !data.Ok {
		return nil, NewErrActionFailed(addTagAction, *data.Msg)
	}

	// set tag in state
	err = c.State().SetTag(data.Tag)
	if err != nil {
		return nil, NewErrActionFailed(addTagAction, err.Error())
	}

	return data.Tag, nil
}

func EditTag(c StatefulEmiter, id int, name, color string) (*state.Tag, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return nil, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(editTagAction, defaultEmitTimeout, &editTagRequest{
		Id:    id,
		Name:  name,
		Color: color,
	})
	if err != nil {
		return nil, NewErrActionFailed(editTagAction, err.Error())
	}

	// unmarshal data
	data := &editTagResponse{}
	if err := decode(response, data); err != nil {
		return nil, NewErrActionFailed(editTagAction, err.Error())
	}

	// check if the response is ok
	if !data.Ok {
		return nil, NewErrActionFailed(editTagAction, *data.Msg)
	}

	// set tag in state
	err = c.State().SetTag(data.Tag)
	if err != nil {
		return nil, NewErrActionFailed(editTagAction, err.Error())
	}

	return data.Tag, nil
}

func DeleteTag(c StatefulEmiter, id int) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(deleteTagAction, defaultEmitTimeout, id)
	if err != nil {
		return NewErrActionFailed(deleteTagAction, err.Error())
	}

	// unmarshal data
	data := &deleteTagResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(deleteTagAction, err.Error())
	}

	// check if the response is ok
	if !data.Ok {
		return NewErrActionFailed(deleteTagAction, *data.Msg)
	}

	// delete tag from state
	err = c.State().DeleteTag(id)
	if err != nil {
		return NewErrActionFailed(deleteTagAction, err.Error())
	}

	return nil
}
