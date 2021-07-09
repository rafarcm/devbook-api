$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);
$('#deletar-usuario').on('click', deletarUsuario);

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
        $("#parar-de-seguir").prop('disabled', false);
    });    
}

function seguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao seguir o usuário!", "error");
        $("#seguir").prop('disabled', false);
    });    
}

function editar(evento) {
    evento.preventDefault();

    $.ajax({
        url: `/editar-usuario`,
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
        .then(function () {
            window.location = "/perfil";
        });
    }).fail(function() {success
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
    });
}

function atualizarSenha(evento) {
    evento.preventDefault();

    if($('#senha-nova').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return;
    }

    $.ajax({
        url: `/atualizar-senha`,
        method: "POST",
        data: {
            atual: $('#senha-atual').val(),
            nova: $('#senha-nova').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Senha atualizada com sucesso!", "success")
        .then(function () {
            window.location = "/perfil";
        });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar senha!", "error");
    });
}

function deletarUsuario() {
    Swal.fire({
        title: "Atenção",
        text: "Tem certeza que deseja excluir a sua conta? Essa é uma ação irreversível",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if(!confirmacao.value) return;

        $.ajax({
            url: "/deletar-usuario",
            method: "DELETE"
        }).done(function() {
            Swal.fire("Sucesso!", "Seu usuário foi excluido com sucesso!", "success")
            .then(function () {
                window.location = "/logout";
            });
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a conta do usuário!", "error");
        });
    });
}