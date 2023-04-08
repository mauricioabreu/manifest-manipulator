package filter

import (
	"bytes"
	"math"

	"github.com/grafov/m3u8"
)

type MasterPlaylist struct {
	Playlist *m3u8.MasterPlaylist
}

type BandwidthFilter struct {
	Min int
	Max int
}

func NewMasterPlaylist(data bytes.Buffer) (*MasterPlaylist, error) {
	playlist, _, err := m3u8.Decode(data, false)
	if err != nil {
		return nil, err
	}
	return &MasterPlaylist{
		Playlist: playlist.(*m3u8.MasterPlaylist),
	}, nil
}

func (p *MasterPlaylist) FilterBandwidth(f BandwidthFilter) {
	max := f.Max
	if max <= 0 {
		max = math.MaxInt
	}
	variants := make([]*m3u8.Variant, 0)
	for _, variant := range p.Playlist.Variants {
		if int(variant.Bandwidth) >= f.Min && int(variant.Bandwidth) <= max {
			variants = append(variants, variant)
		}
	}
	p.Playlist.Variants = variants
}