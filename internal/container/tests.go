//go:build tests

package container

import "context"

func Inject(ctx context.Context, a any) error {
	return inject(ctx, a)
}
