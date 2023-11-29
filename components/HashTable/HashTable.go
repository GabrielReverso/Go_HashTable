package HashTable

import (
	"errors"
	"fmt"
)

// Estrutura Hash que contém um slice de VetorHash, um contador de quantidade, e um vetor para salvar indices adicionados
type Hash struct {
	Indices     []VetorHash
	Referencias []int
	Quantidade  int
}

// Estrutura VetorHash que contém um ponteiro para Dados e um verificador de colisão
type VetorHash struct {
	Dados_Usuario       *Dados
	Verificador_colisao bool
}

// Estrutura Dados que contém informações do usuário e um ponteiro para o próximo Dados
type Dados struct {
	Nome     string
	Endereco string
	Telefone string
	Next     *Dados
}

// Função para criar um novo Hash
func CriaHash() *Hash {

	// Cria um novo Hash com um slice de VetorHash de tamanho inicial 20
	hash_Table := &Hash{Indices: make([]VetorHash, 20), Referencias: make([]int, 0), Quantidade: 0}

	return hash_Table
}

// Função para inserir dados no Hash
func InserirDados(hash_table *Hash, Nome_input string, Endereco_input string, Telefone_input string) {

	// Calcula o índice onde os dados devem ser inseridos
	Indice := Peso_strings(Nome_input, hash_table)

	// Cria um novo Dados com as informações do usuário
	Informacoes := &Dados{Nome: Nome_input, Endereco: Endereco_input, Telefone: Telefone_input}

	// Se o índice for maior do que o tamanho do slice, redimensiona o slice
	if len(hash_table.Indices) <= Indice {
		temporary := make([]VetorHash, len(hash_table.Indices)*2)
		copy(temporary, hash_table.Indices)
		hash_table.Indices = temporary
		Rehash(hash_table, Nome_input)
		Indice = Peso_strings(Nome_input, hash_table)
	}

	// Cria um alias para o VetorHash no índice
	Hash := &hash_table.Indices[Indice]

	// Se o VetorHash no índice não contém dados, insere os dados
	// Se já contém dados, insere os novos dados no final da lista ligada
	if Hash.Dados_Usuario == nil {
		Hash.Dados_Usuario = Informacoes
		hash_table.Referencias = append(hash_table.Referencias, Indice)
	} else {
		current := Hash.Dados_Usuario
		for current.Next != nil {
			if Informacoes.Nome != current.Nome {
				Hash.Verificador_colisao = true
			}
			current = current.Next
		}
		current.Next = Informacoes
	}

	// Incrementa a quantidade de dados no Hash
	hash_table.Quantidade++
}

// Função para buscar dados no Hash
func BuscaHash(hash_table *Hash, Nome_search string) ([]string, error) {

	// Calcula o índice onde os dados devem estar
	Indice := Peso_strings(Nome_search, hash_table)

	if len(hash_table.Indices) <= Indice {
		return nil, errors.New("nenhum dado corresponde ao informado")
	}

	// Começa a busca no primeiro Dados no índice
	current := hash_table.Indices[Indice].Dados_Usuario

	// Se não há dados no índice, imprime uma mensagem
	if current == nil {
		return nil, errors.New("nenhum dado corresponde ao informado")
	}

	// Cria uma slice para armazenar os dados
	var data []string

	// Percorre a lista ligada no índice, adicionando os dados que correspondem ao nome buscado à slice
	for current != nil {
		if current.Nome == Nome_search {
			data = append(data, fmt.Sprintf("%s_%s_%s", current.Nome, current.Telefone, current.Endereco))
		}
		current = current.Next
	}

	// Se nenhum dado for igual ao procurado, returna um erro
	if len(data) == 0 {
		return nil, errors.New("nenhum dado corresponde ao informado")
	}

	// Retorna a slice de dados
	return data, nil
}

// Função para mostrar todos os dados da Hash
func BuscaTodosHash(hash_table *Hash) ([]string, error) {

	// Cria uma slice para armazenar os dados
	var data []string

	if len(hash_table.Referencias) == 0 {
		return nil, errors.New("nenhum dado na Hash")
	}

	for _, Indice := range hash_table.Referencias {
		// Começa a busca no primeiro Dados no índice
		current := hash_table.Indices[Indice].Dados_Usuario

		// Percorre a lista ligada no índice, adicionando os dados que correspondem ao nome buscado à slice
		for current != nil {
			data = append(data, fmt.Sprintf("%s_%s_%s", current.Nome, current.Telefone, current.Endereco))
			current = current.Next
		}
	}

	// Retorna a slice de dados
	return data, nil
}

