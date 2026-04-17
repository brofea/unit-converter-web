<script setup lang="ts">
import { computed, ref } from 'vue'
import ConversionPage from './components/ConversionPage.vue'

type Category = '长度' | '面积' | '体积' | '重量' | '速度' | '温度' | '金钱'

const pages: { key: Category; units: string[] }[] = [
  { key: '长度', units: ['米', '千米', '英寸', '英里'] },
  { key: '面积', units: ['平方米', '平方千米', '公顷', '平方英尺'] },
  { key: '体积', units: ['升', '毫升', '立方米', '加仑'] },
  { key: '重量', units: ['克', '千克', '磅', '盎司'] },
  { key: '速度', units: ['米/秒', '千米/小时', '英里/小时', '节'] },
  { key: '温度', units: ['摄氏度', '华氏度', '开尔文'] },
  { key: '金钱', units: ['CNY', 'USD', 'EUR', 'JPY'] }
]

const activePage = ref<Category>('长度')

const activeConfig = computed(() => {
  return pages.find((page) => page.key === activePage.value) ?? pages[0]!
})

const apiBaseUrl = computed(() => {
  return import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'
})

function switchPage(page: Category): void {
  activePage.value = page
}
</script>

<template>
  <div class="layout">
    <header class="topbar">
      <h1 class="brand">Unit Converter</h1>
      <nav class="nav">
        <button
          v-for="page in pages"
          :key="page.key"
          class="nav-btn"
          :class="{ active: activePage === page.key }"
          @click="switchPage(page.key)"
        >
          {{ page.key }}
        </button>
      </nav>
    </header>

    <main class="content">
      <ConversionPage
        :category="activeConfig.key"
        :unit-options="activeConfig.units"
        :api-base-url="apiBaseUrl"
      />
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
  margin: 0 0 0.8rem;
  font-size: clamp(1.35rem, 1.25rem + 0.7vw, 1.9rem);
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.nav {
  display: flex;
  flex-wrap: wrap;
  gap: 0.55rem;
}

.nav-btn {
  border: 1px solid #8ba2c1;
  border-radius: 10px;
  background: #f8fbff;
  color: #173359;
  padding: 0.4rem 0.8rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 140ms ease;
}

.nav-btn:hover {
  border-color: #2d74d6;
  transform: translateY(-1px);
}

.nav-btn.active {
  background: #2f76da;
  color: #fff;
  border-color: #2f76da;
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
