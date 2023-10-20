/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// TerminalOrdersCompanyLevelApi service
type TerminalOrdersCompanyLevelApi common.Service

// All parameters accepted by TerminalOrdersCompanyLevelApi.CancelOrder
type TerminalOrdersCompanyLevelApiCancelOrderInput struct {
	companyId string
	orderId   string
}

/*
Prepare a request for CancelOrder
@param companyId The unique identifier of the company account.@param orderId The unique identifier of the order.
@return TerminalOrdersCompanyLevelApiCancelOrderInput
*/
func (a *TerminalOrdersCompanyLevelApi) CancelOrderInput(companyId string, orderId string) TerminalOrdersCompanyLevelApiCancelOrderInput {
	return TerminalOrdersCompanyLevelApiCancelOrderInput{
		companyId: companyId,
		orderId:   orderId,
	}
}

/*
CancelOrder Cancel an order

Cancels the terminal products order identified in the path.
Cancelling is only possible while the order has the status **Placed**.
To cancel an order, make a POST call without a request body. The response returns the full order details, but with the status changed to **Cancelled**.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiCancelOrderInput - Request parameters, see CancelOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) CancelOrder(ctx context.Context, r TerminalOrdersCompanyLevelApiCancelOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders/{orderId}/cancel"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.CreateOrder
type TerminalOrdersCompanyLevelApiCreateOrderInput struct {
	companyId            string
	terminalOrderRequest *TerminalOrderRequest
}

func (r TerminalOrdersCompanyLevelApiCreateOrderInput) TerminalOrderRequest(terminalOrderRequest TerminalOrderRequest) TerminalOrdersCompanyLevelApiCreateOrderInput {
	r.terminalOrderRequest = &terminalOrderRequest
	return r
}

/*
Prepare a request for CreateOrder
@param companyId The unique identifier of the company account.
@return TerminalOrdersCompanyLevelApiCreateOrderInput
*/
func (a *TerminalOrdersCompanyLevelApi) CreateOrderInput(companyId string) TerminalOrdersCompanyLevelApiCreateOrderInput {
	return TerminalOrdersCompanyLevelApiCreateOrderInput{
		companyId: companyId,
	}
}

/*
CreateOrder Create an order

Creates an order for payment terminal products for the company identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiCreateOrderInput - Request parameters, see CreateOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) CreateOrder(ctx context.Context, r TerminalOrdersCompanyLevelApiCreateOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.terminalOrderRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.CreateShippingLocation
type TerminalOrdersCompanyLevelApiCreateShippingLocationInput struct {
	companyId        string
	shippingLocation *ShippingLocation
}

func (r TerminalOrdersCompanyLevelApiCreateShippingLocationInput) ShippingLocation(shippingLocation ShippingLocation) TerminalOrdersCompanyLevelApiCreateShippingLocationInput {
	r.shippingLocation = &shippingLocation
	return r
}

/*
Prepare a request for CreateShippingLocation
@param companyId The unique identifier of the company account.
@return TerminalOrdersCompanyLevelApiCreateShippingLocationInput
*/
func (a *TerminalOrdersCompanyLevelApi) CreateShippingLocationInput(companyId string) TerminalOrdersCompanyLevelApiCreateShippingLocationInput {
	return TerminalOrdersCompanyLevelApiCreateShippingLocationInput{
		companyId: companyId,
	}
}

/*
CreateShippingLocation Create a shipping location

Creates a shipping location for the company identified in the path. A shipping location defines an address where orders can be delivered.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiCreateShippingLocationInput - Request parameters, see CreateShippingLocationInput
@return ShippingLocation, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) CreateShippingLocation(ctx context.Context, r TerminalOrdersCompanyLevelApiCreateShippingLocationInput) (ShippingLocation, *http.Response, error) {
	res := &ShippingLocation{}
	path := "/companies/{companyId}/shippingLocations"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.shippingLocation,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.GetOrder
type TerminalOrdersCompanyLevelApiGetOrderInput struct {
	companyId string
	orderId   string
}

/*
Prepare a request for GetOrder
@param companyId The unique identifier of the company account.@param orderId The unique identifier of the order.
@return TerminalOrdersCompanyLevelApiGetOrderInput
*/
func (a *TerminalOrdersCompanyLevelApi) GetOrderInput(companyId string, orderId string) TerminalOrdersCompanyLevelApiGetOrderInput {
	return TerminalOrdersCompanyLevelApiGetOrderInput{
		companyId: companyId,
		orderId:   orderId,
	}
}

