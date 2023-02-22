package constants

import "testing"

func TestGetIDFromUserMsgKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		wantId int64
	}{
		{
			name:   "test1",
			args:   args{key: "user_message_1"},
			wantId: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotId := GetIDFromUserMsgKey(tt.args.key); gotId != tt.wantId {
				t.Errorf("GetIDFromUserMsgKey() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGetIDFromUserLikeListKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		wantId int64
	}{
		{
			name:   "test1",
			args:   args{key: "user_like_list_1"},
			wantId: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotId := GetIDFromUserLikeListKey(tt.args.key); gotId != tt.wantId {
				t.Errorf("GetIDFromUserLikeListKey() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGetIDFromUserFollowListKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		wantId int64
	}{
		{
			name:   "test1",
			args:   args{key: "user_follow_list_1"},
			wantId: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotId := GetIDFromUserFollowListKey(tt.args.key); gotId != tt.wantId {
				t.Errorf("GetIDFromUserFollowListKey() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGetIDFromVideoMsgKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		wantId int64
	}{
		{
			name:   "test1",
			args:   args{key: "video_message_1"},
			wantId: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotId := GetIDFromVideoMsgKey(tt.args.key); gotId != tt.wantId {
				t.Errorf("GetIDFromVideoMsgKey() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGetIDFromUserFollowerListKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		wantId int64
	}{
		{
			name:   "test1",
			args:   args{key: "user_follower_list_1"},
			wantId: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotId := GetIDFromUserFollowerListKey(tt.args.key); gotId != tt.wantId {
				t.Errorf("GetIDFromUserFollowerListKey() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
