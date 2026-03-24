package backupcontroller

import (
	"context"
	"fmt"

	k8upv1 "github.com/k8up-io/k8up/v2/api/v1"
)

func (b *BackupExecutor) validateBackupTarget(ctx context.Context) error {
	switch b.backup.Spec.Backend.BackendInterface.GetType() {
	case k8upv1.BackendTypeS3:
		return b.validateS3Target(ctx)
	case k8upv1.BackendTypeGCS:
		return b.validateGCSTarget(ctx)
	case k8upv1.BackendTypeAzure:
		return b.validateAzureTarget(ctx)
	case k8upv1.BackendTypeSwift:
		return b.validateSwiftTarget(ctx)
	case k8upv1.BackendTypeB2:
		return b.validateB2Target(ctx)
	case k8upv1.BackendTypeRest:
		return b.validateRestTarget(ctx)
	default:
		return fmt.Errorf("unsupported backup target type: %s", b.backup.Spec.Backend.BackendInterface.GetType())
	}
}

func (b *BackupExecutor) validateRestTarget(ctx context.Context) error {
	panic("unimplemented")
}

func (b *BackupExecutor) validateB2Target(ctx context.Context) error {
	panic("unimplemented")
}

func (b *BackupExecutor) validateSwiftTarget(ctx context.Context) error {
	panic("unimplemented")
}

func (b *BackupExecutor) validateAzureTarget(ctx context.Context) error {
	panic("unimplemented")
}

func (b *BackupExecutor) validateGCSTarget(ctx context.Context) error {
	panic("unimplemented")
}

func (b *BackupExecutor) validateS3Target(ctx context.Context) error {
	panic("unimplemented")
}