// Função para mostrar um dado especifico da Hash
func BuscaEspecificoHash(hash_table *Hash, nome string, telefone string, endereco string) ([]string, error) {

	// Cria uma slice para armazenar os dados
	var data []string

	Indice := Peso_strings(nome, hash_table)

	// Começa a busca no primeiro Dados no índice
	current := hash_table.Indices[Indice].Dados_Usuario

	// Percorre a lista ligada no índice, adicionando os dados que correspondem ao nome buscado à slice
	for current != nil {
		if current.Nome == nome && current.Telefone == telefone && current.Endereco == endereco {
			data = append(data, fmt.Sprintf("%s_%s_%s", current.Nome, current.Telefone, current.Endereco))
			return data, nil
		}
		current = current.Next
	}

	// Retorna a string de dados
	return nil, errors.New("nenhum dado na Hash")
}

func Rehash(hash_table *Hash, novoNome string) error {

	Referencia := hash_table.Referencias
	hash_table.Referencias = make([]int, 0)
	hash_table.Quantidade = 0

	i := 0
	max := 100
	for FlagNovoPeso(hash_table, Referencia, novoNome) && i < max {
		i++
	}

	if i < max {
		// Cria um slice temporário para armazenar os dados
		tempDados := make([]*Dados, 0)
		for _, indice := range Referencia {
			Hash_Auxiliar := hash_table.Indices[indice]
			current := Hash_Auxiliar.Dados_Usuario
			for current != nil {
				tempDados = append(tempDados, current)
				current = current.Next
			}
			Hash_Auxiliar.Dados_Usuario = nil
			Hash_Auxiliar.Verificador_colisao = false
		}

		// Limpa a tabela hash
		hash_table.Indices = make([]VetorHash, len(hash_table.Indices))

		// Insere os dados novamente na tabela hash
		for _, dados := range tempDados {
			InserirDados(hash_table, dados.Nome, dados.Endereco, dados.Telefone)
		}

	} else {
		return errors.New("erro ao fazer rehash")
	}

	return nil
}

