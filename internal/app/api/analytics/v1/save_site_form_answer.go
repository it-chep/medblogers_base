package v1

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"
	"medblogers_base/internal/modules/analytics/actions/save_site_form_answer/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/analytics/v1"
)

// SaveSiteFormAnswer сохраняет ответ формы сайта.
func (i *Implementation) SaveSiteFormAnswer(ctx context.Context, req *desc.SaveSiteFormAnswerRequest) (*desc.SaveSiteFormAnswerResponse, error) {
	answer, err := protojson.Marshal(req.GetAnswer())
	if err != nil {
		return nil, nil
	}

	err = i.analytics.Actions.SaveSiteFormAnswer.Do(ctx, dto.SaveSiteFormAnswerRequest{
		FormName: req.GetFormName(),
		Answer:   answer,
		CookieID: req.GetCookieId(),
		Source:   req.GetSource(),
		TG:       req.Tg,
	})
	if err != nil {
		return nil, nil
	}

	return &desc.SaveSiteFormAnswerResponse{}, nil
}