/*
GetOrder Get an order

Returns the details of the terminal products order identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiGetOrderInput - Request parameters, see GetOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) GetOrder(ctx context.Context, r TerminalOrdersCompanyLevelApiGetOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders/{orderId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.ListBillingEntities
type TerminalOrdersCompanyLevelApiListBillingEntitiesInput struct {
	companyId string
	name      *string
}

// The name of the billing entity.
func (r TerminalOrdersCompanyLevelApiListBillingEntitiesInput) Name(name string) TerminalOrdersCompanyLevelApiListBillingEntitiesInput {
	r.name = &name
	return r
}

/*
Prepare a request for ListBillingEntities
@param companyId The unique identifier of the company account.
@return TerminalOrdersCompanyLevelApiListBillingEntitiesInput
*/
func (a *TerminalOrdersCompanyLevelApi) ListBillingEntitiesInput(companyId string) TerminalOrdersCompanyLevelApiListBillingEntitiesInput {
	return TerminalOrdersCompanyLevelApiListBillingEntitiesInput{
		companyId: companyId,
	}
}

/*
ListBillingEntities Get a list of billing entities

Returns the billing entities of the company identified in the path and all merchant accounts belonging to the company.
A billing entity is a legal entity where we charge orders to. An order for terminal products must contain the ID of a billing entity.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiListBillingEntitiesInput - Request parameters, see ListBillingEntitiesInput
@return BillingEntitiesResponse, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) ListBillingEntities(ctx context.Context, r TerminalOrdersCompanyLevelApiListBillingEntitiesInput) (BillingEntitiesResponse, *http.Response, error) {
	res := &BillingEntitiesResponse{}
	path := "/companies/{companyId}/billingEntities"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.name != nil {
		common.ParameterAddToQuery(queryParams, "name", r.name, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.ListOrders
type TerminalOrdersCompanyLevelApiListOrdersInput struct {
	companyId              string
	customerOrderReference *string
	status                 *string
	offset                 *int32
	limit                  *int32
}

// Your purchase order number.
func (r TerminalOrdersCompanyLevelApiListOrdersInput) CustomerOrderReference(customerOrderReference string) TerminalOrdersCompanyLevelApiListOrdersInput {
	r.customerOrderReference = &customerOrderReference
	return r
}

// The order status. Possible values (not case-sensitive): Placed, Confirmed, Cancelled, Shipped, Delivered.
func (r TerminalOrdersCompanyLevelApiListOrdersInput) Status(status string) TerminalOrdersCompanyLevelApiListOrdersInput {
	r.status = &status
	return r
}

// The number of orders to skip.
func (r TerminalOrdersCompanyLevelApiListOrdersInput) Offset(offset int32) TerminalOrdersCompanyLevelApiListOrdersInput {
	r.offset = &offset
	return r
}

// The number of orders to return.
func (r TerminalOrdersCompanyLevelApiListOrdersInput) Limit(limit int32) TerminalOrdersCompanyLevelApiListOrdersInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for ListOrders
@param companyId The unique identifier of the company account.
@return TerminalOrdersCompanyLevelApiListOrdersInput
*/
func (a *TerminalOrdersCompanyLevelApi) ListOrdersInput(companyId string) TerminalOrdersCompanyLevelApiListOrdersInput {
	return TerminalOrdersCompanyLevelApiListOrdersInput{
		companyId: companyId,
	}
}

