package protocol

const (
	EventTypeAchievementAwarded = iota
	EventTypeEntityInteract
	EventTypePortalBuilt
	EventTypePortalUsed
	EventTypeMobKilled
	EventTypeCauldronUsed
	EventTypePlayerDied
	EventTypeBossKilled
	EventTypeAgentCommand
	EventTypeAgentCreated
	EventTypePatternRemoved
	EventTypeSlashCommandExecuted
	EventTypeFishBucketed
	EventTypeMobBorn
	EventTypePetDied
	EventTypeCauldronInteract
	EventTypeComposterInteract
	EventTypeBellUsed
	EventTypeEntityDefinitionTrigger
	EventTypeRaidUpdate
	EventTypeMovementAnomaly
	EventTypeMovementCorrected
	EventTypeExtractHoney
	EventTypeTargetBlockHit
	EventTypePiglinBarter
	EventTypePlayerWaxedOrUnwaxedCopper
	EventTypeCodeBuilderRuntimeAction
	EventTypeCodeBuilderScoreboard
	EventTypeStriderRiddenInLavaInOverworld
	EventTypeSneakCloseToSculkSensor
	EventTypeCarefulRestoration
)

// lookupEvent looks up an Event matching the event type passed. False is
// returned if no such event data exists.
func lookupEvent(eventType int32, x *Event) bool {
	switch eventType {
	case EventTypeAchievementAwarded:
		*x = &AchievementAwardedEvent{}
	case EventTypeEntityInteract:
		*x = &EntityInteractEvent{}
	case EventTypePortalBuilt:
		*x = &PortalBuiltEvent{}
	case EventTypePortalUsed:
		*x = &PortalUsedEvent{}
	case EventTypeMobKilled:
		*x = &MobKilledEvent{}
	case EventTypeCauldronUsed:
		*x = &CauldronUsedEvent{}
	case EventTypePlayerDied:
		*x = &PlayerDiedEvent{}
	case EventTypeBossKilled:
		*x = &BossKilledEvent{}
	case EventTypeAgentCommand:
		*x = &AgentCommandEvent{}

	// PhoenixBuilder specific changes.
	// Author: Liliya233
	case EventTypeAgentCreated:
		*x = &AgentCreatedEvent{}

	case EventTypePatternRemoved:
		*x = &PatternRemovedEvent{}
	case EventTypeSlashCommandExecuted:
		*x = &SlashCommandExecutedEvent{}
	case EventTypeFishBucketed:
		*x = &FishBucketedEvent{}
	case EventTypeMobBorn:
		*x = &MobBornEvent{}
	case EventTypePetDied:
		*x = &PetDiedEvent{}
	case EventTypeCauldronInteract:
		*x = &CauldronInteractEvent{}
	case EventTypeComposterInteract:
		*x = &ComposterInteractEvent{}
	case EventTypeBellUsed:
		*x = &BellUsedEvent{}
	case EventTypeEntityDefinitionTrigger:
		*x = &EntityDefinitionTriggerEvent{}
	case EventTypeRaidUpdate:
		*x = &RaidUpdateEvent{}
	case EventTypeMovementAnomaly:
		*x = &MovementAnomalyEvent{}
	case EventTypeMovementCorrected:
		*x = &MovementCorrectedEvent{}
	case EventTypeExtractHoney:
		*x = &ExtractHoneyEvent{}
	case EventTypeTargetBlockHit:
		*x = &TargetBlockHitEvent{}
	case EventTypePiglinBarter:
		*x = &PiglinBarterEvent{}
	case EventTypePlayerWaxedOrUnwaxedCopper:
		*x = &WaxedOrUnwaxedCopperEvent{}
	case EventTypeCodeBuilderRuntimeAction:
		*x = &CodeBuilderRuntimeActionEvent{}
	case EventTypeCodeBuilderScoreboard:
		*x = &CodeBuilderScoreboardEvent{}
	case EventTypeStriderRiddenInLavaInOverworld:
		*x = &StriderRiddenInLavaInOverworldEvent{}
	case EventTypeSneakCloseToSculkSensor:
		*x = &SneakCloseToSculkSensorEvent{}
	case EventTypeCarefulRestoration:
		*x = &CarefulRestorationEvent{}
	default:
		return false
	}
	return true
}

