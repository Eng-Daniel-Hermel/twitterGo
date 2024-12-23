git add .
git commit -m "Ultimo Commit"
git push --set-upstream origin twitter-main
go build main.go
del main.zip
tar -cf main.tar main
Compress-Archive -Path main.tar -DestinationPath main.zip