package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTagFetchFollowersWithRelationshipIDReader is a Reader for the PostRemoteAPIJTagFetchFollowersWithRelationshipID structure.
type PostRemoteAPIJTagFetchFollowersWithRelationshipIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDOK creates a PostRemoteAPIJTagFetchFollowersWithRelationshipIDOK with default headers values
func NewPostRemoteAPIJTagFetchFollowersWithRelationshipIDOK() *PostRemoteAPIJTagFetchFollowersWithRelationshipIDOK {
	return &PostRemoteAPIJTagFetchFollowersWithRelationshipIDOK{}
}

/*PostRemoteAPIJTagFetchFollowersWithRelationshipIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJTagFetchFollowersWithRelationshipIDOK struct {
	Payload PostRemoteAPIJTagFetchFollowersWithRelationshipIDOKBody
}

func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.fetchFollowersWithRelationship/{id}][%d] postRemoteApiJTagFetchFollowersWithRelationshipIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJTagFetchFollowersWithRelationshipIDOKBody post remote API j tag fetch followers with relationship ID o k body
swagger:model PostRemoteAPIJTagFetchFollowersWithRelationshipIDOKBody
*/
type PostRemoteAPIJTagFetchFollowersWithRelationshipIDOKBody struct {
	models.JTag

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO0 models.JTag
	if err := swag.ReadJSON(raw, &postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO0); err != nil {
		return err
	}
	o.JTag = postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO0

	var postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJTagFetchFollowersWithRelationshipIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO0, err := swag.WriteJSON(o.JTag)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO0)

	postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJTagFetchFollowersWithRelationshipIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j tag fetch followers with relationship ID o k body
func (o *PostRemoteAPIJTagFetchFollowersWithRelationshipIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JTag.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
