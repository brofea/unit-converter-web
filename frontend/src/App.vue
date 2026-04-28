<script setup lang="ts">
import { computed, ref } from 'vue'
import ConversionPage from './components/ConversionPage.vue'

type Category = '长度' | '重量' | '面积' | '体积' | '时间' | '温度'

const pages: Array<{ key: Category; units: Array<{ label: string; value: string }> }> = [
  {
    key: '长度',
    units: [
      { label: '米', value: 'm' },
      { label: '厘米', value: 'cm' },
      { label: '毫米', value: 'mm' },
      { label: '千米', value: 'km' },
      { label: '英寸', value: 'inch' },
      { label: '英尺', value: 'ft' },
      { label: '码', value: 'yd' },
      { label: '英里', value: 'mi' }
    ]
  },
  {
    key: '重量',
    units: [
      { label: '千克', value: 'kg' },
      { label: '克', value: 'g' },
      { label: '毫克', value: 'mg' },
      { label: '磅', value: 'lb' },
      { label: '盎司', value: 'oz' }
    ]
  },
  {
    key: '面积',
    units: [
      { label: '平方米', value: 'm2' },
      { label: '平方厘米', value: 'cm2' },
      { label: '平方毫米', value: 'mm2' },
      { label: '平方千米', value: 'km2' },
      { label: '公顷', value: 'ha' },
      { label: '英亩', value: 'ac' },
      { label: '平方英尺', value: 'sqft' },
      { label: '平方英寸', value: 'sqin' }
    ]
  },
  {
    key: '体积',
    units: [
      { label: '立方米', value: 'm3' },
      { label: '升', value: 'l' },
      { label: '毫升', value: 'ml' },
      { label: '加仑', value: 'gal' },
      { label: '夸脱', value: 'qt' }
    ]
  },
  {
    key: '时间',
    units: [
      { label: '秒', value: 's' },
      { label: '分钟', value: 'min' },
      { label: '小时', value: 'h' },
      { label: '天', value: 'd' }
    ]
  },
  {
    key: '温度',
    units: [
      { label: '摄氏度', value: 'cel' },
      { label: '华氏度', value: 'degf' },
      { label: '开尔文', value: 'kel' }
    ]
  }
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
          type="button"
          @click="switchPage(page.key)"
        >
          {{ page.key }}
        </button>
      </nav>
    </header>

    <main class="content">
      <ConversionPage
        :key="activeConfig.key"
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
  margin: 0 0 0.4rem;
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
