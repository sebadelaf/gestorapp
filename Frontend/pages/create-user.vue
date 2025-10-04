<template>
    <div class="create-user-page">
        <div class="header-section">
            <button @click="goBack" class="back-button">
                <- Volver a usuarios
            </button>
            <h1>Crear nuevo usuario</h1>
        </div>
        <!--Formulario para crear un nuevo usuario-->
        <div class="form-container">
            <form @submit.prevent="handleSubmit" class="user-form">
                <!-- Campo Nombre -->
                <div class="form-group">
                <label for="name">Nombre Completo *</label>
                <input
                    id="name"
                    v-model="form.name"
                    type="text"
                    required
                    placeholder="Ingresa el nombre completo"
                    :disabled="isSubmitting"
                />
                <span v-if="errors.name" class="error-message">{{ errors.name }}</span>
                </div>
                <!-- Campo Email -->
                <div class="form-group">
                <label for="email">Correo Electrónico *</label>
                <input
                    id="email"
                    v-model="form.email"
                    type="email"
                    required
                    placeholder="ejemplo@correo.com"
                    :disabled="isSubmitting"
                />
                </div>
                
                <!--Campo contraseña-->
                <div class="form-group">
                <label for="password">Contraseña*</label>
                <input
                    id="password"
                    v-model="form.password"
                    type="password"
                    required
                    placeholder="1234$"
                    :disabled="isSubmitting"
                />
                <span v-if="errors.password" class="error-message">{{ errors.password }}</span>
                </div>

                 <!-- Mensaje de error general -->
                <div v-if="generalError" class="general-error">
                {{ generalError }}
                </div>

                <!-- Mensaje de éxito -->
                <div v-if="successMessage" class="success-message">
                {{ successMessage }}
                </div>

                 <!-- Botones -->
                <div class="form-actions">
                <button 
                    type="button" 
                    @click="goBack" 
                    class="cancel-button"
                    :disabled="isSubmitting"
                >
                    Cancelar
                </button>
                <button 
                    type="submit" 
                    class="submit-button"
                    :disabled="isSubmitting || !isFormValid"
                >
                    <span v-if="isSubmitting">Creando...</span>
                    <span v-else>✅ Crear Usuario</span>
                </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
    import {ref,computed} from 'vue'
    import type { CreateUserDto } from '~/types/user';
    import { createUser } from '~/services/userService';
    const form=ref<CreateUserDto>({
        name:'',
        email:'',
        password:''
    })
    const errors = ref<Partial<CreateUserDto>>({});
    const generalError = ref<string>('');
    const successMessage = ref<string>('');
    const isSubmitting = ref(false);

    // Validación del formulario
    const isFormValid = computed(() => {
    return form.value.name.trim() !== '' && 
            form.value.email.trim() !== '' && 
            isValidEmail(form.value.email);
    });

    const isValidEmail = (email: string): boolean => {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    };
    const clearErrors = () => {
        errors.value = {};
        generalError.value = '';
    };

    const validateForm=():boolean=>{
        clearErrors();
        let isValid=true
         // Validar nombre
        if (!form.value.name.trim()) {
            errors.value.name = 'El nombre es obligatorio';
            isValid = false;
        } else if (form.value.name.trim().length < 2) {
            errors.value.name = 'El nombre debe tener al menos 2 caracteres';
            isValid = false;
        }
         // Validar email
        if (!form.value.email.trim()) {
            errors.value.email = 'El email es obligatorio';
            isValid = false;
        } else if (!isValidEmail(form.value.email)) {
            errors.value.email = 'Ingresa un email válido';
            isValid = false;
        }
        // Validar la contraseña

        if(!form.value.password.trim()){
            errors.value.password='La contraseña es obligatoria';
            isValid=false
        }

        return isValid
    }

    // funcion para subir el formulario al backend
    const handleSubmit=async()=>{
        if(!validateForm()){
            return;
        }
        isSubmitting.value=true;
        clearErrors()
        try{
            const newUser = await createUser(form.value);
            successMessage.value='Usuario creado exitosamente'
            form.value={name:'',email:'',password:''};
            setTimeout(() => {
            goBack();
            }, 2000);
        } catch(error:any){
            generalError.value=error.message || 'error al crear el usuario'
        } finally{
            isSubmitting.value=false
        }
    }
    const goBack=()=>{
        window.location.href='/users'
    }

</script>

<style scoped>
.create-user-page {
  padding: 2rem;
  max-width: 600px;
  margin: 0 auto;
}

.header-section {
  margin-bottom: 2rem;
}

.header-section h1 {
  color: #8384DF;
  margin: 1rem 0;
  font-size: 2.5rem;
  text-align: center;
}

.back-button {
  background-color: #6c757d;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: background-color 0.2s ease;
  margin-bottom: 1rem;
}

.back-button:hover {
  background-color: #5a6268;
}

.form-container {
  background-color: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border: 1px solid #e9ecef;
}

.user-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group label {
  font-weight: 600;
  color: #333;
  font-size: 0.9rem;
}

.form-group input {
  padding: 0.75rem;
  border: 2px solid #e9ecef;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.2s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #8384DF;
  box-shadow: 0 0 0 3px rgba(131, 132, 223, 0.1);
}

.form-group input:disabled {
  background-color: #f8f9fa;
  cursor: not-allowed;
  opacity: 0.7;
}

.error-message {
  color: #dc3545;
  font-size: 0.875rem;
  font-weight: 500;
}

.general-error {
  background-color: #f8d7da;
  color: #721c24;
  padding: 0.75rem;
  border-radius: 6px;
  border: 1px solid #f5c6cb;
  font-weight: 500;
}

.success-message {
  background-color: #d4edda;
  color: #155724;
  padding: 0.75rem;
  border-radius: 6px;
  border: 1px solid #c3e6cb;
  font-weight: 500;
  text-align: center;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 1rem;
}

.cancel-button {
  background-color: #f8f9fa;
  color: #6c757d;
  border: 2px solid #6c757d;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: all 0.2s ease;
}

.cancel-button:hover:not(:disabled) {
  background-color: #6c757d;
  color: white;
}

.cancel-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.submit-button {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: background-color 0.2s ease;
}

.submit-button:hover:not(:disabled) {
  background-color: #218838;
}

.submit-button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
  opacity: 0.6;
}
</style>