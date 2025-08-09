package main

import (
	"fmt"
	"github.com/Yeah114/blocks"
)

func main() {
	// 示例1: RuntimeID转方块状态字符串
	runtimeID := uint32(282305430) // 假设这是一个有效的runtimeID
	blockStr, found := blocks.RuntimeIDToBlockNameWithStateStr(runtimeID)
	fmt.Printf("示例1 - RuntimeID %d 转换结果: %s (found: %t)\n\n", runtimeID, blockStr, found)

	// 示例2: 方块字符串转RuntimeID
	blockString := `minecraft:leaves["persistent_bit"=false,"update_bit"=false,"old_leaf_type"="jungle"]`
	rtid, found := blocks.SchemBlockStrToRuntimeID(blockString)
	fmt.Printf("示例2 - 方块字符串 %q 转换结果: %d (found: %t)\n\n", blockString, rtid, found)
	blockStr, found = blocks.RuntimeIDToBlockNameWithStateStr(rtid)
	fmt.Printf("示例2.1 - RuntimeID %d 转换结果: %s (found: %t)\n\n", rtid, blockStr, found)

	// 示例3: 旧版名称+数据值转RuntimeID
	legacyName := "minecraft:leaves"
	dataValue := uint16(3)
	legacyRtid, found := blocks.LegacyBlockToRuntimeID(legacyName, dataValue)
	fmt.Printf("示例3 - 旧版方块 %s:%d 转换结果: %d (found: %t)\n\n", 
		legacyName, dataValue, legacyRtid, found)
	blockStr, found = blocks.RuntimeIDToBlockNameWithStateStr(legacyRtid)
	fmt.Printf("示例3.1 - RuntimeID %d 转换结果: %s (found: %t)\n\n", legacyRtid, blockStr, found)

	// 示例4: 使用属性map转换
	properties := map[string]any{
		"coral_color": "blue",
		"dead_bit":    true,
	}
	propRtid, found := blocks.BlockNameAndStateToRuntimeID("minecraft:coral_block", properties)
	fmt.Printf("示例4 - 属性map转换结果: %d (found: %t)\n\n", propRtid, found)

	// 示例5: 处理无效输入
	invalidRtid := uint32(999999)
	_, invalidFound := blocks.RuntimeIDToBlock(invalidRtid)
	fmt.Printf("示例5 - 无效RuntimeID %d 查找结果: %t\n\n", invalidRtid, invalidFound)

	// 示例6: 拆分方块名称和状态
	rtidForSplit := uint32(100) // 使用示例1的ID
	name, state, found := blocks.RuntimeIDToBlockNameAndStateStr(rtidForSplit)
	fmt.Printf("示例6 - 拆分结果 - 名称: %q 状态: %q (found: %t)\n\n", name, state, found)

	// 示例7: 原理图转换(旧版)
	schemBlock := "chest[facing=north]"
	schemRtid, found := blocks.SchemBlockStrToRuntimeID(schemBlock)
	fmt.Printf("示例7 - 原理图方块 %q 转换结果: %d (found: %t)\n", schemBlock, schemRtid, found)

	// 示例8 - Schematic双向转换
	blockID := uint8(17) // log
	dataVal := uint8(3)
	schematicRtid := blocks.SchematicToRuntimeID(blockID, dataVal)
	fmt.Printf("示例8 - Schematic转换 ID:%d Data:%d -> RuntimeID:%d\n", 
		blockID, dataVal, schematicRtid)
	blockStr, found = blocks.RuntimeIDToBlockNameWithStateStr(schematicRtid)
	fmt.Printf("示例8.1 - RuntimeID %d 转换结果: %s (found: %t)\n\n", schematicRtid, blockStr, found)

	// 示例9 - RuntimeID反向转换为Schematic
	reverseBlock, reverseData, reverseFound := blocks.RuntimeIDToSchematic(schematicRtid)
	fmt.Printf("示例9 - 反向转换 RuntimeID:%d -> Block:%d Data:%d (found: %t)\n", 
		schematicRtid, reverseBlock, reverseData, reverseFound)
	
	// 验证双向转换的一致性
	verifyRtid := blocks.SchematicToRuntimeID(reverseBlock, reverseData)
	fmt.Printf("示例9.1 - 验证一致性: %d -> (%d,%d) -> %d (一致: %t)\n\n", 
		schematicRtid, reverseBlock, reverseData, verifyRtid, schematicRtid == verifyRtid)
}