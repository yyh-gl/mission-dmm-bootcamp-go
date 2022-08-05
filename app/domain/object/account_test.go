package object

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_Account_NewAccount(t *testing.T) {
	type (
		args struct {
			username string
			password string
		}

		want struct {
			account Account
			err     error
		}
	)

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "正しいアカウント情報を生成できる",
			args: args{
				username: "username",
				password: "password",
			},
			want: want{
				account: Account{
					Username: "username",
				},
			},
		},
		{
			name: "usernameが10文字より多いときはエラー",
			args: args{
				username: "usernameusername",
				password: "password",
			},
			want: want{
				err: errors.New("username is too long"),
			},
		},
		{
			name: "passwordが5文字より少ないときはエラー",
			args: args{
				username: "username",
				password: "p",
			},
			want: want{
				err: errors.New("password is too short"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccount(tt.args.username, tt.args.password)

			if tt.want.err == nil {
				// got（受け取った値）が想定どおりかみてる
				if got.Username != tt.want.account.Username {
					t.Error("invalid username")
				}
				if err != nil {
					t.Error("error is happened")
				}
			} else {
				if err.Error() != tt.want.err.Error() {
					t.Error("エラーが違う")
				}
			}
		})
	}
}
