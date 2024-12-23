git add .
git commit -m "Ultimo Commit"
git push --set-upstream origin twitter-main
go build main.go
del main.zip
tar -a main.zip twitterGo