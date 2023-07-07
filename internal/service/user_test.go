package service

//func TestUserService_Login(t *testing.T) {
//	type fields struct {
//		userRepo repo.UserRepoInterface
//	}
//	crtl := gomock.NewController(t)
//	userRepoMock := mock_repo.NewMockUserRepoInterface(crtl)
//
//	emailMock := "thinh@gmail.com"
//	hashPwMock := "$2a$10$MxGQRSf19ricWBzJpTu93.wLSpgk0x0Uoau8eObPbNSZMShEdvF66"
//
//	type args struct {
//		email      string
//		password   string
//		expectFunc func()
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name: "case 1 login success",
//			fields: fields{
//				userRepo: userRepoMock,
//			},
//			args: args{
//				email: emailMock,
//				expectFunc: func() {
//					userRepoMock.EXPECT().GetUserByEmail(emailMock).Return(&model.User{
//						Email:    emailMock,
//						Password: hashPwMock,
//					}, nil)
//				},
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &UserService{
//				userRepo: tt.fields.userRepo,
//			}
//			tt.args.expectFunc()
//			if err, _ := s.Login(tt.args.email, tt.args.password); (err != "") != tt.wantErr {
//				t.Errorf("CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestUserService_Login1(t *testing.T) {
//	type fields struct {
//		userRepo repo.UserRepoInterface
//	}
//	crtl := gomock.NewController(t)
//	userRepoMock := mock_repo.NewMockUserRepoInterface(crtl)
//
//	emailMock := "thinh@gmail.com"
//	hashPwMock := "$2a$10$MxGQRSf19ricWBzJpTu93.wLSpgk0x0Uoau8eObPbNSZMShEdvF66"
//
//	type args struct {
//		email      string
//		password   string
//		expectFunc func()
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name: "case 1",
//			fields: fields{
//				userRepo: userRepoMock,
//			},
//			args: args{
//				email: emailMock,
//				expectFunc: func() {
//					userRepoMock.EXPECT().GetUserByEmail(emailMock).Return(&model.User{
//						Email:    emailMock,
//						Password: hashPwMock,
//					}, nil)
//				},
//			},
//			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ4LXVzZXItaWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJ4LXVzZXItcm9sZSI6IiIsImV4cCI6MTY4ODY1NTI0N30.nULgF4WBN1erUPBqzsZswwofgdVNWvHtm38DD-ahbZA",
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &UserService{
//				userRepo: tt.fields.userRepo,
//			}
//			tt.args.expectFunc()
//			got, err := s.Login(tt.args.email, tt.args.password)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Login() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
