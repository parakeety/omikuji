package kuji

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_omikuji(t *testing.T) {
	type args struct {
		i int
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "12/31 normal day",
			args: args{
				i: 4,
				t: time.Date(2019, time.December, 31, 23, 59, 59, 59, time.UTC),
			},
			want: "大凶",
		},
		{
			name: "1/1 lucky day",
			args: args{
				i: 4,
				t: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			want: "大吉",
		},
		{
			name: "1/3 lucky day",
			args: args{
				i: 4,
				t: time.Date(2019, time.January, 3, 23, 59, 59, 59, time.UTC),
			},
			want: "大吉",
		},
		{
			name: "1/4 normal day",
			args: args{
				i: 4,
				t: time.Date(2019, time.January, 4, 0, 0, 0, 0, time.UTC),
			},
			want: "大凶",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			omikuji(recorder, req, tt.args.i, tt.args.t)
			res := recorder.Result()
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Fatal("unexpected status code")
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			var kuji Omikuji
			err = json.Unmarshal(b, &kuji)
			if err != nil {
				t.Fatal(err)
			}

			if tt.want != kuji.Message {
				t.Errorf("expected: %v \n actual: %v", tt.want, kuji.Message)
			}
		})
	}

}
