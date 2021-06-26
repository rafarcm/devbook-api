package repository

import (
	"api/src/models"
	"database/sql"
)

// publicacoes representa um repositório de publicações
type publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *publicacoes {
	return &publicacoes{db}
}

// Criar adiciona uma nova publicação no banco de dados
func (repositorio publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz as publicações que aparecem no feed do usuário
func (repositorio publicacoes) Buscar(ID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`select distinct p.*, u.nick 
		from publicacoes p 
		inner join usuarios u on p.autor_id = u.id 
		inner join seguidores s on p.autor_id = s.usuario_id 
		where p.id = ? or s.seguidor_id = ? 
		order by 1 desc;`,
		ID, ID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// BuscarPorID traz uma publicação de um usuário do banco de dados com o mesmo id passado
func (repositorio publicacoes) BuscarPorID(ID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		`select p.*, u.nick 
		from publicacoes p 
		inner join usuarios u on p.autor_id = u.id 
		where p.id = ?`,
		ID)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Atualizar altera os dados de uma publicação do banco de dados
func (repositorio publicacoes) Atualizar(ID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		"update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio publicacoes) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUsuario traz todas as publicações de um dado usuário
func (repositorio publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`select p.*, u.nick 
		from publicacoes p 
		inner join usuarios u on p.autor_id = u.id 
		where p.autor_id = ?`,
		usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio publicacoes) Curtir(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida na publicação
func (repositorio publicacoes) Descurtir(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes set curtidas = 
		case when curtidas > 0 then curtidas - 1 
		else 0 end
		where id = ?`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
