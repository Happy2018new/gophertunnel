package main

import (
	"fmt"
	"strings"

	dwf_block_general "github.com/TriM-Organization/dream-weaver-factory/block/general"
	_ "github.com/TriM-Organization/dream-weaver-factory/block/std"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// AppendCustomEntity ..
func AppendCustomEntity(pk *packet.AvailableActorIdentifiers) error {
	var m map[string]any

	err := nbt.UnmarshalEncoding(pk.SerialisedEntityIdentifiers, &m, nbt.NetworkLittleEndian)
	if err != nil {
		return fmt.Errorf("AppendCustomEntity: %v", err)
	}

	actors, _ := m["idlist"].([]any)
	actors = append(actors, map[string]any{"id": "rtx:falling_pointed_dripstone_base"})
	actors = append(actors, map[string]any{"id": "rtx:falling_chipped_anvil"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_3"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_7"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_2"})
	actors = append(actors, map[string]any{"id": "rtx:falling_light_gray_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_green_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_black_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_red_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_damaged_anvil"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_6"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_5"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_4"})
	actors = append(actors, map[string]any{"id": "rtx:falling_brown_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_1"})
	actors = append(actors, map[string]any{"id": "rtx:falling_snow_layer_0"})
	actors = append(actors, map[string]any{"id": "rtx:falling_purple_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_pointed_dripstone_tip"})
	actors = append(actors, map[string]any{"id": "rtx:falling_pointed_dripstone_middle"})
	actors = append(actors, map[string]any{"id": "rtx:falling_pointed_dripstone_frustum"})
	actors = append(actors, map[string]any{"id": "rtx:falling_anvil"})
	actors = append(actors, map[string]any{"id": "rtx:falling_suspicious_sand"})
	actors = append(actors, map[string]any{"id": "rtx:falling_suspicious_gravel"})
	actors = append(actors, map[string]any{"id": "rtx:falling_scaffolding"})
	actors = append(actors, map[string]any{"id": "rtx:falling_sand"})
	actors = append(actors, map[string]any{"id": "rtx:falling_yellow_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_red_sand"})
	actors = append(actors, map[string]any{"id": "rtx:falling_gravel"})
	actors = append(actors, map[string]any{"id": "rtx:falling_white_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_dragon_egg"})
	actors = append(actors, map[string]any{"id": "rtx:falling_pink_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_orange_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_magenta_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_lime_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_light_blue_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_gray_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_cyan_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_blue_concrete_powder"})
	actors = append(actors, map[string]any{"id": "rtx:falling_chipped_anvil_rotated"})
	actors = append(actors, map[string]any{"id": "rtx:falling_damaged_anvil_rotated"})
	actors = append(actors, map[string]any{"id": "rtx:falling_anvil_rotated"})
	m["idlist"] = actors

	result, err := nbt.MarshalEncoding(m, nbt.NetworkLittleEndian)
	if err != nil {
		return fmt.Errorf("AppendCustomEntity: %v", err)
	}
	pk.SerialisedEntityIdentifiers = result

	return nil
}

// RedirectFallingBlockID ..
func RedirectFallingBlockID(entityType string, metadata map[uint32]any) (entityID string, found bool) {
	if entityType != "minecraft:falling_block" || metadata == nil {
		return
	}

	blockRuntimeID, ok := metadata[protocol.EntityDataKeyVariant].(int32)
	if !ok {
		return "", false
	}
	blockName, blockStates, found := dwf_block_general.StdRuntimeIDToState(uint32(blockRuntimeID))
	if !found {
		return "", false
	}

	switch blockName {
	case "minecraft:snow_layer":
		entityID = fmt.Sprintf("rtx:falling_snow_layer_%d", blockStates["height"])
	case "minecraft:anvil", "minecraft:chipped_anvil", "minecraft:damaged_anvil":
		direction, _ := blockStates["minecraft:cardinal_direction"].(string)
		if direction == "north" || direction == "south" {
			entityID = fmt.Sprintf("rtx:falling_%s_rotated", strings.TrimPrefix(blockName, "minecraft:"))
		}
	case "minecraft:pointed_dripstone":
		thickness, _ := blockStates["dripstone_thickness"].(string)
		if thickness == "merge" {
			thickness = "tip"
		}
		entityID = fmt.Sprintf("rtx:falling_pointed_dripstone_%v", thickness)
	}
	if len(entityID) == 0 {
		entityID = fmt.Sprintf("rtx:falling_%s", strings.TrimPrefix(blockName, "minecraft:"))
	}

	return entityID, true
}
