## Demo
```sh
12:24:56 > gtree -ts << EOS
- root
  - parent_a
  - parent_b
    - child_a
      - 1
      - 2
        - a
          - 1
    - child_b
      - 1
        - a
  - parent_c
    - child_a
  - parent_d
EOS
root
├── parent_a
├── parent_b
│   ├── child_a
│   │   ├── 1
│   │   └── 2
│   │       └── a
│   │           └── 1
│   └── child_b
│       └── 1
│           └── a
├── parent_c
│   └── child_a
└── parent_d
```
