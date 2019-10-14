#! /bin/bash

# exec) sudo ./init.sh

set -ex

# dockerイメージ作成
docker build -t ansios:v1 .

# コンテナ起動
ansible-playbook gen-container-playbook.yml

# コンテナに新規ユーザーとグループを作成
ansible-playbook -i hosts target-container-playbook.yml --private-key .ssh/id_rsa


# このコンテナに対して、Ansibleモジュールを色々試す用にしてもよさそう


# 以下、作業メモ

## (このブロックの処理は、Dockerfileで完結させた。)
# docker exec -it <ansios:v1 のコンテナID> bash
## vi /etc/ssh/sshd_config を、http://ossfan.net/setup/openssh-02.html　の「(3) 公開鍵認証を許可する設定」を参考にパスワード認証をストップ
## ps auxで、kill </usr/sbin/sshd -D のPID>
## 再び、/usr/sbin/sshd -D &　実行
## exit

# ssh root@0.0.0.0 -p 2222 -i .ssh/id_rsa   OR
# sudo ssh root@172.17.0.2 -p 22 -i .ssh/id_rsa  (memo: http://watermans-linuxtips.blogspot.com/2009/05/ssh.html)
# でsshは出来た。また、

# https://docs.ansible.com/ansible/latest/user_guide/intro_getting_started.html
# の「Your first commands」を参考に、以下を実施した。とても重要
# 1. /etc/ansible/hosts 内に、docker inspect | grep IP で取得したコンテナのIPを追記
# 2. sudo su でrootにアタッチ
# 3. ansible 172.17.0.2 -m ping --private-key .ssh/id_rsa     で、pingpong!
# 4. ansible-playbook -i hosts target-container-playbook.yml --private-key .ssh/id_rsa
# 5. docker exec -it <ansios:v1 のコンテナID> cat /etc/passwd で4.のplaybookで作成したユーザーを確認
