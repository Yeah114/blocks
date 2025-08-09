package blocks

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/Yeah114/blocks/block_set"
	"github.com/Yeah114/blocks/convertor"
	"github.com/Yeah114/blocks/describe"

	"github.com/andybalholm/brotli"
)

//go:embed "nemc.br"
var nemcBlockInfoBytes []byte

//go:embed "bedrock_java_to_translate.br"
var toNemcDataLoadBedrockJavaTranslateInfo []byte

//go:embed "specific_legacy_value_to_translate.br"
var toNemcDataLoadSpecificLegacyValuesTranslateInfo []byte

//go:embed "schem_to_translate.br"
var toNemcDataLoadSchemTranslateInfo []byte

func readAndUnpack(bs []byte) string {
	dataBytes, err := io.ReadAll(brotli.NewReader(bytes.NewBuffer(bs)))
	if err != nil {
		panic(err)
	}
	return string(dataBytes)
}

var MC_CURRENT *block_set.BlockSet
var MCBlocks *block_set.BlockSet

// duplicate in future
var NEMC_BLOCK_VERSION = uint32(0)
var NEMC_AIR_RUNTIMEID = uint32(0)
var AIR_RUNTIMEID = uint32(0)

func initNEMCBlocks() {
	bs := block_set.BlockSetFromStringRecords(readAndUnpack(nemcBlockInfoBytes), 0xFFFFFFFF)
	MC_CURRENT = bs
	MCBlocks = bs
	NEMC_BLOCK_VERSION = bs.Version()
	NEMC_AIR_RUNTIMEID = bs.AirRuntimeID()
	AIR_RUNTIMEID = bs.AirRuntimeID()
}

var DefaultAnyToNemcConvertor *convertor.ToNEMCConvertor
var SchemToNemcConvertor *convertor.ToNEMCConvertor

var quickSchematicMapping [256][16]uint32
var runtimeIDToSchematic map[uint32][2]uint8

func initSchematicBlockCheck(schematicToNemcConvertor *convertor.ToNEMCConvertor) {
	quickSchematicMapping = [256][16]uint32{}
	runtimeIDToSchematic = make(map[uint32][2]uint8)
	
	for i := 0; i < 256; i++ {
		blockName := schematicBlockStrings[i]
		_, found := DefaultAnyToNemcConvertor.TryBestSearchByLegacyValue(describe.BlockNameForSearch(blockName), 0)
		if !found {
			panic(fmt.Errorf("schematic %v 0 not found", blockName))
		}
	}
	for blockI := 0; blockI < 256; blockI++ {
		blockName := schematicBlockStrings[blockI]
		for dataI := 0; dataI < 16; dataI++ {
			rtid, found := schematicToNemcConvertor.TryBestSearchByLegacyValue(describe.BlockNameForSearch(blockName), uint16(dataI))
			if !found || rtid == AIR_RUNTIMEID {
				rtid, _ = schematicToNemcConvertor.TryBestSearchByLegacyValue(describe.BlockNameForSearch(blockName), 0)
			}
			quickSchematicMapping[blockI][dataI] = rtid
			
			// 构建反向映射，优先选择数据值为0的映射
			if existing, exists := runtimeIDToSchematic[rtid]; !exists || dataI == 0 {
				if !exists || existing[1] != 0 {
					runtimeIDToSchematic[rtid] = [2]uint8{uint8(blockI), uint8(dataI)}
				}
			}
		}
	}
	schematicToNemcConvertor = nil
}

func initConvertor() {
	DefaultAnyToNemcConvertor = MC_CURRENT.CreateEmptyConvertor()
	SchemToNemcConvertor = MC_CURRENT.CreateEmptyConvertor()
	schematicToNemcConvertor := MC_CURRENT.CreateEmptyConvertor()
	mcConvertRecords, err := convertor.ReadRecordsFromString(readAndUnpack(toNemcDataLoadBedrockJavaTranslateInfo))
	if err != nil {
		panic(err)
	}
	specificLegacyValuesConvertRecords, err := convertor.ReadRecordsFromString(readAndUnpack(toNemcDataLoadSpecificLegacyValuesTranslateInfo))
	if err != nil {
		panic(err)
	}
	for _, r := range mcConvertRecords {
		DefaultAnyToNemcConvertor.LoadConvertRecord(r, false, true)
		SchemToNemcConvertor.LoadConvertRecord(r, false, true)
		schematicToNemcConvertor.LoadConvertRecord(r, false, true)
	}
	for _, r := range specificLegacyValuesConvertRecords {
		DefaultAnyToNemcConvertor.LoadConvertRecord(r, true, true)
		SchemToNemcConvertor.LoadConvertRecord(r, true, true)
		schematicToNemcConvertor.LoadConvertRecord(r, true, true)
	}
	schemConvertRecords, err := convertor.ReadRecordsFromString(readAndUnpack(toNemcDataLoadSchemTranslateInfo))
	if err != nil {
		panic(err)
	}
	for _, r := range schemConvertRecords {
		DefaultAnyToNemcConvertor.LoadConvertRecord(r, false, false)
		SchemToNemcConvertor.LoadConvertRecord(r, true, true)
	}
	initSchematicBlockCheck(schematicToNemcConvertor)
}

var inited bool

func init() {
	initNEMCBlocks()
	Init()
}

func Init() {
	if inited {
		return
	}
	inited = true
	initConvertor()
}
