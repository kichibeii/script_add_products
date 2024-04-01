package services

import (
	"reflect"
	"testing"

	"github.com/script_add_products/server/commons"
	"github.com/script_add_products/server/domain/repositories"
	"github.com/script_add_products/server/domain/thirdparties"
)

func TestNewSyncProductService(t *testing.T) {
	type args struct {
		opt        commons.Options
		repo       repositories.IRepository
		thirdParty thirdparties.IThirdParty
	}
	tests := []struct {
		name string
		args args
		want ISyncProductService
	}{
		{
			name: "success",
			args: args{},
			want: &SyncProductService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSyncProductService(tt.args.opt, tt.args.repo, tt.args.thirdParty); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSyncProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}
