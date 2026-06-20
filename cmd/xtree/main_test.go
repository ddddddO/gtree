package main

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/ddddddO/gtree"
)

var jsonData = strings.TrimSpace(`
{
    "name": "Alice",
    "age": 30,
    "height": 175.5,
    "is_active": true,
    "metadata": null,
    "roles": ["admin", "editor"],
    "settings": {
        "theme": "dark",
        "notifications": true
    },
    "devices": [
        { "type": "mobile", "os": "ios" },
        { "type": "desktop", "os": "windows" }
    ]
}
`)

var tomlData = strings.TrimSpace(`
name = "Alice"
age = 30
height = 175.5
is_active = true
roles = ["admin", "editor"]

[settings]
theme = "dark"
notifications = true

[[devices]]
type = "mobile"
os = "ios"

[[devices]]
type = "desktop"
os = "windows"
`)

var yamlData = strings.TrimSpace(`
name: "Alice"
age: 30
height: 175.5
is_active: true
metadata: null
roles:
  - "admin"
  - "editor"
settings:
  theme: "dark"
  notifications: true
devices:
  - type: "mobile"
    os: "ios"
  - type: "desktop"
    os: "windows"
`)

func TestOutput(t *testing.T) {
	tests := []struct {
		name           string
		inputData      io.Reader
		inputRoot      *gtree.Node
		inputOmitIndex bool
		want           string
		wantErr        error
	}{
		{
			name:           "JSON",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── [0]
│   │   ├── os
│   │   │   └── ios
│   │   └── type
│   │       └── mobile
│   └── [1]
│       ├── os
│       │   └── windows
│       └── type
│           └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── [0]
│   │   └── admin
│   └── [1]
│       └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "JSON_allow_duplicate",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── [0]
│   │   ├── os
│   │   │   └── ios
│   │   └── type
│   │       └── mobile
│   └── [1]
│       ├── os
│       │   └── windows
│       └── type
│           └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── [0]
│   │   └── admin
│   └── [1]
│       └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "JSON_omit_index",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: true,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── os
│   │   ├── ios
│   │   └── windows
│   └── type
│       ├── mobile
│       └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── admin
│   └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "JSON_allow_duplicate_and_omit_index",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputOmitIndex: true,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── os
│   │   └── ios
│   ├── type
│   │   └── mobile
│   ├── os
│   │   └── windows
│   └── type
│       └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── admin
│   └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "TOML",
			inputData:      strings.NewReader(tomlData),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── [0]
│   │   ├── os
│   │   │   └── ios
│   │   └── type
│   │       └── mobile
│   └── [1]
│       ├── os
│       │   └── windows
│       └── type
│           └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── name
│   └── Alice
├── roles
│   ├── [0]
│   │   └── admin
│   └── [1]
│       └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "TOML_allow_duplicate",
			inputData:      strings.NewReader(tomlData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── [0]
│   │   ├── os
│   │   │   └── ios
│   │   └── type
│   │       └── mobile
│   └── [1]
│       ├── os
│       │   └── windows
│       └── type
│           └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── name
│   └── Alice
├── roles
│   ├── [0]
│   │   └── admin
│   └── [1]
│       └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "TOML_omit_index",
			inputData:      strings.NewReader(tomlData),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: true,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── os
│   │   ├── ios
│   │   └── windows
│   └── type
│       ├── mobile
│       └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── name
│   └── Alice
├── roles
│   ├── admin
│   └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "TOML_allow_duplicate_and_omit_index",
			inputData:      strings.NewReader(tomlData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputOmitIndex: true,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── os
│   │   └── ios
│   ├── type
│   │   └── mobile
│   ├── os
│   │   └── windows
│   └── type
│       └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── name
│   └── Alice
├── roles
│   ├── admin
│   └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "YAML",
			inputData:      strings.NewReader(yamlData),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── [0]
│   │   ├── os
│   │   │   └── ios
│   │   └── type
│   │       └── mobile
│   └── [1]
│       ├── os
│       │   └── windows
│       └── type
│           └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── [0]
│   │   └── admin
│   └── [1]
│       └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "YAML_allow_duplicate",
			inputData:      strings.NewReader(yamlData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── [0]
│   │   ├── os
│   │   │   └── ios
│   │   └── type
│   │       └── mobile
│   └── [1]
│       ├── os
│       │   └── windows
│       └── type
│           └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── [0]
│   │   └── admin
│   └── [1]
│       └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "YAML_omit_index",
			inputData:      strings.NewReader(yamlData),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: true,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── os
│   │   ├── ios
│   │   └── windows
│   └── type
│       ├── mobile
│       └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── admin
│   └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
		{
			name:           "YAML_allow_duplicate_and_omit_index",
			inputData:      strings.NewReader(yamlData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputOmitIndex: true,
			want: strings.TrimPrefix(`
.
├── age
│   └── 30
├── devices
│   ├── os
│   │   └── ios
│   ├── type
│   │   └── mobile
│   ├── os
│   │   └── windows
│   └── type
│       └── desktop
├── height
│   └── 175.5
├── is_active
│   └── true
├── metadata
│   └── <nil>
├── name
│   └── Alice
├── roles
│   ├── admin
│   └── editor
└── settings
    ├── notifications
    │   └── true
    └── theme
        └── dark
`, "\n"),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ret := &bytes.Buffer{}
			gotErr := output(ret, tt.inputData, tt.inputRoot, tt.inputOmitIndex)
			gotOutput := ret.String()

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
			if gotOutput != tt.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.want)
			}

		})
	}
}

func TestOutput_multiRow(t *testing.T) {
	tests := []struct {
		name           string
		inputData      io.Reader
		inputRoot      *gtree.Node
		inputOmitIndex bool
		want           string
		wantErr        error
	}{
		{
			name: "JSON",
			inputData: strings.NewReader(strings.TrimSpace(`
{
  "description": "This is a sample\ncontaining multiple lines\nwithin a single JSON value."
}
`)),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
└── description
    └── This is a sample\ncontaining multiple lines\nwithin a single JSON value.
`, "\n"),
			wantErr: nil,
		},
		{
			name: "TOML",
			inputData: strings.NewReader(strings.TrimSpace(`
[data]
multiline = """
This is a multi-line string
using triple quotes in TOML.
It allows preserving newlines."""

single = "PPPP"
`)),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
└── data
    ├── multiline
    │   └── This is a multi-line string\nusing triple quotes in TOML.\nIt allows preserving newlines.
    └── single
        └── PPPP
`, "\n"),
			wantErr: nil,
		},
		{
			name: "YAML",
			inputData: strings.NewReader(strings.TrimSpace(`
config:
  multiline_literal: |
    line 1
    line 2
    line 3
  multiline_folded: >
    This text is folded
    into a single line
    when parsed.
`)),
			inputRoot:      gtree.NewRoot("."),
			inputOmitIndex: false,
			want: strings.TrimPrefix(`
.
└── config
    ├── multiline_folded
    │   └── This text is folded into a single line when parsed.
    └── multiline_literal
        └── line 1\nline 2\nline 3\n
`, "\n"),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ret := &bytes.Buffer{}
			gotErr := output(ret, tt.inputData, tt.inputRoot, tt.inputOmitIndex)
			gotOutput := ret.String()

			if gotErr != nil {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.wantErr)
				}
			}
			if gotOutput != tt.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.want)
			}
		})
	}
}
