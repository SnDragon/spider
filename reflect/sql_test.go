package reflect

import "testing"

type Student struct {
	Name string
	Age  int
}

type User struct {
	UserId int
}

func TestCreateSQL(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				v: &Student{Name: "longerwu", Age: 27},
			},
			want: "insert into Student values(\"longerwu\",27)",
		},
		{
			name: "t2",
			args: args{
				v: Student{Name: "longerwu", Age: 27},
			},
			want: "insert into Student values(\"longerwu\",27)",
		},
		{
			name: "t3",
			args: args{
				v: User{123},
			},
			want: "insert into User values(123)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSQL(tt.args.v); got != tt.want {
				t.Errorf("CreateSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
