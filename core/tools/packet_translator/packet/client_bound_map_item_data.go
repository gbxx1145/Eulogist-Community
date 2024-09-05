package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"
	"image/color"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type ClientBoundMapItemData struct{}

func (pk *ClientBoundMapItemData) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.ClientBoundMapItemData{}
	input := standard.(*standardPacket.ClientBoundMapItemData)

	p.MapID = input.MapID
	p.UpdateFlags = input.UpdateFlags
	p.Dimension = input.Dimension
	p.LockedMap = input.LockedMap
	p.Origin = neteaseProtocol.BlockPos(input.Origin)
	p.Scale = input.Scale
	p.MapsIncludedIn = input.MapsIncludedIn
	p.Height = input.Height
	p.Width = input.Width
	p.XOffset = input.XOffset
	p.YOffset = input.YOffset

	p.TrackedObjects = packet_translate_struct.ConvertSlice(
		input.TrackedObjects,
		func(from standardProtocol.MapTrackedObject) neteaseProtocol.MapTrackedObject {
			return neteaseProtocol.MapTrackedObject{
				Type:           from.Type,
				EntityUniqueID: from.EntityUniqueID,
				BlockPosition:  neteaseProtocol.BlockPos(from.BlockPosition),
			}
		},
	)
	p.Decorations = packet_translate_struct.ConvertSlice(
		input.Decorations,
		func(from standardProtocol.MapDecoration) neteaseProtocol.MapDecoration {
			return neteaseProtocol.MapDecoration(from)
		},
	)

	if len(input.Pixels) == 0 {
		return &p
	}

	colorMapA := make(map[color.RGBA]int)
	colorIndex := -1
	for _, value := range input.Pixels {
		colorIndex++
		colorMapA[value] = colorIndex
	}

	colorMapB := make(map[int]color.RGBA)
	for key, value := range colorMapA {
		colorMapB[value] = key
	}

	if len(colorMapA) <= 255 {
		pixels := neteaseProtocol.Uint8Pixels{
			Pixels:   make([]uint8, 0),
			ColorMap: make([]neteaseProtocol.Uint8ColorMap, 0),
		}
		for i := 0; i < len(colorMapB); i++ {
			pixels.ColorMap = append(pixels.ColorMap, neteaseProtocol.Uint8ColorMap{
				Colour: colorMapB[i],
				Index:  uint8(i),
			})
		}
		for _, value := range input.Pixels {
			pixels.Pixels = append(pixels.Pixels, uint8(colorMapA[value]))
		}
		p.Pixels = neteaseProtocol.MapPixels{
			IsEmpty: false,
			Data:    &pixels,
		}
	}

	if len(colorMapA) <= 65535 {
		pixels := neteaseProtocol.Uint16Pixels{
			Pixels:   make([]uint16, 0),
			ColorMap: make([]neteaseProtocol.Uint16ColorMap, 0),
		}
		for i := 0; i < len(colorMapB); i++ {
			pixels.ColorMap = append(pixels.ColorMap, neteaseProtocol.Uint16ColorMap{
				Colour: colorMapB[i],
				Index:  uint16(i),
			})
		}
		for _, value := range input.Pixels {
			pixels.Pixels = append(pixels.Pixels, uint16(colorMapA[value]))
		}
		p.Pixels = neteaseProtocol.MapPixels{
			IsEmpty: false,
			Data:    &pixels,
		}
	}

	if len(colorMapA) > 65535 {
		p.Pixels = neteaseProtocol.MapPixels{
			IsEmpty: false,
			Data: &neteaseProtocol.StandardPixels{
				Pixels: input.Pixels,
			},
		}
	}

	return &p
}

func (pk *ClientBoundMapItemData) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.ClientBoundMapItemData{}
	input := netease.(*neteasePacket.ClientBoundMapItemData)

	p.MapID = input.MapID
	p.UpdateFlags = input.UpdateFlags
	p.Dimension = input.Dimension
	p.LockedMap = input.LockedMap
	p.Origin = standardProtocol.BlockPos(input.Origin)
	p.Scale = input.Scale
	p.MapsIncludedIn = input.MapsIncludedIn
	p.Height = input.Height
	p.Width = input.Width
	p.XOffset = input.XOffset
	p.YOffset = input.YOffset

	p.TrackedObjects = packet_translate_struct.ConvertSlice(
		input.TrackedObjects,
		func(from neteaseProtocol.MapTrackedObject) standardProtocol.MapTrackedObject {
			return standardProtocol.MapTrackedObject{
				Type:           from.Type,
				EntityUniqueID: from.EntityUniqueID,
				BlockPosition:  standardProtocol.BlockPos(from.BlockPosition),
			}
		},
	)
	p.Decorations = packet_translate_struct.ConvertSlice(
		input.Decorations,
		func(from neteaseProtocol.MapDecoration) standardProtocol.MapDecoration {
			return standardProtocol.MapDecoration(from)
		},
	)

	if input.Pixels.IsEmpty {
		return &p
	} else {
		p.Pixels = make([]color.RGBA, 0)
	}

	switch data := input.Pixels.Data.(type) {
	case *neteaseProtocol.Uint8Pixels:
		colorMap := make(map[int]color.RGBA)
		for _, value := range data.ColorMap {
			colorMap[int(value.Index)] = value.Colour
		}
		for _, value := range data.Pixels {
			p.Pixels = append(p.Pixels, colorMap[int(value)])
		}
	case *neteaseProtocol.Uint16Pixels:
		colorMap := make(map[int]color.RGBA)
		for _, value := range data.ColorMap {
			colorMap[int(value.Index)] = value.Colour
		}
		for _, value := range data.Pixels {
			p.Pixels = append(p.Pixels, colorMap[int(value)])
		}
	case *neteaseProtocol.StandardPixels:
		p.Pixels = data.Pixels
	}

	return &p
}
