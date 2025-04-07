run-256-75000:
    go run main.go --database vimkimdb --table tbl_256_75000 -u vimkim --dim 256 --rows 75000

run-256-150000:
    go run main.go --database vimkimdb --table tbl_256_150000 -u vimkim --dim 256 --rows 150000

run-256-300000:
    go run main.go --database vimkimdb --table tbl_256_300000 -u vimkim --dim 256 --rows 300000

run-768-75000:
    go run main.go --database vimkimdb --table tbl_768_75000 -u vimkim --dim 768 --rows 75000

run-768-150000:
    go run main.go --database vimkimdb --table tbl_768_150000 -u vimkim --dim 768 --rows 150000

run-768-300000:
    go run main.go --database vimkimdb --table tbl_768_300000 -u vimkim --dim 768 --rows 300000

run-1536-75000:
    go run main.go --database vimkimdb --table tbl_1536_75000 -u vimkim --dim 1536 --rows 75000

run-1536-150000:
    go run main.go --database vimkimdb --table tbl_1536_150000 -u vimkim --dim 1536 --rows 150000

run-1536-300000:
    go run main.go --database vimkimdb --table tbl_1536_300000 -u vimkim --dim 1536 --rows 300000
