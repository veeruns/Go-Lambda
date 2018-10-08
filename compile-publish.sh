GOOS=linux GOARCH=amd64 go build -o alexa-skill alexa-skill.go dyndb.go && zip alexa-skill.zip alexa-skill && aws s3 cp alexa-skill.zip s3://veeruns && aws lambda  update-function-code --function-name golang_lamda --s3-bucket veeruns --s3-key alexa-skill.zip --publish
rm -rf alexa-skill alexa-skill.zip
