package structs

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
)

// TODO(tproxy): docs
type ClusterConfigEntry struct {
	Kind string
	Name string

	TransparentProxy TransparentProxyClusterConfig `alias:"transparent_proxy"`

	Meta           map[string]string `json:",omitempty"`
	EnterpriseMeta `hcl:",squash" mapstructure:",squash"`
	RaftIndex
}

// TODO(tproxy): docs
type TransparentProxyClusterConfig struct {
	CatalogDestinationsOnly bool `alias:"catalog_destinations_only"`
}

func (e *ClusterConfigEntry) GetKind() string {
	return ClusterConfig
}

func (e *ClusterConfigEntry) GetName() string {
	if e == nil {
		return ""
	}

	return e.Name
}

func (e *ClusterConfigEntry) GetMeta() map[string]string {
	if e == nil {
		return nil
	}
	return e.Meta
}

func (e *ClusterConfigEntry) Normalize() error {
	if e == nil {
		return fmt.Errorf("config entry is nil")
	}

	e.Kind = ClusterConfig
	e.Name = ClusterConfigCluster

	e.EnterpriseMeta.Normalize()

	return nil
}

func (e *ClusterConfigEntry) Validate() error {
	if e == nil {
		return fmt.Errorf("config entry is nil")
	}

	if e.Name != ClusterConfigCluster {
		return fmt.Errorf("invalid name (%q), only %q is supported", e.Name, ClusterConfigCluster)
	}

	if err := validateConfigEntryMeta(e.Meta); err != nil {
		return err
	}

	return e.validateEnterpriseMeta() // TODO
}

func (e *ClusterConfigEntry) CanRead(authz acl.Authorizer) bool {
	return true
}

func (e *ClusterConfigEntry) CanWrite(authz acl.Authorizer) bool {
	var authzContext acl.AuthorizerContext
	e.FillAuthzContext(&authzContext)
	return authz.OperatorWrite(&authzContext) == acl.Allow
}

func (e *ClusterConfigEntry) GetRaftIndex() *RaftIndex {
	if e == nil {
		return &RaftIndex{}
	}

	return &e.RaftIndex
}

func (e *ClusterConfigEntry) GetEnterpriseMeta() *EnterpriseMeta {
	if e == nil {
		return nil
	}

	return &e.EnterpriseMeta
}
