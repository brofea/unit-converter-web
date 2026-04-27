<script setup lang="ts">
import { computed } from 'vue'
import ConversionPage from './components/ConversionPage.vue'

const apiBaseUrl = computed(() => {
  return import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'
})
</script>

<template>
  <div class="layout">
    <header class="topbar">
      <h1 class="brand">Unit Converter</h1>
      <p class="subtitle">当前版本仅支持长度转换</p>
    </header>

    <main class="content">
      <ConversionPage :api-base-url="apiBaseUrl" />
    </main>
  </div>
</template>

<style scoped>
:global(body) {
  margin: 0;
  font-family: 'IBM Plex Sans', 'Noto Sans SC', 'Source Han Sans SC', sans-serif;
  color: #10213a;
  background:
    radial-gradient(circle at 90% 10%, rgba(66, 148, 255, 0.22), transparent 40%),
    radial-gradient(circle at 10% 80%, rgba(18, 70, 130, 0.24), transparent 45%),
    linear-gradient(160deg, #eaf3ff 0%, #f7fbff 45%, #e4eef9 100%);
  min-height: 100vh;
}

.layout {
  min-height: 100vh;
}

.topbar {
  position: sticky;
  top: 0;
  z-index: 10;
  padding: 1rem clamp(1rem, 2vw, 2.2rem);
  border-bottom: 1px solid rgba(12, 31, 58, 0.16);
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(8px);
}

.brand {
  margin: 0 0 0.4rem;
  font-size: clamp(1.35rem, 1.25rem + 0.7vw, 1.9rem);
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.subtitle {
  margin: 0;
  color: #406188;
  font-size: 0.96rem;
}

.content {
  box-sizing: border-box;
  width: 100%;
  padding: clamp(1rem, 2.8vw, 2.2rem);
  animation: fade-in 360ms ease;
}

@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
