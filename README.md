# work
- お試し

# ローカルk8sクラスタ環境(work共通)
## WSL側
- k8sリソース作成

## Windows側
- Docker Desktopを起動
- kubectl でk8sリソースをデプロイ
- Dockerイメージのビルドもこちらで

## Tips
- 同一Pod内の複数コンテナから特定のコンテナのログを見たい
    - `kubectl logs <Pod名> <コンテナ名>`
---
# GKE構築フロー(terraform)
## 一旦メモ殴り書き
### クラスター作成
- provider(gettingstart)/resource(公式コピペ)を定義したmain.tfを作成
- terraform init
- terraform applyで以下エラー
```
Error: google: could not find default credentials. See https://developers.google.com/accounts/docs/application-default-credentials for more information.

  on main.tf line 4, in provider "google":
   4: provider "google"
```

- エラーのリンク先とgettingstartのとこ(https://www.terraform.io/docs/providers/google/getting_started.html#adding-credentials の先のリンク)より
    - サービスアカウント作成し(「役割」は..一旦「オーナー」)、JsonKeyファイルをDLし、以下のパスに指定
    - export GOOGLE_CLOUD_KEYFILE_JSON={{path}}
        - export GOOGLE_CLOUD_KEYFILE_JSON=/mnt/c/Users/lbfde/Downloads/work-0a0225cca708.json
- terraform apply で、「yes」入力→以下エラー
```
Error: googleapi: Error 403: Kubernetes Engine API has not been used in project 219308425897 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/container.googleapis.com/overview?project=219308425897 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry., accessNotConfigured
```
- 上記エラーリンク先で「Kubernetes Engine API」を有効化
- terraform apply　成功
    - cluster 作成に6m30s
    - node-pool 作成に1m30s
    - 一服できる

### クラスタの構築ができたので、ローカルからkubectlを使ってk8sをデプロイしていきたい
- kubectl config get-contexts 実行(WSL)。なにもない。
- https://cloud.google.com/kubernetes-engine/docs/quickstart?hl=ja を確認で以下実行
    - `gcloud config set project [PROJECT_ID]`
        - `gcloud config set project work1111`
    - `gcloud config set compute/zone [COMPUTE_ZONE]`
        - `gcloud config set compute/zone us-central1` (ここの指定、regionでないと、後続のコマンドでエラーになったけど。。)
- kubectl config get-contexts →　この時点でも、まだなにもない
- `gcloud container clusters get-credentials [CLUSTER_NAME]`
    - `gcloud container clusters get-credentials  gle-pubsub-cluster`
- kubectl config get-contexts →　gkeのコンテキストが設定されたことを確認
```
CURRENT   NAME                                          CLUSTER                                       AUTHINFO                                      NAMESPACE
*         gke_work1111_us-central1_gle-pubsub-cluster   gke_work1111_us-central1_gle-pubsub-cluster   gke_work1111_us-central1_gle-pubsub-cluster 
```

#### なので、GKE上のクラスタに対してk8sをデプロイできるようになった
- が、現状イメージをローカルから取得するようにしているから、、
    - https://cloud.google.com/container-registry/docs/pushing-and-pulling?hl=ja (push/pullについて)
        - `gcloud auth configure-docker`
        - `docker tag [SOURCE_IMAGE] [HOSTNAME]/[PROJECT-ID]/[IMAGE]` or
        - `docker tag [SOURCE_IMAGE] [HOSTNAME]/[PROJECT-ID]/[IMAGE]:[TAG]`
            - [SOURCE_IMAGE]:ローカル イメージ名またはイメージ ID
            - [HOSTNAME]:gcr.io/us.gcr.io/eu.gcr.io/asia.gcr.io 地域別のストレージバケット
        - `docker push [HOSTNAME]/[PROJECT-ID]/[IMAGE]` or
        - `docker push [HOSTNAME]/[PROJECT-ID]/[IMAGE]:[TAG]`

- イメージをgcrにpushして、マニフェストも書き換えてデプロイしたけど、エラー(ImagePullBackOff)
    - `kubectl describe po pubsubsys` で確認したところ、権限ないよと言われているよう。
        - https://cloud.google.com/container-registry/docs/advanced-authentication　見なさい、と

- TODO: **ここから続き**
    - [ ] ノードからpullを成功させる
    - [ ] 一通りのマニフェストの動作チェック
    - [ ] DB永続化
    - [ ] skaffold
    - [ ] kustamize

- 後処理
    - terraform destroy
        - node-poolの削除に3m30s
        - cluster の削除に4m