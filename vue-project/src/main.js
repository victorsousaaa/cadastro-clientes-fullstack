import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'

createApp(App).mount('#app')
const { createApp, ref } = Vue

createApp({
  setup() {
    const nome = ref('')
    const email = ref('')
    const cpf = ref('')
    const nascimento = ref('')
    const telefone = ref('')

    const salvarCliente = () => {
      console.log('Cliente cadastrado:')
      console.log({
        nome: nome.value,
        email: email.value,
        cpf: cpf.value,
        nascimento: nascimento.value,
        telefone: telefone.value
      })
    }

    return {
      nome,
      email,
      cpf,
      nascimento,
      telefone,
      salvarCliente
    }
  }
}).mount('#app')
