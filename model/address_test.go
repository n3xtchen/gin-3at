package model

import (
	"testing"
)

func TestAddress(t *testing.T) {
	err := DB.AutoMigrate(&Address{})
	if err != nil {
		t.Errorf("Failed to migrate address table: %v", err)
		return
	}
	t.Log("Address table migrated successfully")

	// 创建一个地址记录
	newAddress := &Address{UserID: 1, Name: "n3xtchen", Phone: "1231231234", Address: "123 Test St"}

	result := DB.Create(newAddress)
	if result.Error != nil {
		t.Errorf("Failed to create address: %v", result.Error)
		return
	}

	var address Address
	result = DB.First(&address)
	// 验证地址是否正确
	if result.Error != nil {
		t.Errorf("Failed to retrieve address: %v", result.Error)
		return
	}
	if address.Name != newAddress.Name || address.Phone != newAddress.Phone || address.Address != newAddress.Address {
		t.Errorf("Address mismatch: got %v, want %v", address, newAddress)
		return
	}

	// 删除 address 表
	err = DB.Migrator().DropTable(&Address{})
	if err != nil {
		t.Errorf("Failed to drop address table: %v", err)
		return
	}
	t.Log("Address table dropped successfully")
}