// lookupEventType looks up an event type that matches the Event passed.
func lookupEventType(x Event, eventType *int32) bool {
	switch x.(type) {
	case *AchievementAwardedEvent:
		*eventType = EventTypeAchievementAwarded
	case *EntityInteractEvent:
		*eventType = EventTypeEntityInteract
	case *PortalBuiltEvent:
		*eventType = EventTypePortalBuilt
	case *PortalUsedEvent:
		*eventType = EventTypePortalUsed
	case *MobKilledEvent:
		*eventType = EventTypeMobKilled
	case *CauldronUsedEvent:
		*eventType = EventTypeCauldronUsed
	case *PlayerDiedEvent:
		*eventType = EventTypePlayerDied
	case *BossKilledEvent:
		*eventType = EventTypeBossKilled
	case *AgentCommandEvent:
		*eventType = EventTypeAgentCommand

	// PhoenixBuilder specific changes.
	// Author: Happy2018new
	case *AgentCreatedEvent:
		*eventType = EventTypeAgentCreated

	case *PatternRemovedEvent:
		*eventType = EventTypePatternRemoved
	case *SlashCommandExecutedEvent:
		*eventType = EventTypeSlashCommandExecuted
	case *FishBucketedEvent:
		*eventType = EventTypeFishBucketed
	case *MobBornEvent:
		*eventType = EventTypeMobBorn
	case *PetDiedEvent:
		*eventType = EventTypePetDied
	case *CauldronInteractEvent:
		*eventType = EventTypeCauldronInteract
	case *ComposterInteractEvent:
		*eventType = EventTypeComposterInteract
	case *BellUsedEvent:
		*eventType = EventTypeBellUsed
	case *EntityDefinitionTriggerEvent:
		*eventType = EventTypeEntityDefinitionTrigger
	case *RaidUpdateEvent:
		*eventType = EventTypeRaidUpdate
	case *MovementAnomalyEvent:
		*eventType = EventTypeMovementAnomaly
	case *MovementCorrectedEvent:
		*eventType = EventTypeMovementCorrected
	case *ExtractHoneyEvent:
		*eventType = EventTypeExtractHoney
	case *TargetBlockHitEvent:
		*eventType = EventTypeTargetBlockHit
	case *PiglinBarterEvent:
		*eventType = EventTypePiglinBarter
	case *WaxedOrUnwaxedCopperEvent:
		*eventType = EventTypePlayerWaxedOrUnwaxedCopper
	case *CodeBuilderRuntimeActionEvent:
		*eventType = EventTypeCodeBuilderRuntimeAction
	case *CodeBuilderScoreboardEvent:
		*eventType = EventTypeCodeBuilderScoreboard
	case *StriderRiddenInLavaInOverworldEvent:
		*eventType = EventTypeStriderRiddenInLavaInOverworld
	case *SneakCloseToSculkSensorEvent:
		*eventType = EventTypeSneakCloseToSculkSensor
	case *CarefulRestorationEvent:
		*eventType = EventTypeCarefulRestoration
	default:
		return false
	}
	return true
}

// Event represents an object that holds data specific to an event.
// The data it holds depends on the type.
type Event interface {
	// Marshal encodes/decodes a serialised event data object.
	Marshal(r IO)
}

// AchievementAwardedEvent is the event data sent for achievements.
type AchievementAwardedEvent struct {
	// AchievementID is the ID for the achievement.
	AchievementID int32
}

// Marshal ...
func (a *AchievementAwardedEvent) Marshal(r IO) {
	r.Varint32(&a.AchievementID)
}

// EntityInteractEvent is the event data sent for entity interactions.
type EntityInteractEvent struct {
	// InteractionType ...
	InteractionType int32
	// InteractionEntityType ...
	InteractionEntityType int32
	// EntityVariant ...
	EntityVariant int32
	// EntityColour ...
	EntityColour uint8
}

// Marshal ...
func (e *EntityInteractEvent) Marshal(r IO) {
	r.Varint32(&e.InteractionType)
	r.Varint32(&e.InteractionEntityType)
	r.Varint32(&e.EntityVariant)
	r.Uint8(&e.EntityColour)
}

// PortalBuiltEvent is the event data sent when a portal is built.
type PortalBuiltEvent struct {
	// DimensionID ...
	DimensionID int32
}