// DeleteHash é uma função que remove um usuário específico da tabela hash.
func DeleteHash(hash_table *Hash, Nome_Delete string, Telefone_Delete string) {
	// Calcula a posição na tabela hash onde o usuário deve estar.
	Position := Peso_strings(Nome_Delete, hash_table)
	// Cria um ponteiro para o vetor hash na posição calculada.
	Hash := &hash_table.Indices[Position]

	// Inicializa um contador.
	count := 0

	// Cria um ponteiro para o primeiro usuário na lista ligada na posição calculada.
	current := Hash.Dados_Usuario
	// Cria um ponteiro para o usuário anterior na lista ligada.
	var prev *Dados

	// Se houve uma colisão na posição calculada, o código dentro deste bloco if será executado.
	if Hash.Verificador_colisao {
		// Este loop percorre a lista ligada na posição calculada.
		for current != nil {
			// Se o nome e o telefone do usuário atual correspondem aos dados a serem excluídos e o contador é zero,
			// o próximo usuário na lista ligada se torna o primeiro usuário.
			if current != nil && count == 0 && current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
				Hash.Dados_Usuario = current.Next
				count = 1
				return
			} else {
				// Se o usuário atual é nil, a função retorna.
				if current == nil {
					return
				}
				// Se o nome e o telefone do usuário atual correspondem aos dados a serem excluídos,
				// o próximo usuário na lista ligada se torna o próximo usuário do usuário anterior.
				if current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
					prev.Next = current.Next
					current = current.Next
				}
			}

			// Se o usuário atual é nil, a função retorna.
			if current == nil {
				return
			}
			// O contador é incrementado e os ponteiros para o usuário atual e anterior são atualizados.
			count = 1
			prev = current
			current = current.Next
		}

		// O verificador de colisão é definido como falso.
		Hash.Verificador_colisao = false
		// O ponteiro para o usuário atual é redefinido para o primeiro usuário na lista ligada.
		current = Hash.Dados_Usuario
		// Este loop verifica se ainda há uma colisão na posição calculada.
		for current != nil {
			// Se o nome do próximo usuário é diferente do nome do usuário atual, o verificador de colisão é definido como verdadeiro.
			if current.Next != nil && current.Nome != current.Next.Nome {
				Hash.Verificador_colisao = true
			}
			// O ponteiro para o usuário atual é atualizado para o próximo usuário na lista ligada.
			current = current.Next
		}
	} else {
		// Se não houve uma colisão na posição calculada, o código dentro deste bloco else será executado.
		// Se o próximo usuário na lista ligada é nil, o primeiro usuário na lista ligada é definido como nil e o verificador de colisão é definido como falso.
		if current.Next == nil {
			Hash.Dados_Usuario = nil
			Hash.Verificador_colisao = false
			// Cria um slice auxiliar para armazenar as referências.
			Referencias_auxiliar := make([]int, 0)

			// Este loop percorre o slice de referências na tabela hash.
			for _, conteudo := range hash_table.Referencias {
				// Se o conteúdo atual não é igual à posição calculada, o conteúdo é adicionado ao slice auxiliar.
				if conteudo != Position {
					Referencias_auxiliar = append(Referencias_auxiliar, conteudo)
				}
			}
			// O slice de referências na tabela hash é atualizado para o slice auxiliar.
			hash_table.Referencias = Referencias_auxiliar
			return
		}
		// O ponteiro para o usuário atual é redefinido para o primeiro usuário na lista ligada.
		current = Hash.Dados_Usuario
		// Este loop percorre a lista ligada na posição calculada.
		for current != nil {
			// Se o nome e o telefone do usuário atual correspondem aos dados a serem excluídos e o contador é zero

			// o próximo usuário na lista ligada se torna o primeiro usuário.
			if current != nil && count == 0 && current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
				Hash.Dados_Usuario = current.Next
				count = 1
				return
			} else {
				// Se o usuário atual é nil, a função retorna.
				if current == nil {
					return
				}
				// Se o nome e o telefone do usuário atual correspondem aos dados a serem excluídos,
				// o próximo usuário na lista ligada se torna o próximo usuário do usuário anterior.
				if current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
					prev.Next = current.Next
					current = current.Next
				}
			}

			// Se o usuário atual é nil, a função retorna.
			if current == nil {
				return
			}
			// O contador é incrementado e os ponteiros para o usuário atual e anterior são atualizados.
			count = 1
			prev = current
			current = current.Next
		}
	}
}

// DeleteAllHash é uma função que remove todos os usuários com um nome específico da tabela hash.
func DeleteAllHash(hash_table *Hash, Nome_Delete string) {
	// Calcula a posição na tabela hash onde os usuários devem estar.
	Position := Peso_strings(Nome_Delete, hash_table)
	// Cria um ponteiro para o vetor hash na posição calculada.
	Hash := &hash_table.Indices[Position]

	// Se houve uma colisão na posição calculada, o código dentro deste bloco if será executado.
	if Hash.Verificador_colisao {
		// Cria um ponteiro para o primeiro usuário na lista ligada na posição calculada.
		current := Hash.Dados_Usuario
		// Cria um ponteiro para o usuário anterior na lista ligada.
		var prev *Dados

		// Este loop percorre a lista ligada na posição calculada.
		for current != nil {
			// Se o nome do usuário atual corresponde ao nome a ser excluído, o código dentro deste bloco if será executado.
			if current.Nome == Nome_Delete {
				// Se o usuário anterior não é nil, o próximo usuário na lista ligada se torna o próximo usuário do usuário anterior.
				// Caso contrário, o próximo usuário na lista ligada se torna o primeiro usuário.
				if prev != nil {
					prev.Next = current.Next
				} else {
					Hash.Dados_Usuario = current.Next
				}
			} else {
				// Se o nome do usuário atual não corresponde ao nome a ser excluído, o usuário atual se torna o usuário anterior.
				prev = current
			}
			// Se o usuário atual não é nil, o ponteiro para o usuário atual é atualizado para o próximo usuário na lista ligada.
			if current != nil {
				current = current.Next
			}
		}

		// O ponteiro para o usuário atual é redefinido para o primeiro usuário na lista ligada.
		current = Hash.Dados_Usuario
		// O verificador de colisão é definido como falso.
		hash_table.Indices[Position].Verificador_colisao = false
		// Este loop verifica se ainda há uma colisão na posição calculada.
		for current != nil {
			// Se o nome do próximo usuário é diferente do nome do usuário atual, o verificador de colisão é definido como verdadeiro.
			if current.Next != nil && current.Nome != current.Next.Nome {
				hash_table.Indices[Position].Verificador_colisao = true
			}
			// O ponteiro para o usuário atual é atualizado para o próximo usuário na lista ligada.
			current = current.Next
		}

	} else {
		// Se não houve uma colisão na posição calculada, o código dentro deste bloco else será executado.
		// O primeiro usuário na lista ligada é definido como nil e o verificador de colisão é definido como falso.
		Hash.Dados_Usuario = nil
		Hash.Verificador_colisao = false
		// Cria um slice auxiliar para armazenar as referências.
		Referencias_auxiliar := make([]int, len(hash_table.Referencias))
		// Este loop percorre o slice de referências na tabela hash.
		for _, Conteudo := range hash_table.Referencias {
			// Se o conteúdo atual não é igual à posição calculada, o conteúdo é adicionado ao slice auxiliar

			if Conteudo != Position {
				Referencias_auxiliar = append(Referencias_auxiliar, Conteudo)
			}
		}
		// O slice de referências na tabela hash é atualizado para o slice auxiliar.
		hash_table.Referencias = Referencias_auxiliar
	}
}

