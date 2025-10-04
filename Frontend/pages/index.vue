<template>
    <div class="home-page">
        <div class="hero-section">
            <div class="hero-content">
                <h1>User Manager</h1>
                <p class="hero-description">
                    Sistema de gestion de usuarios
                </p>
                <p class="hero-subtitle">
                    Administra a tus usuarios de una manera eficiente
                </p>
                <!--Botones de acciones-->
                <div class="hero-actions">
                    <NuxtLink to="/users/" class="btn btn-primary">    
                        Ver usuarios
                    </NuxtLink>
                    <NuxtLink to="/create-user" class="btn btn-secondary">
                        Crear usuario
                    </NuxtLink>
                </div>
            </div>
        </div>
    </div>

    <!--Caracteristicas-->
    <div class="features-section">
        <div class="features-container">
            <h2>Caracteristicas principales</h2>
            <div class="features-grid">
                <div class="feature-card">
                    <div class="feature-icon">👁️</div>
                    <h3>Ver usuarios</h3>
                    <p>Consulta la lista completa de los usuarios</p>
                </div>
                <div class="feature-card">
                <div class="feature-icon">➕</div>
                    <h3>Crear Usuarios</h3>
                    <p>Registra nuevos usuarios con un formulario intuitivo y validaciones en tiempo real</p>
                </div>
          
                <div class="feature-card">
                    <div class="feature-icon">🔍</div>
                    <h3>Detalles Completos</h3>
                    <p>Accede a la información detallada de cada usuario de forma individual</p>
                </div>
            </div>
        </div>
    </div>
     <!-- Sección de tecnologías -->
    <div class="tech-section">
      <div class="tech-container">
        <h2>Tecnologías Utilizadas</h2>
        <!-- Grid de tecnologías -->
        <div class="tech-grid">
          <div class="tech-item">
            <span class="tech-name">Nuxt 3</span>
            
          </div>
          <div class="tech-item">
            <span class="tech-name">Vue 3</span>
            
          </div>
          <div class="tech-item">
            <span class="tech-name">TypeScript</span>
            
          </div>
          <div class="tech-item">
            <span class="tech-name">Go</span>
            
          </div>
          <div class="tech-item">
            <span class="tech-name">MongoDB</span>
            
          </div>
        </div>
      </div>
    </div>

</template>
<script setup lang="ts">

</script>

<style scoped>
.home-page {
  min-height: calc(100vh - 140px);
}

.hero-section {
  background: linear-gradient(135deg, #8384DF 0%, #6c6fdd 100%);
  color: white;
  padding: 4rem 2rem;
  text-align: center;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
}

.hero-content h1 {
  font-size: 3rem;
  font-weight: 700;
  margin-bottom: 1rem;
  line-height: 1.2;
}

.hero-description {
  font-size: 1.25rem;
  margin-bottom: 1rem;
  line-height: 1.6;
}

.hero-subtitle {
  font-size: 1rem;
  margin-bottom: 2.5rem;
  opacity: 0.9;
  line-height: 1.6;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 0.75rem 2rem;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
  font-size: 1rem;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

.btn-primary {
  background-color: white;
  color: #8384DF;
}

.btn-primary:hover {
  background-color: #f8f9fa;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.btn-secondary {
  background-color: transparent;
  color: white;
  border: 2px solid white;
}

.btn-secondary:hover {
  background-color: white;
  color: #8384DF;
  transform: translateY(-2px);
}

.features-section {
  padding: 4rem 2rem;
  background-color: white;
}

.features-container {
  max-width: 1200px;
  margin: 0 auto;
}

.features-container h2 {
  text-align: center;
  font-size: 2.5rem;
  color: #333;
  margin-bottom: 3rem;
  font-weight: 700;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
}

.feature-card {
  background-color: #f8f9fa;
  padding: 2rem;
  border-radius: 12px;
  text-align: center;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  border: 1px solid #e9ecef;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
}

.feature-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.feature-card h3 {
  font-size: 1.25rem;
  color: #333;
  margin-bottom: 1rem;
  font-weight: 600;
}

.feature-card p {
  color: #666;
  line-height: 1.6;
}

.tech-section {
  padding: 4rem 2rem;
  background-color: #f8f9fa;
}

.tech-container {
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
}

.tech-container h2 {
  font-size: 2.5rem;
  color: #333;
  margin-bottom: 3rem;
  font-weight: 700;
}

.tech-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1.5rem;
}

.tech-item {
  background-color: white;
  padding: 1.5rem;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  transition: transform 0.3s ease;
}

.tech-item:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.tech-name {
  display: block;
  font-weight: 600;
  color: #8384DF;
  font-size: 1.1rem;
  margin-bottom: 0.5rem;
}

.tech-desc {
  display: block;
  color: #666;
  font-size: 0.9rem;
}
</style>