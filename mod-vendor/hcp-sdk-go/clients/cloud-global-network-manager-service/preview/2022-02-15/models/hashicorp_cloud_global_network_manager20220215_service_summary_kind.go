// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind Kind is the kind of the service that was registered in the Catalog.
// https://developer.hashicorp.com/consul/api-docs/agent/service#kind
//
// - KIND_UNSPECIFIED: The default, unspecified service kind. Kind unknown.
//   - KIND_TYPICAL: A typical, classic Consul service.
//   - KIND_CONNECT_PROXY: A Connect proxy instance.
//   - KIND_MESH_GATEWAY: A mesh gateway service instance.
//   - KIND_TERMINATING_GATEWAY: A terminating gateway service instance.
//   - KIND_INGRESS_GATEWAY: An ingress gateway service instance.
//   - KIND_DESTINATION: A Destination  for the Connect feature.
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ServiceSummary.Kind
type HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind string

func NewHashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind(value HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind) *HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind.
func (m HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind) Pointer() *HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind {
	return &m
}

const (

	// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDUNSPECIFIED captures enum value "KIND_UNSPECIFIED"
	HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDUNSPECIFIED HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind = "KIND_UNSPECIFIED"

	// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDTYPICAL captures enum value "KIND_TYPICAL"
	HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDTYPICAL HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind = "KIND_TYPICAL"

	// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDCONNECTPROXY captures enum value "KIND_CONNECT_PROXY"
	HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDCONNECTPROXY HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind = "KIND_CONNECT_PROXY"

	// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDMESHGATEWAY captures enum value "KIND_MESH_GATEWAY"
	HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDMESHGATEWAY HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind = "KIND_MESH_GATEWAY"

	// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDTERMINATINGGATEWAY captures enum value "KIND_TERMINATING_GATEWAY"
	HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDTERMINATINGGATEWAY HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind = "KIND_TERMINATING_GATEWAY"

	// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDINGRESSGATEWAY captures enum value "KIND_INGRESS_GATEWAY"
	HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDINGRESSGATEWAY HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind = "KIND_INGRESS_GATEWAY"

	// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDDESTINATION captures enum value "KIND_DESTINATION"
	HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindKINDDESTINATION HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind = "KIND_DESTINATION"
)

// for schema
var hashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindEnum []interface{}

func init() {
	var res []HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind
	if err := json.Unmarshal([]byte(`["KIND_UNSPECIFIED","KIND_TYPICAL","KIND_CONNECT_PROXY","KIND_MESH_GATEWAY","KIND_TERMINATING_GATEWAY","KIND_INGRESS_GATEWAY","KIND_DESTINATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindEnum = append(hashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindEnum, v)
	}
}

func (m HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind) validateHashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindEnum(path, location string, value HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud global network manager 20220215 service summary kind
func (m HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudGlobalNetworkManager20220215ServiceSummaryKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud global network manager 20220215 service summary kind based on context it is used
func (m HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