/***************************************AUXILIARES********************************************/

// Função para calcular o peso de uma string
func Peso_strings(nome string, hash_table *Hash) int {

	var Peso int
	Grau := len(nome)

	// Calcula o peso da string
	Somatoria := 0
	for _, Letra := range nome {
		Somatoria += int(Letra) * Grau
		Grau--
	}

	Peso = Somatoria

	// Calcula o índice baseado no peso e no tamanho do slice
	Resto := Peso % (len(hash_table.Indices) + 1)
	return Resto
}

// FlagNovoPeso é uma função que recebe uma tabela hash e retorna um booleano.
func FlagNovoPeso(hash_table *Hash, Referencia []int, novoNome string) bool {

	if Peso_strings(novoNome, hash_table) >= len(hash_table.Indices) {
		temporary := make([]VetorHash, Peso_strings(novoNome, hash_table)+1)
		copy(temporary, hash_table.Indices)
		hash_table.Indices = temporary
		// A função retorna true indicando necessidade de aumentar vetor novamente.
		return true
	}

	// O loop for percorre cada índice na lista de referências da tabela hash.
	for _, indice := range Referencia {
		// Hash é um alias que armazena o valor no índice atual da tabela hash.
		Hash := hash_table.Indices[indice]
		// Se o verificador de colisão do Hash for verdadeiro, o código dentro deste bloco if será executado.
		if Hash.Verificador_colisao {
			current := Hash.Dados_Usuario
			// Este loop for continuará enquanto current não for nil.
			for current != nil {
				NovoIndice := Peso_strings(current.Nome, hash_table)
				// Se o tamanho da lista de índices da tabela hash for menor ou igual a NovoIndice, ajusta o tamanho do vetor.
				if len(hash_table.Indices) <= NovoIndice {
					temporary := make([]VetorHash, NovoIndice+1)
					copy(temporary, hash_table.Indices)
					hash_table.Indices = temporary
					// A função retorna true indicando necessidade de aumentar vetor novamente.
					return true
				}
				// current é atualizado para ser o próximo valor na lista ligada.
				current = current.Next
			}
		} else {
			NovoIndice := Peso_strings(Hash.Dados_Usuario.Nome, hash_table)
			// Se o tamanho da lista de índices da tabela hash for menor ou igual a NovoIndice, o código dentro deste bloco if será executado.
			if len(hash_table.Indices) <= NovoIndice {
				temporary := make([]VetorHash, NovoIndice+1)
				copy(temporary, hash_table.Indices)
				hash_table.Indices = temporary
				// A função retorna true indicando necessidade de aumentar vetor novamente.
				return true
			}
		}
	}
	// Se o loop for terminar sem retornar true, a função retornará false, indicando que não é preciso aumentar vetor novamente.
	return false
}
