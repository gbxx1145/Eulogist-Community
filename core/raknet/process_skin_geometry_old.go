package RaknetConnection

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	_ "embed"
)

type SkinNumberOld json.Number

func (s *SkinNumberOld) MarshalJSON() ([]byte, error) {
	str := string(*s)
	if strings.Contains(str, ".") && !strings.HasSuffix(str, "0") {
		str += "0"
	}
	return json.Marshal(json.Number(str))
}

func (s *SkinNumberOld) UnmarshalJSON(data []byte) error {
	var num json.Number
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}
	*s = SkinNumberOld(num.String())
	return nil
}

type SkinCubeOld struct {
	Inflate SkinNumberOld   `json:"inflate"`
	Mirror  bool            `json:"mirror"`
	Origin  []SkinNumberOld `json:"origin"`
	Size    []SkinNumberOld `json:"size"`
	Uv      []SkinNumberOld `json:"uv"`
}

type SkinGeometryBoneOld struct {
	Cubes         *[]SkinCubeOld  `json:"cubes,omitempty"`
	Name          string          `json:"name"`
	Parent        string          `json:"parent,omitempty"`
	Pivot         []SkinNumberOld `json:"pivot"`
	RenderGroupID int             `json:"render_group_id,omitempty"`
	Rotation      []SkinNumberOld `json:"rotation,omitempty"`
}

type SkinGeometryOld struct {
	Bones               []*SkinGeometryBoneOld `json:"bones"`
	TextureHeight       int                    `json:"textureheight"`
	TextureWidth        int                    `json:"texturewidth"`
	VisibleBoundsHeight SkinNumberOld          `json:"visible_bounds_height,omitempty"`
	VisibleBoundsOffset []SkinNumberOld        `json:"visible_bounds_offset,omitempty"`
	VisibleBoundsWidth  SkinNumberOld          `json:"visible_bounds_width,omitempty"`
}

func ProcessOldGeometry(skin *Skin, geometryName string, skinGeometry json.RawMessage) (err error) {
	/* Layer 2 */
	geometry := &SkinGeometryOld{}
	if err = json.Unmarshal(skinGeometry, geometry); err != nil {
		return fmt.Errorf("ProcessGeometry: %v", err)
	}
	// handle bones
	hasRoot := false
	renderGroupNames := []string{"leftArm", "rightArm"}
	for _, bone := range geometry.Bones {
		// setup parent
		switch bone.Name {
		case "waist", "leftLeg", "rightLeg":
			bone.Parent = "root"
		case "head":
			bone.Parent = "body"
		case "leftArm", "rightArm":
			bone.Parent = "body"
			bone.RenderGroupID = 1
		case "body":
			bone.Parent = "waist"
		case "root":
			hasRoot = true
		}
		// setup render group
		if slices.Contains(renderGroupNames, bone.Parent) {
			bone.RenderGroupID = 1
			renderGroupNames = append(renderGroupNames, bone.Name)
		}
	}
	if !hasRoot {
		geometry.Bones = append(geometry.Bones, &SkinGeometryBoneOld{
			Name: "root",
			Pivot: []SkinNumberOld{
				"0.0",
				"0.0",
				"0.0",
			},
		})
	}
	// return
	skin.SkinGeometry, _ = json.Marshal(map[string]any{
		"format_version": "1.8.0",
		geometryName:     geometry,
	})
	return
}
