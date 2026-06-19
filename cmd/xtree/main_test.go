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
		inputShowIndex bool
		want           string
		wantErr        error
	}{
		{
			name:           "JSON",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot("."),
			inputShowIndex: false,
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
			name:           "JSON_allow_duplicate",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputShowIndex: false,
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
			name:           "JSON_allow_duplicate",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputShowIndex: false,
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
			name:           "JSON_show_index",
			inputData:      strings.NewReader(jsonData),
			inputRoot:      gtree.NewRoot("."),
			inputShowIndex: true,
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
			name:           "TOML",
			inputData:      strings.NewReader(tomlData),
			inputRoot:      gtree.NewRoot("."),
			inputShowIndex: false,
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
			name:           "TOML_allow_duplicate",
			inputData:      strings.NewReader(tomlData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputShowIndex: false,
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
			name:           "TOML_show_index",
			inputData:      strings.NewReader(tomlData),
			inputRoot:      gtree.NewRoot("."),
			inputShowIndex: true,
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
			name:           "YAML",
			inputData:      strings.NewReader(yamlData),
			inputRoot:      gtree.NewRoot("."),
			inputShowIndex: false,
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
			name:           "YAML_allow_duplicate",
			inputData:      strings.NewReader(yamlData),
			inputRoot:      gtree.NewRoot(".", gtree.WithDuplicationAllowed()),
			inputShowIndex: false,
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
			name:           "YAML_show_index",
			inputData:      strings.NewReader(yamlData),
			inputRoot:      gtree.NewRoot("."),
			inputShowIndex: true,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// NOTE: 並行処理で呼び出すと稀に枝が意図しない表示になる。
			//       gtree.XxxxFromYyy関数が平行に走るとグローバルのindex counterのresetが走って別で生成するNodeのindexがずれることによるもの
			//       並行でこの関数群を呼び出すことはそうないと思うので一旦置いとく
			// t.Parallel()

			ret := &bytes.Buffer{}
			gotErr := output(ret, tt.inputData, tt.inputRoot, tt.inputShowIndex)
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
