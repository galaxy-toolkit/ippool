package kuaidaili

import (
	"context"
	"reflect"
	"testing"

	"github.com/galaxy-toolkit/ippool/domain/model"
)

func TestCrawler_Crawl(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.IP
		wantErr bool
	}{
		{"", fields{ctx: context.TODO()}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Crawler{
				ctx: tt.fields.ctx,
			}
			got, err := c.Crawl()
			if (err != nil) != tt.wantErr {
				t.Errorf("Crawl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Crawl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *Crawler
	}{
		{"", args{ctx: context.TODO()}, &Crawler{ctx: context.TODO()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
