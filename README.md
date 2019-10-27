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
    - `gcloud container clusters get-credentials  gke-pubsub-cluster`
        - (また、例えば、PC再起動した後で、コンテキストが設定されていても、上記コマンド実行しないとダメ(kubectl効かない)っぽい?)
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
            - 上記リンクの最後のリンク先に、ぽいことが記載されている(要はノードからプライベートなgcrへのアクセス権が無いよ、と)
                - https://cloud.google.com/container-registry/docs/using-with-google-cloud-platform　の「Google Kubernetes Engine クラスタにスコープを設定する」
                    - terraform設定ファイルに、以下設定を記載し、terraform apply
                        - https://www.terraform.io/docs/providers/google/r/container_cluster.html#oauth_scopes
                            - kubectl apply で、プライベートgcrからpull成功

#### Pod/ReplicaSet/Deploymentの動作チェック
- [x] 02_k8s_gke/pod-pubsub.yml
- [x] 02_k8s_gke/rs-pubsub.yml
- [x] 02_k8s_gke/deployment
    - `kubectl logs <sub Pod> - f` で確認すると、pub２台分のNotifyを確認

#### TODO: DBの永続化をする
https://tkzo.jp/blog/use-persistent-volumes-in-gke/ の内容を踏襲する。

- 永続ディスクをterraformで作成
  - 一旦、GCPコンソールから作成する。
- 作成した永続ディスクとk8sのpvオブジェクトと紐づけ兼pvとpvcの作成
    - まず、PersistentVolumeClaim を作成する。
- DBデプロイメントとpvcの紐づけ
- デプロイ
をやる。

##### 確認手順
- [x] 永続化前
    1. 永続化前のdeploy-svc-db.ymlをデプロイして、`INSERT INTO tags (name, users_id) VALUES ('kube-pd-test', 2);` を実行
    2. `kubectl delete pod <DBのPod>` で殺したあと、復活するPod内で`SELECT * FROM tags WHERE users_id = 2;` で、1.でInsertしたタグが**消えていること**

- [x] 永続化後
    1. 永続化後のdeploy-svc-db.ymlをデプロイして、`INSERT INTO tags (name, users_id) VALUES ('kube-pd-test', 2);` を実行
    2. `kubectl delete pod <DBのPod>` で殺したあと、復活するPod内で`SELECT * FROM tags WHERE users_id = 2;` で、1.でInsertしたタグが**残っていること**

- 以下、DB永続化作業ログ
    - https://tkzo.jp/blog/use-persistent-volumes-in-gke/ と同様の手順でk8sオブジェクトを作成(マウントパスは、postgresqlデータがデフォルトで格納される`/var/lib/postgresql/data`)し、デプロイしたがDBコンテナで以下エラー
```
initdb: error: directory "/var/lib/postgresql/data" exists but is not empty
It contains a lost+found directory, perhaps due to it being a mount point.
Using a mount point directly as the data directory is not recommended.
Create a subdirectory under the mount point.
```

- 以下を参考に、マニフェストファイルを書き換え、デプロイ
    - https://stackoverflow.com/questions/51168558/how-to-mount-a-postgresql-volume-using-aws-ebs-in-kubernete/51174380
    - https://hub.docker.com/_/postgres　「PGDATA」
        - dbマニフェストで、PGDATA: `/var/lib/postgresql/data/pgdata` を設定
        - (DBのPodはRunningし始めた。が、変更したpostgresのデータ格納先とマウント先が異なるので、Pod削除したら、データ保存されないのでは？)
- とりあえず、起動したDBコンテナに入り、永続化検証
  - 上記の「永続化後」の内容を実行する
  - 永続化が成功していた！
    - `マウントパスとpostgresqlのデータ格納先の階層が異なること(データ格納先の方が一つ下の階層)が要因か？`
- ReplicaSetを削除した場合は？
    - データは残っていた
- Deployment削除後、再度applyで残っていれば、永続化自体はOK、と思われる
    - データの確認はできた。永続化成功。
- `PersistentVolume`オブジェクトで、pvcが削除された場合に永続ボリュームが削除されない設定(`persistentVolumeReclaimPolicy: Retain`)が効いているか確認する
    - DBのdeployment及び、pv/pvcを削除し、それらを再デプロイする。
        - データ(`kube-pd-test`) を確認、これにてDone
---
- 関連メモ
    - `persistentVolumeReclaimPolicy: Delete`(PVCが削除された際に永続ディスクを削除)を設定した場合に、DBのdeployment及び、pv/pvcを削除し、それらを再デプロイした時どうなるか？
        - DBのdeployment及び、pv/pvcを削除し -> `GCPコンソールの「ディスク」から消えた`
        - 一旦ここまでで終了する

###### ちょっとメモ
1. kubectl apply -f <DBのdeployment>
2. kubectl delete <DBの`Pod`>
3. kubectl get po で以下
```
NAME                      READY   STATUS        RESTARTS   AGE
ps-db-7bbfd8f56-4b9ln     1/1     Running       0          22s
ps-db-7bbfd8f56-v76qh     1/1     Terminating   0          38m ←しばらくすると消える
```

---
- TODO: **ここから続き**
    - [ ] 各種k8sオブジェクトの作成
        - [ ] StatefullSet
        - [ ] Job
        - [ ] CronJob
    - [ ] Namespace
    - [ ] Label/Selector
    - [ ] Annotation
    - [ ] skaffold
    - [ ] kustamize
    - [ ] GKEのネットワーク設計的な


---

### 後処理
- `terraform destroy`
    - node-poolの削除に3m30s
    - cluster の削除に4m