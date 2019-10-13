#! /bin/bash

set -ex

docker build -t ansios:v1 .

# ansible-playbook gen-container-playbook.yml
# docker exec -it <ansios:v1 のコンテナID> bash
## vi /etc/ssh/sshd_config を、http://ossfan.net/setup/openssh-02.html　の「(3) 公開鍵認証を許可する設定」を参考にパスワード認証をストップ
## ps auxで、kill </usr/sbin/sshd -D のPID>
## 再び、/usr/sbin/sshd -D &　実行
## exit
# ssh root@0.0.0.0 -p 2222 -i .ssh/id_rsa   OR
# sudo ssh root@172.17.0.2 -p 22 -i .ssh/id_rsa 
# でsshは出来たけれど、
# NOTE: 生成したコンテナに対して、playbook(target-container-playbook.yml)の実行も出来なければ、ansible 172.17.0.2 -m ping も出来ていない
