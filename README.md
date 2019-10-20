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