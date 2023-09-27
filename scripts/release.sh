echo "Qual é a tag?[ex: v1.0.0] "; read tag
echo "Digite a descricao da tag: "; read description

git tag -a $tag -m "$description" && git push --tags origin master