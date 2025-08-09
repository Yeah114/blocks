package main

import (
	"testing"
	"github.com/Yeah114/blocks"
)

func BenchmarkSchematicToRuntimeID(b *testing.B) {
	blocks.Init() // 确保初始化完成
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 测试常见的方块转换
		blocks.SchematicToRuntimeID(1, 0)   // stone
		blocks.SchematicToRuntimeID(2, 0)   // grass
		blocks.SchematicToRuntimeID(4, 0)   // cobblestone
		blocks.SchematicToRuntimeID(5, 0)   // planks
		blocks.SchematicToRuntimeID(17, 0)  // log
		blocks.SchematicToRuntimeID(17, 1)  // log with data
		blocks.SchematicToRuntimeID(35, 14) // wool with different color
	}
}

func BenchmarkRuntimeIDToSchematic(b *testing.B) {
	blocks.Init() // 确保初始化完成
	
	// 预先获取一些RuntimeID用于测试
	rtid1 := blocks.SchematicToRuntimeID(1, 0)   // stone
	rtid2 := blocks.SchematicToRuntimeID(2, 0)   // grass
	rtid3 := blocks.SchematicToRuntimeID(4, 0)   // cobblestone
	rtid4 := blocks.SchematicToRuntimeID(17, 3)  // log with data
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blocks.RuntimeIDToSchematic(rtid1)
		blocks.RuntimeIDToSchematic(rtid2)
		blocks.RuntimeIDToSchematic(rtid3)
		blocks.RuntimeIDToSchematic(rtid4)
	}
}