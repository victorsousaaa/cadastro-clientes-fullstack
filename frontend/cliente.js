const { createApp, ref, computed, onMounted } = Vue;

createApp({
  setup() {
    const clientes = ref([]);
    const termoBusca = ref('');
    const ordenacao = ref({ campo: '', direcao: 'asc' });
    const editandoCliente = ref(null);

    function formatarCPF(cpf) {
      if (!cpf) return '';
      cpf = cpf.replace(/\D/g, '');
      return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, '$1.$2.$3-$4');
    }

    function formatarTelefone(telefone) {
      if (!telefone) return '';
      telefone = telefone.replace(/\D/g, '');
      if (telefone.length === 11) {
        return telefone.replace(/(\d{2})(\d{5})(\d{4})/, '($1) $2-$3');
      }
      return telefone;
    }

    async function carregarClientes() {
      try {
        const response = await fetch('http://localhost:8080/api/clientes');
        if (response.ok) {
          clientes.value = await response.json();
          console.log('Clientes carregados da API');
        } else {
          throw new Error('API não disponível');
        }
      } catch (error) {
        console.log('Carregando do localStorage');
        const clientesSalvos = JSON.parse(localStorage.getItem('clientes')) || [];
        clientes.value = clientesSalvos;
      }
    }

    const clientesFiltrados = computed(() => {
      let resultado = clientes.value;

      if (termoBusca.value) {
        const termo = termoBusca.value.toLowerCase();
        resultado = resultado.filter(cliente => 
          cliente.nome.toLowerCase().includes(termo) ||
          cliente.email.toLowerCase().includes(termo)
        );
      }

      if (ordenacao.value.campo) {
        resultado.sort((a, b) => {
          let valorA = a[ordenacao.value.campo];
          let valorB = b[ordenacao.value.campo];
          
          if (typeof valorA === 'string') {
            valorA = valorA.toLowerCase();
            valorB = valorB.toLowerCase();
          }

          if (ordenacao.value.direcao === 'asc') {
            return valorA > valorB ? 1 : -1;
          } else {
            return valorA < valorB ? 1 : -1;
          }
        });
      }

      return resultado;
    });

    function ordenarPor(campo) {
      if (ordenacao.value.campo === campo) {
        ordenacao.value.direcao = ordenacao.value.direcao === 'asc' ? 'desc' : 'asc';
      } else {
        ordenacao.value.campo = campo;
        ordenacao.value.direcao = 'asc';
      }
    }

    function limparBusca() {
      termoBusca.value = '';
    }

    function editarCliente(cliente) {
      editandoCliente.value = { ...cliente };
    }

    function cancelarEdicao() {
      editandoCliente.value = null;
    }

    async function salvarEdicao() {
      try {
        const response = await fetch(`http://localhost:8080/api/clientes/${editandoCliente.value.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(editandoCliente.value)
        });

        if (response.ok) {
          await carregarClientes();
          editandoCliente.value = null;
          alert('✅ Cliente atualizado com sucesso na API!');
        } else {
          throw new Error('Erro na API');
        }
      } catch (error) {
        const clientesSalvos = JSON.parse(localStorage.getItem('clientes')) || [];
        const index = clientesSalvos.findIndex(c => c.cpf === editandoCliente.value.cpf);
        if (index !== -1) {
          clientesSalvos[index] = editandoCliente.value;
          localStorage.setItem('clientes', JSON.stringify(clientesSalvos));
          clientes.value = clientesSalvos;
          editandoCliente.value = null;
          alert('✅ Cliente atualizado com sucesso!');
        }
      }
    }

    async function removerCliente(cliente) {
      if (confirm(`Tem certeza que deseja remover ${cliente.nome}?`)) {
        try {
          const response = await fetch(`http://localhost:8080/api/clientes/${cliente.id}`, {
            method: 'DELETE'
          });

          if (response.ok) {
            await carregarClientes();
            alert('✅ Cliente removido com sucesso da API!');
          } else {
            throw new Error('Erro na API');
          }
        } catch (error) {
          const clientesSalvos = JSON.parse(localStorage.getItem('clientes')) || [];
          const novaLista = clientesSalvos.filter(c => c.cpf !== cliente.cpf);
          localStorage.setItem('clientes', JSON.stringify(novaLista));
          clientes.value = novaLista;
          alert('✅ Cliente removido com sucesso!');
        }
      }
    }

    onMounted(() => {
      carregarClientes();
    });

    return {
      clientes,
      clientesFiltrados,
      termoBusca,
      ordenacao,
      editandoCliente,
      formatarCPF,
      formatarTelefone,
      ordenarPor,
      limparBusca,
      editarCliente,
      cancelarEdicao,
      salvarEdicao,
      removerCliente
    };
  }
}).mount('#app');