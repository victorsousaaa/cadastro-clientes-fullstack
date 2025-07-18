<template>
  <div class="box">
    <form @submit.prevent="salvarCliente">
      <fieldset>
        <legend><b>CADASTRO DE CLIENTES</b></legend>
        <div class="inputbox">
          <input type="text" v-model="nome" class="inputUser" required />
          <label>NOME COMPLETO</label>
        </div>
        <div class="inputbox">
          <input type="email" v-model="email" class="inputUser" required />
          <label>EMAIL</label>
        </div>
        <div class="inputbox">
          <input type="text" v-model="cpf" class="inputUser" required />
          <label>CPF</label>
        </div>
        <div class="inputbox">
          <input type="date" v-model="nascimento" class="inputUser" required />
          <label>DATA DE NASCIMENTO</label>
        </div>
        <div class="inputbox">
          <input type="tel" v-model="telefone" class="inputUser" />
          <label>TELEFONE</label>
        </div>
        <button type="submit" class="btn-submit">Cadastrar</button>
      </fieldset>
    </form>
  </div>
</template>

<script setup>
const nome = ref('');
const email = ref('');
const cpf = ref('');
const nascimento = ref('');
const telefone = ref('');

function validarCPF(cpf) {
  // Lógica simples de validação (pode ser melhorada)
  return cpf.length === 11;
}

async function salvarCliente() {
  const cpfFormatado = cpf.value.replace(/\D/g, '');

  if (!validarCPF(cpfFormatado)) {
    alert('CPF inválido. Verifique e tente novamente.');
    return;
  }

  const clientesSalvos = JSON.parse(localStorage.getItem('clientes') || '[]');
  const emailExistente = clientesSalvos.find(c => c.email === email.value);
  const cpfExistente = clientesSalvos.find(c => c.cpf === cpfFormatado);

  if (emailExistente || cpfExistente) {
    alert('Email ou CPF já cadastrados!');
    return;
  }

  const cliente = {
    nome: nome.value,
    email: email.value,
    cpf: cpfFormatado,
    nascimento: nascimento.value,
    telefone: telefone.value
  };

  clientesSalvos.push(cliente);
  localStorage.setItem('clientes', JSON.stringify(clientesSalvos));
  
  // Redirecionar para página de clientes
  await navigateTo('/clientes');
}
</script>

<style scoped>
/* Estilos do formulário serão importados do CSS global */
</style>
