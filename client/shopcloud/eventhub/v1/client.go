package client

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/Talk-Point/shopcloud/client/middleware"
	eventhubv1 "github.com/Talk-Point/shopcloud/gen/proto/shopcloud/eventhub/v1"
	v1connect "github.com/Talk-Point/shopcloud/gen/proto/shopcloud/eventhub/v1/eventhubv1connect"
)

type Client struct {
	client v1connect.EventHubDispatcherServiceClient
}

func (c *Client) GetWorkspaces(next string, page_size int32) (*eventhubv1.ListWorkspacesResponse, error) {
	res, err := c.client.ListWorkspaces(
		context.Background(),
		connect.NewRequest(&eventhubv1.ListWorkspacesRequest{
			PageSize:  page_size,
			PageToken: next,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) GetWorkspace(id string) (*eventhubv1.Workspace, error) {
	res, err := c.client.GetWorkspace(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetWorkspaceRequest{
			Id: id,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Workspace, nil
}

func (c *Client) GetWorkspaceByName(name string) (*eventhubv1.Workspace, error) {
	res, err := c.client.GetWorkspaceByName(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetWorkspaceByNameRequest{
			Name: name,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Workspace, nil
}

func (c *Client) GetOrCreateWorkspace(name string) (*eventhubv1.Workspace, error) {
	res, err := c.client.GetOrCreateWorkspace(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetOrCreateWorkspaceRequest{
			Name: name,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Workspace, nil
}

func (c *Client) CreateWorkspace(name string) (*eventhubv1.Workspace, error) {
	res, err := c.client.CreateWorkspace(
		context.Background(),
		connect.NewRequest(&eventhubv1.CreateWorkspaceRequest{
			Name: name,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Workspace, nil
}

func (c *Client) ListWorkspaces(next string, page_size int32) (*eventhubv1.ListWorkspacesResponse, error) {
	res, err := c.client.ListWorkspaces(
		context.Background(),
		connect.NewRequest(&eventhubv1.ListWorkspacesRequest{
			PageSize:  page_size,
			PageToken: next,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) DeleteWorkspace(id string) error {
	_, err := c.client.DeleteWorkspace(
		context.Background(),
		connect.NewRequest(&eventhubv1.DeleteWorkspaceRequest{
			Id: id,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddUsersToWorkspace(id string, users []string) (*eventhubv1.AddUsersToWorkspaceResponse, error) {
	res, err := c.client.AddUsersToWorkspace(
		context.Background(),
		connect.NewRequest(&eventhubv1.AddUsersToWorkspaceRequest{
			Id:    id,
			Users: users,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) RemoveUsersFromWorkspace(id string, users []string) (*eventhubv1.RemoveUsersFromWorkspaceResponse, error) {
	res, err := c.client.RemoveUsersFromWorkspace(
		context.Background(),
		connect.NewRequest(&eventhubv1.RemoveUsersFromWorkspaceRequest{
			Id:    id,
			Users: users,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) PublishEvent(workspace string, topic string, payload string) (string, error) {
	res, err := c.client.PublishEvent(
		context.Background(),
		connect.NewRequest(&eventhubv1.PublishEventRequest{
			Workspace: workspace,
			Topic:     topic,
			Payload:   payload,
		}),
	)
	if err != nil {
		return "", err
	}

	return res.Msg.Id, nil
}

func (c *Client) DispatchEventToSubscriptions(event string, next string) ([]string, error) {
	res, err := c.client.DispatchEventToSubscriptions(
		context.Background(),
		connect.NewRequest(&eventhubv1.DispatchEventToSubscriptionsRequest{
			Event:         event,
			PageSize:      100,
			NextPageToken: next,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Subscriptions, nil
}

func (c *Client) GetEvent(id string) (*eventhubv1.Event, error) {
	res, err := c.client.GetEvent(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetEventRequest{
			Id: id,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) GetSubscription(id string) (*eventhubv1.Subscription, error) {
	res, err := c.client.GetSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetSubscriptionRequest{
			Id: id,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Subscription, nil
}

func (c *Client) CreateSubscription(workspace string, name string, topic string, endpoint string) (*eventhubv1.Subscription, error) {
	res, err := c.client.CreateSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.CreateSubscriptionRequest{
			Name:      name,
			Workspace: workspace,
			Topics:    []string{topic},
			Endpoint:  endpoint,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Subscription, nil
}

func (c *Client) DeleteSubscription(id string) error {
	_, err := c.client.DeleteSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.DeleteSubscriptionRequest{
			Id: id,
		}),
	)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) ListSubscriptions(next string, page_size int32) (*eventhubv1.ListSubscriptionsResponse, error) {
	res, err := c.client.ListSubscriptions(
		context.Background(),
		connect.NewRequest(&eventhubv1.ListSubscriptionsRequest{
			PageSize:  page_size,
			PageToken: next,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) AddSubscriptionTopics(id string, topics []string) error {
	_, err := c.client.AddSubscriptionTopics(
		context.Background(),
		connect.NewRequest(&eventhubv1.AddSubscriptionTopicsRequest{
			Id:     id,
			Topics: topics,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveSubscriptionTopics(id string, topics []string) error {
	_, err := c.client.RemoveSubscriptionTopics(
		context.Background(),
		connect.NewRequest(&eventhubv1.RemoveSubscriptionTopicsRequest{
			Id:     id,
			Topics: topics,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SetSubscriptionTopics(id string, topics []string) error {
	_, err := c.client.SetSubscriptionTopics(
		context.Background(),
		connect.NewRequest(&eventhubv1.SetSubscriptionTopicsRequest{
			Id:     id,
			Topics: topics,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ActivateSubscription(id string) error {
	_, err := c.client.ActivateSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.ActivateSubscriptionRequest{
			Id: id,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeactivateSubscription(id string) error {
	_, err := c.client.DeactivateSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.DeactivateSubscriptionRequest{
			Id: id,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PauseSubscription(id string) error {
	_, err := c.client.PauseSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.PauseSubscriptionRequest{
			Id: id,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ResumeSubscription(id string) error {
	_, err := c.client.ResumeSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.ResumeSubscriptionRequest{
			Id: id,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetDelivery(id string) (*eventhubv1.Delivery, error) {
	res, err := c.client.GetDelivery(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetDeliveryRequest{
			Id: id,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Delivery, nil
}

func (c *Client) ListDeliveries(next string, page_size int32) (*eventhubv1.ListDeliveriesResponse, error) {
	res, err := c.client.ListDeliveries(
		context.Background(),
		connect.NewRequest(&eventhubv1.ListDeliveriesRequest{
			PageSize:  page_size,
			PageToken: next,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) DeleteDelivery(id string) (*eventhubv1.DeleteDeliveryResponse, error) {
	res, err := c.client.DeleteDelivery(
		context.Background(),
		connect.NewRequest(&eventhubv1.DeleteDeliveryRequest{
			Id: id,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) PushDelivery(id string) (bool, error) {
	res, err := c.client.PushDelivery(
		context.Background(),
		connect.NewRequest(&eventhubv1.PushDeliveryRequest{
			Id: id,
		}),
	)
	if err != nil {
		return false, err
	}

	return res.Msg.IsSuccess, nil
}

func (c *Client) CreateDelivery(subscription string, event string) (string, error) {
	res, err := c.client.CreateDelivery(
		context.Background(),
		connect.NewRequest(&eventhubv1.CreateDeliveryRequest{
			Subscription: subscription,
			Event:        event,
		}),
	)
	if err != nil {
		return "", err
	}

	return res.Msg.Id, nil
}

func (c *Client) GetSubscriptionByName(workspace string, name string) (*eventhubv1.Subscription, error) {
	res, err := c.client.GetSubscriptionByName(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetSubscriptionByNameRequest{
			Workspace: workspace,
			Name:      name,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Subscription, nil
}

func (c *Client) GetOrCreateSubscription(workspace string, name string, topic string, endpoint string) (*eventhubv1.Subscription, error) {
	res, err := c.client.GetOrCreateSubscription(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetOrCreateSubscriptionRequest{
			Workspace: workspace,
			Name:      name,
			Topics:    []string{topic},
			Endpoint:  endpoint,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Subscription, nil
}

func (c *Client) GetGateway(id string) (*eventhubv1.Gateway, error) {
	res, err := c.client.GetGateway(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetGatewayRequest{
			Id: id,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Gateway, nil
}

func (c *Client) GetGatewayByCode(code string, workspace string) (*eventhubv1.Gateway, error) {
	res, err := c.client.GetGatewayByCode(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetGatewayByCodeRequest{
			Code:      code,
			Workspace: workspace,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Gateway, nil
}

func (c *Client) GetGatewayByHash(hash string) (*eventhubv1.Gateway, error) {
	res, err := c.client.GetGatewayByHash(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetGatewayByHashRequest{
			Hash: hash,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Gateway, nil
}

func (c *Client) ListGateways(next string, page_size int32) (*eventhubv1.ListGatewaysResponse, error) {
	res, err := c.client.ListGateways(
		context.Background(),
		connect.NewRequest(&eventhubv1.ListGatewaysRequest{
			PageSize:  page_size,
			PageToken: next,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func (c *Client) CreateGateway(workspace string, name string, topic string, auth string) (*eventhubv1.Gateway, error) {
	res, err := c.client.CreateGateway(
		context.Background(),
		connect.NewRequest(&eventhubv1.CreateGatewayRequest{
			Name:      name,
			Workspace: workspace,
			Topic:     topic,
			Auth:      auth,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Gateway, nil
}

func (c *Client) DeleteGateway(id string) error {
	_, err := c.client.DeleteGateway(
		context.Background(),
		connect.NewRequest(&eventhubv1.DeleteGatewayRequest{
			Id: id,
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetOrCreateGateway(workspace string, name string, topic string, auth string) (*eventhubv1.Gateway, error) {
	res, err := c.client.GetOrCreateGateway(
		context.Background(),
		connect.NewRequest(&eventhubv1.GetOrCreateGatewayRequest{
			Name:      name,
			Workspace: workspace,
			Topic:     topic,
			Auth:      auth,
		}),
	)
	if err != nil {
		return nil, err
	}

	return res.Msg.Gateway, nil
}

func (c *Client) PublishToGateway(hash string, payload string) (string, error) {
	res, err := c.client.PublishToGateway(
		context.Background(),
		connect.NewRequest(&eventhubv1.PublishToGatewayRequest{
			Hash:    hash,
			Payload: payload,
		}),
	)
	if err != nil {
		return "", err
	}

	return res.Msg.Id, nil
}

func NewClient(endpoint string, authToken string) *Client {
	interceptors := connect.WithInterceptors(
		middleware.ClientTokenAuthInterceptor(authToken),
	)
	client := v1connect.NewEventHubDispatcherServiceClient(
		http.DefaultClient,
		endpoint,
		connect.WithGRPC(),
		interceptors,
	)
	return &Client{
		client: client,
	}
}

func NewLocalClient(authToken string) *Client {
	return NewClient("http://localhost:8080", authToken)
}

func NewCloudClient(endpoint string, authToken string) *Client {
	return NewClient(fmt.Sprintf("https://%s", endpoint), authToken)
}
