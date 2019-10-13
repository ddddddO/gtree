FROM centos:centos7 
#RUN yum clean all \
#	&& yum -y groupinstall core \
#    && yum -y groupinstall base

# ssh接続用コンテナ
# https://qiita.com/TakashiOshikawa/items/081d4780abdc21d63b1a
RUN yum -y update && yum clean all
RUN yum -y install openssh-server openssh-clients

WORKDIR /root 
RUN mkdir /root/.ssh 
COPY /.ssh/id_rsa.pub /root/.ssh/authorized_keys 
RUN chmod 0600 /root/.ssh/authorized_keys 
EXPOSE 22

#CMD /usr/sbin/sshd -d