/*
ListOrders Get a list of orders

Returns a lists of terminal products orders for the company identified in the path.
To filter the list, use one or more of the query parameters.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiListOrdersInput - Request parameters, see ListOrdersInput
@return TerminalOrdersResponse, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) ListOrders(ctx context.Context, r TerminalOrdersCompanyLevelApiListOrdersInput) (TerminalOrdersResponse, *http.Response, error) {
	res := &TerminalOrdersResponse{}
	path := "/companies/{companyId}/terminalOrders"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.customerOrderReference != nil {
		common.ParameterAddToQuery(queryParams, "customerOrderReference", r.customerOrderReference, "")
	}
	if r.status != nil {
		common.ParameterAddToQuery(queryParams, "status", r.status, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.ListShippingLocations
type TerminalOrdersCompanyLevelApiListShippingLocationsInput struct {
	companyId string
	name      *string
	offset    *int32
	limit     *int32
}

// The name of the shipping location.
func (r TerminalOrdersCompanyLevelApiListShippingLocationsInput) Name(name string) TerminalOrdersCompanyLevelApiListShippingLocationsInput {
	r.name = &name
	return r
}

// The number of locations to skip.
func (r TerminalOrdersCompanyLevelApiListShippingLocationsInput) Offset(offset int32) TerminalOrdersCompanyLevelApiListShippingLocationsInput {
	r.offset = &offset
	return r
}

// The number of locations to return.
func (r TerminalOrdersCompanyLevelApiListShippingLocationsInput) Limit(limit int32) TerminalOrdersCompanyLevelApiListShippingLocationsInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for ListShippingLocations
@param companyId The unique identifier of the company account.
@return TerminalOrdersCompanyLevelApiListShippingLocationsInput
*/
func (a *TerminalOrdersCompanyLevelApi) ListShippingLocationsInput(companyId string) TerminalOrdersCompanyLevelApiListShippingLocationsInput {
	return TerminalOrdersCompanyLevelApiListShippingLocationsInput{
		companyId: companyId,
	}
}

/*
ListShippingLocations Get a list of shipping locations

Returns the shipping locations for the company identified in the path and all merchant accounts belonging to the company.
A shipping location includes the address where orders can be delivered, and an ID which you need to specify when ordering terminal products.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiListShippingLocationsInput - Request parameters, see ListShippingLocationsInput
@return ShippingLocationsResponse, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) ListShippingLocations(ctx context.Context, r TerminalOrdersCompanyLevelApiListShippingLocationsInput) (ShippingLocationsResponse, *http.Response, error) {
	res := &ShippingLocationsResponse{}
	path := "/companies/{companyId}/shippingLocations"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.name != nil {
		common.ParameterAddToQuery(queryParams, "name", r.name, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.ListTerminalModels
type TerminalOrdersCompanyLevelApiListTerminalModelsInput struct {
	companyId string
}

/*
Prepare a request for ListTerminalModels
@param companyId The unique identifier of the company account.
@return TerminalOrdersCompanyLevelApiListTerminalModelsInput
*/
func (a *TerminalOrdersCompanyLevelApi) ListTerminalModelsInput(companyId string) TerminalOrdersCompanyLevelApiListTerminalModelsInput {
	return TerminalOrdersCompanyLevelApiListTerminalModelsInput{
		companyId: companyId,
	}
}