// Marshal ...
func (p *PortalBuiltEvent) Marshal(r IO) {
	r.Varint32(&p.DimensionID)
}

// PortalUsedEvent is the event data sent when a portal is used.
type PortalUsedEvent struct {
	// FromDimensionID ...
	FromDimensionID int32
	// ToDimensionID ...
	ToDimensionID int32
}

// Marshal ...
func (p *PortalUsedEvent) Marshal(r IO) {
	r.Varint32(&p.FromDimensionID)
	r.Varint32(&p.ToDimensionID)
}

// MobKilledEvent is the event data sent when a mob is killed.
type MobKilledEvent struct {
	// KillerEntityUniqueID ...
	KillerEntityUniqueID int64
	// VictimEntityUniqueID ...
	VictimEntityUniqueID int64
	// KillerEntityType ...
	KillerEntityType int32
	// EntityDamageCause ...
	EntityDamageCause int32
	// VillagerTradeTier -1 if not a trading actor.
	VillagerTradeTier int32
	// VillagerDisplayName Empty if not a trading actor.
	VillagerDisplayName string
}

// Marshal ...
func (m *MobKilledEvent) Marshal(r IO) {
	r.Varint64(&m.KillerEntityUniqueID)
	r.Varint64(&m.VictimEntityUniqueID)
	r.Varint32(&m.KillerEntityType)
	r.Varint32(&m.EntityDamageCause)
	r.Varint32(&m.VillagerTradeTier)
	r.String(&m.VillagerDisplayName)
}

// CauldronUsedEvent is the event data sent when a cauldron is used.
type CauldronUsedEvent struct {
	// PotionID ...
	PotionID int32
	// Colour ...
	Colour int32
	// FillLevel ...
	FillLevel int32
}

// Marshal ...
func (c *CauldronUsedEvent) Marshal(r IO) {
	r.Varint32(&c.PotionID)
	r.Varint32(&c.Colour)
	r.Varint32(&c.FillLevel)
}

// PlayerDiedEvent is the event data sent when a player dies.
type PlayerDiedEvent struct {
	// AttackerEntityID ...
	AttackerEntityID int32
	// AttackerVariant ...
	AttackerVariant int32
	// EntityDamageCause ...
	EntityDamageCause int32
	// InRaid ...
	InRaid bool
}

// Marshal ...
func (p *PlayerDiedEvent) Marshal(r IO) {
	r.Varint32(&p.AttackerEntityID)
	r.Varint32(&p.AttackerVariant)
	r.Varint32(&p.EntityDamageCause)
	r.Bool(&p.InRaid)
}

// BossKilledEvent is the event data sent when a boss dies.
type BossKilledEvent struct {
	// BossEntityUniqueID ...
	BossEntityUniqueID int64
	// PlayerPartySize ...
	PlayerPartySize int32
	// InteractionEntityType ...
	InteractionEntityType int32
}

// Marshal ...
func (b *BossKilledEvent) Marshal(r IO) {
	r.Varint64(&b.BossEntityUniqueID)
	r.Varint32(&b.PlayerPartySize)
	r.Varint32(&b.InteractionEntityType)
}

// AgentCommandEvent is an event used in Education Edition.
type AgentCommandEvent struct {
	// AgentResult ...
	AgentResult int32
	// DataValue ...
	DataValue int32
	// Command ...
	Command string
	// DataKey ...
	DataKey string
	// Output ...
	Output string
}

// Marshal ...
func (a *AgentCommandEvent) Marshal(r IO) {
	r.Varint32(&a.AgentResult)
	r.Varint32(&a.DataValue)
	r.String(&a.Command)
	r.String(&a.DataKey)
	r.String(&a.Output)
}

// PhoenixBuilder specific struct.
// Author: Liliya233
//
// AgentCreatedEvent does not have any data.
type AgentCreatedEvent struct{}

// PhoenixBuilder specific func.
// Author: Liliya233
//
// Marshal ...
func (a *AgentCreatedEvent) Marshal(r IO) {}

// PatternRemovedEvent is the event data sent when a pattern is removed.
type PatternRemovedEvent struct{}

// Marshal ...
func (p *PatternRemovedEvent) Marshal(r IO) {}

