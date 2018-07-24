wget -O dictionary/dict.txt.gz https://cc-cedict.org/editor/editor_export_cedict.php?c=gz
gunzip dictionary/dict.txt.gz
go run dictionary/main.go
rm dictionary/dict.txt