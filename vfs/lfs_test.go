package vfs

import (
	"testing"
)

// TestSerializedIndex 测试 serializedIndex 函数
func TestSerializedIndex(t *testing.T) {
	// 创建一个测试的 INode
	inode := &INode{
		RegionID:  1234,
		Position:  5678,
		Length:    100,
		ExpiredAt: 1617181723,
		CreatedAt: 1617181623,
	}

	// 计算预期的字节切片
	expectedLength := 48

	// 调用 serializeIndex
	result, err := serializedIndex(1001, inode)
	if err != nil {
		t.Fatalf("serialized index failed: %v", err)
	}

	// 检查返回的字节切片长度
	if len(result) != expectedLength {
		t.Errorf("expected result length %d, got %d", expectedLength, len(result))
	}

	// 验证内容字段进行反序列化并检查
	inum, node, err := deserializedIndex(result)

	if err != nil {
		t.Errorf("failed to deserialized: %v", err)
	}

	// 验证字段是否一致
	if inum != 1001 {
		t.Errorf("expected inum %d, got %d", 1001, inum)
	}
	if node.RegionID != inode.RegionID {
		t.Errorf("expected RegionID %d, got %d", inode.RegionID, node.RegionID)
	}
	if node.Position != inode.Position {
		t.Errorf("expected Offset %d, got %d", inode.Position, node.RegionID)
	}
	if node.Length != inode.Length {
		t.Errorf("expected Length %d, got %d", inode.Length, node.Length)
	}
	if node.ExpiredAt != inode.ExpiredAt {
		t.Errorf("expected ExpiredAt %d, got %d", inode.ExpiredAt, node.ExpiredAt)
	}
	if node.CreatedAt != inode.CreatedAt {
		t.Errorf("expected CreatedAt %d, got %d", inode.CreatedAt, node.CreatedAt)
	}

}
