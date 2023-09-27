set /p tag=Digite a tag ex: v1.0.0: 
set /p description=Digite a descricao da tag:  

git tag -a %tag% -m "%description%" && git push --tags origin master