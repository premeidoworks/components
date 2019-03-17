package kanatasupport

import (
	"github.com/premeidoworks/kanata/api"

	"github.com/gogo/protobuf/proto"
)

type gogoprotobufMarshalImpl struct {
}

func (gogoprotobufMarshalImpl) MarshalAcquireRequest(a *api.AcquireRequest) ([]byte, error) {
	ar := &AcquireRequest{
		Queue: a.Queue,
	}
	return proto.Marshal(ar)
}

func (gogoprotobufMarshalImpl) UnmarshalAcquireRequest(data []byte) (*api.AcquireRequest, error) {
	ar := new(AcquireRequest)
	err := proto.Unmarshal(data, ar)
	if err != nil {
		return nil, err
	}
	return &api.AcquireRequest{
		Queue: ar.Queue,
	}, nil
}

func (gogoprotobufMarshalImpl) MarshalAcquireResponse(a *api.AcquireResponse) ([]byte, error) {
	ar := &AcquireResponse{
		MessageList: make([]*Message, len(a.MessageList)),
	}
	for i := 0; i < len(a.MessageList); i++ {
		ar.MessageList[i] = &Message{
			MsgBody: a.MessageList[i].MsgBody,
			MsgId:   &a.MessageList[i].MsgId,
		}
	}
	return proto.Marshal(ar)
}

func (gogoprotobufMarshalImpl) UnmarshalAcquireResponse(data []byte) (*api.AcquireResponse, error) {
	ar := new(AcquireResponse)
	err := proto.Unmarshal(data, ar)
	if err != nil {
		return nil, err
	}
	result := &api.AcquireResponse{
		MessageList: make([]*struct {
			MsgId   string
			MsgBody []byte
		}, len(ar.MessageList)),
	}
	for i := 0; i < len(ar.MessageList); i++ {
		ar.MessageList[i] = &Message{
			MsgBody: ar.MessageList[i].MsgBody,
			MsgId:   ar.MessageList[i].MsgId,
		}
	}
	return result, nil
}

func (gogoprotobufMarshalImpl) MarshalPublishRequest(p *api.PublishRequest) ([]byte, error) {
	pr := &PublishRequest{
		Topic:       p.Topic,
		MessageList: make([]*Message, len(p.MessageList)),
	}
	for i := 0; i < len(pr.MessageList); i++ {
		pr.MessageList[i] = &Message{
			MsgBody:  p.MessageList[i].MsgBody,
			MsgId:    &p.MessageList[i].MsgId,
			MsgOutId: &p.MessageList[i].MsgOutId,
		}
	}
	return proto.Marshal(pr)
}

func (gogoprotobufMarshalImpl) UnmarshalPublishRequest(data []byte) (*api.PublishRequest, error) {
	pr := new(PublishRequest)
	err := proto.Unmarshal(data, pr)
	if err != nil {
		return nil, err
	}
	result := &api.PublishRequest{
		Topic: pr.Topic,
		MessageList: make([]*struct {
			MsgId    string
			MsgOutId string
			MsgBody  []byte
		}, len(pr.MessageList)),
	}
	for i := 0; i < len(pr.MessageList); i++ {
		result.MessageList[i] = &struct {
			MsgId    string
			MsgOutId string
			MsgBody  []byte
		}{
			MsgBody:  pr.MessageList[i].MsgBody,
			MsgId:    *pr.MessageList[i].MsgId,
			MsgOutId: *pr.MessageList[i].MsgOutId,
		}
	}
	return result, nil
}

func (gogoprotobufMarshalImpl) MarshalPublishResponse(p *api.PublishResponse) ([]byte, error) {
	pr := &PublishResponse{
		SuccessIdList: make([]*SuccessMessageId, len(p.SuccessIdList)),
		FailIdList:    make([]*FailMessageId, len(p.FailIdList)),
	}
	for i := 0; i < len(pr.SuccessIdList); i++ {
		pr.SuccessIdList[i] = &SuccessMessageId{
			MsgOutId: &p.SuccessIdList[i].MsgOutId,
			MsgId:    &p.SuccessIdList[i].MsgId,
		}
	}
	for i := 0; i < len(pr.FailIdList); i++ {
		pr.FailIdList[i] = &FailMessageId{
			MsgOutId: &p.FailIdList[i].MsgOutId,
			MsgId:    &p.FailIdList[i].MsgId,
		}
	}
	return proto.Marshal(pr)
}

func (gogoprotobufMarshalImpl) UnmarshalPublishResponse(data []byte) (*api.PublishResponse, error) {
	pr := new(PublishResponse)
	err := proto.Unmarshal(data, pr)
	if err != nil {
		return nil, err
	}
	result := &api.PublishResponse{
		SuccessIdList: make([]*struct {
			MsgId    string
			MsgOutId string
		}, len(pr.SuccessIdList)),
		FailIdList: make([]*struct {
			MsgId    string
			MsgOutId string
			Code     string
		}, len(pr.FailIdList)),
	}
	for i := 0; i < len(pr.SuccessIdList); i++ {
		result.SuccessIdList[i] = &struct {
			MsgId    string
			MsgOutId string
		}{
			MsgId:    *pr.SuccessIdList[i].MsgId,
			MsgOutId: *pr.SuccessIdList[i].MsgOutId,
		}
	}
	for i := 0; i < len(pr.FailIdList); i++ {
		result.FailIdList[i] = &struct {
			MsgId    string
			MsgOutId string
			Code     string
		}{
			MsgOutId: *pr.FailIdList[i].MsgOutId,
			MsgId:    *pr.FailIdList[i].MsgId,
			Code:     *pr.FailIdList[i].Code,
		}
	}
	return result, nil
}
