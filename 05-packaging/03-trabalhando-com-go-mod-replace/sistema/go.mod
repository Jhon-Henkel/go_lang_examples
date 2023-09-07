module github.com/Jhon-Henkel/go_lang_examples/05-packaging/03-trabalhando-com-go-mod-replace/sistema

go 1.21.0

replace github.com/Jhon-Henkel/go_lang_examples/05-packaging/03-trabalhando-com-go-mod-replace/math => ../math

require github.com/Jhon-Henkel/go_lang_examples/05-packaging/03-trabalhando-com-go-mod-replace/math v0.0.0-00010101000000-000000000000