/*
ListTerminalModels Get a list of terminal models

Returns a list of payment terminal models that the company identified in the path has access to.
The response includes the terminal model ID, which can be used as a query parameter when getting a list of terminals or a list of products for ordering.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiListTerminalModelsInput - Request parameters, see ListTerminalModelsInput
@return TerminalModelsResponse, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) ListTerminalModels(ctx context.Context, r TerminalOrdersCompanyLevelApiListTerminalModelsInput) (TerminalModelsResponse, *http.Response, error) {
	res := &TerminalModelsResponse{}
	path := "/companies/{companyId}/terminalModels"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.ListTerminalProducts
type TerminalOrdersCompanyLevelApiListTerminalProductsInput struct {
	companyId       string
	country         *string
	terminalModelId *string
	offset          *int32
	limit           *int32
}

// The country to return products for, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. For example, **US**
func (r TerminalOrdersCompanyLevelApiListTerminalProductsInput) Country(country string) TerminalOrdersCompanyLevelApiListTerminalProductsInput {
	r.country = &country
	return r
}

// The terminal model to return products for. Use the ID returned in the [GET &#x60;/terminalModels&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/companies/{companyId}/terminalModels) response. For example, **Verifone.M400**
func (r TerminalOrdersCompanyLevelApiListTerminalProductsInput) TerminalModelId(terminalModelId string) TerminalOrdersCompanyLevelApiListTerminalProductsInput {
	r.terminalModelId = &terminalModelId
	return r
}

// The number of products to skip.
func (r TerminalOrdersCompanyLevelApiListTerminalProductsInput) Offset(offset int32) TerminalOrdersCompanyLevelApiListTerminalProductsInput {
	r.offset = &offset
	return r
}

// The number of products to return.
func (r TerminalOrdersCompanyLevelApiListTerminalProductsInput) Limit(limit int32) TerminalOrdersCompanyLevelApiListTerminalProductsInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for ListTerminalProducts
@param companyId The unique identifier of the company account.
@return TerminalOrdersCompanyLevelApiListTerminalProductsInput
*/
func (a *TerminalOrdersCompanyLevelApi) ListTerminalProductsInput(companyId string) TerminalOrdersCompanyLevelApiListTerminalProductsInput {
	return TerminalOrdersCompanyLevelApiListTerminalProductsInput{
		companyId: companyId,
	}
}

/*
ListTerminalProducts Get a list of terminal products

Returns a country-specific list of payment terminal packages and parts that the company identified in the path has access to.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiListTerminalProductsInput - Request parameters, see ListTerminalProductsInput
@return TerminalProductsResponse, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) ListTerminalProducts(ctx context.Context, r TerminalOrdersCompanyLevelApiListTerminalProductsInput) (TerminalProductsResponse, *http.Response, error) {
	res := &TerminalProductsResponse{}
	path := "/companies/{companyId}/terminalProducts"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.country != nil {
		common.ParameterAddToQuery(queryParams, "country", r.country, "")
	}
	if r.terminalModelId != nil {
		common.ParameterAddToQuery(queryParams, "terminalModelId", r.terminalModelId, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TerminalOrdersCompanyLevelApi.UpdateOrder
type TerminalOrdersCompanyLevelApiUpdateOrderInput struct {
	companyId            string
	orderId              string
	terminalOrderRequest *TerminalOrderRequest
}

func (r TerminalOrdersCompanyLevelApiUpdateOrderInput) TerminalOrderRequest(terminalOrderRequest TerminalOrderRequest) TerminalOrdersCompanyLevelApiUpdateOrderInput {
	r.terminalOrderRequest = &terminalOrderRequest
	return r
}

/*
Prepare a request for UpdateOrder
@param companyId The unique identifier of the company account.@param orderId The unique identifier of the order.
@return TerminalOrdersCompanyLevelApiUpdateOrderInput
*/
func (a *TerminalOrdersCompanyLevelApi) UpdateOrderInput(companyId string, orderId string) TerminalOrdersCompanyLevelApiUpdateOrderInput {
	return TerminalOrdersCompanyLevelApiUpdateOrderInput{
		companyId: companyId,
		orderId:   orderId,
	}
}

/*
UpdateOrder Update an order

Updates the terminal products order identified in the path.
Updating is only possible while the order has the status **Placed**.

The request body only needs to contain what you want to change.
However, to update the products in the `items` array, you must provide the entire array. For example, if the array has three items:
 To remove one item, the array must include the remaining two items. Or to add one item, the array must include all four items.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersCompanyLevelApiUpdateOrderInput - Request parameters, see UpdateOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersCompanyLevelApi) UpdateOrder(ctx context.Context, r TerminalOrdersCompanyLevelApiUpdateOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/companies/{companyId}/terminalOrders/{orderId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.terminalOrderRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
