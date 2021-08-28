package mycrypto

import "testing"

func TestPassword_Check(t *testing.T) {
	type args struct {
		encPwd string
	}
	tests := []struct {
		name string
		pwd  Password
		args args
		want bool
	}{
		{name: "check ok: 123123", pwd: Password("123123"), args: args{encPwd: Password("123123").Encrypt(GetSalt())}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pwd.Check(tt.args.encPwd); got != tt.want {
				t.Errorf("Password.Check() = %v, want %v %v", got, tt.want, tt.args)
			}
		})
	}
}
