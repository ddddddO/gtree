# CloudVPC/サブネットを作成し、インスタンスに紐づける
- [x] https://www.topgate.co.jp/gcp31-gcp-subnet を実施する。
    - VPCネットワークを作成しファイアウォールでsshルールを追加していない状態で、VMインスタンスと作成したネットワークを紐づけ、GCPコンソールのSSHボタンからSSH接続が出来ないことを確認
```
接続しています...
接続できませんでした。再試行しています（2/3）...
VM のシリアル コンソール出力に、接続の問題のトラブルシューティングに役立つ詳細情報が含まれている場合があります。この問題で考えられる他の原因については、ヘルプ ドキュメントをご覧ください。
OS ログインに移行することにより、鍵転送回数を大幅に改善できます。
```
↓
```
接続できませんでした
ポート 22 で VM に接続できません。この問題の考えられる原因の詳細についてご確認ください。
```

- VM削除し、作成したネットワークにファイアウォールでsshルールを追加後、VM作成(作成済みネットワークと紐づける)し、sshできることを確認。
- また、作成済みVPNネットワークに、ファイアウォールでICMPを許可するルールを追加後、別端末から`ping <gceインスタンスの外部IP>`で疎通を確認。
(ICMP許可ルール追加前は`100% packet loss`。また、**ルール追加後にVMの削除/再起動も不要だった。**)

- [ ] terraformで上記の構成を作成する。
    - 上記の構成を手動で作成できるか不安なので答え用として`terraformer`でtfファイルを作成する(https://qiita.com/andromeda/items/fda67a65bbb56f21e6bd)
        - `terraformer import google --resources=networks,firewalls,instances,subnetworks --regions=asia-east1 --projects=work1111` 実行前の準備
            - `export GOOGLE_APPLICATION_CREDENTIALS=/mnt/c/Users/lbfde/Downloads/work-0a0225cca708.json`
            - `cp /usr/bin/terraform-provider-google_v2.17.0_x4 ~/.terraform.d/plugins/linux_amd64/`
        - `work/gcp/terraform/vpc-by-terraformer` 以下に`generated`ディレクトリの作成を確認。
```
20:07:22 > tree generated/
generated/
└── google
    └── work1111
        ├── firewalls
        │   └── asia-east1
        │       ├── compute_firewall.tf
        │       ├── outputs.tf
        │       ├── provider.tf
        │       ├── terraform.tfstate
        │       └── variables.tf
        ├── instances
        │   └── asia-east1
        │       ├── compute_instance.tf
        │       ├── outputs.tf
        │       ├── provider.tf
        │       └── terraform.tfstate
        ├── networks
        │   └── asia-east1
        │       ├── compute_network.tf
        │       ├── outputs.tf
        │       ├── provider.tf
        │       └── terraform.tfstate
        └── subnetworks
            └── asia-east1
                ├── compute_subnetwork.tf
                ├── outputs.tf
                ├── provider.tf
                └── terraform.tfstate
```
- GCPコンソール上から作成したリソースを削除し、生成された各networks,subnetworks,firewalls,instancesディレクト配下で`terraform init` -> `terraform apply` で、リソース作成ができたことを確認。しかし、instancesのみ、`Error: migration error: found metadata key in unexpected format: metadata.%` で作成できず。