<template>
  <div class="box">
    <h2>Clientes Cadastrados</h2>
    
    <!-- Campo de busca -->
    <div style="margin-bottom: 20px;">
      <input 
        v-model="termoBusca" 
        type="text" 
        placeholder="Buscar por nome ou email..." 
        style="width: 300px; padding: 8px; border: 1px solid #ccc; border-radius: 4px; font-size: 14px;"
      >
      <button v-if="termoBusca" @click="limparBusca" style="margin-left: 10px;">Limpar</button>
    </div>
    
    <table>
      <thead>
        <tr>
          <th @click="ordenarPor('nome')" style="cursor: pointer;">
            Nome {{ ordemAtual === 'nome' ? (crescente ? '↑' : '↓') : '' }}
          </th>
          <th @click="ordenarPor('email')" style="cursor: pointer;">
            Email {{ ordemAtual === 'email' ? (crescente ? '↑' : '↓') : '' }}
          </th>
          <th>CPF</th>
          <th>Nascimento</th>
          <th>Telefone</th>
          <th>Ações</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(cliente, index) in clientesFiltrados" :key="cliente.cpf">
          <td v-if="!cliente.editando">{{ cliente.nome }}</td>
          <td v-else><input v-model="cliente.nome" type="text"></td>
          
          <td v-if="!cliente.editando">{{ cliente.email }}</td>
          <td v-else><input v-model="cliente.email" type="email"></td>
          
          <td>{{ cliente.cpf }}</td>
          
          <td v-if="!cliente.editando">{{ cliente.nascimento }}</td>
          <td v-else><input v-model="cliente.nascimento" type="date"></td>
          
          <td v-if="!cliente.editando">{{ cliente.telefone }}</td>
          <td v-else><input v-model="cliente.telefone" type="tel"></td>
          
          <td>
            <button v-if="!cliente.editando" @click="editarCliente(index)" style="margin-right: 5px;">Editar</button>
            <button v-if="cliente.editando" @click="salvarEdicao(index)" style="margin-right: 5px;">Salvar</button>
            <button v-if="cliente.editando" @click="cancelarEdicao(index)" style="margin-right: 5px;">Cancelar</button>
            <button @click="removerCliente(index)" style="background: #ff4444; color: white;">Remover</button>
          </td>
        </tr>
      </tbody>
    </table>
    
    <NuxtLink to="/">← Voltar ao cadastro</NuxtLink>
  </div>
</template>

<script setup>
const clientes = ref([]);
const ordemAtual = ref('');
const crescente = ref(true);
const clientesBackup = ref([]);
const termoBusca = ref('');

onMounted(() => {
  // Busca os clientes cadastrados no localStorage
  const dados = localStorage.getItem('clientes');
  if (dados) {
    try {
      clientes.value = JSON.parse(dados);
    } catch (e) {
      clientes.value = [];
    }
  }
});

const clientesOrdenados = computed(() => {
  if (!ordemAtual.value) return clientes.value;
  
  return [...clientes.value].sort((a, b) => {
    const valorA = a[ordemAtual.value].toLowerCase();
    const valorB = b[ordemAtual.value].toLowerCase();
    
    if (crescente.value) {
      return valorA < valorB ? -1 : valorA > valorB ? 1 : 0;
    } else {
      return valorA > valorB ? -1 : valorA < valorB ? 1 : 0;
    }
  });
});

const clientesFiltrados = computed(() => {
  const clientesParaFiltrar = clientesOrdenados.value;
  
  if (!termoBusca.value) return clientesParaFiltrar;
  
  const termo = termoBusca.value.toLowerCase();
  return clientesParaFiltrar.filter(cliente => 
    cliente.nome.toLowerCase().includes(termo) || 
    cliente.email.toLowerCase().includes(termo)
  );
});

function ordenarPor(campo) {
  if (ordemAtual.value === campo) {
    crescente.value = !crescente.value;
  } else {
    ordemAtual.value = campo;
    crescente.value = true;
  }
}

function limparBusca() {
  termoBusca.value = '';
}

function editarCliente(index) {
  // Fazer backup do cliente original
  const clienteOriginal = clientesFiltrados.value[index];
  const indexReal = clientes.value.findIndex(c => c.cpf === clienteOriginal.cpf);
  clientesBackup.value[indexReal] = { ...clientes.value[indexReal] };
  clientes.value[indexReal].editando = true;
}

function salvarEdicao(index) {
  const clienteOriginal = clientesFiltrados.value[index];
  const indexReal = clientes.value.findIndex(c => c.cpf === clienteOriginal.cpf);
  clientes.value[indexReal].editando = false;
  localStorage.setItem('clientes', JSON.stringify(clientes.value));
  delete clientesBackup.value[indexReal];
}

function cancelarEdicao(index) {
  // Restaurar dados originais
  const clienteOriginal = clientesFiltrados.value[index];
  const indexReal = clientes.value.findIndex(c => c.cpf === clienteOriginal.cpf);
  if (clientesBackup.value[indexReal]) {
    Object.assign(clientes.value[indexReal], clientesBackup.value[indexReal]);
    delete clientesBackup.value[indexReal];
  }
  clientes.value[indexReal].editando = false;
}

function removerCliente(index) {
  if (confirm('Tem certeza que deseja remover este cliente?')) {
    const clienteOriginal = clientesFiltrados.value[index];
    const indexReal = clientes.value.findIndex(c => c.cpf === clienteOriginal.cpf);
    clientes.value.splice(indexReal, 1);
    localStorage.setItem('clientes', JSON.stringify(clientes.value));
  }
}
</script>

<style scoped>
/* Estilos da tabela serão importados do CSS global */
</style>
