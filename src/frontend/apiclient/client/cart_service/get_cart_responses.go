// Code generated by go-swagger; DO NOT EDIT.

package cart_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/siriscac/microservices-demo/fontend/apiclient/models"
)

// GetCartReader is a Reader for the GetCart structure.
type GetCartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCartOK creates a GetCartOK with default headers values
func NewGetCartOK() *GetCartOK {
	return &GetCartOK{}
}

/*GetCartOK handles this case with default header values.

A successful response.
*/
type GetCartOK struct {
	Payload *models.HipstershopCart
}

func (o *GetCartOK) Error() string {
	return fmt.Sprintf("[GET /carts/{user_id}][%d] getCartOK  %+v", 200, o.Payload)
}

func (o *GetCartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HipstershopCart)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
