// Code generated by go-swagger; DO NOT EDIT.

package recommendation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/siriscac/microservices-demo/fontend/apiclient/models"
)

// ListRecommendationsReader is a Reader for the ListRecommendations structure.
type ListRecommendationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRecommendationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRecommendationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRecommendationsOK creates a ListRecommendationsOK with default headers values
func NewListRecommendationsOK() *ListRecommendationsOK {
	return &ListRecommendationsOK{}
}

/*ListRecommendationsOK handles this case with default header values.

A successful response.
*/
type ListRecommendationsOK struct {
	Payload *models.HipstershopListRecommendationsResponse
}

func (o *ListRecommendationsOK) Error() string {
	return fmt.Sprintf("[GET /recommendations/{user_id}][%d] listRecommendationsOK  %+v", 200, o.Payload)
}

func (o *ListRecommendationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HipstershopListRecommendationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
