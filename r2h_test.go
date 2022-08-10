package r2h

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name            string
		args            args
		wantResult      string
		wantIsCompleted bool
	}{
		{
			name: "single_romaji_letter",
			args: args{
				s: "a",
			},
			wantResult:      "あ",
			wantIsCompleted: true,
		},
		{
			name: "single_non_romaji_letter",
			args: args{
				s: "x",
			},
			wantResult:      "x",
			wantIsCompleted: false,
		},
		{
			name: "single_hiragana_letter",
			args: args{
				s: "あ",
			},
			wantResult:      "あ",
			wantIsCompleted: false,
		},
		{
			name: "all_vowel_letters",
			args: args{
				s: "aiueo",
			},
			wantResult:      "あいうえお",
			wantIsCompleted: true,
		},
		{
			name: "all_consonant_letters",
			args: args{
				s: "ksthmyrw",
			},
			wantResult:      "ksthmyrw",
			wantIsCompleted: false,
		},
		{
			name: "uppercase_and_lowercase",
			args: args{
				s: "Aragaki Yui",
			},
			wantResult:      "あらがき ゆい",
			wantIsCompleted: true,
		},
		{
			name: "romaji_and_alpha",
			args: args{
				s: "Triendl Reina",
			},
			// The correct answer is "とりんどる れいな", but since it is partly converted into romaji, it cannot be written correctly.
			wantResult:      "Tりえんdl れいな",
			wantIsCompleted: false,
		},
		{
			name: "short_sentence",
			args: args{
				s: "awayachiisanagomikarahamassugunakagenobouga,nanamenimizunonakaninarandetachimashita.",
			},
			wantResult:      "あわやちいさなごみからはまっすぐなかげのぼうが、ななめにみずのなかにならんでたちました。",
			wantIsCompleted: true,
		},
		{
			name: "double_vowel",
			args: args{
				s: "yuuki",
			},
			wantResult:      "ゆうき",
			wantIsCompleted: true,
		},
		{
			name: "double_consonant",
			args: args{
				s: "gakki",
			},
			wantResult:      "がっき",
			wantIsCompleted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotIsCompleted := Convert(tt.args.s)
			if gotResult != tt.wantResult {
				t.Errorf("Convert() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotIsCompleted != tt.wantIsCompleted {
				t.Errorf("Convert() gotIsCompleted = %v, want %v", gotIsCompleted, tt.wantIsCompleted)
			}
		})
	}
}


func TestConvertStrict(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name            string
		args            args
		wantResult      string
		wantErr bool
	}{
		{
			name: "single_romaji_letter",
			args: args{
				s: "a",
			},
			wantResult:      "あ",
			wantErr: false,
		},
		{
			name: "single_non_romaji_letter",
			args: args{
				s: "x",
			},
			wantResult:      "",
			wantErr: true,
		},
		{
			name: "single_hiragana_letter",
			args: args{
				s: "あ",
			},
			wantResult:      "",
			wantErr: true,
		},
		{
			name: "all_vowel_letters",
			args: args{
				s: "aiueo",
			},
			wantResult:      "あいうえお",
			wantErr: false,
		},
		{
			name: "all_consonant_letters",
			args: args{
				s: "ksthmyrw",
			},
			wantResult:      "",
			wantErr: true,
		},
		{
			name: "uppercase_and_lowercase",
			args: args{
				s: "Aragaki Yui",
			},
			wantResult:      "あらがき ゆい",
			wantErr: false,
		},
		{
			name: "romaji_and_alpha",
			args: args{
				s: "Triendl Reina",
			},
			wantResult:      "",
			wantErr: true,
		},
		{
			name: "short_sentence",
			args: args{
				s: "awayachiisanagomikarahamassugunakagenobouga,nanamenimizunonakaninarandetachimashita.",
			},
			wantResult:      "あわやちいさなごみからはまっすぐなかげのぼうが、ななめにみずのなかにならんでたちました。",
			wantErr: false,
		},
		{
			name: "double_vowel",
			args: args{
				s: "yuuki",
			},
			wantResult:      "ゆうき",
			wantErr: false,
		},
		{
			name: "double_consonant",
			args: args{
				s: "gakki",
			},
			wantResult:      "がっき",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ConvertStrict(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertStrict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ConvertStrict() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
