package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type StartGame struct{}

func (pk *StartGame) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.StartGame{}
	input := standard.(*standardPacket.StartGame)

	p.EntityUniqueID = input.EntityUniqueID
	p.EntityRuntimeID = input.EntityRuntimeID
	p.PlayerGameMode = input.PlayerGameMode
	p.PlayerPosition = input.PlayerPosition
	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.WorldSeed = input.WorldSeed
	p.SpawnBiomeType = input.SpawnBiomeType
	p.UserDefinedBiomeName = input.UserDefinedBiomeName
	p.Dimension = input.Dimension
	p.Generator = input.Generator
	p.WorldGameMode = input.WorldGameMode
	p.Difficulty = input.Difficulty
	p.WorldSpawn = neteaseProtocol.BlockPos(input.WorldSpawn)
	p.AchievementsDisabled = input.AchievementsDisabled
	p.EditorWorld = input.EditorWorld
	p.CreatedInEditor = input.CreatedInEditor
	p.ExportedFromEditor = input.ExportedFromEditor
	p.DayCycleLockTime = input.DayCycleLockTime
	p.EducationEditionOffer = input.EducationEditionOffer
	p.EducationFeaturesEnabled = input.EducationFeaturesEnabled
	p.EducationProductID = input.EducationProductID
	p.RainLevel = input.RainLevel
	p.LightningLevel = input.LightningLevel
	p.ConfirmedPlatformLockedContent = input.ConfirmedPlatformLockedContent
	p.MultiPlayerGame = input.MultiPlayerGame
	p.LANBroadcastEnabled = input.LANBroadcastEnabled
	p.XBLBroadcastMode = input.XBLBroadcastMode
	p.PlatformBroadcastMode = input.PlatformBroadcastMode
	p.CommandsEnabled = input.CommandsEnabled
	p.TexturePackRequired = input.TexturePackRequired
	p.ExperimentsPreviouslyToggled = input.ExperimentsPreviouslyToggled
	p.BonusChestEnabled = input.BonusChestEnabled
	p.StartWithMapEnabled = input.StartWithMapEnabled
	p.PlayerPermissions = input.PlayerPermissions
	p.ServerChunkTickRadius = input.ServerChunkTickRadius
	p.HasLockedBehaviourPack = input.HasLockedBehaviourPack
	p.HasLockedTexturePack = input.HasLockedTexturePack
	p.FromLockedWorldTemplate = input.FromLockedWorldTemplate
	p.MSAGamerTagsOnly = input.MSAGamerTagsOnly
	p.FromWorldTemplate = input.FromWorldTemplate
	p.WorldTemplateSettingsLocked = input.WorldTemplateSettingsLocked
	p.OnlySpawnV1Villagers = input.OnlySpawnV1Villagers
	p.PersonaDisabled = input.PersonaDisabled
	p.CustomSkinsDisabled = input.CustomSkinsDisabled
	p.EmoteChatMuted = input.EmoteChatMuted
	p.BaseGameVersion = input.BaseGameVersion
	p.LimitedWorldWidth = input.LimitedWorldWidth
	p.LimitedWorldDepth = input.LimitedWorldDepth
	p.NewNether = input.NewNether
	p.EducationSharedResourceURI = neteaseProtocol.EducationSharedResourceURI(input.EducationSharedResourceURI)
	p.LevelID = input.LevelID
	p.WorldName = input.WorldName
	p.TemplateContentIdentity = input.TemplateContentIdentity
	p.Trial = input.Trial
	p.PlayerMovementSettings = neteaseProtocol.PlayerMovementSettings(input.PlayerMovementSettings)
	p.Time = input.Time
	p.EnchantmentSeed = input.EnchantmentSeed
	p.MultiPlayerCorrelationID = input.MultiPlayerCorrelationID
	p.ServerAuthoritativeInventory = input.ServerAuthoritativeInventory
	p.GameVersion = input.GameVersion
	p.PropertyData = input.PropertyData
	p.ServerBlockStateChecksum = input.ServerBlockStateChecksum
	p.ClientSideGeneration = input.ClientSideGeneration
	p.WorldTemplateID = input.WorldTemplateID
	p.ChatRestrictionLevel = input.ChatRestrictionLevel
	p.DisablePlayerInteractions = input.DisablePlayerInteractions
	p.UseBlockNetworkIDHashes = input.UseBlockNetworkIDHashes
	p.ServerAuthoritativeSound = input.ServerAuthoritativeSound

	forceExperimentalGameplay, has := input.ForceExperimentalGameplay.Value()
	if has {
		p.ForceExperimentalGameplay = neteaseProtocol.Option(forceExperimentalGameplay)
	}

	p.GameRules = ConvertSlice(
		input.GameRules,
		func(from standardProtocol.GameRule) neteaseProtocol.GameRule {
			return neteaseProtocol.GameRule(from)
		},
	)
	p.Experiments = ConvertSlice(
		input.Experiments,
		func(from standardProtocol.ExperimentData) neteaseProtocol.ExperimentData {
			return neteaseProtocol.ExperimentData(from)
		},
	)
	p.Blocks = ConvertSlice(
		input.Blocks,
		func(from standardProtocol.BlockEntry) neteaseProtocol.BlockEntry {
			return neteaseProtocol.BlockEntry(from)
		},
	)
	p.Items = ConvertSlice(
		input.Items,
		func(from standardProtocol.ItemEntry) neteaseProtocol.ItemEntry {
			return neteaseProtocol.ItemEntry(from)
		},
	)

	p.Unknown1 = false
	p.Unknown2 = ""
	p.Unknown3 = ""
	p.Unknown4 = false
	p.Unknown5 = false
	p.Unknown6 = 0
	p.Unknown7 = 0
	p.Unknown8 = 0
	p.Unknown9 = 0
	p.Unknown10 = 0
	p.Unknown11 = false
	p.Unknown12 = false
	p.Unknown13 = false
	p.Unknown14 = 0
	p.Unknown15 = false
	p.Unknown16 = false
	p.Unknown17 = false
	p.Unknown18 = false
	p.Unknown19 = false
	p.Unknown20 = false
	p.Unknown21 = false
	p.Unknown22 = false
	p.Unknown23 = false
	p.Unknown24 = make([]byte, 0)
	p.Unknown25 = false
	p.Unknown26 = false
	p.Unknown27 = false
	p.Unknown28 = ""
	p.Unknown29 = false
	p.Unknown30 = false

	return &p
}

