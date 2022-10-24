package helper

import (
	"strconv"
	"testing"
	"time"

	"github.com/alibaba/ilogtail/pkg/protocol"
	"github.com/stretchr/testify/assert"
)

func TestReviseFileOffset(t *testing.T) {
	var mock protocol.Log
	mock.Contents = append(mock.Contents, &protocol.Log_Content{Key: FileOffsetKey, Value: "1000"})
	mock.Time = uint32(time.Now().Unix())

	tests := []struct {
		name   string
		log    *protocol.Log
		offset int64
		want   int64
	}{
		{
			name:   "revise_file",
			log:    &mock,
			offset: 101,
			want:   1101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviseFileOffset(tt.log, 101)
			cont := getFileOffsetTag(tt.log)
			assert.True(t, cont != nil)
			off, err := strconv.ParseInt(cont.Value, 10, 64)
			assert.True(t, err == nil)
			assert.Equal(t, tt.want, off)
		})
	}
}

func getFileOffsetTag(log *protocol.Log) *protocol.Log_Content {
	for _, cont := range log.Contents {
		if cont.Key == FileOffsetKey {
			return cont
		}
	}
	return nil
}
