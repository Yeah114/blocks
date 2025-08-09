import nbtlib
import json

nbt = nbtlib.load("netease_block_states.nbt")

blocks = []
# 直接遍历NBT标签中的blocks列表（保留原始NBT类型信息）
for block in nbt['blocks']:
    block_entry = {
        "data": block['val'].unpack(),  # 仅对值进行解包
        "name": block['name'].unpack(),
        "name_hash": abs(block['name_hash'].unpack()),
        "network_id": block['network_id'].unpack(),
        "states": []
    }
    
    # 遍历states的NBT标签，通过nbtlib的类型判断原始类型
    for state_name, state_tag in block['states'].items():
        # 解包获取值（保持原始类型）
        state_value = state_tag.unpack()
        
        # 直接判断NBT标签类型（保留原始类型信息）
        if isinstance(state_tag, nbtlib.Int):
            state_type = "int"
        elif isinstance(state_tag, nbtlib.Byte):
            state_type = "byte"
        elif isinstance(state_tag, nbtlib.String):
            state_type = "string"
        else:
            state_type = "unknown"  # 处理其他可能的类型
        
        state_entry = {
            "name": state_name,
            "type": state_type,
            "value": state_value
        }
        block_entry["states"].append(state_entry)
    
    blocks.append(block_entry)

output = {"blocks": blocks}

with open("block_palette_2.12.json", "w") as f:
    json.dump(output, f, indent="\t")