// SlashCommandExecutedEvent is the event data sent when a slash command is executed.
type SlashCommandExecutedEvent struct {
	// CommandName ...
	CommandName string
	// SuccessCount ...
	SuccessCount int32
	// MessageCount indicates the amount of OutputMessages present.
	MessageCount int32
	// OutputMessages is a list of messages joint with ;.
	OutputMessages string
}

// Marshal ...
func (s *SlashCommandExecutedEvent) Marshal(r IO) {
	r.Varint32(&s.SuccessCount)
	r.Varint32(&s.MessageCount)
	r.String(&s.CommandName)
	r.String(&s.OutputMessages)
}

// FishBucketedEvent is the event data sent when a fish is bucketed.
type FishBucketedEvent struct {
	// Pattern ...
	Pattern int32
	// Preset ...
	Preset int32
	// BucketedEntityType ...
	BucketedEntityType int32
	// Release ...
	Release bool
}

// Marshal ...
func (f *FishBucketedEvent) Marshal(r IO) {
	r.Varint32(&f.Pattern)
	r.Varint32(&f.Preset)
	r.Varint32(&f.BucketedEntityType)
	r.Bool(&f.Release)
}

// MobBornEvent is the event data sent when a mob is born.
type MobBornEvent struct {
	// EntityType ...
	EntityType int32
	// Variant ...
	Variant int32
	// Colour ...
	Colour uint8
}

// Marshal ...
func (m *MobBornEvent) Marshal(r IO) {
	r.Varint32(&m.EntityType)
	r.Varint32(&m.Variant)
	r.Uint8(&m.Colour)
}

// PetDiedEvent is the event data sent when a pet dies.
type PetDiedEvent struct{}

// Marshal ...
func (p *PetDiedEvent) Marshal(r IO) {}

// CauldronInteractEvent is the event data sent when a cauldron is interacted with.
type CauldronInteractEvent struct {
	// BlockInteractionType ...
	BlockInteractionType int32
	// ItemID ...
	ItemID int32
}

// Marshal ...
func (c *CauldronInteractEvent) Marshal(r IO) {
	r.Varint32(&c.BlockInteractionType)
	r.Varint32(&c.ItemID)
}

// ComposterInteractEvent is the event data sent when a composter is interacted with.
type ComposterInteractEvent struct {
	// BlockInteractionType ...
	BlockInteractionType int32
	// ItemID ...
	ItemID int32
}

// Marshal ...
func (c *ComposterInteractEvent) Marshal(r IO) {
	r.Varint32(&c.BlockInteractionType)
	r.Varint32(&c.ItemID)
}

// BellUsedEvent is the event data sent when a bell is used.
type BellUsedEvent struct {
	// ItemID ...
	ItemID int32
}

// Marshal ...
func (b *BellUsedEvent) Marshal(r IO) {
	r.Varint32(&b.ItemID)
}

// EntityDefinitionTriggerEvent is an event used for an unknown purpose.
type EntityDefinitionTriggerEvent struct {
	// EventName ...
	EventName string
}

// Marshal ...
func (e *EntityDefinitionTriggerEvent) Marshal(r IO) {
	r.String(&e.EventName)
}

// RaidUpdateEvent is an event used to update a raids progress client side.
type RaidUpdateEvent struct {
	// CurrentRaidWave ...
	CurrentRaidWave int32
	// TotalRaidWaves ...
	TotalRaidWaves int32
	// WonRaid ...
	WonRaid bool
}

// Marshal ...
func (ra *RaidUpdateEvent) Marshal(r IO) {
	r.Varint32(&ra.CurrentRaidWave)
	r.Varint32(&ra.TotalRaidWaves)
	r.Bool(&ra.WonRaid)
}

// MovementAnomalyEvent is an event used for updating the other party on movement data.
type MovementAnomalyEvent struct {
	// EventType ...
	EventType uint8
	// CheatingScore ...
	CheatingScore float32
	// AveragePositionDelta ...
	AveragePositionDelta float32
	// TotalPositionDelta ...
	TotalPositionDelta float32
	// MinPositionDelta ...
	MinPositionDelta float32
	// MaxPositionDelta ...
	MaxPositionDelta float32
}

