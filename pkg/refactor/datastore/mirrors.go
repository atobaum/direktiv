package datastore

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MirrorConfig holds configuration data that are needed to create a mirror (pulling mirror credentials, urls, keys
// and any other details).
type MirrorConfig struct {
	Namespace string `json:"-"`

	URL                  string `json:"url"`
	GitRef               string `json:"gitRef"`
	GitCommitHash        string `json:"gitCommitHash"`
	PublicKey            string `json:"publicKey"`
	PrivateKey           string `json:"-"`
	PrivateKeyPassphrase string `json:"-"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Insecure bool `json:"insecure"`
}

// MirrorProcess different statuses.
const (
	ProcessStatusComplete  = "complete"
	ProcessStatusPending   = "pending"
	ProcessStatusExecuting = "executing"
	ProcessStatusFailed    = "failed"
)

// MirrorProcess different types.
const (
	// Indicates initial mirroring process.
	ProcessTypeInit = "init"

	// Indicates re-mirroring process.
	ProcessTypeSync = "sync"

	// Indicates dry run process.
	ProcessTypeDryRun = "dryrun"
)

// MirrorProcess represents an instance of mirroring process that happened or is currently happened. For every mirroring
// process gets executing, a MirrorProcess instance should be created with datastore.MirrorStore.
type MirrorProcess struct {
	ID        uuid.UUID
	Namespace string

	Status string
	Typ    string

	EndedAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// MirrorStore *doesn't* lunch any mirroring process. MirrorStore is only responsible for fetching and setting datastore.MirrorConfig and
// datastore.MirrorProcess from datastore.
type MirrorStore interface {
	// CreateConfig stores a new config in the store.
	CreateConfig(ctx context.Context, config *MirrorConfig) (*MirrorConfig, error)

	// UpdateConfig updates a config in the store.
	UpdateConfig(ctx context.Context, config *MirrorConfig) (*MirrorConfig, error)

	// GetConfig gets config by namespace from the store.
	GetConfig(ctx context.Context, namespace string) (*MirrorConfig, error)

	GetAllConfigs(ctx context.Context) ([]*MirrorConfig, error)
	// CreateProcess stores a new process in the store.
	CreateProcess(ctx context.Context, process *MirrorProcess) (*MirrorProcess, error)

	// UpdateProcess update a process in the store.
	UpdateProcess(ctx context.Context, process *MirrorProcess) (*MirrorProcess, error)

	// GetProcess gets a process by id from the store.
	GetProcess(ctx context.Context, id uuid.UUID) (*MirrorProcess, error)

	// GetProcessesByNamespace gets all processes that belong to a namespace from the store.
	GetProcessesByNamespace(ctx context.Context, namespace string) ([]*MirrorProcess, error)

	// GetUnfinishedProcesses gets all processes that haven't completed from the store.
	GetUnfinishedProcesses(ctx context.Context) ([]*MirrorProcess, error)

	// DeleteOldProcesses deletes all old processes.
	DeleteOldProcesses(ctx context.Context, before time.Time) error
}
