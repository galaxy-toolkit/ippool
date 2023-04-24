package pool

import (
	"context"
	"testing"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/internal/global"
)

func init() {
	global.InitConfig("") // 配置
	global.InitPostgres() // 数据库
}

func TestIPInsert(t *testing.T) {
	data := []*model.IP{
		{Address: "1.1.1.1", Status: "normal", Latency: 100, Source: "a.com"},
		{Address: "2.1.1.1", Status: "normal", Latency: 100, Source: "dd.com"},
		{Address: "3.1.1.1", Status: "invalid", Latency: 100, Source: "dd.com"},
	}
	if err := Use(context.TODO()).IP.InsertOne(data[0]); err != nil {
		t.Fatal(err)
	}
	t.Log(data[0].ID)

	if err := Use(context.TODO()).IP.InsertMany(data[1:2]); err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestIPGet(t *testing.T) {
	ip, err := Use(context.TODO()).IP.GetOneByID(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)

	ips, err := Use(context.TODO()).IP.GetManyByIDs([]int64{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ips)
}

func TestIPUpdate(t *testing.T) {
	updateData := map[string]any{
		"status": "invalid",
	}

	rows, err := Use(context.TODO()).IP.UpdateByID(3, updateData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rows)

	ip, err := Use(context.TODO()).IP.GetOneByID(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}

func TestIPDelete(t *testing.T) {
	rows, err := Use(context.TODO()).IP.DeleteManyByIDs([]int64{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rows)
}