func (pk *StartGame) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.StartGame{}
	input := netease.(*neteasePacket.StartGame)

	p.EntityUniqueID = input.EntityUniqueID
	p.EntityRuntimeID = input.EntityRuntimeID
	p.PlayerGameMode = input.PlayerGameMode
	p.PlayerPosition = input.PlayerPosition
	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.WorldSeed = input.WorldSeed
	p.SpawnBiomeType = input.SpawnBiomeType
	p.UserDefinedBiomeName = input.UserDefinedBiomeName
	p.Dimension = input.Dimension
	p.Generator = input.Generator
	p.WorldGameMode = input.WorldGameMode
	p.Difficulty = input.Difficulty
	p.WorldSpawn = standardProtocol.BlockPos(input.WorldSpawn)
	p.AchievementsDisabled = input.AchievementsDisabled
	p.EditorWorld = input.EditorWorld
	p.CreatedInEditor = input.CreatedInEditor
	p.ExportedFromEditor = input.ExportedFromEditor
	p.DayCycleLockTime = input.DayCycleLockTime
	p.EducationEditionOffer = input.EducationEditionOffer
	p.EducationFeaturesEnabled = input.EducationFeaturesEnabled
	p.EducationProductID = input.EducationProductID
	p.RainLevel = input.RainLevel
	p.LightningLevel = input.LightningLevel
	p.ConfirmedPlatformLockedContent = input.ConfirmedPlatformLockedContent
	p.MultiPlayerGame = input.MultiPlayerGame
	p.LANBroadcastEnabled = input.LANBroadcastEnabled
	p.XBLBroadcastMode = input.XBLBroadcastMode
	p.PlatformBroadcastMode = input.PlatformBroadcastMode
	p.CommandsEnabled = input.CommandsEnabled
	p.TexturePackRequired = input.TexturePackRequired
	p.ExperimentsPreviouslyToggled = input.ExperimentsPreviouslyToggled
	p.BonusChestEnabled = input.BonusChestEnabled
	p.StartWithMapEnabled = input.StartWithMapEnabled
	p.PlayerPermissions = input.PlayerPermissions
	p.ServerChunkTickRadius = input.ServerChunkTickRadius
	p.HasLockedBehaviourPack = input.HasLockedBehaviourPack
	p.HasLockedTexturePack = input.HasLockedTexturePack
	p.FromLockedWorldTemplate = input.FromLockedWorldTemplate
	p.MSAGamerTagsOnly = input.MSAGamerTagsOnly
	p.FromWorldTemplate = input.FromWorldTemplate
	p.WorldTemplateSettingsLocked = input.WorldTemplateSettingsLocked
	p.OnlySpawnV1Villagers = input.OnlySpawnV1Villagers
	p.PersonaDisabled = input.PersonaDisabled
	p.CustomSkinsDisabled = input.CustomSkinsDisabled
	p.EmoteChatMuted = input.EmoteChatMuted
	p.BaseGameVersion = input.BaseGameVersion
	p.LimitedWorldWidth = input.LimitedWorldWidth
	p.LimitedWorldDepth = input.LimitedWorldDepth
	p.NewNether = input.NewNether
	p.EducationSharedResourceURI = standardProtocol.EducationSharedResourceURI(input.EducationSharedResourceURI)
	p.LevelID = input.LevelID
	p.WorldName = input.WorldName
	p.TemplateContentIdentity = input.TemplateContentIdentity
	p.Trial = input.Trial
	p.PlayerMovementSettings = standardProtocol.PlayerMovementSettings(input.PlayerMovementSettings)
	p.Time = input.Time
	p.EnchantmentSeed = input.EnchantmentSeed
	p.MultiPlayerCorrelationID = input.MultiPlayerCorrelationID
	p.ServerAuthoritativeInventory = input.ServerAuthoritativeInventory
	p.GameVersion = input.GameVersion
	p.PropertyData = input.PropertyData
	p.ClientSideGeneration = input.ClientSideGeneration
	p.WorldTemplateID = input.WorldTemplateID
	p.ChatRestrictionLevel = input.ChatRestrictionLevel
	p.DisablePlayerInteractions = input.DisablePlayerInteractions
	p.UseBlockNetworkIDHashes = input.UseBlockNetworkIDHashes
	p.ServerAuthoritativeSound = input.ServerAuthoritativeSound

	forceExperimentalGameplay, has := input.ForceExperimentalGameplay.Value()
	if has {
		p.ForceExperimentalGameplay = standardProtocol.Option(forceExperimentalGameplay)
	}

	p.GameRules = ConvertSlice(
		input.GameRules,
		func(from neteaseProtocol.GameRule) standardProtocol.GameRule {
			return standardProtocol.GameRule(from)
		},
	)
	p.Experiments = ConvertSlice(
		input.Experiments,
		func(from neteaseProtocol.ExperimentData) standardProtocol.ExperimentData {
			return standardProtocol.ExperimentData(from)
		},
	)
	p.Blocks = ConvertSlice(
		input.Blocks,
		func(from neteaseProtocol.BlockEntry) standardProtocol.BlockEntry {
			return standardProtocol.BlockEntry(from)
		},
	)
	p.Items = ConvertSlice(
		input.Items,
		func(from neteaseProtocol.ItemEntry) standardProtocol.ItemEntry {
			return standardProtocol.ItemEntry(from)
		},
	)

	p.ServerBlockStateChecksum = 0

	return &p
}
