<template>
    <div class="users-page">
        <div class="header-section">
            <h1>Administracion de usuarios</h1>
            <button @click="goToCreateUser" class="add-user-button">
                agregar nuevo usuario
            </button>
        </div>
        <!-- Estado de carga -->
        <div v-if="loading" class="loading-state">
            <p>Cargando usuarios...</p>
        </div>
        <!-- Estado de error -->
        <div v-else-if="error" class="error-state">
            <p>Error al cargar usuarios: {{ error }}</p>
            <button @click="retryLoad" class="retry-button">Reintentar</button>
        </div>

        <!--Tabla de los usuarios creados-->
        <div v-else-if="users.length >0 " class="table-container">
            <table class="users-table">
                <thead>
                    <tr>
                        <th>nombre</th>
                        <th>email</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="user in users" :key="user.id">
                        <td>{{ user.name || 'N/A' }}</td>
                        <td>{{ user.email || 'N/A' }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
         <!-- Estado vacío -->
        <div v-else class="empty-state">
            
        </div>
    </div>
</template>

<script setup lang="ts">
    import {ref,onMounted} from 'vue'
    import { getUsers } from '../services/userService';
    import type { User } from '../types/user';
    const users=ref<User[]>([]);
    const loading=ref(true);
    const error=ref<string | null>(null);

    /*carga de los usuarios */
    const loadUsers = async () => {
    try {
        loading.value = true;
        error.value = null;
        users.value = await getUsers();
    } catch (err: any) {
        if (err.message.includes('fetch')) {
        error.value = 'No se puede conectar con el servidor. Verifica que el backend esté ejecutándose en http://localhost:8080';
        } else {
        error.value = err.message || 'Error desconocido al cargar usuarios';
        }
    } finally {
        loading.value = false;
    }
        }
    const retryLoad = () => {
    loadUsers();
    };
    // Función para navegar al formulario de crear usuario
    const goToCreateUser = () => {
      window.location.href = '/create-user';
    };
    // Cargar usuarios al montar el componente
    onMounted(() => {
        loadUsers();
    });   
</script>

<style scoped>
.users-page {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  flex-wrap: wrap;
  gap: 1rem;
}

h1 {
  color: #8384DF;
  font-size: 2.5rem;
  margin: 0;
}

.add-user-button {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: background-color 0.2s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.add-user-button:hover {
  background-color: #218838;
}

.loading-state,
.error-state,
.empty-state {
  text-align: center;
  padding: 3rem;
  background-color: #f8f9fa;
  border-radius: 8px;
  margin: 2rem 0;
}

.loading-state p {
  color: #6c757d;
  font-size: 1.1rem;
}

.error-state p {
  color: #dc3545;
  font-size: 1.1rem;
  margin-bottom: 1rem;
}

.retry-button {
  background-color: #8384DF;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}

.retry-button:hover {
  background-color: #6c6fdd;
}

.empty-state p {
  color: #6c757d;
  font-size: 1.1rem;
}

.table-container {
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.users-table {
  width: 100%;
  border-collapse: collapse;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.users-table th {
  background-color: #8384DF;
  color: white;
  padding: 12px 15px;
  text-align: left;
}

.users-table td {
  padding: 12px 15px;
  border-bottom: 1px solid #e0e0e0;
}

.users-table tr:nth-child(even) {
  background-color: #f8f8ff;
}

.users-table tr:hover {
  background-color: #f0f0ff;
}

.users-table tr:last-child td {
  border-bottom: none;
}

.details-button {
  background-color: #8384DF;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.875rem;
  font-weight: 500;
  transition: background-color 0.2s ease;
}

.details-button:hover:not(:disabled) {
  background-color: #6c6fdd;
}

.details-button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
  opacity: 0.6;
}
</style>