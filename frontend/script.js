const { createApp, ref } = Vue;

createApp({
  setup() {
    const nome = ref('');
    const email = ref('');
    const cpf = ref('');
    const nascimento = ref('');
    const telefone = ref('');

    function validarEmail(email) {
      const regex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
      return regex.test(email);
    }

    function aplicarMascaraCPF(event) {
      console.log('Função CPF chamada!', event.target.value);
      let valor = event.target.value;
      valor = valor.replace(/\D/g, '');
      valor = valor.replace(/(\d{3})(\d)/, '$1.$2');
      valor = valor.replace(/(\d{3})(\d)/, '$1.$2');
      valor = valor.replace(/(\d{3})(\d{1,2})$/, '$1-$2');
      cpf.value = valor;
      console.log('CPF formatado:', valor);
    }

    function aplicarMascaraTelefone(event) {
      console.log('Função telefone chamada!', event.target.value);
      let valor = event.target.value;
      valor = valor.replace(/\D/g, '');
      valor = valor.replace(/(\d{2})(\d)/, '($1) $2');
      valor = valor.replace(/(\d{5})(\d)/, '$1-$2');
      telefone.value = valor;
      console.log('Telefone formatado:', valor);
    }

    function validarCPF(cpf) {
      cpf = cpf.replace(/\D/g, '');
      
      if (cpf.length !== 11 || /^(\d)\1+$/.test(cpf)) {
        return false;
      }

      let soma = 0;
      for (let i = 0; i < 9; i++) {
        soma += parseInt(cpf.charAt(i)) * (10 - i);
      }
      let resto = 11 - (soma % 11);
      if (resto === 10 || resto === 11) resto = 0;
      if (resto !== parseInt(cpf.charAt(9))) return false;

      soma = 0;
      for (let i = 0; i < 10; i++) {
        soma += parseInt(cpf.charAt(i)) * (11 - i);
      }
      resto = 11 - (soma % 11);
      if (resto === 10 || resto === 11) resto = 0;
      if (resto !== parseInt(cpf.charAt(10))) return false;

      return true;
    }

    async function salvarCliente() {
      if (!nome.value.trim()) {
        alert('❌ Nome é obrigatório!');
        return;
      }

      if (!validarEmail(email.value)) {
        alert('❌ Email inválido! Digite um email válido (ex: nome@dominio.com)');
        return;
      }

      const cpfLimpo = cpf.value.replace(/\D/g, '');
      if (!validarCPF(cpfLimpo)) {
        alert('❌ CPF inválido! Verifique e tente novamente.');
        return;
      }

      if (!nascimento.value) {
        alert('❌ Data de nascimento é obrigatória!');
        return;
      }

      const cliente = {
        nome: nome.value.trim(),
        email: email.value.trim(),
        cpf: cpfLimpo,
        nascimento: nascimento.value,
        telefone: telefone.value.replace(/\D/g, '')
      };

      try {
        const response = await fetch('http://localhost:8080/api/clientes', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(cliente)
        });

        if (response.ok) {
          alert('✅ Cliente cadastrado com sucesso na API!');
          nome.value = '';
          email.value = '';
          cpf.value = '';
          nascimento.value = '';
          telefone.value = '';
          window.location.href = 'cliente.html';
        } else {
          const error = await response.json();
          alert(`❌ Erro na API: ${error.error || 'Erro desconhecido'}`);
        }
      } catch (error) {
        console.log('API não disponível, salvando no localStorage');
        
        const clientesSalvos = JSON.parse(localStorage.getItem('clientes')) || [];
        
        const emailExistente = clientesSalvos.find(c => c.email === email.value.trim());
        const cpfExistente = clientesSalvos.find(c => c.cpf === cpfLimpo);

        if (emailExistente) {
          alert('❌ Email já cadastrado!');
          return;
        }

        if (cpfExistente) {
          alert('❌ CPF já cadastrado!');
          return;
        }

        clientesSalvos.push(cliente);
        localStorage.setItem('clientes', JSON.stringify(clientesSalvos));
        
        alert('✅ Cliente cadastrado com sucesso!');
        
        nome.value = '';
        email.value = '';
        cpf.value = '';
        nascimento.value = '';
        telefone.value = '';
        
        window.location.href = 'cliente.html';
      }
    }

    return {
      nome,
      email,
      cpf,
      nascimento,
      telefone,
      salvarCliente,
      aplicarMascaraCPF,
      aplicarMascaraTelefone
    };
  }
}).mount('#app');

