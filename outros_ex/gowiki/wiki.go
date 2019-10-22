package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//struct Page sera o conteudo da pagina
//titulo (Title) e corpo (Body)
//Title do tipo String, e Body é do tipo []byte por causa da biblioteca "io/ioutil"
//A struct descreve como a pagina sera guardada na memoria
//[] pode significar um array de tamanho indefinido
type Page struct {
	Title string
	Body  []byte
}

//metodo salvar (save)
//recebe p que aponta para a struct Page, nao aceita parametros
//e retorna um valor do tipo error
//ess metodo salva o corpo da mensagem (Body) em um arquivo
//o nome do arquivo eh o titulo (Title) concatenado .txt
//o retorno eh do tipo error porque eh o tipo que WriteFile retorna
//se tudo estiver correto o retorno sera nil(nulo)
//se algo der errado enquanto grava retorna error
//o ocal 0600 indica que o arquivo sera criado com permissao de leitura e gravacao apenas para o usuario atual
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//metodo carregar pagina
//title eh parametro de entrada da funcao loadPage
//ReadFile retorna um []byte e error
//[]byte esta sendo atribuido a body e o error esta sendo tratado pela variavel err
//se err for diferente de nil (nulo) retorna nulo para body e error
//senao retorna title para Title e body para Body (?)
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//o uso do _(sublinhado) que ignora o erro nao eh uma boa pratica, fique atento!!!
//title recebe e eh atribuido o endereco da aplicacao mais a subdivisao '/view/' -> localhost:8080/view/'title'
//nesse caso a func loadPage recebe title como parametro
//retorna p que eh as struct Page e, neste caso, estamos ignorando se o retorno for nil(nulo)
//o Fprintf escreve o Title e o Body do txt com uma formatacao HTML simples e grava o conteudo em 'w'
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

//a funcao template.ParseFiles le o conteudo do arquivo edit.html
//o metodo t.Execute executa o template, gravando o HTML em 'w' fazendo as devidas adaptacoes de acordo com a struct Page 'p'
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func saveHandler() {

}

//o viewHandler inicializa o http  e manipula as solicitacoes no caminho '/view/'
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//	http.HandleFunc("/save", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
