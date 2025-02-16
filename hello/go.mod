module example.com/hello

go 1.21.6

replace example.com/depedency => ../depedency

require example.com/depedency v0.0.0-00010101000000-000000000000

require github.com/google/go-cmp v0.6.0 // indirect