// Marshal ...
func (m *MovementAnomalyEvent) Marshal(r IO) {
	r.Uint8(&m.EventType)
	r.Float32(&m.CheatingScore)
	r.Float32(&m.AveragePositionDelta)
	r.Float32(&m.TotalPositionDelta)
	r.Float32(&m.MinPositionDelta)
	r.Float32(&m.MaxPositionDelta)
}

// MovementCorrectedEvent is an event used to correct movement anomalies.
type MovementCorrectedEvent struct {
	// PositionDelta ...
	PositionDelta float32
	// CheatingScore ...
	CheatingScore float32
	// ScoreThreshold ...
	ScoreThreshold float32
	// DistanceThreshold ...
	DistanceThreshold float32
	// DurationThreshold ...
	DurationThreshold int32
}

// Marshal ...
func (m *MovementCorrectedEvent) Marshal(r IO) {
	r.Float32(&m.PositionDelta)
	r.Float32(&m.CheatingScore)
	r.Float32(&m.ScoreThreshold)
	r.Float32(&m.DistanceThreshold)
	r.Varint32(&m.DurationThreshold)
}

// ExtractHoneyEvent is an event used to extract honey from a hive.
type ExtractHoneyEvent struct{}

// Marshal ...
func (*ExtractHoneyEvent) Marshal(r IO) {}

// TargetBlockHitEvent is an event used when a target block is hit by a arrow.
type TargetBlockHitEvent struct {
	// RedstoneLevel ...
	RedstoneLevel int32
}

// Marshal ...
func (t *TargetBlockHitEvent) Marshal(r IO) {
	r.Varint32(&t.RedstoneLevel)
}

// PiglinBarterEvent is called when a player drops gold ingots to a piglin to initiate a trade for an item.
type PiglinBarterEvent struct {
	// ItemID ...
	ItemID int32
	// WasTargetingBarteringPlayer ...
	WasTargetingBarteringPlayer bool
}

// Marshal ...
func (p *PiglinBarterEvent) Marshal(r IO) {
	r.Varint32(&p.ItemID)
	r.Bool(&p.WasTargetingBarteringPlayer)
}

const (
	WaxNotOxidised   = uint16(0xa609)
	WaxExposed       = uint16(0xa809)
	WaxWeathered     = uint16(0xaa09)
	WaxOxidised      = uint16(0xac09)
	UnWaxNotOxidised = uint16(0xae09)
	UnWaxExposed     = uint16(0xb009)
	UnWaxWeathered   = uint16(0xb209)
	UnWaxOxidised    = uint16(0xfa0a)
)

// WaxedOrUnwaxedCopperEvent is an event sent by the server when a copper block is waxed or unwaxed.
type WaxedOrUnwaxedCopperEvent struct {
	Type uint16
}

// Marshal ...
func (w *WaxedOrUnwaxedCopperEvent) Marshal(r IO) {
	r.Uint16(&w.Type)
}

// CodeBuilderRuntimeActionEvent is an event sent by the server when a code builder runtime action is performed.
type CodeBuilderRuntimeActionEvent struct {
	// Action ...
	Action string
}

// Marshal ...
func (c *CodeBuilderRuntimeActionEvent) Marshal(r IO) {
	r.String(&c.Action)
}

// CodeBuilderScoreboardEvent is an event sent by the server when a code builder scoreboard is updated.
type CodeBuilderScoreboardEvent struct {
	// ObjectiveName ...
	ObjectiveName string
	// Score ...
	Score int32
}

// Marshal ...
func (c *CodeBuilderScoreboardEvent) Marshal(r IO) {
	r.String(&c.ObjectiveName)
	r.Varint32(&c.Score)
}

// StriderRiddenInLavaInOverworldEvent is an event sent by the server when a strider is ridden in lava in the overworld.
type StriderRiddenInLavaInOverworldEvent struct{}

// Marshal ...
func (s *StriderRiddenInLavaInOverworldEvent) Marshal(r IO) {}

// SneakCloseToSculkSensorEvent is an event sent by the server when a player sneaks close to a sculk sensor.
type SneakCloseToSculkSensorEvent struct{}

// Marshal ...
func (*SneakCloseToSculkSensorEvent) Marshal(r IO) {}

// CarefulRestorationEvent is an event sent by the server when a player performs a careful restoration.
type CarefulRestorationEvent struct{}

// Marshal ...
func (c *CarefulRestorationEvent) Marshal(r IO) {}
