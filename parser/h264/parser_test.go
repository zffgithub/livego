package h264

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestH264SeqDemux(t *testing.T) {
	at := assert.New(t)
	seq := []byte{
		0x01, 0x4d, 0x00, 0x1e, 0xff, 0xe1, 0x00, 0x17, 0x67, 0x4d, 0x00,
		0x1e, 0xab, 0x40, 0x5a, 0x12, 0x6c, 0x09, 0x28, 0x28, 0x28, 0x2f,
		0x80, 0x00, 0x01, 0xf4, 0x00, 0x00, 0x61, 0xa8, 0x4a, 0x01, 0x00,
		0x04, 0x68, 0xde, 0x31, 0x12,
	}
	d := NewParser()
	w := bytes.NewBuffer(nil)
	err := d.Parse(seq, true, w)
	at.Equal(err, nil)
	at.Equal(d.specificInfo, []byte{0x00, 0x00, 0x00, 0x01, 0x67, 0x4d, 0x00,
		0x1e, 0xab, 0x40, 0x5a, 0x12, 0x6c, 0x09, 0x28, 0x28, 0x28, 0x2f,
		0x80, 0x00, 0x01, 0xf4, 0x00, 0x00, 0x61, 0xa8, 0x4a, 0x00, 0x00, 0x00, 0x01, 0x68, 0xde, 0x31, 0x12})
}

func TestH264AnnexbDemux(t *testing.T) {
	at := assert.New(t)
	nalu := []byte{
		0x00, 0x00, 0x00, 0x01, 0x67, 0x4d, 0x00, 0x1e, 0xab, 0x40, 0x5a, 0x12, 0x6c, 0x09, 0x28, 0x28,
		0x28, 0x2f, 0x80, 0x00, 0x01, 0xf4, 0x00, 0x00, 0x61, 0xa8, 0x4a, 0x00, 0x00, 0x00, 0x01, 0x68,
		0xde, 0x31, 0x12, 0x00, 0x00, 0x00, 0x01, 0x65, 0x23,
	}
	d := NewParser()
	w := bytes.NewBuffer(nil)
	err := d.Parse(nalu, false, w)
	at.Equal(err, nil)
	at.Equal(w.Len(), 41)
}

func TestH264NalueSizeException(t *testing.T) {
	at := assert.New(t)
	nalu := []byte{
		0x00, 0x00, 0x10,
	}
	d := NewParser()
	w := bytes.NewBuffer(nil)
	err := d.Parse(nalu, false, w)
	at.Equal(err, errors.New("video data not match"))
}

func TestH264Mp4Demux(t *testing.T) {
	at := assert.New(t)
	nalu := []byte{
		0x00, 0x00, 0x00, 0x17, 0x67, 0x4d, 0x00, 0x1e, 0xab, 0x40, 0x5a, 0x12, 0x6c, 0x09, 0x28, 0x28,
		0x28, 0x2f, 0x80, 0x00, 0x01, 0xf4, 0x00, 0x00, 0x61, 0xa8, 0x4a, 0x00, 0x00, 0x00, 0x04, 0x68,
		0xde, 0x31, 0x12, 0x00, 0x00, 0x00, 0x02, 0x65, 0x23,
	}
	d := NewParser()
	w := bytes.NewBuffer(nil)
	err := d.Parse(nalu, false, w)
	at.Equal(err, nil)
	at.Equal(w.Len(), 47)
	at.Equal(w.Bytes(), []byte{0x00, 0x00, 0x00, 0x01, 0x09, 0xf0, 0x00, 0x00, 0x00, 0x01, 0x67, 0x4d, 0x00, 0x1e, 0xab, 0x40, 0x5a, 0x12, 0x6c, 0x09, 0x28, 0x28,
		0x28, 0x2f, 0x80, 0x00, 0x01, 0xf4, 0x00, 0x00, 0x61, 0xa8, 0x4a, 0x00, 0x00, 0x00, 0x01, 0x68,
		0xde, 0x31, 0x12, 0x00, 0x00, 0x00, 0x01, 0x65, 0x23})
}

func TestH264Mp4DemuxException1(t *testing.T) {
	at := assert.New(t)
	nalu := []byte{
		0x00, 0x00, 0x00, 0x29, 0x00, 0x00, 0x00,
	}
	d := NewParser()
	w := bytes.NewBuffer(nil)

	err := d.Parse(nalu, false, w)
	at.Equal(err, naluBodyLenError)
}

func TestH264Mp4DemuxException2(t *testing.T) {
	at := assert.New(t)
	nalu := []byte{
		0x00, 0x00, 0x00, 0x29, 0x00, 0x00, 0x00, 0x17, 0x67, 0x4d, 0x00, 0x1e, 0xab, 0x40, 0x5a, 0x12, 0x6c, 0x09, 0x28, 0x28,
		0x28, 0x2f, 0x80, 0x00, 0x01, 0xf4, 0x00, 0x00, 0x61, 0xa8, 0x4a, 0x00, 0x00, 0x00,
	}
	d := NewParser()
	w := bytes.NewBuffer(nil)
	err := d.Parse(nalu, false, w)
	at.Equal(err, naluBodyLenError)
}
