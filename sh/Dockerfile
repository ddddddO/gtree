FROM debian

COPY /entry.sh /
RUN mkdir tttttmp

# ENTRYPOINTのshellにCMDのコマンドを渡し、shell内の`exec $@`でCMDを実行する(ENTRYPOINT は[]で囲わないとダメっぽい)
# https://qiita.com/uehaj/items/e6dd013e28593c26372d
ENTRYPOINT ["/entry.sh"]  
CMD ["ls",  "-al"]
